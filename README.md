# Interstate

HTTP backend proxy for cross-state dependendies in Terraform / OpenTofu

# Usage

Set the following environment variables:

- `S3_BUCKET`: The name of your S3 bucket where backend states will be stored
- `PORT` (optional): The port on which the server will run (default: 8080)

Run for local development

```
go mod download
go build -o interstate cmd/server/main.go
./interstate
```

The service will start on http://localhost:8080 (or the port specified in the PORT environment variable).

# API Reference

## Backends

List all backends

- `GET /api/backends`
- Returns a list of all backends

Create a new backend

- POST /api/backends
- Creates a new backend
- Request body: `{ "name": "backend-name" }`

Get a specific backend

- GET `/api/backends/{backendId}`
- Retrieves details of a specific backend

Update a backend

- `PUT /api/backends/{backendId}`
- Updates an existing backend
- Request body: { "name": "new-backend-name" }

Delete a backend

- `DELETE /api/backends/{backendId}`
- Deletes a specific backend

## Backend State

Get backend state

- `GET /api/backends/{backendId}/state`
- Retrieves the current state (terraform.tfstate) of a specific backend

Update backend state

- `POST /api/backends/{backendId}/state`
- Updates the state of a specific backend
- Request body: `Terraform state JSON`

Reset backend state

- `DELETE /api/backends/{backendId}/state`
- Resets or clears the state of a specific backend
