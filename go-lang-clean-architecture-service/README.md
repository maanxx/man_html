# go-lang
golang clean architecture service

## project structure

*** module golang ***
```bash

    go mod init name app // tạo tên app

    go mod tidy // kéo pkg go

```
## command
- make server: run server

## CLI Cobra https://github.com/spf13/cobra

## service context
```base
go get -u github.com/teoit/gosctx@latest
```

## cấu hình env

```bash
cp .env.sample .env

ln -s $(pwd)/.env /.env 

```
