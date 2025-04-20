# Local SSL Certificates

This directory is for storing local development SSL certificates. Do not commit these certificates to version control.

## Generating Certificates

### macOS

```bash
# Install mkcert
brew install mkcert

# Install the local CA
mkcert -install

# Generate certificates
mkcert -cert-file local-cert.pem -key-file local-key.pem "localhost" "*.localhost" "api.localhost"
```

### Linux

```bash
# Ubuntu/Debian
sudo apt install libnss3-tools
sudo apt install mkcert

# Or build from source
go install filippo.io/mkcert@latest

# Install the local CA
mkcert -install

# Generate certificates
mkcert -cert-file local-cert.pem -key-file local-key.pem "localhost" "*.localhost" "api.localhost"
```

### Windows

```powershell
# Using chocolatey
choco install mkcert

# Or using scoop
scoop install mkcert

# Install the local CA
mkcert -install

# Generate certificates
mkcert -cert-file local-cert.pem -key-file local-key.pem "localhost" "*.localhost" "api.localhost"
```

## Certificate Files

After generation, you should have two files in this directory:
- `local-cert.pem` - The SSL certificate
- `local-key.pem` - The private key

These files are used by Traefik to provide HTTPS for local development.

## Important Notes

1. Never commit these certificates to version control
2. The certificates are only for local development
3. Production environments will use Let's Encrypt for automatic SSL certification
4. If you get certificate warnings, ensure you've run `mkcert -install`
5. The certificates are valid for:
   - localhost
   - *.localhost (including api.localhost)

## Troubleshooting

If you see certificate warnings:
1. Ensure mkcert is installed
2. Run `mkcert -install` again
3. Regenerate the certificates
4. Restart your browser
5. Restart the Docker containers
