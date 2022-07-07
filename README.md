# jat

It's `cat`, but for JWTs. A command-line tool to parse JWT payloads, written in Go.

[![.github/workflows/build.yml](https://github.com/mathcale/jat/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/mathcale/jat/actions/workflows/build.yml)

## Usage

```sh
jat <some-jwt>
```

## Install

1. Navigate to the [Releases](https://github.com/mathcale/jat/releases) page;
2. Download the corresponding version for your OS;
3. Decompress the archive and place its content on a directory of your choice (preferably `$HOME/.local/bin`);
4. _(Optional)_ Add this directory to PATH, if it's not already;
5. Done!

## Running locally

1. Clone this repository;
2. Compile the program with `make build`;
3. A fresh new binary will be outputted to the `bin` directory, and you can run it with `./bin/jat`;
