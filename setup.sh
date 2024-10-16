./.github/scripts/init-temp-keys.sh
chmod 777 keys/*.pem
# Note this might be `podman compose` on some systems
docker compose -f docker-compose.yaml up -d

# Use wait-for-it to wait until Keycloak is ready (on port 8888)
./wait-for-it.sh localhost:8888 --timeout=60 --strict -- echo "Keycloak is up"

# Once Keycloak is up, run the Go commands to initialize it (first go)
go run ./service provision keycloak
go run ./service start