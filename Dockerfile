FROM golang:1.22.1 as build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/bin/server .
COPY --from=build /app/.env .

EXPOSE 9000

ENTRYPOINT ["./server"]