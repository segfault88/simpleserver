# Simple Go server

HTTPS server that reads a cert from the current directory and returns "hello, world!" and 200 to any requests.

To make a self signed cert:

```
/usr/local/opt/openssl/bin/openssl req -x509 -nodes -newkey rsa:2048 -keyout server.rsa.key -out server.rsa.crt -days 3650
ln -sf server.rsa.key server.key
ln -sf server.rsa.crt server.crt
```

Build a go binary (for linux) with:

```
GOOS=linux GOARCH=amd64 go build simpleserver.go
```
