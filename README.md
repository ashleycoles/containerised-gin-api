# Containerised Gin API

A simple API built using Go with Gin and an SQL Database, containerised using docker.

Run `docker compose up --build` to get everything running. This does the following:
1. Builds the binary
2. Runs air to rebuild the binary when you make code changes
3. Sets up MariaDB and automatically creates the required database and it's structure