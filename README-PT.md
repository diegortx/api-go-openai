# API de Chat em Go

Uma API RESTful que fornece funcionalidade de chat utilizando os modelos GPT da OpenAI.

## Funcionalidades

- Endpoint de chat para processamento de mensagens com integração OpenAI
- Gerenciamento de conversas com histórico
- Validação de requisições e tratamento abrangente de erros
- Suporte a CORS (Cross-Origin Resource Sharing) para clientes web
- Registro detalhado de requisições com informações de tempo
- Endpoint de verificação de saúde com timestamp
- Armazenamento de conversas thread-safe usando mutex
- Limite máximo de histórico de mensagens para otimizar contexto
- Funcionalidade de limpar conversas

## Endpoints da API

### POST /api/v1/chat
Processa uma mensagem de chat e obtém uma resposta gerada por IA.

**Corpo da Requisição:**

```json
{
    "user_id": "user123",
    "message": "Qual é a capital da França?",
    "context": {
        "system_prompt": "Você é um assistente especializado em geografia.",
        "metadata": {
            "language": "pt",
            "topic": "geografia"
        },
        "temperature": 0.7
    }
}
```

**Resposta:**

```json
{
    "success": true,
    "response": "A capital da França é Paris."
}
```

### DELETE /api/v1/chat/{userID}
Limpa o histórico de conversas de um usuário específico.

**Resposta:**

```json
{
    "success": true,
    "response": "Conversa limpa com sucesso"
}
```

### GET /health
Verifica se o serviço está em execução.

**Resposta:**

```json
{
    "status": "ok",
    "time": "2024-03-21T10:00:00Z"
}
```

## Configuração

1. Clone o repositório:
```bash
git clone <repository-url>
cd go-api
```

2. Copie o arquivo de ambiente de exemplo e configure suas configurações:
```bash
cp .env.example .env
# Edite o .env com sua chave de API da OpenAI
```

3. Instale as dependências:
```bash
go mod tidy
```

4. Execute a aplicação:
```bash
# Desenvolvimento com recarga automática
air

# Ou execução padrão
go run main.go
```

## Variáveis de Ambiente

| Variável | Descrição | Obrigatório | Padrão |
|----------|-----------|-------------|---------|
| PORT | Porta do servidor | Não | 8080 |
| OPENAI_API_KEY | Chave da API OpenAI | Sim | - |

## Estrutura do Projeto

```
.
├── config/         # Gerenciamento de configuração
├── controllers/    # Manipuladores de requisição
├── models/         # Estruturas de dados
├── routes/         # Definições de rotas
├── services/       # Lógica de negócio
├── .env           # Configuração de ambiente
├── .env.example   # Exemplo de configuração
├── main.go        # Ponto de entrada da aplicação
└── README.md      # Documentação
```

## Exemplos de Requisição

### Enviar uma Mensagem de Chat
```bash
curl -X POST http://localhost:8080/api/v1/chat \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "message": "Qual é a capital da França?",
    "context": {
        "system_prompt": "Você é um assistente prestativo.",
        "temperature": 0.7
    }
  }'
```

### Limpar Histórico de Chat
```bash
curl -X DELETE http://localhost:8080/api/v1/chat/user123
```

## Tratamento de Erros

A API retorna códigos de status HTTP apropriados e mensagens de erro:

```json
{
    "success": false,
    "error": "Mensagem de erro aqui"
}
```

Códigos de status comuns:
- 400: Requisição Inválida (entrada inválida)
- 401: Não Autorizado (chave de API inválida)
- 500: Erro Interno do Servidor

## Desenvolvimento

### Recarga Automática
O projeto usa Air para recarga automática durante o desenvolvimento:

1. Instale o Air:
```bash
go install github.com/cosmtrek/air@latest
```

2. Execute com recarga automática:
```bash
air
```

## Contribuindo

1. Faça um fork do repositório
2. Crie sua branch de feature (`git checkout -b feature/recurso-incrivel`)
3. Faça commit de suas alterações (`git commit -m 'Adiciona algum recurso incrível'`)
4. Faça push para a branch (`git push origin feature/recurso-incrivel`)
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo LICENSE para detalhes.

## Agradecimentos

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [OpenAI API](https://platform.openai.com/docs/api-reference)

