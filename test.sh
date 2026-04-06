#!/bin/bash
git checkout pkg/config/config.go
$(go env GOPATH)/bin/golines -w pkg/config/config.go
git diff pkg/config/config.go
