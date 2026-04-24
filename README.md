# 🛡️ JWT Authentication API (Go + PostgreSQL)

API de autenticação de alto desempenho construída com **Golang**, focada em segurança, escalabilidade e organização seguindo as melhores práticas de backend.

---

## 🚀 Tecnologias

| Tecnologia | Descrição |
| :--- | :--- |
| **Go (Golang)** | Linguagem principal de alta performance |
| **Gin Gonic** | Web framework rápido e produtivo |
| **PostgreSQL** | Banco de dados relacional robusto |
| **GORM** | ORM para manipulação eficiente de dados |
| **JWT (v5)** | Autenticação baseada em tokens seguros |
| **Godotenv** | Gestão de variáveis de ambiente |

---

## 📌 Funcionalidades

* ✅ **Auth:** Registro de usuários e Login com geração de JWT.
* ✅ **Segurança:** Middleware para proteção de rotas privadas.
* ✅ **Dados:** Validação rigorosa com DTOs (Data Transfer Objects).
* ✅ **Persistência:** Integração completa com PostgreSQL via GORM.
* ✅ **Arquitetura:** Estrutura de pastas organizada para fácil manutenção.

---

## 📂 Estrutura de Pastas

```bash
jwt-authentication-go/
├── cmd/api/            # Ponto de entrada da aplicação
├── internal/
│   ├── app/            # Handlers (Controllers) e DTOs
│   ├── domain/         # Modelos de dados (Entities)
│   └── database/       # Configuração da conexão com DB
├── middleware/         # Lógica de proteção JWT
└── .env                # Variáveis sensíveis
