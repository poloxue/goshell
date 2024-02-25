# Shell in Go

This project is a simple, custom shell implemented in Go. It aims to provide a basic understanding of how shells operate, including parsing commands, handling input/output, and executing system commands. This shell supports basic command execution, along with built-in commands for navigating the filesystem.

## Features

- Command parsing and execution
- Built-in command: `cd` (Change Directory)
- Support for executing system commands
- Redirecting command output to stdout and stderr
- Interactive command-line interface with a fixed prompt ([Directory]$)

## Getting Started

### Prerequisites

Ensure you have Go installed on your system. This project was developed using Go version 1.15, but it should be compatible with most Go versions.

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/custom-shell-go.git
cd custom-shell-go
```

Build the shell from source:

```bash
go build -o myshell
```

Run the shell:

```bash
./myshell
```

## Usage

Once the shell is running, you will be greeted with a prompt where you can type commands. For example:

- To change directories, use the `cd` command:

```bash
$ cd /path/to/directory
```

- To execute a system command, simply type the command name followed by any arguments:

```bash
$ ls -l
```

- To exit the shell, type `exit`:

```bash
$ exit
```

## Built-in Commands

Currently, the shell supports the following built-in command:

- `cd`: Change the current working directory.

## Extending the Shell

To add new features or commands to the shell, modify the `Shell` struct and implement additional methods as needed. For example, to add a new built-in command, update the `builtinCmds` map and add a corresponding execution function.

