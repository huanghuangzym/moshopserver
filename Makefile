

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: manager

# Run tests

# Build manager binary
manager: generate fmt vet
	go build -ldflags "-X main.build=`git rev-parse HEAD`" main.go

build: fmt  vet
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"'  -o mshop main.go

# Run against the configured Kubernetes cluster in ~/.kube/config

# Install CRDs into a cluster

# Uninstall CRDs from a cluster

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config

# Generate manifests e.g. CRD, RBAC etc.

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
