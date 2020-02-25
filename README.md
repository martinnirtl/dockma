# Dockma CLI [WIP]

Level up your docker-compose game during development!

- No more navigating to your docker-compose file
- Create service host definitions dynamically whether they run in docker or locally
- Launch your defined services with an interactive CLI
- Switch between your docker(-compose) based projects
- Add custom commands by writing your own scripts

## Install

### macOS

### Linux

### Windows

## Setup

## Usage

## Commands

## Scripts

## Contribute

## TODOs

### Code-Related

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

- finish makefile
  - bind version, commit, date to vars provided on 'go build'
- use bats for system tests (https://github.com/sstephenson/bats)
