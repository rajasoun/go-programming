#!/usr/bin/env bash

read -p 'Module Name: ' module

echo -e "Creating $module"
mkdir -p "$module"/{bin,pkg,src}
cp hello-world/.gitignore           "$module/"
cp hello-world/Makefile             "$module/"
cp hello-world/go.mod               "$module/"
cp hello-world/src/hello.go         "$module/src/$module.go"
sed  -i "s/hello-world/$module/g"   "$module/go.mod"
sed  -i "s/Hello, world!/$module/g" "$module/src/$module.go"


echo -e "Done !!!"