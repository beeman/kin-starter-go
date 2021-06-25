FROM golang:1.16-alpine

WORKDIR /usr/local/bin

COPY . .

RUN go install

CMD ["go", "run", "main.go"]
