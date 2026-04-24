# 🛡️ JWT Authentication API (Go + PostgreSQL)

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin Framework](https://img.shields.io/badge/Gin-008080?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white)

API de autenticação robusta e escalável desenvolvida em **Go**, utilizando o framework **Gin** e o ORM **GORM**. Este projeto foca em segurança de backend, implementando o fluxo completo de tokens JWT.

---

## 🚀 Tecnologias Utilizadas

* **Linguagem:** Go (Golang)
* **Web Framework:** Gin Gonic
* **Banco de Dados:** PostgreSQL
* **ORM:** GORM
* **Autenticação:** JWT (JSON Web Tokens)
* **Configuração:** Godotenv para variáveis de ambiente

---

## 📌 Funcionalidades Principal

* ✅ **Registro de Usuários:** Cadastro seguro com criptografia de senha.
* ✅ **Login & JWT:** Geração de tokens de acesso para sessões seguras.
* ✅ **Middleware de Autenticação:** Proteção de rotas privadas contra acessos não autorizados.
* ✅ **Validação de Dados:** Uso de DTOs para garantir a integridade dos inputs.
* ✅ **Arquitetura Limpa:** Organização de pastas seguindo padrões modernos de desenvolvimento Go.

---

## 📂 Estrutura do Projeto

```bash
jwt-authentication-go/
├── cmd/
│   └── api/                # Ponto de entrada (main.go)
├── internal/
│   ├── app/
│   │   ├── handlers/       # Controladores das rotas
│   │   └── dto/            # Data Transfer Objects
│   ├── domain/
│   │   └── models/         # Modelos do Banco de Dados
│   └── database/           # Configuração e conexão DB
├── middleware/             # Interceptor de segurança JWT
├── .env.example            # Template de variáveis de ambiente
├── go.mod                  # Gerenciador de módulos Go
└── README.md
