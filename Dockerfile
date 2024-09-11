FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build .


FROM scratch

WORKDIR /app

COPY --from=builder /app/random_error .

CMD ["/app/random_error"]

