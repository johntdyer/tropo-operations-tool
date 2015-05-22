#!/bin/sh
CGO_ENABLED=0

echo "Building Tropo Operations Tool"
# Build for OSX
gb -o builds/tot.osx tot.go papi.go papi_structs.go utils.go config.go commands.go actions.go address_helper.go sip.go tables.go version.go
if [ $? -eq 0 ]; then
  echo "Success Build artifact - builds/tot.osx"
else
  echo "Build error"
  exit $?
fi

# Build for Linux
GOOS=linux
GOARCH=amd64
CGO_ENABLED=0
go build -o builds/tot.linux tot.go papi.go papi_structs.go utils.go config.go commands.go actions.go address_helper.go sip.go tables.go version.go
if [ $? -eq 0 ]; then
  echo "Success Build artifact - builds/tot.linux"
else
  echo "Build error"
  exit $?
fi


chmod +x builds/tot.osx
chmod +x builds/tot.linux
