# Tech Challenge - Sistema de Autoatendimento para Lanchonete

Este projeto é parte do **Tech Challenge - Fase 01**, e tem como objetivo desenvolver um sistema de controle de pedidos para uma lanchonete.

## 📋 Índice

- [Visão Geral](#-visão-geral)
- [Tecnologias](#-tecnologias)
- [Estrutura do Projeto](#-estrutura-do-projeto)
- [Configuração](#️-configuração)
- [Execução](#-execução)
- [Documentação da API](#-documentação-da-api)
- [Testes](#-testes)
- [Documentação DDD](#-documentação-do-sistema-ddd-com-event-storming)

## 🎯 Visão Geral

Este projeto implementa um sistema completo de autoatendimento para lanchonetes que inclui:

- **Gestão de Clientes**: Cadastro e consulta de clientes por CPF
- **Catálogo de Produtos**: CRUD completo com categorização (lanches, acompanhamentos, bebidas, sobremesas)
- **Sistema de Pedidos**: Criação, acompanhamento e atualização de status dos pedidos
- **Integração de Pagamento**: Integração com MercadoPago para processamento de pagamentos
- **Painel Administrativo**: Monitoramento de pedidos em andamento

## 🛠 Tecnologias

- **Go** - Linguagem de programação
- **Gin** - Framework web
- **MySQL 8.0** - Banco de dados
- **Docker** - Containerização
- **Swagger** - Documentação da API
- **MercadoPago SDK** - Processamento de pagamentos

## 📁 Estrutura do Projeto

O projeto segue a **arquitetura hexagonal** (ports and adapters) organizando o código em camadas bem definidas:

```
├── cmd/
│   └── server/              # Ponto de entrada da aplicação
│       └── main.go
├── internal/
│   ├── domain/              # Regras de negócio e entidades
│   │   ├── entities/        # Entidades do domínio
│   │   └── ports/           # Interfaces (contratos)
│   │       ├── input/       # Portas de entrada (services)
│   │       │   └── services/
│   │       └── output/      # Portas de saída (repositories)
│   │           └── repositories/
│   ├── application/         # Casos de uso e serviços
│   │   ├── services/        # Implementação dos serviços
│   │   └── dto/             # Data Transfer Objects
│   ├── adapters/           # Adaptadores (HTTP, Repository)
│   │   ├── http/
│   │   │   ├── handlers/    # Controllers HTTP
│   │   │   └── router/      # Configuração de rotas
│   │   └── repositories/
│   │       └── persistence/ # Implementação dos repositórios
│   └── infrastructure/     # Configurações e conexões externas
│       ├── database/
│       │   └── mysql/
│       └── mercadopago/
├── docs/                   # Documentação Swagger
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
```

## ⚙️ Configuração

### Pré-requisitos

- **Docker** - Para execução do projeto
- **Git** - Para clonar o repositório

### Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
# Banco de Dados
DB_USER=root
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=3306
DB_NAME=fast_food_db

# MercadoPago
ACCESSTOKEN=seu_token_mercadopago_aqui

# Servidor
PORT=8080
```

### Instalação

1. Clone o repositório:
```bash
git clone https://github.com/samuellalvs/soat_tech_challenge_fast_food.git
cd soat_tech_challenge_fast_food
```

> **Nota**: Não é necessário instalar dependências ou gerar documentação Swagger manualmente. O Docker se encarrega de tudo automaticamente durante o build.

## 🚀 Execução

1. **Construa e inicie os containers**:
```bash
docker compose up --build
```

2. **A aplicação estará disponível em**: `http://localhost:8080`

3. **Para parar os containers**:
```bash
docker compose down
```

### Acesso rápido aos serviços

- **API**: http://localhost:8080
- **Swagger**: http://localhost:8080/swagger/index.html
- **Health Check**: http://localhost:8080/health

## 📚 Documentação da API

### Swagger

A documentação completa da API está disponível através do **Swagger**:

1. **Com a aplicação rodando**, acesse:
   - **URL**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Endpoints Principais

#### 👤 Clientes
- `POST /api/v1/customers` - Cadastrar novo cliente
- `GET /api/v1/customers/{cpf}` - Buscar cliente pelo CPF

#### 🍔 Produtos
- `POST /api/v1/products` - Criar novo produto
- `GET /api/v1/products` - Listar todos os produtos
- `GET /api/v1/products/category/{category}` - Listar por categoria
- `PUT /api/v1/products` - Atualizar produto
- `DELETE /api/v1/products/{id}` - Remover produto

#### 🧾 Pedidos
- `POST /api/v1/orders` - Criar novo pedido
- `GET /api/v1/orders/{id}` - Buscar pedido por ID
- `PATCH /api/v1/orders/{id}/status` - Atualizar status do pedido

#### 💳 Pagamentos
- `POST /api/v1/checkout` - Processar pagamento

#### 📊 Administração
- `GET /api/v1/admin/orders/active` - Listar pedidos em andamento

### Exemplo de Uso

```bash
# Criar um cliente
curl -X POST http://localhost:8080/api/v1/customers \
  -H "Content-Type: application/json" \
  -d '{"first_name":"João","last_name":"Silva","email":"joao@email.com","cpf":"123.456.789-00"}'

# Criar um produto
curl -X POST http://localhost:8080/api/v1/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Big Burger","description":"Hambúrguer artesanal","price":"25.90","category":"burger"}'

# Criar um pedido
curl -X POST http://localhost:8080/api/v1/orders \
  -H "Content-Type: application/json" \
  -d '{"customer_id":1,"cpf":"123.456.789-00","status":"received","items":[{"product_id":1,"quantity":2,"price":25.90}]}'
```

## 🧪 Testes

### Executar Testes

```bash
docker run --rm -v $(pwd):/app -w /app golang:1.24-alpine go test ./...
```

## 📚 Documentação do sistema (DDD) com Event Storming

As versões da documentação do DDD estão disponíveis através dos links do Miro:

- [**DDD - Versão 1** - Iniciação](https://miro.com/app/board/uXjVIDaCt8I=/)
- [**DDD - Versão 2** - Evolução](https://miro.com/app/board/uXjVI26PK8k=/)
- [**DDD - Versão 3** - Final](https://miro.com/app/board/uXjVIzM5S5Q=/)
