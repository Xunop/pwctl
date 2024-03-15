# PWCTL Demo

This is a simple demo of the [PWCTL](https://wiki.postgresql.org/wiki/GSoC_2024#pwctl_CLI).

Notes:
- This is a simple demo, it's not a complete implementation.
- The `pwctl` is a CLI tool to interact with the [pwatch3](https://github.com/cybertec-postgresql/pgwatch3#quick-start).
- The commands are not implemented, it's just a demo of how to use the `cobra` library to create a CLI tool.

The `pwatch3` **api documentation** in `apidoc/` directory.

## Getting Started

### Prerequisites

You Need to run the [pwatch3](https://github.com/cybertec-postgresql/pgwatch3#quick-start)

### Clone and Run

```bash
git clone https://github.com/Xunop/pwctl.git
cd pwctl
go run main.go
```

## Usage

```bash
pwctl -h
```
