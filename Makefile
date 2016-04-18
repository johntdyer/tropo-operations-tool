SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=tropo-operations-tool

VERSION=$(shell cat version.go | awk '{FS="="}{print $2}' | sed 's_"__g')

LDFLAGS=-ldflags "-X main.Build=`git rev-parse HEAD`"

.DEFAULT_GOAL: $(BINARY)

packages/$(BINARY): $(SOURCES)
		# go build ${LDFLAGS} -o packages/${BINARY} spark-pd-oncall.go pagerduty.go spark.go version.go file-utils.go
		gox \
		-osarch="!darwin/386"    \
		-osarch="!linux/386"     \
		-osarch="!openbsd/386"   \
		-osarch="!openbsd/amd64" \
		-osarch="!freebsd/amd64" \
		-osarch="!freebsd/arm"   \
		-osarch="!freebsd/386"   \
		-osarch="!windows/386"   \
		-osarch="!netbsd/arm"    \
		-osarch="!linux/arm"     \
		-osarch="!netbsd/386"    \
		-osarch="!netbsd/amd64"  \
		${LDFLAGS}               \
		-output packages/{{.Dir}}_{{.OS}}_{{.Arch}}

.PHONY: clean
clean:
		@rm -vf $(BINARY)_*_386 $(BINARY)_*_amd64 $(BINARY)_*_arm $(BINARY)_*.exe
