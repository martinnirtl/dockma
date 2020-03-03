#!/bin/bash

ENV="$(dockma env home)"

echo "Opening vscode at: $ENV"
code $ENV
