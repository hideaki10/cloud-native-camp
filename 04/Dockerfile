FROM golang:1.16 as builder

WORKDIR /app

COPY . ./

RUN go build -o ./httpServer-run ./httpServer.go

FROM gcr.io/distroless/base

EXPOSE 8080

COPY --from=builder app/httpServer-run /.

CMD ["./httpServer-run"]