# Tech Challenge - Sistema de Autoatendimento para Lanchonete

Este projeto é parte do **Tech Challenge - Fase 01**, e tem como objetivo desenvolver um sistema de controle de pedidos para uma lanchonete em expansão, focado em autoatendimento, gestão de pedidos e controle administrativo.

---

## ✅ Checklist de Endpoints da API

### 👤 Customers
- [ ] `POST /customers` — Cadastrar novo cliente
- [ ] `GET /customers/{cpf}` — Buscar cliente pelo CPF

### 🍔 Products
- [ ] `POST /products` — Criar novo produto
- [ ] `PUT /products/{id}` — Atualizar produto existente
- [ ] `DELETE /products/{id}` — Remover produto
- [ ] `GET /products` — Listar todos os produtos
- [ ] `GET /products?category={category}` — Listar produtos por categoria (`burger`, `side`, `drink`, `dessert`)

### 🧾 Orders (Checkout)
- [ ] `POST /checkout` — Criar novo pedido (enviar para fila, simular pagamento)
- [ ] `GET /orders` — Listar todos os pedidos
- [ ] `GET /orders/{id}` — Buscar detalhes do pedido por ID
- [ ] `PATCH /orders/{id}/status` — Atualizar status do pedido (`received`, `preparing`, `ready`, `completed`)

### 📊 Admin / Monitoramento
- [ ] `GET /admin/orders/active` — Listar pedidos em andamento
- [ ] `GET /admin/orders/wait-time` — Consultar tempo médio de espera dos pedidos

### 📦 Categories (Opcional)
- [ ] `GET /categories` — Listar categorias de produtos (`burger`, `side`, `drink`, `dessert`)
