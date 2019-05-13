# go-echo-server

![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/dotlou/go-echo-server.svg?style=for-the-badge) ![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/dotlou/go-echo-server.svg?style=for-the-badge) ![Golang v1.12](https://img.shields.io/badge/Golang-v1.12-green.svg?style=for-the-badge)

Simple echo HTTP server with a fake auth endpoint

## Endpoints

### `/`

Useful for echoing back what it got, such as headers, method, body.

Supported methods:

- OPTIONS
- GET
- POST


### `/fakeAuth`

Useful for using as a test nginx authentication sub-request. Will pass if given an `Authorization` header of `Bearer valid-key`, will fail otherwise

Supported methods:

- OPTIONS
- GET

## k8s

Some kubernetes resource files are provided so that this can run behind an nginx-ingress controller, with the auth-url configured for `/` to use `/fakeAuth` as the authentication path.