FROM golang:1.12 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go test -covermode=atomic -coverpkg=all ./...

RUN CGO_ENABLED=0 GOOS=linux go build

FROM scratch

ARG PORT=8080

COPY --from=builder /app/hue-remote-v2 /app/
EXPOSE $PORT
ENTRYPOINT ["/app/hue-remote-v2"]
