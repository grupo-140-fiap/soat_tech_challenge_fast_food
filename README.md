# Tech Challenge - Sistema de Autoatendimento para Lanchonete

Este projeto é parte do **Tech Challenge - Fase 01**, e tem como objetivo desenvolver um sistema de controle de pedidos para uma lanchonete em expansão, focado em autoatendimento, gestão de pedidos e controle administrativo.

---

## ✅ Checklist de Endpoints da API

### 👤 Customers
- [x] `POST /customers` — Cadastrar novo cliente
- [x] `GET /customers/{cpf}` — Buscar cliente pelo CPF

#### Exemple
```bash
curl -i -X POST http://localhost:8080/api/v1/customers -d '{"first_name":"Test1","last_name":"Test2","email":"test@test.com","cpf":"xxx.xxx.xxx"}'

curl -i -X GET http://localhost:8080/api/v1/customers/xxx.xxx.xxx-xx
```

### 🍔 Products
- [ ] `POST /products` — Criar novo produto
- [ ] `PUT /products`  — Atualizar produto existente
- [ ] `DELETE /products/{id}` — Remover produto
- [ ] `GET /products` — Listar todos os produtos
- [ ] `GET /products?category={category}` — Listar produtos por categoria (`burger`, `side`, `drink`, `dessert`)

#### Exemple
```bash
curl -X POST http://localhost:8080/api/v1/products -H "Content-Type: application/json" -d '{"name":"Pizza","description":"queijo","price":"40","category":"burger"}'

curl -X GET http://localhost:8080/api/v1/producs/12

curl -i -XPUT http://localhost:8080/api/v1/products -d '{"id":1, "name":"Pizza-u","description":"queijo","price":"40","category":"burger"}'

curl -X DELETE http://localhost:8080/api/v1/products/1

curl -X GET http://localhost:8080/api/v1/products/category/burger
```

### 🧾 Orders (Checkout)
- [ ] `POST /checkout` — Criar novo pedido (enviar para fila, simular pagamento)
- [ ] `GET /orders` — Listar todos os pedidos
- [ ] `GET /orders/{id}` — Buscar detalhes do pedido por ID
- [ ] `PATCH /orders/{id}/status` — Atualizar status do pedido (`received`, `preparing`, `ready`, `completed`)

### 📊 Admin / Monitoramento
- [ ] `GET /admin/orders/active` — Listar pedidos em andamento

### 📦 Categories (Opcional)
- [ ] `GET /categories` — Listar categorias de produtos (`burger`, `side`, `drink`, `dessert`)
