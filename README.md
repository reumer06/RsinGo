# RsinGo

RsinGo is a cross platform reverse shell client that provides remote command execution over a network socket.

> **NOTE:** This project is intended solely for educational purposes. Run this tool only on systems you own or have explicit permission to test.

## Quick Start

### Run from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/reumer06/RsinGo.git
   cd RsinGo
   ```

2. Execute the client:
   ```bash
   go run src/main.go
   ```

> **Note:** If no active listener is available on the target port, the application will exit after the connection times out.

### Pre-built Binaries

You can download the pre-compiled binary directly from the [v1.0 Release](https://github.com/reumer06/RsinGo/releases/download/v1.0/rs.exe).

## Build

To compile the binary locally:

```powershell
go build -o bin/rsingo.exe ./src
```

## License

This project is licensed under the terms of the MIT License. See the [LICENSE](https://www.google.com/search?q=LICENSE) file for details.
