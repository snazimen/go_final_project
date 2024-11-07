FROM golang:1.21.0

WORKDIR /app

COPY . .

RUN go mod download

RUN  go build -o /snazimen ./cmd/main.go

CMD ["/snazimen"]