# Alert Checker

## Overview
Alert Checker is a Go application designed to monitor and respond to various alert conditions. It utilizes several external libraries and tools to handle tasks such as environment variable management, Redis operations, and Telegram bot interactions.

## Prerequisites
- Go 1.21.6 or higher
- Protocol Buffers (protoc)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/TropicalDog17/alert-checker.git
cd alert-checker
```

2. Install the required Go modules:
```bash
go mod tidy
```

## Configuration
Create a `.env` file in the root directory of the project and add the necessary environment variables as specified in the `.gitignore` file (lines 1).

## Usage

### Compiling Protobuf Files
To compile the Protocol Buffer files, run the following command:

```bash
make proto
```

### Running the Application
To start the application, use:

```bash
make run
```
