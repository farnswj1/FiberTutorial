# Fiber Tutorial
This is a simple web app powered with Fiber.

## Setup
The project uses the following:
- Go
- Fiber
- Redis
- Docker
- Docker Compose

For additional information on project specifications, see the `go.mod` file in `app`.

### Setting up the API
In the `app` directory, create a `.env` file that contains the following environment variables:
```
CORS_ALLOWED_ORIGINS=http://localhost,http://127.0.0.1
REDIS_URL=redis://redis:6379/0
```

## Building
The project uses Docker. Ensure Docker and Docker Compose are installed before continuing.

To build, run `docker compose build`

## Running
To run the web API, run `docker compose up -d`, then go to http://localhost using your web browser.
