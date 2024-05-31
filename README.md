### POST /purchase - Create a new purchase

```shell
curl --location 'localhost:8080/purchase' \
--header 'x-api-key: test-x-api-key' \
--header 'Content-Type: application/json' \
--data '{
  "id": "123",
  "description": "Purchase description",
  "resultUrl": "http://example.com/result",
  "amount": {
    "value": 100.0,
    "currency": "USD"
  },
  "card": "1234561234561234",
  "payer": {
    "id": "456",
    "name": {
      "firstName": "John",
      "lastName": "Doe"
    }
  }
}'
```

### GET /purchase/{id} - Get purchase by id

```shell
curl --location 'localhost:8080/purchase/123' \
--header 'x-api-key: test-key'
```
