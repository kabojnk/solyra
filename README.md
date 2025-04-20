# Solyra

A photographer's companion application with a Go backend API and React PWA frontend.

## Project Structure

- `/app` - Frontend React PWA
- `/server` - Backend Go API
- `/certs` - Local SSL certificates
- `/traefik` - Traefik reverse proxy configuration

## Prerequisites

- Docker and Docker Compose
- Node.js (for local frontend development)
- Go 1.21+ (for local backend development)
- mkcert (for local SSL certificates)

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/yourusername/solyra.git
cd solyra
```

2. Set up local SSL certificates:
```bash
# Install mkcert (if not already installed)
# macOS
brew install mkcert
# Linux
sudo apt install mkcert
# Windows (using chocolatey)
choco install mkcert

# Create and install the local CA
mkcert -install

# Generate certificates
mkdir certs
mkcert -cert-file certs/local-cert.pem -key-file certs/local-key.pem "localhost" "*.localhost" "api.localhost"
```

3. Create necessary .env files:
   - Copy `server/.env.example` to `server/.env`
   - Create `app/.env` for frontend configuration
   - Create root `.env` for Traefik configuration

4. Start the development environment:
```bash
docker compose up -d
```

## Accessing the Application

- Frontend: https://localhost
- API: https://api.localhost
- API Documentation: https://api.localhost/swagger/index.html

## Development

### Frontend Development
The frontend uses Vite with React, TypeScript, and PWA support. Source code is in the `/app` directory.

### Backend Development
The backend is written in Go with fiber framework. Source code is in the `/server` directory.

### Environment Variables

- Root `.env`: Global and Traefik configuration
- `app/.env`: Frontend-specific variables
- `server/.env`: Backend and database configuration

## Production Deployment

For production deployment:

1. Update the domain in root `.env`:
```env
DOMAIN=your-domain.com
ACME_EMAIL=your-email@example.com
```

2. Ensure your domain's DNS points to your server

3. Deploy using docker compose:
```bash
docker compose up -d
```

The application will automatically obtain SSL certificates through Let's Encrypt.
