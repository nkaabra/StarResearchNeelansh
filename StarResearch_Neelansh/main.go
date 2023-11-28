package main

import (
	Demand "StarResearch_Neelansh/addDemand"
	Demandproto "StarResearch_Neelansh/addDemand/proto"
	Inventory "StarResearch_Neelansh/addInventory"            // Import the addInventory package
	Inventoryproto "StarResearch_Neelansh/addInventory/proto" // Import the addInventory proto package
	DemandListing "StarResearch_Neelansh/listDemand"
	DemandListingproto "StarResearch_Neelansh/listDemand/proto"
	listInventory "StarResearch_Neelansh/listInventory"
	listInventoryproto "StarResearch_Neelansh/listInventory/proto"
	ReadCSV "StarResearch_Neelansh/readCSV"

	//Init "StarResearch_Neelansh/Init"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Starting application...")

	customerDemands, sellerInventory, err := ReadCSV.ReadCSVData("trimmed.csv")
	if err != nil {
		log.Fatalf("Failed to read CSV data: %v", err)
	}

	server := grpc.NewServer()
	customerDemandServer := Demand.NewEcommerceServiceServer(customerDemands, sellerInventory)
	Demandproto.RegisterEcommerceServiceServer(server, customerDemandServer)

	// Register the Inventory service
	inventoryServer := Inventory.NewInventoryServiceServer(sellerInventory, customerDemands)
	Inventoryproto.RegisterInventoryServiceServer(server, inventoryServer)

	// Register the Inventory service from listInventory.go
	listInventoryServer := listInventory.NewInventoryListingServer(sellerInventory)
	listInventoryproto.RegisterInventoryListingServiceServer(server, listInventoryServer)

	listDemandServer := DemandListing.NewDemandListingServer(customerDemands) // Create a listDemand server instance
	DemandListingproto.RegisterDemandListingServiceServer(server, listDemandServer)
	reflection.Register(server)

	// Define the network address to listen on (e.g., ":50051").
	listenAddr := ":50051"

	// Create a listener to listen on the network address.
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start the gRPC server.
	fmt.Printf("Server listening on %s\n", listenAddr)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
