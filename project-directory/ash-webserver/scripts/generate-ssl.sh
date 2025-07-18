#!/bin/bash

# Create SSL directory
mkdir -p ssl

# Generate private key
openssl genrsa -out ssl/ash.it.com.key 2048

# Generate certificate signing request
openssl req -new -key ssl/ash.it.com.key -out ssl/ash.it.com.csr -subj "/C=US/ST=State/L=City/O=Organization/CN=ash.it.com"

# Generate self-signed certificate
openssl x509 -req -in ssl/ash.it.com.csr -signkey ssl/ash.it.com.key -out ssl/ash.it.com.crt -days 365

# Clean up CSR
rm ssl/ash.it.com.csr

echo "SSL certificates generated successfully!"
