# gopher-url-ip

[English](README.md)

Uma ferramenta de linha de comando simples, escrita em Go, que consulta os endereços IP e os servidores de nomes de um host.

## Requisitos

- Go 1.26.5 ou superior

## Instalação

```sh
git clone https://github.com/GusMartins499/gopher-url-ip.git
cd gopher-url-ip
go build -o gopher-url-ip .
```

Ou execute direto, sem compilar:

```sh
go run . <comando> --host <host>
```

## Uso

```sh
./gopher-url-ip <comando> --host <host>
```

### Comandos

| Comando       | Atalho | Descrição                                           |
| ------------- | ------ | --------------------------------------------------- |
| `get-ip`      | `g`    | Exibe os endereços IP (IPv4 e IPv6) do host.        |
| `get-servers` |        | Exibe os servidores de nomes (registros NS) do host. |
| `help`        | `h`    | Mostra a lista de comandos ou a ajuda de um comando. |

### Flags

| Flag     | Descrição                                                     |
| -------- | ------------------------------------------------------------- |
| `--host` | O host a ser consultado, sem o protocolo (ex.: `google.com`). |

### Exemplos

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

Os endereços retornados variam conforme a sua localização e o seu resolvedor de DNS.

## Feito com

- [urfave/cli](https://github.com/urfave/cli) para construir a interface de linha de comando
- O pacote [`net`](https://pkg.go.dev/net) da biblioteca padrão do Go para as consultas de DNS
