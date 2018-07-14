# AiBlocks
Decentralized applications and smart contracts powered by artificial intelligence. Grows smarter with each block. 

## Origin
AiBlocks was originally a fork of Ethereum Classic. An attempt to bring AI to decentralized applications.

Starting in 2015, our team has been developing an artificial intelligence protocol specifically designed to grow and and adapt to the latest blockchain technology. We have since integrated our AI protocol into a whole new system of AI powered decentralized applications, running on our growing network - AiBlocks.

## Our Goal

Our main goal is to see our network of applications grow and adapt, improving the collective artificial intelligence with each new block, transaction, and every AI made decision.

In a world that sees artificial intelligence playing a huge future role in pretty much all aspects of technology, AiBlocks will be at the very core of many of the first big AI powered applications.

With AiBlocks, users have already created some pretty amazing applications, tutorials and examples of practical AI use; shared in AiBlocks's "Developer Center", open for anyone to browse through or publish their own. learn new cutting edge AI programming.

## Install

### :rocket: From a release binary
The simplest way to get started running a node is to visit our [Releases page](https://github.com/aiblocks/aiblocks/releases) and download a zipped executable binary (matching your operating system, of course), then moving the unzipped file `geth` to somewhere in your `$PATH`. Now you should be able to open a terminal and run `$ geth help` to make sure it's working.

### :hammer: Building the source

If your heart is set on the bleeding edge, install from source. However, please be advised that you may encounter some strange things, and we can't prioritize support beyond the release versions. Recommended for developers only.

#### Dependencies
Building geth requires both Go >=1.9 and a C compiler. On Linux systems,
a C compiler can, for example, by installed with `sudo apt-get install
build-essential`. On Mac: Coming Soon.

#### Get source and package dependencies
```
$ go get -v github.com/aiblocks/aiblocks/...`
```

#### Install and build command executables

Executables installed from source will, by default, be installed in `$GOPATH/bin/`.

##### With go:

- the full suite of utilities:
```
$ go install github.com/aiblocks/aiblocks/...`
```

- just __geth__:
```
$ go install github.com/aiblocks/aiblocks/`
```

##### With make:
```
$ cd $GOPATH/src/github.com/aiblocks/aiblocks/
```

- the full suite of utilities:
```
$ make install
```

- just __geth__:
```
$ make install_geth
```

> For further `make` information, use `make help` to see a list and description of available make
> commands.

## Contribution

Thank you for considering to help out with the source code!

The core values of democratic engagement, transparency, and integrity run deep with us. We welcome contributions from everyone, and are grateful for even the smallest of fixes.  :clap:

This project was originally forked from Ethereum Classic, which was migrated from the now hard-forked [Ethereum (ETHF) Github project] (https://github.com/ethereum)

If you'd like to contribute to AiBlocks, please fork, fix, commit and send a pull request for the maintainers to review and merge into the main code base. If you wish to submit more complex changes, please check up with the core devs first at contact@aiblocks.io

Please see the [Wiki](https://github.com/ethereumproject/go-ethereum/wiki) for more details on configuring your environment, managing project dependencies, and testing procedures.

## License

The AiBlocks library (i.e. all code outside of the `cmd` directory) is licensed under the [GNU Lesser General Public License v3.0](http://www.gnu.org/licenses/lgpl-3.0.en.html), also included in our repository in the `COPYING.LESSER` file.

The AiBlocks binaries (i.e. all code inside of the `cmd` directory) is licensed under the [GNU General Public License v3.0](http://www.gnu.org/licenses/gpl-3.0.en.html), also included in our repository in the `COPYING` file.
