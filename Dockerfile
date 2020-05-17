FROM golang:1.13
WORKDIR /app
ADD . /app
CMD ["/usr/local/go/bin/go", "run", "/app/src/main.go"]
EXPOSE 8080