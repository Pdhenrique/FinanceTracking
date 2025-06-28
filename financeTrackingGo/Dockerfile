FROM golang:1.24.3

WORKDIR /app

# baixar dependências
COPY go.mod go.sum ./
RUN go mod download

# copiar código e compilar
COPY . .
RUN go build -o main ./cmd/app

# **dar permissão de execução ao binário**
RUN chmod +x /app/main

# iniciar o app
CMD ["./main"]
