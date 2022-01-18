FROM golang:1.17-alpine3.13 as builder
WORKDIR /go/src/repotmicro
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o build/reportmicro

FROM scratch
COPY --from=builder /go/src/reportmicro/build/reportmicro /usr/bin/reportmicro
ENTRYPOINT [ "/usr/bin/reportmicro" ]