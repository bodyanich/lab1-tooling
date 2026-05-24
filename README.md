# Lab 1-2 — Go Tooling and CI

![Go CI](https://github.com/bborodiy41/lab1-tooling/actions/workflows/ci.yml/badge.svg)

## Опис

Цей проєкт створено для лабораторних робіт №1 та №2 з предмету "Інструментальні засоби розробки програмного забезпечення".

Мета роботи — налаштувати професійне середовище розробки на Go, включаючи:

- структуру Go-проєкту;
- unit-тести;
- table-driven tests;
- статичний аналіз коду через golangci-lint;
- автоматизацію команд через Makefile;
- роботу з Git та GitHub.

## Структура проєкту

```text
lab1-tooling/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── calculator.go
│   └── calculator_test.go
├── bin/
├── .golangci.yml
├── .gitignore
├── Makefile
├── README.md
├── go.mod
└── go.sum
```

## Lab 2 — Dependencies and GitHub Actions

У цій лабораторній роботі було додано керування зовнішніми залежностями через Go Modules та налаштовано CI через GitHub Actions.

### Додані залежності

- `go.uber.org/zap` — структуроване логування;
- `github.com/spf13/viper` — читання конфігурації з `config.yaml`.

### Запуск програми

```bash
go run ./cmd/app
```

## Запуск тестів
go test -v ./...

Або через Makefile:

make test

## Запуск тестів з race detector
go test -race -v ./...
## Запуск лінтера
golangci-lint run

Або через Makefile:

make lint
## Форматування коду
make fmt
## Збірка проєкту
make build

## Після збірки виконуваний файл буде створено у папці:

bin/
## Повна перевірка проєкту
make all

## Команда make all послідовно виконує:

форматування коду;
лінтинг;
запуск тестів з race detector;
збірку виконуваного файлу.

## Автор

Студент: Бородій Богдан Сергійович