# GOROAPI 

Rick and Morty album using microservices in Go and Node.js

![img](.screenshots/goroapi.webp)


# Install

## Using Docker

### Docker compose
```sh
docker compose up -d
```

### Bash script
```sh
bash scripts/build.sh
```

## Manual

### Client
```sh
cd client
```
```sh
npm i
```
```sh
npm run dev
```
or
```sh
npm run build
```
```sh
node dist index.js
```

### Server
```sh
go run cmd/web/main.go
```
or
```sh
go build -o goroapi cmd/web/main.go
```