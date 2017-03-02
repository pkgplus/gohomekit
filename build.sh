#!/bin/sh
version=v0.0.1
cd target

function build_tar_gz()
{
    echo "build on $GOOS-$GOARCH ..."
    go build ../gohomekit.go
    
    echo "tar package ..."
    tar zcvf "./gohomekit-$version-$GOOS-$GOARCH.tar.gz" ./conf ./gohomekit
    rm -f ./gohomekit
}

function build_zip()
{
    echo "build on $GOOS-$GOARCH ..."
    go build ../gohomekit.go
    
    echo "zip package ..."
    if [ "$GOOS" == "windows" ];then
        zip "./gohomekit-$version-$GOOS-$GOARCH.zip" ./conf ./gohomekit.exe
        rm -f ./gohomekit.exe
    else
        zip "./gohomekit-$version-$GOOS-$GOARCH.zip" ./conf ./gohomekit
        rm -f ./gohomekit
    fi
}

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=arm
build_tar_gz

export GOARCH=amd64
build_tar_gz

export GOOS=darwin
export GOARCH=amd64
build_tar_gz

export GOOS=windows
export GOARCH=386
build_zip

echo "over!!!"
