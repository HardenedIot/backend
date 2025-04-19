# Бекенд

![Логотип](./image.png)

## Стек технологий

Язык программирования: Go

Фреймворки и библиотеки:

- Gin Framework
- GORM + Postgres driver
- Mongo driver
- Validator
- Golang-JWT
- Testify

## API, Документация для разработчиков

OpenAPI спецификация доступна [здесь](./openapi.yaml).

## Юнит-тестирование

Проверить работу эндпоинтов можно следующим образом:

```shell
cd backend/src
go test ./tests
```

## Ссылка на собранный контейнер для развертывания

```shell
ghcr.io/hardenediot/backend:latest
```

## Инструкция по самостоятельной сборке Docker образа

1. Клонируйте репозиторий:
   ```shell
   git clone https://github.com/hardenediot/backend.git
   cd backend/
   ```

2. Запустите следующую команду:

```shell
docker build -t hardenediot/backend:latest .
```

## Развертывание приложения

См. [репозиторий деплоя](https://github.com/hardenediot/deploy).

## Лицензия

Этот проект лицензирован под лицензией GPL-3.0. Подробнее - см. файл [LICENSE](./LICENSE).
