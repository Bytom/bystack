Bystack
======

[![Build Status](https://travis-ci.org/Bystack/bystack.svg)](https://travis-ci.org/Bystack/bystack) [![AGPL v3](https://img.shields.io/badge/license-AGPL%20v3-brightgreen.svg)](./LICENSE)

**Official golang implementation of the Bystack protocol.**

Automated builds are available for stable releases and the unstable master branch. Binary archives are published at https://github.com/Bytom/bystack/releases.

## What is Bystack?

Bystack is software designed to operate and connect to highly scalable blockchain networks confirming to the Bystack Blockchain Protocol, which allows partipicants to define, issue and transfer digitial assets on a multi-asset shared ledger. Please refer to the [White Paper](https://github.com/Bystack/wiki/blob/master/White-Paper/%E6%AF%94%E5%8E%9F%E9%93%BE%E6%8A%80%E6%9C%AF%E7%99%BD%E7%9A%AE%E4%B9%A6-%E8%8B%B1%E6%96%87%E7%89%88.md) for more details.

In the current state `bystack` is able to:

- Manage key, account as well as asset
- Send transactions, i.e., issue, spend and retire asset

## Building from source

### Requirements

- [Go](https://golang.org/doc/install) version 1.8 or higher, with `$GOPATH` set to your preferred directory

### Installation

Ensure Go with the supported version is installed properly:

```bash
$ go version
$ go env GOROOT GOPATH
```

- Get the source code

``` bash
$ git clone https://github.com/Bytom/bystack.git $GOPATH/src/github.com/bytom/bystack
```

- Build source code

``` bash
$ cd $GOPATH/src/github.com/bytom/bystack
$ make bystackd    # build bystackd
```

When successfully building the project, the `bystackd` and binary should be present in `cmd/bystackd` directory, respectively.

### Executables

The Bystack project comes with several executables found in the `cmd` directory.

| Command      | Description                                                  |
| ------------ | ------------------------------------------------------------ |
| **bystackd**   | bystackd command can help to initialize and launch bystack domain by custom parameters. `bystackd --help` for command line options. |

## Running bystack

Currently, bystack is still in active development and a ton of work needs to be done, but we also provide the following content for these eager to do something with `bystack`. This section won't cover all the commands of `bystackd` at length, for more information, please the help of every command, e.g., `bystackd help`.

### Initialize

First of all, initialize the node:

Download testnet.zip or mainnet.zip

```bash
$ cd ./build/bin/bystackd
$ ./bystackd --datadir {{datadir}}  init genesis.json
```

After that, you'll see `config.toml` generated, then launch the node.

### launch

``` bash
$ ./bystackd --config config.toml --datadir {{datadir}}  --cache 8000
```

## Contributing

Thank you for considering helping out with the source code! Any contributions are highly appreciated, and we are grateful for even the smallest of fixes!

If you run into an issue, feel free to [bystack issues](https://github.com/Bytom/bystack/issues/) in this repository. We are glad to help!

## License

[AGPL v3](./LICENSE)