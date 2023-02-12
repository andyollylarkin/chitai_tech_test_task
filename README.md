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


## RPS benchmark
```shell
ab -n 1000 -c10 http://localhost:8080/books
```
>Server Software:        \
>Server Hostname:        localhost\
>Server Port:            8080\
>
>Document Path:          /books\
>Document Length:        199 bytes\
>
>Concurrency Level:      10\
>Time taken for tests:   50.620 seconds\
>Complete requests:      1000\
>Failed requests:        0\
>Total transferred:      323000 bytes\
>HTML transferred:       199000 bytes\
>Requests per second:    19.76 [#/sec] (mean)\
>Time per request:       506.201 [ms] (mean)\
>Time per request:       50.620 [ms] (mean, across all concurrent requests)\
>Transfer rate:          6.23 [Kbytes/sec] received\
>
>Connection Times (ms)\
>min  mean[+/-sd] median   max\
>Connect:        0    0   0.1      0       1\
>Processing:   500  501   0.5    501     503\
>Waiting:      500  501   0.5    501     503\
>Total:        500  501   0.5    501     503\
