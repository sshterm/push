# Push Notifications

SSH Term APP Apple Push Notifications

## Build
```
docker run --rm -v "$PWD":/app -w /app golang:latest go build -trimpath -ldflags="-s -w" -tags 'netcgo' -o lib/server server.go && upx lib/server
```