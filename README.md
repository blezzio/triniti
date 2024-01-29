# [Triniti.cc](https://triniti.cc)

A (definitely not over-engineered) URL shortener service.

## Prequisites

You will need these things to set up this project on a local machine.

### Required

- [Go](https://go.dev/) is the programming language this service is written in.
- [Docker](https://www.docker.com/) to set up infrastructure like PostgreSQL without the need for complicated installation.

### Optional

- [TailwindCSS](https://tailwindcss.com/) is optional, and used to generate the CSS styling.
- [NodeJS](https://nodejs.org/en) is required by TailwindCSS, so install it only if you are intending to use TailwindCSS.
- [GNU Make](https://www.gnu.org/software/make/) is optional, you can copy the command from the `Makefile` and run it.

## Setup Local Environment

After installing all of the prerequisites, Run these commands.

```bash
# install dependencies
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# setup
$ docker compose up -d
$ export PORT=:4444
$ export DATABASE_URL=postgresql://postgres:password@localhost:5432/postgres

$ migrate -database 'postgres://postgres:password@localhost:5432/postgres?sslmode=disable' -source file://migrations up
$ go run main.go
```

If the above commands run without trouble the service will now be running at [http://localhost:4444](http://localhost:4444). If you want, you can also use the terminal to interact with the service.

```bash
# to shorten a URL
$ curl 'localhost:4444/https:/www.google.com/'

# to redirect to the full URL
$ curl localhost:4444/I6_V2m
```

## Build with

- [Go](https://go.dev/)
- [TailwindCSS](https://tailwindcss.com/)

## Deploy with

- [Supabase](https://supabase.com/)
- [Render](https://render.com/)
