#!/bin/bash

# Run database migration scripts

echo "Running database migration scripts..."
go run ./cmd/migrate.go
echo "Database migration complete."
