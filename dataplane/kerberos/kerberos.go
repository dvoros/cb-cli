package kerberos

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1env"
	"github.com/hortonworks/cb-cli/dataplane/api-freeipa/client/v1kerberos"
	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/oauth"
	"github.com/hortonworks/dp-cli-common/utils"
	"github.com/urfave/cli"
)

var Header = []string{"Name", "Description", "Type", "Environments", "ID"}

type kerberos struct {
	Name          string `json:"Name" yaml:"Name"`
	Description   string `json:"Description" yaml:"Description"`
	Type          string `json:"Type" yaml:"Type"`
	EnvironmentID string
}

type kerberosOutDescribe struct {
	*kerberos
	ID string `json:"ID" yaml:"ID"`
}

type freeIpaKerberosClient interface {
	CreateKerberosConfigForEnvironment(params *v1kerberos.CreateKerberosConfigForEnvironmentParams) (*v1kerberos.CreateKerberosConfigForEnvironmentOK, error)
	GetKerberosConfigForEnvironment(params *v1kerberos.GetKerberosConfigForEnvironmentParams) (*v1kerberos.GetKerberosConfigForEnvironmentOK, error)
	DeleteKerberosConfigForEnvironment(params *v1kerberos.DeleteKerberosConfigForEnvironmentParams) error
}

func (k *kerberos) DataAsStringArray() []string {
	return []string{k.Name, k.Description, k.Type, k.EnvironmentID}
}

func (k *kerberosOutDescribe) DataAsStringArray() []string {
	return append(k.kerberos.DataAsStringArray(), k.ID)
}

func GetVerifyKdcTrustFlag(c *cli.Context) bool {
	return !c.Bool(fl.FlKerberosDisableVerifyKdcTrust.Name)
}

func CreateAdKerberos(c *cli.Context) error {
	admin := c.String(fl.FlKerberosAdmin.Name)
	verifyKdcTrust := GetVerifyKdcTrustFlag(c)
	domain := c.String(fl.FlKerberosDomain.Name)
	nameServers := c.String(fl.FlKerberosNameServers.Name)
	password := c.String(fl.FlKerberosPassword.Name)
	tcpAllowed := c.Bool(fl.FlKerberosTcpAllowed.Name)
	principal := c.String(fl.FlKerberosPrincipal.Name)
	url := c.String(fl.FlKerberosUrl.Name)
	adminURL := c.String(fl.FlKerberosAdminUrl.Name)
	realm := c.String(fl.FlKerberosRealm.Name)
	ldapURL := c.String(fl.FlKerberosLdapUrl.Name)
	containerDn := c.String(fl.FlKerberosContainerDn.Name)

	adRequest := model.ActiveDirectoryKerberosV1Descriptor{
		VerifyKdcTrust: verifyKdcTrust,
		Domain:         domain,
		NameServers:    nameServers,
		Password:       &password,
		TCPAllowed:     &tcpAllowed,
		Principal:      &principal,
		URL:            &url,
		AdminURL:       &adminURL,
		Realm:          &realm,
		LdapURL:        &ldapURL,
		ContainerDn:    &containerDn,
		Admin:          admin,
	}

	kerberosRequest := CreateKerberosRequest(c)
	kerberosRequest.ActiveDirectory = &adRequest
	return SendCreateKerberosRequest(c, &kerberosRequest)
}

func CreateMitKerberos(c *cli.Context) error {
	admin := c.String(fl.FlKerberosAdmin.Name)
	verifyKdcTrust := GetVerifyKdcTrustFlag(c)
	domain := c.String(fl.FlKerberosDomain.Name)
	nameServers := c.String(fl.FlKerberosNameServers.Name)
	password := c.String(fl.FlKerberosPassword.Name)
	tcpAllowed := c.Bool(fl.FlKerberosTcpAllowed.Name)
	principal := c.String(fl.FlKerberosPrincipal.Name)
	url := c.String(fl.FlKerberosUrl.Name)
	adminURL := c.String(fl.FlKerberosAdminUrl.Name)
	realm := c.String(fl.FlKerberosRealm.Name)

	mitRequest := model.MITKerberosV1Descriptor{
		VerifyKdcTrust: verifyKdcTrust,
		Domain:         domain,
		NameServers:    nameServers,
		Password:       &password,
		TCPAllowed:     &tcpAllowed,
		Principal:      &principal,
		URL:            &url,
		AdminURL:       &adminURL,
		Realm:          &realm,
		Admin:          admin,
	}

	kerberosRequest := CreateKerberosRequest(c)
	kerberosRequest.Mit = &mitRequest
	return SendCreateKerberosRequest(c, &kerberosRequest)
}

