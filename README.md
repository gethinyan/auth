# Server

## Getting Start

```bash
$ git clone -b server https://github.com/gethinyan/auth.git
$ cd auth
$ source ./sendgrid.env
$ go run auth.go
# Then, visit the address http:/localhost:9090/hello.
```

## Initial Postgres

```bash
$ brew install postgresql
$ brew services start postgresql
$ psql postgres

CREATE ROLE auth WITH LOGIN PASSWORD '123456';
CREATE DATABASE auth;
ALTER ROLE auth Superuser;
```

## Initial Redis

```bash
$ brew install redis
$ brew services start redis
$ redis-cli

keys *
```

## OpenAPI

```bash
# go-swagger generate api doc and serve it.
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
$ swagger generate spec -o ./swagger.yml
$ swagger serve ./swagger.yml
```
