FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build .


FROM scratch
ARG USERNAME=1234
USER $USERNAME
WORKDIR /app

COPY --from=builder --chown=$USERNAME /app/random_error .

CMD ["/app/random_error"]

