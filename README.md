# echo-server

A simple Go-based HTTP server that logs request details and environment variables, designed for Kubernetes.

## Features
- Logs request headers
- Retrieves `HOSTNAME` and `NODE_NAME` from environment variables
- Built with Go for performance and efficiency

## Build and Run Locally

```sh
go build -o server
./server
```

## Build and Run in Docker

```sh
docker build -t echo-server .
docker run -p 8080:8080 echo-server
```
