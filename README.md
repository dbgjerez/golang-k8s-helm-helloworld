# Intro
Application used to learn golang or anothers frameworks and tools like Kubernetes, OCP, etc

# Build
## Golang

With Golang: 

```bash
go build app.go
```

## Image

I use podman to build the imagen, you can use podman, docker or whatever you want. I have used my docker's user (b0rr3g0), you should change $USER for your own user.

```bash
podman build -t $USER/golang-hello-world .
```

Another option is to use the Makefile, previus change the parameters:

```bash
make build
```

# Run
## Golang
```bash
go run app.go
```

## Docker
It can be used my image (b0rr3g0/golang-hello-world) or your own image. 
```bash
docker run -p 8080:8080 b0rr3g0/golang-hello-world
```

## Helm/Kustomize
[](https://github.com/dbgjerez/kustomize-vs-helm)

# Endpoints
|Endpoint|Comment|
|--|--|
|/api/v1/health|Healtcheck|
|/api/v1/greetings|Hello world message|
|/metrics|Prometheus metrics exporter|
|not found|404 error code|
