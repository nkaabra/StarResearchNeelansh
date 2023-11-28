# Multi-stage build
FROM golang:latest as builder

WORKDIR /go/src/app

# Explicitly copy all files into the container
COPY . .

EXPOSE 50051

RUN GOARCH=amd64 go build -o main .

# RUN GOARCH=amd64 CGO_ENABLED=0 go build -o main .

# # Final stage
# FROM scratch

# COPY --from=builder /go/src/app/main /app/main
CMD ["./main"]



# # Base image for amd64
# FROM golang:latest AS amd64

# # Set the working directory inside the container
# WORKDIR /go/src/app

# # Copy all the files into the container
# COPY . .

# EXPOSE 50051

# # Build the application for amd64
# RUN GOARCH=amd64 go build -o main .

# # Base image for arm64
# FROM golang:latest AS arm64

# # Set the working directory inside the container
# WORKDIR /go/src/app

# # Copy all the files into the container
# COPY . .

# EXPOSE 50051

# # Build the application for arm64
# RUN GOARCH=arm64 go build -o main .

# # Create a manifest list
# FROM golang:latest AS manifest
# COPY --from=amd64 /go/src/app/main /go/src/app/amd64
# COPY --from=arm64 /go/src/app/main /go/src/app/arm64



# # # Base image
# # FROM golang:latest

# # # Set the working directory inside the container
# # WORKDIR /go/src/app

# # # Copy all the files into the container
# # COPY . .

# # EXPOSE 50051

# # # Build the application
# # RUN go build -o main .

# # # Command to run the application
# # CMD ["./main"]


