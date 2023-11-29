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

## Using Cloudlab

After instantiating an experiment, and ssh'ing into a cloudLab node, please follow the following steps:

(1) sudo apt update
(2) sudo apt install apt-transport-https ca-certificates curl software-properties-common
(3) curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
(4) echo "deb [signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
(5) sudo apt update
(6) sudo apt install docker-ce docker-ce-cli containerd.io
(7) sudo usermod -aG docker $USER
(8) sudo docker login
(9) Enter username and password
(10) sudo docker pull nkaabra/myecommerceapp:latest
(11) sudo docker run nkaabra/myecommerceapp:latest & 
(12) ip addr show (run this on the cloudlab machine)
Running the command above will provide you with details regarding the IP addresses being used. 
(13) Find the IP Address from the terminal output 

Sample output: 
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: eno1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether ec:b1:d7:85:4a:62 brd ff:ff:ff:ff:ff:ff
    altname enp9s0
   # inet 128.110.217.103/21 metric 1024 brd 128.110.223.255 scope global eno1
       valid_lft forever preferred_lft forever
    inet6 fe80::eeb1:d7ff:fe85:4a62/64 scope link 
       valid_lft forever preferred_lft forever
3: eno1d1: <BROADCAST,MULTICAST> mtu 1500 qdisc noop state DOWN group default qlen 1000
    link/ether ec:b1:d7:85:4a:63 brd ff:ff:ff:ff:ff:ff
    altname enp9s0d1
4: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:74:fa:5c:35 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet6 fe80::42:74ff:fefa:5c35/64 scope link 
       valid_lft forever preferred_lft forever
12: vethaf2f5c0@if11: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
    link/ether 6e:9a:21:09:26:04 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet6 fe80::6c9a:21ff:fe09:2604/64 scope link 
       valid_lft forever preferred_lft forever


The bolded line contains the IP: 128.110.217.103
You can now use the public IP Address to interact with the application
For eg: 172.17.0.1 

(13) grpcurl -plaintext -d '{"product_name": "A", "quantity": 5}' 172.17.0.1:50051 addDemand.EcommerceService/MakeDemand











