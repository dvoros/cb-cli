// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCluster deletes cluster on a specific stack

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) DeleteCluster(params *DeleteClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCluster",
		Method:             "DELETE",
		PathPattern:        "/v1/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
FailureReportCluster failures report

Endpoint to report the failed nodes in the given cluster. If recovery mode for the node's hostgroup is AUTO then autorecovery would be started. If recovery mode for the node's hostgroup is MANUAL, the nodes will be marked as unhealthy.
*/
func (a *Client) FailureReportCluster(params *FailureReportClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFailureReportClusterParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "failureReportCluster",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{id}/cluster/failurereport",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &FailureReportClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetCluster retrieves cluster by stack id

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetCluster(params *GetClusterParams) (*GetClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCluster",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterOK), nil

}

/*
GetConfigsCluster gets cluster properties with blueprint outputs

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetConfigsCluster(params *GetConfigsClusterParams) (*GetConfigsClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigsClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConfigsCluster",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{id}/cluster/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetConfigsClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfigsClusterOK), nil

}

/*
GetFullCluster retrieves cluster by stack id

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetFullCluster(params *GetFullClusterParams) (*GetFullClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFullClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFullCluster",
		Method:             "GET",
		PathPattern:        "/v1/stacks/{id}/cluster/full",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetFullClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFullClusterOK), nil

}

/*
GetPrivateCluster retrieves cluster by stack name private

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetPrivateCluster(params *GetPrivateClusterParams) (*GetPrivateClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivateCluster",
		Method:             "GET",
		PathPattern:        "/v1/stacks/user/{name}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateClusterOK), nil

}

/*
GetPublicCluster retrieves cluster by stack name public

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) GetPublicCluster(params *GetPublicClusterParams) (*GetPublicClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicCluster",
		Method:             "GET",
		PathPattern:        "/v1/stacks/account/{name}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicClusterOK), nil

}

/*
PostCluster creates cluster for stack

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) PostCluster(params *PostClusterParams) (*PostClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postCluster",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostClusterOK), nil

}

/*
PutCluster updates cluster by stack id

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) PutCluster(params *PutClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putCluster",
		Method:             "PUT",
		PathPattern:        "/v1/stacks/{id}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutClusterV2 updates cluster by stack name

Clusters are materialised Hadoop services on a given infrastructure. They are built based on a Blueprint (running the components and services specified) and on a configured infrastructure Stack. Once a cluster is created and launched, it can be used the usual way as any Hadoop cluster. We suggest to start with the Cluster's Ambari UI for an overview of your cluster.
*/
func (a *Client) PutClusterV2(params *PutClusterV2Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterV2Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putClusterV2",
		Method:             "PUT",
		PathPattern:        "/v2/stacks/{name}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutClusterV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RepairCluster repairs the cluster

Removing the failed nodes and starting new nodes to substitute them.
*/
func (a *Client) RepairCluster(params *RepairClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairClusterParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairCluster",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{id}/cluster/manualrepair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
UpgradeCluster upgrades the ambari version

Ambari is used to provision the Hadoop clusters.
*/
func (a *Client) UpgradeCluster(params *UpgradeClusterParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeClusterParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upgradeCluster",
		Method:             "POST",
		PathPattern:        "/v1/stacks/{id}/cluster/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpgradeClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
