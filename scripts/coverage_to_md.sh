#!/bin/bash
echo "## Test Coverage Report"
echo '```'
go tool cover -func=coverage.out
echo '```'