func CreateFreeIpaKerberos(c *cli.Context) error {
	admin := c.String(fl.FlKerberosAdmin.Name)
	verifyKdcTrust := GetVerifyKdcTrustFlag(c)
	domain := c.String(fl.FlKerberosDomain.Name)
	nameServers := c.String(fl.FlKerberosNameServers.Name)
	password := c.String(fl.FlKerberosPassword.Name)
	tcpAllowed := c.Bool(fl.FlKerberosTcpAllowed.Name)
	principal := c.String(fl.FlKerberosPrincipal.Name)
	url := c.String(fl.FlKerberosUrl.Name)
	adminURL := c.String(fl.FlKerberosAdminUrl.Name)
	realm := c.String(fl.FlKerberosRealm.Name)

	freeIpaRequest := model.FreeIPAKerberosV1Descriptor{
		VerifyKdcTrust: verifyKdcTrust,
		Domain:         domain,
		NameServers:    nameServers,
		Password:       &password,
		TCPAllowed:     &tcpAllowed,
		Principal:      &principal,
		URL:            &url,
		AdminURL:       &adminURL,
		Realm:          &realm,
		Admin:          admin,
	}

	kerberosRequest := CreateKerberosRequest(c)
	kerberosRequest.FreeIpa = &freeIpaRequest
	return SendCreateKerberosRequest(c, &kerberosRequest)
}

func getEnvirontmentCrnByName(c *cli.Context) string {
	environment := c.String(fl.FlEnvironmentName.Name)
	envClient := oauth.Environment(*oauth.NewEnvironmentClientFromContext(c)).Environment
	resp, err := envClient.V1env.GetEnvironmentV1(v1env.NewGetEnvironmentV1Params().WithName(environment))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	environmentDetails := resp.Payload
	if len(environmentDetails.ID) == 0 {
		errmsg := fmt.Sprintf("Failed to get the CRN of environment: %s", environment)
		utils.LogErrorMessageAndExit(errmsg)
	}
	return environmentDetails.ID
}

func CreateKerberosRequest(c *cli.Context) model.CreateKerberosConfigV1Request {
	kerberosName := c.String(fl.FlName.Name)
	description := c.String(fl.FlDescriptionOptional.Name)
	environment := getEnvirontmentCrnByName(c)

	kerberosRequest := &model.CreateKerberosConfigV1Request{
		Name:          &kerberosName,
		Description:   &description,
		EnvironmentID: &environment,
	}

	return *kerberosRequest
}

func SendCreateKerberosRequest(c *cli.Context, request *model.CreateKerberosConfigV1Request) error {
	defer utils.TimeTrack(time.Now(), "create kerberos")
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	freeIpaClient := oauth.FreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	return SendCreateKerberosRequestImpl(freeIpaClient.V1kerberos, workspaceID, request, output.Write)
}

func SendCreateKerberosRequestImpl(kerberosClient freeIpaKerberosClient, workspaceID int64, request *model.CreateKerberosConfigV1Request, writer func([]string, utils.Row)) error {
	resp, err := kerberosClient.CreateKerberosConfigForEnvironment(v1kerberos.NewCreateKerberosConfigForEnvironmentParams().WithBody(request))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	kerberosResponse := resp.Payload
	writeResponse(writer, kerberosResponse)
	log.Infof("[SendCreateKerberosRequestImpl] kerberos created with name: %s, id: %s", kerberosResponse.Name, kerberosResponse.ID)
	return nil
}

func GetKerberos(c *cli.Context) error {
	defer utils.TimeTrack(time.Now(), "describe a kerberos")
	output := utils.Output{Format: c.String(fl.FlOutputOptional.Name)}
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	environmentName := getEnvirontmentCrnByName(c)
	freeIpaClient := oauth.FreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	return GetKerberosImpl(freeIpaClient.V1kerberos, workspaceID, environmentName, output.Write)
}

func GetKerberosImpl(kerberosClient freeIpaKerberosClient, workspaceID int64, environmentName string, writer func([]string, utils.Row)) error {
	log.Infof("[GetKerberos] describe kerberos config from environment: %s", environmentName)
	resp, err := kerberosClient.GetKerberosConfigForEnvironment(v1kerberos.NewGetKerberosConfigForEnvironmentParams().WithEnvironmentID(&environmentName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	kerberosResponse := resp.Payload
	writeResponse(writer, kerberosResponse)
	return nil
}

func DeleteKerberos(c *cli.Context) error {
	defer utils.TimeTrack(time.Now(), "delete a kerberos")
	workspaceID := c.Int64(fl.FlWorkspaceOptional.Name)
	environmentName := getEnvirontmentCrnByName(c)
	freeIpaClient := oauth.FreeIpa(*oauth.NewFreeIpaClientFromContext(c)).FreeIpa
	return DeleteKerberosImpl(freeIpaClient.V1kerberos, workspaceID, environmentName)
}

func DeleteKerberosImpl(kerberosClient freeIpaKerberosClient, workspaceID int64, environmentName string) error {
	log.Infof("[DeleteKerberosImpl] delete kerberos config by name: %s", environmentName)
	err := kerberosClient.DeleteKerberosConfigForEnvironment(v1kerberos.NewDeleteKerberosConfigForEnvironmentParams().WithEnvironmentID(&environmentName))
	if err != nil {
		utils.LogErrorAndExit(err)
	}
	log.Infof("[DeleteKerberosImpl] kerberos config deleted from environment: %s", environmentName)
	return nil
}

func writeResponse(writer func([]string, utils.Row), kerberosResponse *model.DescribeKerberosConfigV1Response) {
	writer(append(Header), &kerberosOutDescribe{
		&kerberos{
			Name:          kerberosResponse.Name,
			Description:   utils.SafeStringConvert(kerberosResponse.Description),
			EnvironmentID: *kerberosResponse.EnvironmentID,
			Type:          utils.SafeStringConvert(kerberosResponse.Type),
		},
		kerberosResponse.ID})
}
