# GoLocal - API

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](/backend/)
[![Postgres](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)](/backend/internal/database/database.go)
[![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)](/backend/docker-compose.yml)

**REST API** server which powers `Flutter App`. Server provides _CRUD_ operations on database, supports _email sending_, creating _account_ by providers such as _Google_, _Facebook_. Verification is handled by **JWT**(JSON Web Tokens). Potential images are stored in [media](/backend/media/) directory.

## Environment Configuration

Set the following environment variables in your `.env` file using [.env.sample](/backend/.env.sample)

### Hint

```bash
cp .env.sample .env
```

**Remember**: you should insert `.env` file in backend directory.

## MakeFile

Down below, there are some commands which will help you to run the app.

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

## Tips

- if firewall is bugging you on Windows, you can add an [Inbound Rule](https://stackoverflow.com/a/65393403)
