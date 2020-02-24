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

- implement arguement based command execution
  - check all commands for arg count check and command usage definition
  - add 'dynamic' bash completion
- build command/solution to change selection without downing env in between
- format error messages which get logged with prefix 'Error: ' (not handed over to config.Save func) and include ':'
- finish makefile
  - bind version, commit, date to vars provided on 'go build'
- add logging for verbose output
  - inspect cmd only for external commands
  - verboseFlag shall ignore timebridgers
- write long command descriptions (eg. doctl help completion)
- rethink git-pull setting and updating it becoming git-based project
- use table instead of long lists (dockma profile list --services)
- support swarm mode and move towards kubernetes ?

- use bats for system tests (https://github.com/sstephenson/bats)
- build "assessment" tool for docker-compose files
- write unit tests
