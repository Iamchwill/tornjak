package api

import (
	"errors"

	tornjakTypes "github.com/spiffe/tornjak/pkg/agent/types"
)

/*

Agent

ListAgents(ListAgentsRequest) returns (ListAgentsResponse);
BanAgent(BanAgentRequest) returns (google.protobuf.Empty);
DeleteAgent(DeleteAgentRequest) returns (google.protobuf.Empty);
CreateJoinToken(CreateJoinTokenRequest) returns (spire.types.JoinToken);

Entries

ListEntries(ListEntriesRequest) returns (ListEntriesResponse);
BatchCreateEntry(BatchCreateEntryRequest) returns (BatchCreateEntryResponse);
GetEntry(GetEntryRequest) returns (spire.types.Entry);

*/

type ListSelectorsRequest struct{}
type ListSelectorsResponse tornjakTypes.AgentInfoList

// ListSelectors returns list of agents from the local DB with the following info
// spiffeid string
// plugin   string
func (s *Server) ListSelectors(inp ListSelectorsRequest) (*ListSelectorsResponse, error) {
	resp, err := s.Db.GetAgentSelectors()
	if err != nil {
		return nil, err
	}
	return (*ListSelectorsResponse)(&resp), nil
}

type RegisterSelectorRequest tornjakTypes.AgentInfo

// DefineSelectors registers an agent to the local DB with the following info
// spiffeid string
// plugin   string
func (s *Server) DefineSelectors(inp RegisterSelectorRequest) error {
	sinfo := tornjakTypes.AgentInfo(inp)
	if len(sinfo.Spiffeid) == 0 {
		return errors.New("agent's info missing mandatory field - Spiffeid")
	}
	return s.Db.CreateAgentEntry(sinfo)
}

type ListAgentMetadataRequest tornjakTypes.AgentMetadataRequest
type ListAgentMetadataResponse tornjakTypes.AgentInfoList

// ListAgentMetadata takes in list of agent spiffeids
// and returns list of those agents from the local DB with following info
// spiffeid string
// plugin string
// cluster string
// if no metadata found, no row is included
// if no spiffeids are specified, all agent metadata is returned
func (s *Server) ListAgentMetadata(inp ListAgentMetadataRequest) (*ListAgentMetadataResponse, error) {
	inpReq := tornjakTypes.AgentMetadataRequest(inp)
	resp, err := s.Db.GetAgentsMetadata(inpReq)
	if err != nil {
		return nil, err
	}
	return (*ListAgentMetadataResponse)(&resp), nil
}

type ListClustersRequest struct{}
type ListClustersResponse tornjakTypes.ClusterInfoList

// ListClusters returns list of clusters from the local DB with the following info
// name string
// details json
func (s *Server) ListClusters(inp ListClustersRequest) (*ListClustersResponse, error) {
	retVal, err := s.Db.GetClusters()
	if err != nil {
		return nil, err
	}
	return (*ListClustersResponse)(&retVal), nil
}

type RegisterClusterRequest tornjakTypes.ClusterInput

// DefineCluster registers cluster to local DB
func (s *Server) DefineCluster(inp RegisterClusterRequest) error {
	cinfo := tornjakTypes.ClusterInfo(inp.ClusterInstance)
	if len(cinfo.Name) == 0 {
		return errors.New("cluster definition missing mandatory field - Name")
	} else if len(cinfo.PlatformType) == 0 {
		return errors.New("cluster definition missing mandatory field - PlatformType")
	} else if len(cinfo.EditedName) > 0 {
		return errors.New("cluster definition attempts renaming on create cluster - EditedName")
	}
	return s.Db.CreateClusterEntry(cinfo)
}

type EditClusterRequest tornjakTypes.ClusterInput

// EditCluster registers cluster to local DB
func (s *Server) EditCluster(inp EditClusterRequest) error {
	cinfo := tornjakTypes.ClusterInfo(inp.ClusterInstance)
	if len(cinfo.Name) == 0 {
		return errors.New("cluster definition missing mandatory field - Name")
	} else if len(cinfo.PlatformType) == 0 {
		return errors.New("cluster definition missing mandatory field - PlatformType")
	} else if len(cinfo.EditedName) == 0 {
		return errors.New("cluster definition missing mandatory field - EditedName")
	}
	return s.Db.EditClusterEntry(cinfo)
}

type DeleteClusterRequest tornjakTypes.ClusterInput

// DeleteCluster deletes cluster with name cinfo.Name and assignment to agents
func (s *Server) DeleteCluster(inp DeleteClusterRequest) error {
	cinfo := tornjakTypes.ClusterInfo(inp.ClusterInstance)
	if len(cinfo.Name) == 0 {
		return errors.New("input missing mandatory field - Name")
	}
	return s.Db.DeleteClusterEntry(cinfo.Name)
}

type EditAgentRequest tornjakTypes.AgentInfo

// EditAgent updates the server information in the local DB.
func (s *Server) EditAgent(inp EditAgentRequest) error {
	sinfo := tornjakTypes.AgentInfo(inp.AgentInstance)
	// Validate input fields
	if len(inp.Spiffeid) == 0 {
		return errors.New("agent missing mandatory field - Spiffeid")
	}
	if len(inp.Plugin) == 0 {
		return errors.New("agent missing mandatory field - Plugin")
	}
	if len(inp.Cluster) == 0 {
		return errors.New("agent missing mandatory field - Cluster")
	}
	if len(inp.EditedCluster) == 0 {
		return errors.New("agent missing mandatory field - EditedCluster")
	}
	err := s.Db.EditAgentEntry(sinfo)
	if err != nil {
		return err
	}

	// Return nil if everything was successful
	return nil
}

// // EntryInput defines the request structure for editing server details.
// type EntryInput struct {
// 	ServerID  string `json:"server_id"`
// 	Name      string `json:"name"`
// 	IPAddress string `json:"ip_address"`
// 	Cluster   string `json:"cluster"`
// 	Platform  string `json:"platform"`
// }

// // EditServer updates the server information in the local DB.
// func (s *Server) EditServer(inp EntryInput) error {
// 	// Validate input fields
// 	if len(inp.ServerID) == 0 {
// 		return errors.New("server missing mandatory field - Server ID")
// 	}
// 	if len(inp.Name) == 0 {
// 		return errors.New("server missing mandatory field - Name")
// 	}
// 	if len(inp.IPAddress) == 0 {
// 		return errors.New("server missing mandatory field - IP address")
// 	}
// 	if len(inp.Cluster) == 0 {
// 		return errors.New("server missing mandatory field - Cluster")
// 	}

// 	// Create the server info struct
// 	serverInfo := tornjakTypes.ServerInfo{
// 		ServerID:  inp.ServerID,
// 		Name:      inp.Name,
// 		IPAddress: inp.IPAddress,
// 		Cluster:   inp.Cluster,
// 		Platform:  inp.Platform,
// 	}

// 	err := s.Db.UpdateServerEntry(serverInfo)
// 	if err != nil {
// 		return err
// 	}

// 	// Return nil if everything was successful
// 	return nil
// }
