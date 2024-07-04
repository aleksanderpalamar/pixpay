# PixPay Gateway

## Descrição
PixPay Gateway é uma API Gateway para processamento de pagamentos utilizando Pix. Esta aplicação é construída com Golang e o framework Gin, e utiliza PostgreSQL como banco de dados, Docker para containerização e Terraform para gerenciar a infraestrutura.

## Funcionalidades
1. **Roteamento de Solicitações**: Encaminha solicitações dos clientes para o serviço apropriado no backend e facilita o balanceamento de carga.
2. **Autenticação e Autorização**: Verifica se as solicitações são autenticadas e autorizadas antes de encaminhá-las para os serviços internos.
3. **Agregação de Respostas**: Combina respostas de múltiplos serviços backend em uma única resposta para o cliente.
4. **Transformação de Mensagens**: Modifica as solicitações e respostas conforme necessário (ex.: transformação de formatos de dados, adição/removal de cabeçalhos).
5. **Gerenciamento de Taxas (Rate Limiting)**: Controla o número de solicitações que um cliente pode fazer em um determinado período para prevenir abuso e sobrecarga do sistema.
6. **Monitoramento e Logging**: Coleta métricas e logs das solicitações para monitoramento, análise e troubleshooting.
7. **Cache**: Armazena em cache as respostas de serviços backend para melhorar o desempenho e reduzir a carga nos serviços.

## Estrutura do Projeto

```shell
pixpay-gateway/
├── api/
│ ├── handlers/
│ ├── middleware/
│ ├── routes/
│ └── main.go
├── config/
│ ├── config.go
│ └── config.yaml
├── db/
│ ├── migrations/
│ └── db.go
├── docker/
│ ├── Dockerfile
│ ├── docker-compose.yaml
├── scripts/
│ └── init.sql
├── terraform/
│ ├── main.tf
│ ├── variables.tf
│ └── outputs.tf
├── .env
├── go.mod
├── go.sum
└── README.md
```

## Requisitos
- Go 1.22+
- Docker
- Docker Compose
- Terraform
- PostgreSQL

## Configuração
### Banco de Dados
1. Crie o banco de dados PostgreSQL utilizando o script `scripts/init.sql`.
2. Configure as variáveis de ambiente no arquivo `.env` com as credenciais do banco de dados.

### Docker
1. Construa e inicie os containers:
   ```bash
   docker-compose up --build
   ```
2. Acesse o servidor Docker:
   ```bash
   docker-compose exec pixpay-gateway bash
   ```
3. Inicie o serviço:
   ```bash
   go run cmd/api/main.go
   ```
### Terraform
1. Inicialize o Terraform:
   ```bash
   terraform init
   ```
2. Configure as variáveis de ambiente no arquivo `terraform/variables.tf`.
3. Execute o Terraform:
   ```bash
   terraform apply
   ```

## Comandos

### Criar Banco de Dados
```bash
docker-compose exec pixpay-gateway psql -U postgres -d postgres -f scripts/init.sql
```

### Iniciar Servidor
```bash
go run cmd/api/main.go
```

### Iniciar Servidor Docker
```bash
docker-compose exec pixpay-gateway go run cmd/api/main.go
```

## Licença
Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para obter detalhes.