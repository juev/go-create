.DEFAULT_GOAL = all

BINARY = {{ .Name }}
VET_REPORT = vet.report
GOARCH = amd64

VERSION?={{ .Version }}
BUILD_TIME?=$(shell date -u '+%Y%m%d-%H%M')
COMMIT?=$(shell git rev-parse HEAD)
BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-s -w \
	-X main.version=${VERSION} \
	-X main.commit=${COMMIT} \
	-X main.branch=${BRANCH}\
	-X main.buildTime=${BUILD_TIME}"

# Build the project
all: clean vet linux freebsd darwin windows

linux:
	@mkdir -p bin
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o bin/${BINARY}-linux-${GOARCH} ./src

freebsd:
	@mkdir -p bin
	GOOS=freebsd GOARCH=${GOARCH} go build ${LDFLAGS} -o bin/${BINARY}-freebsd-${GOARCH} ./src

darwin:
	@mkdir -p bin
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o bin/${BINARY}-darwin-${GOARCH} ./src

windows:
	@mkdir -p bin
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o bin/${BINARY}-windows-${GOARCH}.exe ./src

vet:
	@go vet ./src/... > ${VET_REPORT} 2>&1

fmt:
	@go fmt $(go list ./src/... | grep -v /vendor/)

clean:
	@rm -f ${VET_REPORT}
	@rm -rf bin

run:
	go run main.go

.PHONY: all linux darwin windows vet fmt clean run
