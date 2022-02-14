#!/usr/bin/env bash

read -p 'Module Name: ' module

echo -e "Creating $module"
mkdir -p "$module"/{bin,pkg,src}
cp hello/.gitignore           "$module/"
cp hello/Makefile             "$module/"
cp hello/go.mod               "$module/"
cp hello/src/hello.go         "$module/src/$module.go"
sed  -i "s/hello/$module/g"   "$module/go.mod"
sed  -i "s/Hello, world!/$module/g" "$module/src/$module.go"


echo -e "Done !!!"