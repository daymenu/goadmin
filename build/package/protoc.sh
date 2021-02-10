#!/bin/bash

protocPath="/home/hanjian/workspace/go/goadmin/api/"


for dir in `ls $protocPath`
do
    for file in `ls $protocPath$dir/*.proto`
    do
        echo $file
        `protoc -I. --proto_path=$protocPath$dir --go_out=$protocPath$dir --go_opt=paths=source_relative $file`
    done
done