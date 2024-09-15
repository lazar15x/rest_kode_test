Собираем проект

```sh
go run ./cmd/api
```

или

Собираем и запускаем контейнер

```sh
docker build -t kode_test .
docker run -p 8080:8080 kode_test
```

Для проверки API запускаем проект и в консоли пишем:

```sh
sh api_requests.sh
```
