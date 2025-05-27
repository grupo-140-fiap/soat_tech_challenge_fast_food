# Build stage
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
# CGO_ENABLED=0: Desativa o suporte ao CGO, resultando em um binário estático que não depende de bibliotecas C externas
# GOOS=linux: Especifica o sistema operacional alvo como Linux, garantindo compatibilidade com containers
# -o main: Define o nome do arquivo de saída como 'main'
# ./cmd/server: Caminho do pacote principal que contém a função main()
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# Final stage
FROM alpine:latest

# Instalar certificados CA e curl para healthcheck
RUN apk --no-cache add ca-certificates curl

WORKDIR /app

# Criar usuário não-root
RUN adduser -D -g '' appuser

# Copiar binário do builder
COPY --from=builder /app/main .

# Mudar para usuário não-root
USER appuser

# Expor porta
EXPOSE 8080

# Healthcheck
HEALTHCHECK --interval=30s --timeout=3s \
  CMD curl -f http://localhost:8080/health || exit 1

# Comando para executar a aplicação
CMD ["./main"]
