#!/bin/bash
set -e

cp omap.go ~/go/src/local/jgbaldwinbrown/omap
go build test.go
