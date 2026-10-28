# Ciclobits PCP

Sistema de Planejamento e Controle da Produção (PCP) desenvolvido para uma indústria fictícia de bicicletas, como projeto acadêmico da disciplina de **Análise e Projeto de Sistemas** do curso de Ciência da Computação da UNISC.

## Sobre o projeto

A **Ciclobits** é uma empresa fictícia especializada na fabricação e montagem de bicicletas urbanas, mountain bikes e bicicletas infantis.

O sistema tem como objetivo auxiliar no planejamento, organização e acompanhamento da produção, centralizando informações sobre pedidos de clientes, produtos, matérias-primas, estoque, recursos produtivos e ordens de produção.

A aplicação permitirá acompanhar o andamento da fabricação, calcular a necessidade de materiais, verificar a disponibilidade de estoque e organizar a programação da produção, considerando os prazos de entrega e a capacidade produtiva.

## Tecnologias utilizadas

O projeto será desenvolvido utilizando as seguintes tecnologias:

- **Backend:** Go (Golang)
- **Frontend:** React com TypeScript
- **Banco de dados:** SQLite
- **Comunicação:** API REST
- **Controle de versão:** Git e GitHub

## Pré-requisitos

Para executar o projeto localmente, será necessário ter instalado:

- **Go:** versão 1.22 ou superior
- **Node.js:** versão 20 ou superior
- **npm:** gerenciador de pacotes do Node.js
- **Git:** para clonar o repositório

O banco de dados SQLite será inicializado automaticamente pelo backend, não sendo necessária a instalação de um servidor de banco de dados.

## Estrutura do projeto

```text
velotech-pcp/
├── backend/
│   ├── cmd/
│   ├── internal/
│   └── data/
├── frontend/
└── README.md
```

- `backend/`: API REST, regras de negócio, persistência e inicialização do banco de dados.
- `frontend/`: interface web desenvolvida com React e TypeScript.

## Base de dados

A aplicação utilizará arquivos JSON para a carga inicial do banco de dados, contendo informações fictícias sobre:

- Produtos e suas estruturas de fabricação;
- Materiais e componentes;
- Estoque de materiais;
- Recursos produtivos;
- Pedidos de clientes;
- Ordens de produção.

Na primeira execução, o backend será responsável por criar as tabelas e importar os dados para o SQLite. As alterações realizadas posteriormente serão persistidas no banco de dados.

## Execução do projeto

As instruções para instalação das dependências, configuração do ambiente e execução do backend e frontend serão adicionadas conforme o desenvolvimento da aplicação.

## Uso de Inteligência Artificial

O projeto conta com o apoio do ChatGPT (OpenAI) na elaboração da base inicial de dados fictícios, no planejamento da arquitetura e na documentação.

Os materiais gerados com auxílio de IA deverão ser revisados e validados pela equipe antes de sua utilização no projeto.

## Status do projeto

**Em desenvolvimento.**

Projeto acadêmico desenvolvido para a disciplina de Análise e Projeto de Sistemas — UNISC, 2026.
