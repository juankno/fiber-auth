# Fiber Auth

Fiber Auth is a lightweight authentication service built with the [Fiber](https://gofiber.io/) web framework for Go.

## Features

- User registration and login
- JWT-based authentication
- Secure password hashing
- Middleware for protected routes

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/fiber-auth.git
   cd fiber-auth
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up database:
   - Ensure you have a PostgreSQL database running.
   - Update the database connection string in `config/config.go`.

```bash
    docker run --name fiber-auth -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=fiber-auth -p 5432:5432 -d postgres
```

3. Run the application:

```bash
go run main.go
```

## Usage

- Register a new user:

```http
POST /register
Content-Type: application/json

{
   "username": "example",
   "password": "password123"
}
```

- Login:

  ```http
  POST /login
  Content-Type: application/json

  {
     "username": "example",
     "password": "password123"
  }
  ```

- Access protected routes:
  Add the `Authorization: Bearer <token>` header to your requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
