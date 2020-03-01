# Dockma CLI ![build](https://github.com/martinnirtl/dockma/workflows/build/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/martinnirtl/dockma)](https://goreportcard.com/report/github.com/martinnirtl/dockma) ![GitHub](https://img.shields.io/github/license/martinnirtl/dockma) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/martinnirtl/dockma)

Level up your docker-compose game during development with [**dockma**](#features)!

```
A fast and flexible CLI tool to boost your productivity during development in docker-compose based environments.

Usage:
  dockma [command]

Available Commands:
  completion  Generate shell completion code
  config      Dockma configuration details
  down        Downs active environment
  env         Environments reflect docker-compose based projects
  help        Help about any command
  init        Initialize dockma
  inspect     Print detailed output of previously executed external command
  logs        Logs output of all or only selected services
  profile     Manage profiles (named service selections)
  ps          List running services of active environment
  restart     Restart all or only selected services
  script      Run script (.sh) located in scripts dir of active environment
  up          Runs active environment with profile or service selection
  version     Print the version number of dockma.

Flags:
  -h, --help      help for dockma
  -v, --verbose   verbose output

Use "dockma [command] --help" for more information about a command.
```

## Content

- [Features](#features)
- [Install](#install)
  - [macOS](#macos)
  - [Linux](#linux)
  - [Windows](#windows)
  - [Build from source](#build-from-source)
- [Setup](#setup)
- [Usage](#usage)
- [Roadmap](#roadmap)
  - [Features/Todos](#features/todos)
  - [Other](#other)
- [Contribute](#contribute)

## Features

The following list of features outlines the main features of dockma and how it can improve your workflow:

- No more navigating to your docker-compose files
- Develop locally while the rest of your micro-services run inside docker
- Launch your defined services with an interactive CLI or directly from the command line
- Switch between your docker-compose based projects quickly
- Extend dockma by custom scripts and bring your own specific functionality (e.g. database import/export scripts)

## Install

Dockma CLI gets built and released for the following operating systems. If your platform is not supported, you can also [build **dockma** from source](#build-dockma-from-source).

### macOS

The recommended way to install dockma on macOS is to use the **dockma homebrew-tap**:

```
brew install martinnirtl/dockma
```

Another option is to **use** the **install script**:

```
curl TODO
```

### Linux

On Linux, you can either **download** the binary directly **from [github](https://github.com/martinnirtl/dockma/releases)** or you can use the install script:

```
curl TODO
```

### Windows

The only way to get **dockma** for windows at the moment, is to **download it from [gitub](https://github.com/martinnirtl/dockma/releases)** directly or [**build it from source**](#build-dockma-from-source).

### Build Dockma from Source

To build dockma from source, you have to **install [go](https://golang.org/doc/install)** first. Afterwards just clone the repository by running `git clone https://github.com/martinnirtl/dockma.git` and execute `make build` in the project directory. This will generate a **dockma binary** in _builds_ named **dockma**.

**Linux or macOS:** You can `cp` the binary to _/usr/local/bin_ or create a symlink with `ln -s` to make it globally available on the command line.

## Setup

To initialize **dockma** on your system just run the following command:

// dockma init GIF

Afterwards you usually would continue with adding the first docker-compose based project to **dockma**:

// dockma env init dockma/examples/hello-world GIF

### Shell Completion

Shell completion is supported for bash, zsh and powershell. To add it to your dockma installation, run the dockma completion command (example for macOS with brew's bash-completion):

// dockma completion bash > $DOCKMA_HOME/completion && ln -s $DOCKMA_HOME/completion /usr/local/etc/bash_completion.d/dockma

**Please note** that, using `DOCKMA_HOME` variable assumes you already set it somewhere.

## Usage

As already mentioned in [setup](#setup), the usual thing you would start with is adding a so-called environment (mainly represents a dir containing docker-compose file).

// dockma env init

After adding it, you can check your configured environments:

// dockma envs list --path

Now let's say you want to add a small feature in the service `backend`. You would run `backend` locally with `yarn run dev`. Rest of your services would be running in docker:

// dockma up

**Important:** Make sure you map all required ports required for `backend` to work to your localhost in docker-compose.override.yaml.

Finally, after `dockma up` was executed successfully, you can check if everything is up and running by:

// dockma logs

For a more detailed tutorial, have a look into [examples](https://github.com/martinnirtl/dockma/tree/master/examples).

## Roadmap

### Features/Todos

- build command/solution to change selection without downing env in between
- add support for env var prefix (eg. 'DOCKMA\_') and make configurable per environment (subcommand 'env set' is taken and should prob be changed to 'env act(ivate)')
- think about wrapping viper completely in config package
- add logging for verbose output
  - inspect cmd only for external commands
  - verboseFlag shall ignore timebridgers
- use tabby and tablewriter where it makes sense
- write long command descriptions (eg. doctl help completion)
- rethink git-pull setting and updating it becoming git-based project
- use table instead of long lists (dockma profile list --services)
- add 'dynamic' bash completion
- support swarm mode and move towards kubernetes ?
- build "assessment" tool for docker-compose files
- write unit tests 😂

### Other

- use bats for system tests (https://github.com/sstephenson/bats)
- show build tags etc. on top of readme

## Contribute

Everybody is welcome to contribute new features or bugfixes. To do so, please first create a issue with the respective [template](https://github.com/martinnirtl/dockma/issues/new/choose). Fork the repository, add your code and submit your changes linked to the feature request or bug report as a pull request.

After reviewing it, your changes get merged into the project.
