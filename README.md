# Analyser
The analyser is a cli to help provide information to the user about different HTTP headers in requests.

This repo is also a step-by-step example of how to build a gRPC server in Go and a cli with a client,
which calls the server endpoints and outputs information from the responses in the server.


## Pre-requisites
Have Go and `protoc` already installed.

## Building and running the server

```shell
make server-run
```

## Building and running the cli

To build the cli, you need to run:
```shell
make cli-build
```

You need to have the server running before you can use the cli, otherwise you will get an error.
After this you can run `analyser` and view the options available to you via the `help` command. 



