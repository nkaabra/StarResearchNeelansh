package listDemand

import (
	DemandListingproto "StarResearch_Neelansh/listDemand/proto"
	"context"
)

// DemandListingServer implements the DemandListing.DemandListingService server.
type DemandListingServer struct {
	customerDemands                                            map[string]int32
	DemandListingproto.UnimplementedDemandListingServiceServer // Embed the Unimplemented server to get default implementations.

}

// NewDemandListingServer creates a new instance of DemandListingServer.
func NewDemandListingServer(customerDemands map[string]int32) *DemandListingServer {
	return &DemandListingServer{
		customerDemands: customerDemands,
	}
}

// ListDemands returns a list of current customer demands.
func (s *DemandListingServer) ListDemands(ctx context.Context, req *DemandListingproto.ListDemandsRequest) (*DemandListingproto.ListDemandsResponse, error) {
	demands := make([]*DemandListingproto.DemandInfo, 0)

	// Iterate over customerDemands and convert the data into DemandInfo messages.
	for productName, quantity := range s.customerDemands {
		demands = append(demands, &DemandListingproto.DemandInfo{
			ProductName: productName,
			Quantity:    quantity,
		})
	}

	return &DemandListingproto.ListDemandsResponse{
		Demands: demands,
	}, nil
}
