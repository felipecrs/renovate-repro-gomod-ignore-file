# Reproduction of Go Mod Protobuf Issue

This project demonstrates a scenario where a Go module depends on a locally generated Protobuf module.

## Structure

- `go.mod`: Root module.
- `main.go`: Imports the generated code and `github.com/pkg/errors`.
- `api/`: Directory for the generated module.
  - `go.mod`: Module definition for the generated code.
  - `api.proto`: Protobuf definition.

The project uses Go 1.24+ tool dependencies to manage `protoc-gen-go`.

## Reproduction Steps

1.  **Initial State**: The project is set up, but the Go code from Protobuf is not yet generated.

2.  **Fail Case**: Run `go get ./...`.

    ```bash
    go get ./...
    ```

    This should fail because the `example.com/repro/api` package does not exist yet (no Go files in `api/`).

3.  **Generate Code**: Run `go generate` to generate the Go code.

    ```bash
    go generate ./...
    ```
    
    This command will:
    1.  Download `protoc` 33.2 from GitHub (if not already present).
    2.  Unzip it to a local `.tools` directory.
    3.  Run `protoc` using the downloaded binary and the `protoc-gen-go` tool managed by `go.mod`.

4.  **Pass Case**: Run `go get ./...` again.

    ```bash
    go get ./...
    ```

    This should now succeed.
