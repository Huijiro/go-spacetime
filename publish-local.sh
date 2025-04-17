#!/bin/bash

spacetime logout
spacetime login --server-issued-login local
spacetime publish -y --server local --bin-path=$(dirname "$0")/spacetime.wasm

echo -e "\nChecking Database...\n"

DB_HASH=$(spacetime list 2>/dev/null | tail -1)
spacetime logs $DB_HASH
