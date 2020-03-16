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

The following list of features outlines the main features of **dockma** and how it can improve your workflow:

- No more navigating to your docker-compose files
- Develop locally while the rest of your micro-services run inside docker
- Launch your defined services with an interactive CLI or directly from the command line
- Switch between your docker-compose based projects quickly
- Extend **dockma** by custom scripts and bring your own specific functionality (e.g. database import/export scripts)

## Install

**Dockma CLI** gets built and released for the following operating systems:

- Linux
- macOS
- Windows

If **dockma** is not built and released automatically for your platform, you can also [build **dockma** from source](#build-dockma-from-source).

### macOS

The recommended way to install **dockma** on macOS is to use the **dockma homebrew-tap**:

```
brew install martinnirtl/tap/dockma
```

### Linux

To install **dockma** on linux, you can **download** the binary manually **from [github releases](https://github.com/martinnirtl/dockma/releases)** or you can also use utilities like `curl`:

```
curl -OL https://github.com/martinnirtl/dockma/releases/download/v<version>/dockma-<version>-linux-i386.tar.gz
```

Make sure to replace `<version>` by the version you want to download!

Afterwards just extract the binary by:

```
tar xf ~/dockma-<version>-linux-i386.tar.gz
```

To make **dockma** globally available on the command line, run the following command (maybe as `sudo`):

```
mv ./dockma /usr/local/bin
```

Another option is installing **dockma** via linux homebrew (see [homebrew installation macOS](#macos))

### Windows

To get **dockma** for windows, you can **download it from [github releases](https://github.com/martinnirtl/dockma/releases)** or [**build it from source**](#build-dockma-from-source).

If you have `wget` installed, you can also download it from the command line:

```
wget https://github.com/martinnirtl/dockma/releases/download/v<version>/dockma-<version>-windows-i386.zip
```

Make sure to replace `<version>` by the version you want to download!

### Build Dockma from Source

To build **dockma** from source, you have to **setup [go](https://golang.org/doc/install)** first. Afterwards just clone the repository by running `git clone https://github.com/martinnirtl/dockma.git` and execute `make build` in the project directory. This will generate a **dockma binary** in _builds_ named **dockma**.

**Linux or macOS:** You can `cp` the binary to _/usr/local/bin_ or create a symlink with `ln -s` to make it globally available on the command line.

## Setup

To initialize **dockma** on your system just run the following command:

![Dockma init command GIF](https://raw.githubusercontent.com/martinnirtl/dockma/master/assets/gifs/dockma_init.gif)

Afterwards you usually would continue with adding the first docker-compose based project to **dockma**:

![Dockma env init command GIF](https://raw.githubusercontent.com/martinnirtl/dockma/master/assets/gifs/dockma_env_init_hello-world.gif)

### Shell Completion

Shell completion is supported for bash, zsh and powershell. To add it to your **dockma** installation, run the `dockma completion` command (example for macOS with brew's bash-completion):

![Dockma completion command GIF](https://raw.githubusercontent.com/martinnirtl/dockma/master/assets/gifs/dockma_completion_bash.gif)

## Usage

As already mentioned in [setup](#setup), the usual thing you would start with is adding a so-called environment (directory containing docker-compose file). Below you find a [docker-compose example](https://github.com/martinnirtl/dockma/tree/master/examples/getting-started-env), which lists three services:

- backend service is based on the popular http-bin container
- middleware-service uses the backend service
- polling-service is uses the middleware-service as API

The **dockma** specific part in docker-compose files are the variables (see `BACKEND_HOST` and `MIDDLEWARE_HOST` variables), which get set by **dockma** automatically depending on the service selection on `dockma up`. This way all services' addresses get set correctly whether running inside docker or locally on your machine. The variable names are created dynamically following the pattern `<SERVICE-NAME>_HOST` and are always all upper case.

```yaml
version: "3"

services:
  backend:
    image: kennethreitz/httpbin:latest

  middleware:
    build: ../middleware-service
    image: middleware-service:local
    environment:
      - PORT=3500
      - BACKEND_BASEURL=http://${BACKEND_HOST} # dynamic address resolution by dockma

  polling:
    build: ../polling-service
    image: polling-service:local
    environment:
      - PORT=4000
      - POLL_INTERVAL_MS=5000
      - API_BASEURL=http://${MIDDLEWARE_HOST}:3500 # dynamic address resolution by dockma
```

The following command shows now how to add the [getting-started](https://github.com/martinnirtl/dockma/tree/master/examples/getting-started-env) environment to **dockma**:

![Dockma env init command GIF](https://raw.githubusercontent.com/martinnirtl/dockma/master/assets/gifs/dockma_env_init_getting-started.gif)

After adding it, you can check your configured environments:

![Dockma envs list command GIF](https://raw.githubusercontent.com/martinnirtl/dockma/master/assets/gifs/dockma_envs_list.gif)

Now let's say you want to add a small feature to the [middleware-service](https://github.com/martinnirtl/dockma/tree/master/examples/middleware-service). You would run the `middleware-service` locally with `npm run dev` and do your changes. Rest of your services would run in docker with `dockma up`:

![Dockma up command GIF](https://raw.githubusercontent.com/martinnirtl/dockma/master/assets/gifs/dockma_up.gif)

```yaml
# Important: Make sure you map all ports required for middleware-service to work to localhost in docker-compose.override.yml. Otherwise middleware-service can't reach the API.

version: "3"

services:
  backend:
    ports:
      # - "HOST:CONTAINER"
      - "3000:80"
```

You may have noticed, that after `dockma up` was executed successfully, we also checked if all docker containers were up and running by running `dockma logs -f`.

For a more detailed tutorial, have a look into [examples](https://github.com/martinnirtl/dockma/tree/master/examples) or get started by exploring `dockma help`.

## Roadmap

### Features/Todos

| Description                                                                                                                         | Progress | Expected Finishing |
| ----------------------------------------------------------------------------------------------------------------------------------- | -------- | ------------------ |
| build command/solution to change selection without downing env in between                                                           |          |                    |
| think about wrapping viper completely in config package                                                                             |          |                    |
| add logging for verbose output (inspect cmd only for external commands, verboseFlag shall ignore timebridgers)                      |          |                    |
| use tabby and tablewriter where it makes sense                                                                                      |          |                    |
| write long command descriptions (eg. doctl help completion)                                                                         |          |                    |
| rethink git-pull setting and updating it becoming git-based project                                                                 |          |                    |
| use table instead of long lists (dockma profile list --services)                                                                    |          |                    |
| add 'dynamic' bash completion                                                                                                       |          |                    |
| support swarm mode and move towards kubernetes ?                                                                                    |          |                    |
| build "assessment" tool for docker-compose files                                                                                    |          |                    |
| write unit tests                                                                                                                    |          |                    |
| add support for env var prefix (eg. 'DOCKMA\_') and make configurable per environment (subcommand 'env set' should prob be changed) |          |                    |

### Other

| Description                                                     | Progress | Expected Finishing |
| --------------------------------------------------------------- | -------- | ------------------ |
| use bats for system tests (https://github.com/sstephenson/bats) |          |                    |

## Contribute

Everybody is welcome to contribute new features or bugfixes. To do so, please first create a issue with the respective [template](https://github.com/martinnirtl/dockma/issues/new/choose). Fork the repository, add your code and submit your changes linked to the feature request or bug report as a pull request.

After reviewing it, your changes get merged into the project.
