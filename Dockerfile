# Dockerfile

# Etapa de build
FROM golang:1.22.4 AS builder

WORKDIR /app

# Cachear as dependências do módulo
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código-fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/api/main.go

# Etapa final
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /app/api /api

# Porta em que a aplicação vai rodar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["/api"]