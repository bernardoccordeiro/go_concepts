# Golang Concepts

This repository is supposed to be used as a personal space to learn the Go language. Concepts will be separated in their own modules, to help with organization.

## Some useful commands

The first thing that needs to be grasped, is that when creating a new module in Go, it needs to be initialized. It is done via the `go mod init` command. For instance:

`go mod init example.com/my-module`

This will create a go.mod file in the module directory, and allow it to be run via the `go run .` command.

The `go mod tidy` command is also important, as it will clean up the go.mod file, remove any unused dependencies, as well as add dependencies that are not yet tracked by the module.

Another command worth mentioning is the `go mod edit` command, which allows you to edit the go.mod file in a variety of ways. For instance, as per the Go Tour tutorials, if you want to change how the module tries to locate a specific dependency, you can use `go mod edit -replace` command. For instance, something like:

`go mod edit -replace example.com/greetings=../greetings`

This would make Go search for the example.com/greetings dependency in the parent directory, instead of as a published module.

More information might be added here with time.

To know more about a function, it is possible to run the `go doc` command, such as:

`go doc fmt Prinln`