package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/client/accountpreferences"
	"github.com/hortonworks/hdc-cli/client/blueprints"
	"github.com/hortonworks/hdc-cli/client/cluster"
	"github.com/hortonworks/hdc-cli/client/constraints"
	"github.com/hortonworks/hdc-cli/client/credentials"
	"github.com/hortonworks/hdc-cli/client/events"
	"github.com/hortonworks/hdc-cli/client/ldap"
	"github.com/hortonworks/hdc-cli/client/networks"
	"github.com/hortonworks/hdc-cli/client/rdsconfigs"
	"github.com/hortonworks/hdc-cli/client/recipes"
	"github.com/hortonworks/hdc-cli/client/securitygroups"
	"github.com/hortonworks/hdc-cli/client/sssd"
	"github.com/hortonworks/hdc-cli/client/stacks"
	"github.com/hortonworks/hdc-cli/client/templates"
	"github.com/hortonworks/hdc-cli/client/topologies"
	"github.com/hortonworks/hdc-cli/client/usages"
	"github.com/hortonworks/hdc-cli/client/users"
	"github.com/hortonworks/hdc-cli/client/util"
)

// Default cloudbreak HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new cloudbreak HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Cloudbreak {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost", "/", []string{"http", "https"})
	return New(transport, formats)
}

// New creates a new cloudbreak client
func New(transport client.Transport, formats strfmt.Registry) *Cloudbreak {
	cli := new(Cloudbreak)
	cli.Transport = transport

	cli.Accountpreferences = accountpreferences.New(transport, formats)

	cli.Blueprints = blueprints.New(transport, formats)

	cli.Cluster = cluster.New(transport, formats)

	cli.Constraints = constraints.New(transport, formats)

	cli.Credentials = credentials.New(transport, formats)

	cli.Events = events.New(transport, formats)

	cli.Ldap = ldap.New(transport, formats)

	cli.Networks = networks.New(transport, formats)

	cli.Rdsconfigs = rdsconfigs.New(transport, formats)

	cli.Recipes = recipes.New(transport, formats)

	cli.Securitygroups = securitygroups.New(transport, formats)

	cli.Sssd = sssd.New(transport, formats)

	cli.Stacks = stacks.New(transport, formats)

	cli.Templates = templates.New(transport, formats)

	cli.Topologies = topologies.New(transport, formats)

	cli.Usages = usages.New(transport, formats)

	cli.Users = users.New(transport, formats)

	cli.Util = util.New(transport, formats)

	return cli
}

// Cloudbreak is a client for cloudbreak
type Cloudbreak struct {
	Accountpreferences *accountpreferences.Client

	Blueprints *blueprints.Client

	Cluster *cluster.Client

	Constraints *constraints.Client

	Credentials *credentials.Client

	Events *events.Client

	Ldap *ldap.Client

	Networks *networks.Client

	Rdsconfigs *rdsconfigs.Client

	Recipes *recipes.Client

	Securitygroups *securitygroups.Client

	Sssd *sssd.Client

	Stacks *stacks.Client

	Templates *templates.Client

	Topologies *topologies.Client

	Usages *usages.Client

	Users *users.Client

	Util *util.Client

	Transport client.Transport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Cloudbreak) SetTransport(transport client.Transport) {
	c.Transport = transport

	c.Accountpreferences.SetTransport(transport)

	c.Blueprints.SetTransport(transport)

	c.Cluster.SetTransport(transport)

	c.Constraints.SetTransport(transport)

	c.Credentials.SetTransport(transport)

	c.Events.SetTransport(transport)

	c.Ldap.SetTransport(transport)

	c.Networks.SetTransport(transport)

	c.Rdsconfigs.SetTransport(transport)

	c.Recipes.SetTransport(transport)

	c.Securitygroups.SetTransport(transport)

	c.Sssd.SetTransport(transport)

	c.Stacks.SetTransport(transport)

	c.Templates.SetTransport(transport)

	c.Topologies.SetTransport(transport)

	c.Usages.SetTransport(transport)

	c.Users.SetTransport(transport)

	c.Util.SetTransport(transport)

}
