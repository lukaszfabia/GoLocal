# GoLocal - API

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](/backend/)
[![Postgres](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)](/backend/internal/database/database.go)
[![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)](/backend/docker-compose.yml)

**REST API** server which powers `Flutter App`. Server provides _CRUD_ operations on database, supports _email sending_, creating _account_ by providers such as _Google_, _Facebook_. Verification is handled by **JWT**(JSON Web Tokens). Potential images are stored in [media](/backend/media/) directory.

**Not implemented**(but in plans): ML server(service) on **Fast API** which will help to make event recommendations for users.

## Environment Configuration

Set the following environment variables in your `.env` file:

```bash
# ============ Server Configuration ============
# Application port and environment
PORT=                 # Application port (e.g., 3000)
APP_ENV=              # Environment (e.g., development, production)

# ============ Database Configuration ============
# Database connection settings
DB_HOST=              # Database host (e.g., localhost)
DB_PORT=              # Database port (e.g., 5432)
DB_DATABASE=          # Database name
DB_USERNAME=          # Database username
DB_PASSWORD=          # Database password

# ============ pgAdmin Configuration ============
# pgAdmin settings for database management
PGADMIN_EMAIL=        # pgAdmin login email
PGADMIN_PASSWORD=     # pgAdmin login password
PGADMIN_PORT=         # Port for accessing pgAdmin

# ============ Google OAuth 2.0 ============
# Google OAuth credentials for user authentication
GOOGLE_CLIENT_ID=     # Google OAuth Client ID
GOOGLE_CLIENT_SECRET= # Google OAuth Client Secret

# ============ Session Management ============
# Session secret for secure session handling
SESSION_SECRET=       # Secret key for session encryption

# ============ JWT Configuration ============
# JWT (JSON Web Token) configuration for secure authentication
JWT_SECRET=           # Secret key for JWT signing and verification

# ============ Email Sending Configuration ============
# Gmail credentials for sending emails from the application
GMAIL_MAIL=           # Gmail address used for sending emails
GMAIL_PASSWORD=       # Gmail password for authentication
```

**Remember**: you should insert `.env` file in backend directory.

## MakeFile

Down below, there're some commands which will help you to run the app.

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB Container

```bash
make docker-down
```

DB Integrations Test:

```bash
make itest
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```
