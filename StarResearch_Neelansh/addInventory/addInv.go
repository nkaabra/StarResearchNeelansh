// addInventory.go

// package inventory

// import (
// 	invent "StarResearch_Neelansh/addInventory/proto"
// 	ReadCSV "StarResearch_Neelansh/readCSV"
// 	"context"
// 	"fmt"
// )

// // InventoryServiceServer implements the Inventory.InventoryService server.
// type InventoryServiceServer struct {
// 	sellerInventory map[int]ReadCSV.SellerProduct
// 	invent.UnimplementedInventoryServiceServer
// }

// // NewInventoryServiceServer creates a new instance of InventoryServiceServer.
// func NewInventoryServiceServer(sellerInventory map[int]ReadCSV.SellerProduct) *InventoryServiceServer {
// 	return &InventoryServiceServer{
// 		sellerInventory: sellerInventory,
// 	}
// }

// // AddInventory adds a product and its price to the seller's inventory.
// func (s *InventoryServiceServer) AddInventory(ctx context.Context, req *invent.AddInventoryRequest) (*invent.AddInventoryResponse, error) {
// 	productName := req.ProductName
// 	price := req.Price

// 	// Check if the product already exists in sellerInventory.
// 	_, exists := s.sellerInventory[productName]
// 	if exists {
// 		return &invent.AddInventoryResponse{Success: false, Message: "Product already exists in inventory."}, nil
// 	}

// 	// Add the product to sellerInventory.
// 	s.sellerInventory[productName] = price

// 	fmt.Printf("Added product '%s' to inventory with price %.2f\n", productName, price)
// 	fmt.Printf("Seller Inventory: %v\n", s.sellerInventory)
// 	return &invent.AddInventoryResponse{Success: true, Message: "Inventory updated successfully."}, nil
// }

// addInv.go

package inventory

import (
	invent "StarResearch_Neelansh/addInventory/proto"
	ReadCSV "StarResearch_Neelansh/readCSV"
	context "context"
	fmt "fmt"
	"sync"
)

// InventoryServiceServer implements the Inventory.InventoryService server.
type InventoryServiceServer struct {
	sellerInventory map[string]ReadCSV.SellerProduct
	customerDemands map[string]int32
	invent.UnimplementedInventoryServiceServer
	mu sync.Mutex
}

// NewInventoryServiceServer creates a new instance of InventoryServiceServer.
func NewInventoryServiceServer(sellerInventory map[string]ReadCSV.SellerProduct, customerDemands map[string]int32) *InventoryServiceServer {
	return &InventoryServiceServer{
		sellerInventory: sellerInventory,
		customerDemands: customerDemands,
	}
}

// AddInventory adds a product and its price to the seller's inventory.
func (s *InventoryServiceServer) AddInventory(ctx context.Context, req *invent.AddInventoryRequest) (*invent.AddInventoryResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sellerProduct := req.GetSellerProduct()
	//check if demand already exists for the product in the customer's demand

	quantity, exists2 := s.customerDemands[sellerProduct.Name]
	is_in := false
	if exists2 {
		//check if the quantity of the product is greater than the demand
		if quantity <= sellerProduct.Quantity {
			//return a false message, say that demand has been matched, and remove from customer demands
			delete(s.customerDemands, sellerProduct.Name)
			sellerProduct.Quantity -= int32(quantity)
			is_in = true
			//return &invent.AddInventoryResponse{Success: false, Message: "Demand has been matched."}, nil
		} else {
			//if the quantity of the product is greater than the demand, then reduce the quantity of the product in the customer demand
			s.customerDemands[sellerProduct.Name] -= int32(sellerProduct.Quantity)

			return &invent.AddInventoryResponse{Success: false, Message: "Demand has been partially matched."}, nil
		}
	}
	// Check if the product already exists in sellerInventory.
	_, exists := s.sellerInventory[sellerProduct.Name]
	if exists {
		//if the product exists, increase the quantity of the product
		s.sellerInventory[(sellerProduct.Name)] = ReadCSV.SellerProduct{
			Name:     sellerProduct.Name,
			Quantity: int(sellerProduct.Quantity) + s.sellerInventory[(sellerProduct.StockCode)].Quantity,
			Price:    sellerProduct.Price,
		}
		if is_in {
			return &invent.AddInventoryResponse{Success: true, Message: "Inventory updated successfully. Demand has been matched."}, nil
		} else {
			return &invent.AddInventoryResponse{Success: false, Message: "Product already exists in inventory."}, nil
		}
	}

	// Add the product to sellerInventory.
	s.sellerInventory[(sellerProduct.Name)] = ReadCSV.SellerProduct{
		Name:     sellerProduct.Name,
		Quantity: int(sellerProduct.Quantity),
		Price:    sellerProduct.Price,
	}

	fmt.Printf("Added product '%s' to inventory with quantity %d and price %.2f\n", sellerProduct.Name, sellerProduct.Quantity, sellerProduct.Price)
	//fmt.Printf("Seller Inventory: %v\n", s.sellerInventory)
	if is_in {
		return &invent.AddInventoryResponse{Success: true, Message: "Inventory updated successfully. Demand has been matched."}, nil
	} else {
		return &invent.AddInventoryResponse{Success: true, Message: "Inventory updated successfully."}, nil
	}
}
