FROM golang:1.22-bookworm

WORKDIR /app

COPY . . 

RUN go mod init teste

RUN go build -o math 

CMD ["./math"]