// package Demand

// import (
// 	demandService "StarResearch_Neelansh/addDemand/proto"
// 	ReadCSV "StarResearch_Neelansh/readCSV"
// 	"context"
// 	"fmt"
// )

// type EcommerceServiceServer struct {
// 	// Add any data dependencies here, such as customerDemands and sellerInventory.
// 	customerDemands map[string]int32
// 	sellerInventory map[string]ReadCSV.SellerProduct
// 	demandService.UnimplementedEcommerceServiceServer
// }

// func NewEcommerceServiceServer(customerDemands map[string]int32, sellerInventory map[string]ReadCSV.SellerProduct) *EcommerceServiceServer {
// 	return &EcommerceServiceServer{
// 		customerDemands: customerDemands,
// 		sellerInventory: sellerInventory,
// 	}
// }

// func (s *EcommerceServiceServer) MakeDemand(ctx context.Context, req *demandService.DemandRequest) (*demandService.DemandResponse, error) {

// 	productName := req.ProductName
// 	quantity := req.Quantity

// 	// Check if the product is available in sellerInventory.
// 	availableQuantity, existsInSeller := s.sellerInventory[productName]
// 	if existsInSeller && int32(availableQuantity.Quantity) >= quantity {
// 		//remove the quantity from sellerInventory
// 		currQuantity := s.sellerInventory[productName].Quantity
// 		currQuantity -= int(quantity)
// 		s.sellerInventory[productName] = ReadCSV.SellerProduct{
// 			StockCode: s.sellerInventory[productName].StockCode,
// 			Name:      productName,
// 			Quantity:  currQuantity,
// 			Price:     s.sellerInventory[productName].Price,
// 		}
// 		return &demandService.DemandResponse{Success: true, Message: "Product found in invetory!!"}, nil
// 	}

// 	// Check if there is an existing demand for the same product.
// 	_, exists := s.customerDemands[productName]

// 	if exists {
// 		// Increment the existing demand.
// 		s.customerDemands[productName] += quantity
// 		fmt.Printf("Customer placed a demand for %d units of %s.\n", quantity, productName)
// 		return &demandService.DemandResponse{Success: true, Message: "Demand existed. Demand successfully updated."}, nil
// 	} else {
// 		// Create a new entry in customerDemands.
// 		s.customerDemands[productName] = quantity
// 		fmt.Printf("Customer placed a demand for %d units of %s.\n", quantity, productName)
// 		return &demandService.DemandResponse{Success: true, Message: "New Demand successfully placed."}, nil
// 	}
// 	// Print the customerDemands map.

// }

package Demand

import (
	"context"
	"fmt"
	"sync"

	demandService "StarResearch_Neelansh/addDemand/proto"
	ReadCSV "StarResearch_Neelansh/readCSV"
)

type EcommerceServiceServer struct {
	// Add any data dependencies here, such as customerDemands and sellerInventory.
	mu              sync.Mutex
	customerDemands map[string]int32
	sellerInventory map[string]ReadCSV.SellerProduct
	demandService.UnimplementedEcommerceServiceServer
}

func NewEcommerceServiceServer(customerDemands map[string]int32, sellerInventory map[string]ReadCSV.SellerProduct) *EcommerceServiceServer {
	return &EcommerceServiceServer{
		customerDemands: customerDemands,
		sellerInventory: sellerInventory,
	}
}

func (s *EcommerceServiceServer) MakeDemand(ctx context.Context, req *demandService.DemandRequest) (*demandService.DemandResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	productName := req.ProductName
	quantity := req.Quantity

	// Check if the product is available in sellerInventory.
	availableQuantity, existsInSeller := s.sellerInventory[productName]
	if existsInSeller && int32(availableQuantity.Quantity) >= quantity {
		// remove the quantity from sellerInventory
		currQuantity := s.sellerInventory[productName].Quantity
		currQuantity -= int(quantity)
		s.sellerInventory[productName] = ReadCSV.SellerProduct{
			StockCode: s.sellerInventory[productName].StockCode,
			Name:      productName,
			Quantity:  currQuantity,
			Price:     s.sellerInventory[productName].Price,
		}
		return &demandService.DemandResponse{Success: true, Message: "Product found in inventory!!"}, nil
	}

	// Check if there is an existing demand for the same product.
	_, exists := s.customerDemands[productName]

	if exists {
		// Increment the existing demand.
		s.customerDemands[productName] += quantity
		fmt.Printf("Customer placed a demand for %d units of %s.\n", quantity, productName)
		return &demandService.DemandResponse{Success: true, Message: "Demand existed. Demand successfully updated."}, nil
	} else {
		// Create a new entry in customerDemands.
		s.customerDemands[productName] = quantity
		fmt.Printf("Customer placed a demand for %d units of %s.\n", quantity, productName)
		return &demandService.DemandResponse{Success: true, Message: "New Demand successfully placed."}, nil
	}
}
