# AGENTS.md

Guia para agentes de IA (Claude Code, Codex, Copilot etc.) trabalhando neste repositório.

## Visão geral

**Ciclobits PCP** — sistema de Planejamento e Controle da Produção para uma fábrica fictícia de bicicletas. Projeto acadêmico (Análise e Projeto de Sistemas — UNISC, 2026). Estágio inicial: boa parte dos arquivos ainda são esqueletos.

- **Backend:** Go + Gin + GORM, banco SQLite (driver puro-Go `github.com/glebarez/sqlite`, sem CGO).
- **Frontend:** Vue + JavaScript (pasta `frontend/` ainda vazia).
- **Contrato da API:** `docs/openapi.yaml` — **fonte da verdade** para rotas, schemas e regras de negócio.

Idioma: código de domínio, comentários, commits e documentação em **português**.

## Estrutura

```text
backend/
├── cmd/api/main.go              # entrypoint: carrega .env, abre DB, monta App, sobe Gin
├── internal/
│   ├── bootstrap/
│   │   ├── app.go               # struct App com todos os módulos (injeção de dependência manual)
│   │   └── modules/             # um arquivo por módulo; NewX(db) monta controller/use cases/repo
│   ├── routes/routes.go         # rotas sob /api/v1
│   ├── db/
│   │   ├── database.go          # NewDb() a partir de DB_DSN
│   │   └── migrations/up.go     # programa de migração (package main, ainda vazio)
│   └── app/
│       ├── controller/          # handlers HTTP (Gin)
│       ├── use_case/            # regras de negócio (package usecase)
│       ├── entity/              # entidades de domínio
│       └── persistence/
│           ├── model/           # modelos GORM
│           ├── mapper/          # conversão model <-> entity
│           └── repository/      # acesso a dados
└── data/seed/*.json             # carga inicial do banco
docs/openapi.yaml                # contrato da API
```

A modelagem do banco ainda **não está definida**: não documente tabelas/colunas aqui nem no README até a equipe fechar o modelo.

### Fluxo de uma feature

`routes` → `controller` → `use_case` → `repository` → `model` (GORM), com `mapper` convertendo `model` ↔ `entity`. Ao criar um módulo/feature:

1. Model em `persistence/model/<nome>_model.go`, entity em `entity/<nome>_entity.go`, mapper em `persistence/mapper/`.
2. Repository, use case (um arquivo por caso de uso, ex.: `create_product.go`) e controller.
3. Instanciar tudo em `bootstrap/modules/<nome>.go` e expor no `App`.
4. Registrar as rotas em `routes/routes.go`, seguindo exatamente o `openapi.yaml`.

## Comandos

Executar a partir de `backend/`:

```bash
cp .env.example .env        # definir DB_DSN, ex.: DB_DSN=ciclobits.db
go run ./cmd/api            # sobe a API em :8080 (GET /api/v1/ping)
go build ./...
go vet ./...
go test ./...
gofmt -w .
```

O `main.go` encerra com `log.Fatal` se não encontrar o `.env`. Arquivos `.env` e `*.db` são ignorados pelo git.

## Convenções da API (resumo do openapi.yaml)

- Base: `http://localhost:8080/api/v1`.
- Campos JSON em `snake_case`, sem acentos, iguais aos dos arquivos em `data/seed/`.
- Datas `AAAA-MM-DD`; quantidades nunca negativas; campos `readOnly` são calculados pelo backend.
- Erros no schema `Erro`: `400` malformado, `404` não encontrado, `409` conflito, `422` regra de negócio.
- "Hoje" = data de referência do sistema (`GET /dashboard` → `data_referencia`), **não** o relógio do servidor.

Regras de negócio centrais:

- Necessidade de material = qtd. a produzir × `quantidade` × (1 + `perda_percentual` / 100).
- Saldo livre = `quantidade_disponivel` − `quantidade_reservada`.
- Qtd. a produzir de um pedido = `quantidade` − Σ `quantidade_produzida` das ordens vinculadas.
- Reserva de estoque: atende primeiro o pedido com `data_entrega` mais próxima.
- "Atrasada" não é status persistido: é calculado (`data_fim_prevista` < data de referência e ordem não concluída/cancelada).

## Pontos de atenção

- O banco é **SQLite**: não usar sintaxe específica de PostgreSQL (ex.: o operador regex `~` em `CHECK`).
- Pacote de UUID é `github.com/google/uuid`.
- Timestamps dos models seguem o padrão do GORM: `CreatedAt`, `UpdatedAt` e `DeletedAt gorm.DeletedAt` (soft delete).
- Ao alterar o contrato, atualize `docs/openapi.yaml` junto com o código.
- Dados de seed foram gerados com apoio de IA e devem ser revisados pela equipe.
