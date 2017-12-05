# hcloud: Command-line interface for Hetzner Cloud

[![Build status](https://travis-ci.org/hetznercloud/cli.svg?branch=master)](https://travis-ci.org/hetznercloud/cli)

`hcloud` is a command-line interface for interacting with Hetzner Cloud.

**This tool is still in development and lacking a lot of features.**

We do not accept pull requests at the moment.

## Installation

    go get github.com/hetznercloud/cli/cmd/hcloud

## Usage

Configure the `hcloud` program to use your token:

    hcloud configure

See `hcloud help` for a list of commands.

## Example

```
$ hcloud server list            
ID       NAME                    STATUS    IPV4
210216   test1                   running   78.46.122.12
210729   ubuntu-8gb-nbg1-dc3-1   running   94.130.177.158
```