#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Define paths
GO_APP_NAME="sysmetrics" # Name of the Go binary
CORE_FOLDER="../core"    # Go application folder
UI_FOLDER="../ui"        # Electron/React app folder
DIST_FOLDER="../dist"    # Final distribution folder
UI_SRC_FOLDER="$UI_FOLDER/src"
GO_BUILD_OUTPUT="$UI_SRC_FOLDER/$GO_APP_NAME"

# Step 1: Clean existing dist folder
echo "Cleaning up existing dist folder..."
if [ -d "$DIST_FOLDER" ]; then
  rm -rf "$DIST_FOLDER"
fi
mkdir "$DIST_FOLDER"
echo "Dist folder cleaned and ready."

# Step 2: Build the Go application
echo "Building Go application..."
cd $CORE_FOLDER
GOOS=darwin GOARCH=amd64 go build -o "$GO_BUILD_OUTPUT" main.go
echo "Go application built and moved to $GO_BUILD_OUTPUT"

# Step 3: Build the Electron application
echo "Building Electron application..."
cd $UI_FOLDER
npm install # Ensure dependencies are installed
npm run make
echo "Electron application built successfully."

# Step 4: Move Electron output to dist
echo "Moving Electron build to dist..."
mv out/* "$DIST_FOLDER"
echo "Electron build moved to $DIST_FOLDER"

# Step 5: Move Go binary to dist
echo "Copying Go binary to dist..."
cp "$GO_BUILD_OUTPUT" "$DIST_FOLDER"
echo "Go binary copied to $DIST_FOLDER"

echo "Build process completed successfully!"
