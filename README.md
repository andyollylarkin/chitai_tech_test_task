# chitai_tech_test_task

Тестовое задание для компании Читай Технологии. Описание задания: https://pastebin.com/Bq1u901B

> Использована [собственная библиотека](https://github.com/andyollylarkin/go-app-lifecycle) для контроля за жизненым
> циклом приложения

## Запустить проект
```shell
go run ./cmd/main.go
```

## Запросы
```shell
# Получить список книг
curl --request GET \
  --url http://0.0.0.0:8080/books
```

```shell
# Удалить книгу
curl --request DELETE \
  --url http://0.0.0.0:8080/books \
  --header 'Content-Type: application/json' \
  --data '{"book_Id": 1}'
```