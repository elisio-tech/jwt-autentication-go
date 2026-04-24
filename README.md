JWT Authentication API (Go + PostgreSQL)

API de autenticação construída com Golang, Gin, JWT e PostgreSQL, com foco em segurança e boas práticas de backend.

🚀 Tecnologias
Go (Golang)
Gin Web Framework
PostgreSQL
JWT (golang-jwt/jwt/v5)
GORM 
godotenv

📌 Funcionalidades
Registro de usuário
Login com JWT
Middleware de autenticação
Proteção de rotas privadas
Validação de token
Conexão com banco PostgreSQL

jwt-authentication-go/
│
├── middleware/
│   └── jwt.go              # Middleware de autenticação JWT
├── cmd/
│   └── api/
│       └── main.go         # Ponto de entrada da aplicação
├── internal/
│   ├── app/
│   │   ├── handlers/       # Handlers HTTP
│   │   └── dto/            # DTOs (Data Transfer Objects)
│   ├── domain/
│   │   └── models/
│       └── user.go         # Model User
│   └── database/
│       └── db.go           # Conexão com banco de dados
├── .env                    # Variáveis de ambiente
├── go.mod                  # Dependências Go
└── README.md

🚀 Func
✅ Registro de usuários ( POST /register)
✅ Login ( POST /login) com JWT
✅ Perfil protegido ( GET /profile) - requer JWT
✅ Token de atualização ( POST /refresh)
✅ Middleware JWT para rotas protegidas
✅ Validação de dados com DTOs
✅ PostgreSQL com GORM
✅ Variáveis ​​de ambiente (.env)

Instalacao
# Clonar o repositório
git clone github.com/elisio-tech/jwt-autentication-go/
cd jwt-authentication-go

# Instalar dependências
go mod tidy

# Copiar .env.example para .env
cp .env.example .env

ENV
# Servidor
PORT=8080
APP_NAME="JWT Auth API"

# Banco de dados
DB_HOST=localhost
DB_PORT=5432
DB_NAME=jwt_auth
DB_USER=postgres
DB_PASSWORD=your_password

criar banco de dado
-- Criar banco
CREATE DATABASE jwt_auth;
-- Tabela users (criada automaticamente pelo GORM)

Executar
# Development
go run cmd/api/main.go

# Production
go build -o bin/api cmd/api/main.go
./bin/api

ENDPOINT
# 1. Registrar
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@test.com","username": "Jesus", "password":"12345678"}'

# 2. Login
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@test.com","password":"12345678"}'

# 3. Perfil (com JWT)
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer SEU_JWT_TOKEN"


Feito com ❤️ em Go 🚀




