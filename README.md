## Prerequisites

- `Protocol Buffer Compiler`

## Run

- Rename `data.db.example` to `data.db`
- Run the command `go mod tidy` to install dependencies
- Run the command `go run cmd/grpcServer` to start `gRPC` server at port `50051`
- Use [evans](https://github.com/ktr0731/evans) to test the available services
