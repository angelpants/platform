# Use wait-for-it to wait until Keycloak is ready (on port 8888)
./wait-for-it.sh localhost:8888 --timeout=60 --strict -- echo "Keycloak is up"

# Once Keycloak is up, run the Go commands to initialize it (first go)
go run ./service start