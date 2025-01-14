#!/bin/bash
echo "## Test Coverage Report"
echo '```'
go test -cover ./handler . 
echo '```'