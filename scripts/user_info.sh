#!/bin/bash

cd -- "$( dirname -- "$( readlink -f -- "$0" )" )" || exit 1

# URL приложения
BASE_URL="http://localhost:8080/api"
source token.sh
# Получение информации о пользователе
echo "Getting user information..."
# Отправляем GET-запрос на эндпоинт /info с JWT-токеном
INFO_RESPONSE=$(curl -s -X GET "$BASE_URL/info" -H "Authorization: Bearer $TOKEN")
# Выводим информацию о пользователе
echo "User info: $INFO_RESPONSE"
