# StarResearchNeelansh

# E-commerce Microservices Application in Go

Welcome to the E-commerce Microservices Application, a project built using Golang, protocol buffers, and gRPC. This application showcases a microservices architecture, demonstrating the power and flexibility of the Go programming language for building scalable and efficient e-commerce systems.

## Key Features:

- **Microservices Architecture:** The application is structured into microservices, each responsible for specific functionalities such as handling demands, managing inventory, and providing listings.

- **gRPC Communication:** Leveraging gRPC for communication between microservices ensures efficient and language-agnostic communication, enhancing the modularity of the system.

- **Protocol Buffers:** The use of protocol buffers provides a language-agnostic data serialization format, optimizing data exchange between microservices.

The structure of the directory consists of smallers directories for each microservice, along with a trimmed.CSV file containing a real world dataset obtained from Kaggle, trimmed for our specific needs.

Each smaller directory for the microservices contains a folder with the proto files (protocol buffers) such as the .proto file defining the semantics, along with the generated grpc-proto files.

## Starting the Application
To start the application, please follow the following steps: 

(1) Download the docker image via 
'''bash
docker pull nkaabra/myecommerceapp
(2) Use the docker run command to start a container based on the downloaded image. 

The dockerFile has been made such that doing so will automatically run the application for you.

## Interacting with the application
To interact with the application, one can use grpCurl commands from the terminal to simulate a client/seller. 

Sample commands: (Please take care to replace localHost with the IP Address of the machine you are running the docker image on)
grpcurl -plaintext -d '{"product_name": "A", "quantity": 5}' localhost:50051 addDemand.EcommerceService/MakeDemand

grpcurl -plaintext localhost:50051 listDemand.DemandListingService/ListDemands

grpcurl -plaintext -d '{
  "seller_product": {
    "stockCode": "8243A",
    "name": "Product Name",
    "quantity": 10,
     "price": 60.0
  }
}' localhost:50051 Inventory.InventoryService/AddInventory

grpcurl -plaintext -d '{}' localhost:50051 listInventory.InventoryListingService/ListInventory

Ghz services are also provided for load testing, however load balancing application has not been implemented yet.

To use Ghz : 
ghz -n 1000 -c 10 --insecure --proto=demand.proto --call=addDemand.EcommerceService.MakeDemand --data '{"product_name": "A", "quantity": 5}' localhost:50051

ghz -n 1000 -c 10 --insecure --proto=inventory.proto --call=Inventory.InventoryService.AddInventory --data '{"seller_product": {"stockCode": "ABC123", "name": "ProductA", "quantity": 10, "price": 50.0}}' localhost:50051







