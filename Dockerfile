FROM golang:1.12 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

FROM scratch
COPY --from=builder /app/hue-remote-v2 /app/
EXPOSE 8080
ENTRYPOINT ["/app/hue-remote-v2"]
