FROM golang:1.19

WORKDIR /app
COPY . .

RUN go mod init auth-service
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth

EXPOSE 8080

CMD ["/auth"]