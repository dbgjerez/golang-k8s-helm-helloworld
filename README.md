# Intro
Application used to learn golang or anothers frameworks and tools like Kubernetes, OCP, etc

# Build
## Golang
```bash
go build app.go
```

## Docker
To build a docker image is mandatory to have installed Docker or use something like buildah. I have used my docker's user (b0rr3g0), you should change $USER for your own user.
```bash
echo "docker build -t $USER/golang-hello-world ."
```

# Run
## Golang
```bash
go run app.go
```

## Docker
It can be used my image (b0rr3g0/hello-world) or your own image. 
```bash
docker run -p 8080:8080 b0rr3g0/golang-hello-world
```

## Helm
[](https://github.com/dbgjerez/ms-helm-chart)

# Endpoints
|Endpoint|Comment|
|--|--|
|/api/v1/health|Healtcheck|
|/api/v1/grettings|Hello world message|
|not found|404 error code|
