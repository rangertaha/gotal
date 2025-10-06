#!/bin/bash

# This script bumps the version number in version.go

# File containing the version info
version_file="internal/version.go"

# Get current version
current_version=$(grep -e 'VERSION = ".*"' $version_file | cut -d'"' -f2)

# Split version into components
IFS='.' read -r major minor patch <<< "$current_version"

# Increment patch version
new_patch=$((patch + 1))
new_version="$major.$minor.$new_patch"

# Update version in file
sed -i "s/VERSION = \".*\"/VERSION = \"$new_version\"/" $version_file

echo "Bumped version from $current_version to $new_version"
