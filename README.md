# Tech Challenge - Sistema de Autoatendimento para Lanchonete

Este projeto é parte do **Tech Challenge - Fase 01**, e tem como objetivo desenvolver um sistema de controle de pedidos para uma lanchonete em expansão, focado em autoatendimento, gestão de pedidos e controle administrativo.

## 📌 Objetivo

Criar um sistema backend monolítico que permita:

- Realizar e acompanhar pedidos de forma eficiente
- Facilitar o autoatendimento por parte dos clientes
- Gerenciar produtos e categorias via painel administrativo
- Integrar pagamento via QRCode (Mercado Pago)
- Controlar o andamento dos pedidos e tempo de preparo

## 🚀 Tecnologias Utilizadas

- **Go (Golang)**
- **Arquitetura Hexagonal**
- **Swagger (Documentação das APIs)**
- **Docker / Docker Compose**
- **Banco de Dados** (à escolha do desenvolvedor)

## 🔧 Funcionalidades Principais

### APIs

- Cadastro e identificação de clientes (via CPF)
- CRUD de produtos
- Listagem de produtos por categoria
- Envio de pedidos para a fila (checkout fake)
- Listagem e acompanhamento de pedidos

### Outras funcionalidades

- Integração com Mercado Pago (QR Code para pagamento)
- Painel de acompanhamento de pedidos para cliente e cozinha
- Painel administrativo para gerenciar produtos e categorias

## 🐳 Como rodar o projeto localmente

1. **Clone o repositório**
   ```bash
   git clone https://github.com/seuusuario/tech-challenge.git
   cd tech-challenge
