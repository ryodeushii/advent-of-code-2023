#!/bin/bash

name=${1}
mod_base_name="github.com/ryodeushii/advent-of-code-2023/"

echo "Creating project for $name"

# install airr globally to watch changes
echo "Installing Golang Air to watch for changes"
go install github.com/cosmtrek/air@latest

# create directory for "day-*" from arg
cp -R boilerplate $name
# go to new dir and init go module
cd $name && go mod init "$mod_base_name$name" && cd --

