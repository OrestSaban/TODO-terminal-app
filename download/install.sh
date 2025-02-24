#!/bin/bash

# Define the version and platform
VERSION="v1.0"
PLATFORM="linux-amd64" # Change to "darwin-amd64" for macOS

# Define the download URL
URL="https://github.com/OrestSaban/TODO-terminal-app/releases/download/$VERSION/todo-darwin-amd64"

# Download the binary
echo "Downloading todo tool..."
curl -L -o todo $URL

# Make the binary executable
chmod +x todo

# Move the binary to /usr/local/bin
echo "Installing todo tool to /usr/local/bin..."
sudo mv todo /usr/local/bin/

# Verify installation
if command -v todo &> /dev/null; then
    echo "Installation successful! You can now use the 'todo' command."
else
    echo "Installation failed. Please check your PATH and try again."
fi