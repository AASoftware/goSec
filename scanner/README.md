# goSec

This is a simple network scanner written in Go.

## Description

This tool allows you to scan for open ports on a specified host using TCP connections. It supports various scanning options, including standard scan, scanning all ports, and specifying individual ports.

**Note:** This tool is intended for educational purposes and should be used responsibly. Do not use it for any illegal or malicious activities.

## Installation

To use this tool, you need to have Go installed on your system.

1. Clone the repository:

git clone https://github.com/AASoftware/goSec.git


2. Navigate to the project directory:

cd goSec

3. Compile the code:

go build main.go -o scanner.exe


## Usage

Once you have compiled the code, you can use the `main` executable to run the network scanner.

### Command-line Options

- `-h`: Specify the hostname or IP address of the target host.
- `-s`: Perform a standard scan, scanning the first 1500 ports.
- `-a`: Scan all ports (1-65535).
- `-p`: Specify individual ports to scan, separated by commas (e.g., `22,80,135`).
- `-t`: Number of concurrent threads for scanning (maximum 50, default 20).
- `--help`: Show usage information.

### Examples

1. Perform a standard scan on a specified host:

./scanner.exe -h Hostname -t50 -a for all ports

