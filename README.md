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

- walk through all cmds
  - check user interactions (logs and surveys)
    - rethink the way of information output
    - unify success and error messages
- finish makefile
- implement arguement based command execution
  - add 'dynamic' bash completion via ValidArgs
  - add func for Args field (take valid args only but also no args for interactive mode)
- add logging for verbose output (see doctl ?)
  - inspect cmd only for external commands
  - verboseFlag shall ignore timebridgers
- write long command descriptions (eg. doctl help completion)
- rethink git-pull setting and updating it becoming git-based project
- bind version, commit, date to vars provided on 'go build'
- support swarm mode and move towards kubernetes ?

- use bats for system tests (https://github.com/sstephenson/bats)
- build "assessment" tool for docker-compose files
- write unit tests
