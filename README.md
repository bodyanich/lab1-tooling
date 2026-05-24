# Lab 1 — Go Tooling

## Опис

Цей проєкт створено для лабораторної роботи №1 з предмету "Інструментальні засоби розробки програмного забезпечення".

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

## Запуск програми
go run ./cmd/app

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