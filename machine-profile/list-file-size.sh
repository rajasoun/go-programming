#!/usr/bin/env bash

function print_file_size(){
    file=$1
    filesize=$(ls -lh $file | awk '{print  $5}')
    #filesize=$(stat -c %s $file)
    #kb=$(($filesize/1024))
    echo -e "$file : $filesize"
}

print_file_size "src/machine-profile.go"
print_file_size "bin/machine-profile-linux"
