#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CERT_DIR="$SCRIPT_DIR/traefik"
CERT_FILE="$CERT_DIR/cert.pem"
KEY_FILE="$CERT_DIR/key.pem"

# Check if certificates already exist
if [ -f "$CERT_FILE" ] && [ -f "$KEY_FILE" ]; then
  echo "Certificates already exist at:"
  echo "  - $CERT_FILE"
  echo "  - $KEY_FILE"
  exit 0
fi

# Check if mkcert is installed
if ! command -v mkcert &>/dev/null; then
  echo "Error: mkcert is not installed."
  echo "Please install it first:"
  echo "  - Arch/Manjaro: sudo pacman -S mkcert"
  echo "  - Ubuntu/Debian: sudo apt install mkcert"
  echo "  - macOS: brew install mkcert"
  exit 1
fi

# Install local CA if not already done
echo "Installing local CA..."
mkcert -install

# Create cert directory if it doesn't exist
mkdir -p "$CERT_DIR"

# Generate certificates
echo "Generating certificates..."
cd "$CERT_DIR"
mkcert -cert-file cert.pem -key-file key.pem "*.coli.localhost" "localhost"

echo "Certificates created successfully at:"
echo "  - $CERT_FILE"
echo "  - $KEY_FILE"
