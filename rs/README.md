# Teal - Rust Port

This is a Rust port of the [Teal project](https://github.com/Layr-Labs/teal), originally written in Go.

## Structure

The port maintains the same overall structure as the Go project:

- `aggregator`: Service for aggregating BLS signatures 
- `api`: gRPC API definitions
- `node`: Node implementation with server and service components
- `common`: Shared types and utilities

## Building

To build the project:

```bash
cargo build
```

## Running

To run the node:

```bash
cargo run --bin node
```

Options:
- `-p, --port <PORT>`: Port to listen on (default: 9000)

## Implementation Notes

- This port is a conceptual translation from Go to Rust
- The BLS cryptography is implemented with stub functions for simplicity
- In a production implementation, a proper BLS library would be used
- The gRPC service is manually implemented rather than generated from proto files

## Future Work

- Implement proper BLS cryptography using blst or similar libraries
- Generate gRPC code from proto files using tonic-build
- Add tests and benchmarks 