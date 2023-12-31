# LockAndRoll

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

My own CLI secret manager.

Features:
 - Get quick access to secrets in the CLI
 - Use different backend systems to store secrets (e.g. MacOS Keychain)
 - Shortcut to export secrets as enviromental variables

To Do:
 - Support non-JSON secrets
 - Implement dump/import functions
 - Write tests

## Install

```shell
go install github.com/svalentino/lockandroll
```

## Usage

Examples:

```shell
# Create secrets
> lockandroll create test '{"secret1":"xyz.."}'
Secret 'test' created successfully.

# List secrets
❯ lockandroll list
test

# Retrieve secret value
❯ lockandroll get test
secret1="xyz.."

# Export secret values
❯ eval $(lockandroll get test -e)
❯ printenv | grep secret1
secret1=xyz..
```

Full Help:

```
> lockandroll -h

Simple tool to handle secrets in CLI

Usage:
  lockandroll [flags]
  lockandroll [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create and save new secret
  delete      Delete a secret
  get         Retrieve the value of a secret
  help        Help about any command
  list        List all stored secrets
  update      Update the value of a secret

Flags:
  -b, --backend string   backend to use (default macos-keychain) (default "macos-keychain")
  -c, --config string    config file (default is $HOME/.lockandroll.yaml)
  -h, --help             help for lockandroll
  -v, --verbose          Verbose logging

Use "lockandroll [command] --help" for more information about a command.
```

