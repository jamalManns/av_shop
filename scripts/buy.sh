#!/bin/bash

cd -- "$( dirname -- "$( readlink -f -- "$0" )" )" || exit 1

# URL приложения
BASE_URL="http://localhost:8080/api"

# Шаг 1: Регистрация нового пользователя и получение JWT-токена
echo "Step 1: Registering a new user..."
USERNAME="testuser7"

# Отправляем POST-запрос на эндпоинт /auth
AUTH_RESPONSE=$(
  curl -s -X POST "$BASE_URL/auth" \
  -H "Content-Type: application/json" \
  -d "{\"username\": \"$USERNAME\"}"
)

# Проверяем успешность регистрации
# echo "response: $AUTH_RESPONSE"
if echo "$AUTH_RESPONSE" | jq -e '.token' >/dev/null; then
  TOKEN=$(echo "$AUTH_RESPONSE" | jq -r '.token')
  echo "User registered successfully. Token: $TOKEN"
  echo "TOKEN=$TOKEN" > token.sh
else
  echo "Failed to register user. Response: $AUTH_RESPONSE"
  source token.sh
  echo "loaded previous token: $TOKEN"
fi

# Шаг 2: Покупка товара
echo "Step 2: Purchasing an item..."

# Выбираем товар для покупки (например, t-shirt)
ITEM="t-shirt"

# Отправляем GET-запрос на эндпоинт /buy/{item} с JWT-токеном
PURCHASE_RESPONSE=$(
  curl -s -X GET "$BASE_URL/buy/$ITEM" \
  -H "Authorization: Bearer $TOKEN"
)

# echo "response: $PURCHASE_RESPONSE"
# Проверяем успешность покупки
if echo "$PURCHASE_RESPONSE" | jq -e '.message' >/dev/null && [[ $(echo "$PURCHASE_RESPONSE" | jq -r '.message') == "Item purchased successfully" ]]; then
  echo "Item purchased successfully."
else
  echo "Failed to purchase item. Response: $PURCHASE_RESPONSE"
  exit 1
fi
