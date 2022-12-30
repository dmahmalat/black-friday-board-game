FROM golang:alpine AS builder
COPY . /src
WORKDIR /src
RUN go build -o /bin/bfbg ./cmd/bfbg/main.go

FROM scratch
COPY --from=builder /bin/bfbg /bin/bfbg
COPY --from=builder /src/config/config.yaml /bin/config.yaml
WORKDIR /bin
ENTRYPOINT [ "bfbg" ]