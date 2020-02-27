# Dockma CLI [WIP]

Official Website and Documentation on [dockma.dev](https://dockma.dev).

<!-- Level up your docker-compose game during development!

- No more navigating to your docker-compose file
- Create service host definitions dynamically whether they run in docker or locally
- Launch your defined services with an interactive CLI
- Switch between your docker(-compose) based projects
- Add custom commands by writing your own scripts -->

```
A fast and flexible CLI tool to boost your productivity during development in docker-compose based environments. Full documentation is available on https://dockma.dev.

Usage:
  dockma [command]

Available Commands:
  completion  Generate shell completion code
  config      Dockma configuration details
  down        Downs active environment
  env         Environments reflect docker-compose based projects
  help        Help about any command
  init        Initialize the Dockma CLI.
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

- [Install](#install)
  - [macOS](#macos)
  - [Linux](#linux)
  - [Windows](#windows)
  - [Build from source](#build-from-source)
- [Setup](#setup)
- [Usage](#usage)
  - [Commands](#commands)
- [Roadmap](#roadmap)
  - [Features/Todos](#features/todos)
  - [Other](#other)
- [Contribute](#contribute)

## Install

### macOS

### Linux

### Windows

### Build from source

## Setup

## Usage

### Commands

## Roadmap

### Features/Todos

- build command/solution to change selection without downing env in between
- add support for env var prefix (eg. 'DOCKMA\_') and make configurable per environment (subcommand 'env set' is taken and should prob be changed to 'env act(ivate)')
- think about wrapping viper completely in config package
- add logging for verbose output
  - inspect cmd only for external commands
  - verboseFlag shall ignore timebridgers
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
