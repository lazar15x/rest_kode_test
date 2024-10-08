# Входим в аккаунт
# получаем статичный токен
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{
  "username": "admin",
  "password": "admin"
}'

# Получение списка заметок пользователя
curl -X GET http://localhost:8080/lk/notes \
-H "Authorization: 3rte433gggr4"

# Создаем заметку, в ответ сервер присылает созданную заметку
echo '{"title": "Название", "description": "Новая тестовая заметка"}' | iconv -t UTF-8 | curl -X POST http://localhost:8080/lk/notes \
-H "Authorization: 3rte433gggr4" \
-H "Content-Type: application/json; charset=UTF-8" \
--data @-

# Проверка неправильно введеного пароля или логина
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{
  "username": "0000",
  "password": "0000"
}'

# Проверка отправки/получения заметок без токена
curl -X GET http://localhost:8080/lk/notes

