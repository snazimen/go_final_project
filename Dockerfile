FROM golang:1.23.0

WORKDIR /app

COPY . .

RUN go mod download

RUN  go build -o /snazimen ./cmd/main.go

EXPOSE 7540

CMD ["/snazimen"]