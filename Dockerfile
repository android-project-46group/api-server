# Build stage
FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
RUN go build -o populate db/data/populate/main.go 

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/populate .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY migrations ./migrations
COPY ./db/data/blogs ./db/data/blogs
COPY ./db/data/info ./db/data/info

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
