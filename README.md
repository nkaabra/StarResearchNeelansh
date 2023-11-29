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
To

