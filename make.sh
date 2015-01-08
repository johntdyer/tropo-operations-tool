#!/bin/sh
CGO_ENABLED=0

# Build for OSX
go build -o builds/tropo.osx tropo.go papi.go papi_structs.go utils.go config.go commands.go actions.go address_helper.go sip.go tables.go version.go
if [ $? -eq 0 ]; then
  echo "Success Build artifact - builds/tropo.osx"
else
  echo "Build error"
  exit $?
fi

# Build for Linux
GOOS=linux
GOARCH=amd64
CGO_ENABLED=0
go build -o builds/tropo.linux tropo.go papi.go papi_structs.go utils.go config.go commands.go actions.go address_helper.go sip.go tables.go version.go
if [ $? -eq 0 ]; then
  echo "Success Build artifact - builds/tropo.linux"
else
  echo "Build error"
  exit $?
fi


