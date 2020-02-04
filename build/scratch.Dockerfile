############################
FROM golang:alpine AS builder
RUN apk update && apk add git && apk add ca-certificates
RUN adduser -D -g '' appuser
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bill18go
############################
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app /app
WORKDIR /app

#USER appuser
EXPOSE 8090

ENTRYPOINT ["./bill18go"]
CMD ["-conf=conf.yaml"]

# компиляция образа
#docker build -t puzanovma/bill18go-scratch -f scratch.Dockerfile .
#docker run --rm -it -e "PORT=5000" -v $$(pwd)/logs:/app/logs:rw -p 5000:5000 puzanovma/bill18go-scratch