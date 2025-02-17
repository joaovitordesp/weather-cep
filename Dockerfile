FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o /weather-cep ./cmd/api

EXPOSE 8080

CMD ["/weather-cep"] 