package listInventory

import (
	proto "StarResearch_Neelansh/listInventory/proto" // Import the generated protobuf package
	ReadCSV "StarResearch_Neelansh/readCSV"           // Import the ReadCSV package
	"context"
)

// InventoryListingServer implements the InventoryListing.InventoryListingService server.
type InventoryListingServer struct {
	sellerInventory map[string]ReadCSV.SellerProduct
	proto.UnimplementedInventoryListingServiceServer
}

// NewInventoryListingServer creates a new instance of InventoryListingServer.
func NewInventoryListingServer(sellerInventory map[string]ReadCSV.SellerProduct) *InventoryListingServer {
	return &InventoryListingServer{
		sellerInventory: sellerInventory,
	}
}

// ListInventory returns a list of current inventory items.
func (s *InventoryListingServer) ListInventory(ctx context.Context, req *proto.ListInventoryRequest) (*proto.ListInventoryResponse, error) {
	inventoryItems := make([]*proto.SellerProduct, 0)

	// Iterate over sellerInventory and convert the data into proto.InventoryItem messages.
	for _, item := range s.sellerInventory {
		inventoryItem := &proto.SellerProduct{
			StockCode: int32(item.StockCode),
			Name:      item.Name,
			Quantity:  int32(item.Quantity),
			Price:     float32(item.Price),
		}
		inventoryItems = append(inventoryItems, inventoryItem)
	}

	return &proto.ListInventoryResponse{
		InventoryItems: inventoryItems,
	}, nil
}
