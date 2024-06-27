# PixPay - Payment Gateway for Pix

This project is a Payment API developed using Golang. The API supports payment transactions using the Pix system and uses PostgreSQL as its database.

## Project Structure

```bash
/pixpay
|-- /cmd
|   |-- /api
|       |-- main.go
|-- /config
|   |-- config.go
|-- /internal
|   |-- /api
|   |   |-- /handler
|   |   |   |-- payment_handler.go
|   |   |-- /router
|   |       |-- router.go
|   |-- /core
|   |   |-- /payment
|   |       |-- payment.go
|   |-- /repository
|       |-- /postgres
|           |-- payment_repository.go
|-- /pkg
|   |-- /database
|   |   |-- postgres.go
|   |-- /logger
|       |-- logger.go
|-- go.mod
|-- go.sum
```

## Getting Started

**Prerequisites**

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [PostgreSQL](https://www.postgresql.org/)
- [Go Modules](https://blog.golang.org/using-go-modules)

**Installation**

1. Clone the repository

   ```bash
   git clone https://github.com/aleksanderpalamar/pixpay.git
   ```
2. Install dependencies

   ```bash
   go mod download
   ```
   or
   ```bash
   go mod tidy
   ```
3. Set up envirenment variables:

Create a `.env` file in the root directory of the project and add the following variables:

```bash
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
```

4. Run the application

```bash
go run cmd/api/main.go
```

5. Using Docker

```bash
docker compose up --build
```

### Directory Structure

- `cmd`: Contains the entry point for the application.
- `config`: Configuration files for and variables handling.
- `internal`: Internal application code, not exposed outside the module.
  - `api`: API related code (handlers, routers, etc.).
  - `core`: Core business logic.
  - `repository`: Database access code.
- `pkg`: Reusable packages and utilities.
  - `database`: Database initialization and configuration.
  - `logger`: Logger initialization and configuration.

## API Endpoints

- Create Payment
  - URL: `/payments`
  - Method: `POST`
  - Description: Creates a new payment.
- Get Payment
  - URL: `/payments/:id`
  - Method: `GET`
  - Description: Retrieves a payment by its ID.

### Example Request

- Create Payment

```bash	
curl -X POST http://localhost:8080/payments -H "Content-Type: application/json" -d '{"amount": 100, "recipient": "recipient_id"}'
```

- Get Payment

```bash
curl http://localhost:8080/payments/{id}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.




