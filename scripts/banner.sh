#!/bin/bash

# This script updates the banner in the version.go file with the current date and git commit hash


# Get the current date in YYYY-MM-DD format
current_date=$(date +%Y-%m-%d)

# Get the latest git commit hash (first 7 chars)
commit_hash=$(git rev-parse --short HEAD)

# File containing the version info
version_file="internal/version.go"

# Update the COMPILED date and COMMIT hash
sed -i "s/COMPILED = \".*\"/COMPILED = \"$current_date\"/" $version_file
sed -i "s/COMMIT = \".*\"/COMMIT = \"$commit_hash\"/" $version_file

echo "Updated banner with:"
echo "Date: $current_date"
echo "Commit: $commit_hash"
