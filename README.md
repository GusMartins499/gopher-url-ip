# gopher-url-ip

[Português (Brasil)](README.pt-br.md)

A simple command-line tool, written in Go, that looks up the IP addresses and name servers of a host.

## Requirements

- Go 1.26.5 or later

## Installation

```sh
git clone https://github.com/GusMartins499/gopher-url-ip.git
cd gopher-url-ip
go build -o gopher-url-ip .
```

Or run it directly without building:

```sh
go run . <command> --host <host>
```

## Usage

```sh
./gopher-url-ip <command> --host <host>
```

### Commands

| Command       | Alias | Description                                   |
| ------------- | ----- | --------------------------------------------- |
| `get-ip`      | `g`   | Prints the IP addresses (IPv4 and IPv6) of the host. |
| `get-servers` |       | Prints the name servers (NS records) of the host.    |
| `help`        | `h`   | Shows the list of commands or help for one command.  |

### Flags

| Flag     | Description                                   |
| -------- | --------------------------------------------- |
| `--host` | The host to look up, without the protocol (e.g. `google.com`). |

### Examples

```sh
$ ./gopher-url-ip get-ip --host google.com
142.250.79.46
2800:3f0:4001:80a::200e

$ ./gopher-url-ip get-servers --host google.com
ns1.google.com.
ns2.google.com.
ns3.google.com.
ns4.google.com.
```

The addresses returned depend on your location and DNS resolver.

## Built with

- [urfave/cli](https://github.com/urfave/cli) for building the command-line interface
- Go's standard [`net`](https://pkg.go.dev/net) package for DNS lookups
