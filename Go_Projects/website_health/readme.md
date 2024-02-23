# Health Check

## Overview
This is a simple command-line tool developed in Go that performs health checks on websites to determine if they are up or down. It utilizes TCP connections to verify the availability of the specified domain and port.

## Features
- Check whether a website is up or down.
- Ability to specify a custom port for the health check.

## Installation
1. Clone or download the repository.
2. Ensure you have Go installed on your system.
3. Navigate to the project directory.
4. Run `go build` to build the executable.

## Usage
### Command Line Arguments
- `--domain (-d)`: Specifies the domain to check the health of. This argument is required.
- `--port (-p)`: Specifies the port to check the health of. If not provided, the default port `80` will be used.

### Example
```bash
./healthcheck --domain example.com --port 80
