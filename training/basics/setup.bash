#!/usr/bin/env bash 

mkdir -p "$GOPATH/src/github.com/AgarwalConsulting"
git clone https://github.com/AgarwalConsulting/Go-Training.git "$GOPATH/src/github.com/AgarwalConsulting/Go-Training"
cd "$GOPATH/src/github.com/AgarwalConsulting/Go-Training"

go install github.com/AgarwalConsulting/Go-Training@latest
ln -s "$GOPATH/src/github.com/AgarwalConsulting/Go-Training"  Go-Training

Go-Training
