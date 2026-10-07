# Wallet Service

Serviço responsável pelo gerenciamento de carteiras financeiras para um sistema de apostas.

---

# Objetivo

O Wallet Service é responsável por:

- Criar carteiras de jogadores
- Consultar saldo de carteiras
- Registrar transações financeiras
- Registrar histórico financeiro (Ledger)
- Processar apostas (BET)
- Processar ganhos (WIN)
- Registrar movimentações de saldo

Todos os valores monetários são armazenados em centavos utilizando o tipo `BIGINT`, evitando problemas de precisão causados por números de ponto flutuante (`float`).

## Exemplos

| Valor Monetário | Valor Armazenado |
|---------------|---------------|
| R$ 10,00 | 1000 |
| R$ 25,50 | 2550 |
| R$ 100,99 | 10099 |

---

# Tecnologias Utilizadas

- Go
- PostgreSQL
- Docker
- Docker Compose
- UUID
- REST API
- Git
- VS Code

---

# Estrutura do Projeto

```text
wallet-service/

├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── database/
│   │   └── postgres.go
│   │
│   ├── dto/
│   │   ├── create_wallet_request.go
│   │   ├── transaction_request.go
│   │   └── wallet_response.go
│   │
│   ├── entity/
│   │   ├── wallet.go
│   │   ├── ledger.go
│   │   └── wager_transaction.go
│   │
│   ├── repository/
│   │   ├── wallet_repository.go
│   │   └── postgres_wallet_repository.go
│   │
│   ├── service/
│   │   └── wallet_service.go
│   │
│   ├── handler/
│   │   └── wallet_handler.go
│   │
│   └── routes/
│       └── routes.go
│
├── migrations/
│   ├── 001_create_wallets.sql
│   ├── 002_create_wallet_ledger.sql
│   └── 003_create_wager_transactions.sql
│
├── tests/
│
├── .env
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

# Requisitos

Antes de executar o projeto, é necessário instalar:

## Docker

Verificar instalação:

```bash
docker --version
```

---

## Docker Compose

Verificar instalação:

```bash
docker compose version
```

---

## Go

Verificar instalação:

```bash
go version
```

Exemplo:

```bash
go version go1.27.1 windows/amd64
```

---

## Git

Verificar instalação:

```bash
git --version
```

---

# Configuração do Banco de Dados

Banco utilizado:

```text
PostgreSQL 16
```

Configurações:

| Configuração | Valor |
|-------------|---------|
| Host | localhost |
| Porta | 5432 |
| Usuário | postgres |
| Senha | postgres |
| Banco | wallet |

---

# Arquivo docker-compose.yml

```yaml
version: "3.9"

services:

  postgres:
    image: postgres:16

    container_name: wallet-postgres

    restart: always

    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: wallet

    ports:
      - "5432:5432"

    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

---

# Subindo o Banco de Dados

Na raiz do projeto execute:

```bash
docker compose up -d
```

Verifique se o container foi iniciado:

```bash
docker ps
```

Resultado esperado:

```text
wallet-postgres
```

---

# Acessando o PostgreSQL

Para acessar o banco:

```bash
docker exec -it wallet-postgres psql -U postgres -d wallet
```

Resultado esperado:

```text
wallet=#
```

---

# Executando as Migrations

As migrations estão na pasta:

```text
migrations/
```

Executar na seguinte ordem:

```text
001_create_wallets.sql
002_create_wallet_ledger.sql
003_create_wager_transactions.sql
```

---

# Estrutura das Tabelas

## wallets

Tabela responsável por armazenar o estado atual da carteira.

Campos:

```text
id
player_id
balance
currency
version
created_at
updated_at
```

---

## wallet_ledger

Tabela responsável por armazenar o histórico financeiro das operações.

Campos:

```text
id
wallet_id
transaction_id
transaction_type
amount
balance_before
balance_after
created_at
```

---

## wager_transactions

Tabela responsável por armazenar as transações recebidas do provedor.

Campos:

```text
id
wallet_id
provider_transaction_id
transaction_type
amount
status
created_at
updated_at
```

---

# Executando a Aplicação

Na raiz do projeto execute:

```bash
go run cmd/api/main.go
```

Se tudo estiver correto:

```text
server running on :8080
```

---

# Endpoints

## Criar Carteira

### Método

```http
POST /wallets
```

### Exemplo de Requisição

```json
{
  "playerId": "11111111-1111-1111-1111-111111111111",
  "currency": "BRL"
}
```

### Exemplo de Resposta

```json
{
  "id": "f1e47a2d-80f8-4cb5-bf77-5af5d6cc4a9e",
  "playerId": "11111111-1111-1111-1111-111111111111",
  "balance": 0,
  "currency": "BRL",
  "version": 1
}
```

---

# Fluxo de BET

Saldo inicial:

```text
10000
```

Representando:

```text
R$ 100,00
```

Aposta:

```text
2500
```

Representando:

```text
R$ 25,00
```

Saldo final:

```text
7500
```

Representando:

```text
R$ 75,00
```

---

# Fluxo de WIN

Saldo inicial:

```text
7500
```

Representando:

```text
R$ 75,00
```

Ganho:

```text
3000
```

Representando:

```text
R$ 30,00
```

Saldo final:

```text
10500
```

Representando:

```text
R$ 105,00
```

---

# Testando no Postman

## Criar Wallet

Método:

```http
POST
```

URL:

```http
http://localhost:8080/wallets
```

Body:

```json
{
  "playerId": "11111111-1111-1111-1111-111111111111",
  "currency": "BRL"
}
```

---

# Comandos Úteis

## Baixar dependências

```bash
go mod tidy
```

---

## Compilar o projeto

```bash
go build ./...
```

---

## Executar aplicação

```bash
go run cmd/api/main.go
```

---

## Verificar containers

```bash
docker ps
```

---

## Parar containers

```bash
docker compose down
```

---

# Roadmap

Próximas funcionalidades que serão implementadas:

- GET Wallet
- Atualização de saldo
- BET
- WIN
- Registro automático no Ledger
- Tratamento de erros
- Middleware
- Logs
- Testes unitários
- Dockerfile da API
- Integração com SQS
- Outbox Pattern
- Inbox Pattern
- Keycloak
- OAuth2

---

# Autor

Lucas de Souza Resende

Projeto desenvolvido para estudo de Go, PostgreSQL, Docker e arquitetura de microsserviços financeiros.

---

# Licença

Projeto criado para fins educacionais e processo seletivo técnico.