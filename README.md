### Data types:
* `string`
* `int`
* `float`
* `bool`

### POST /purchase - Create a new purchase

```shell
curl --location 'localhost:8080/purchase' \
--header 'x-api-key: test-x-api-key' \
--header 'Content-Type: application/json' \
--data '{
  "id": "123",
  "description": "Purchase description",
  "amount": 100.0,
  "card": "1234561234561234",
  "payer": {
    "id": "456",
    "name": {
      "first": "John",
      "last": "Doe"
    }
  },
  "paymentSystem": {
    "url": "http://localhost:8000",
    "apiKey": "test_api_key"
  }
}'
```

### GET /purchase/{id} - Get purchase by id

```shell
curl --location 'localhost:8080/purchase/123' \
--header 'x-api-key: test-key'
```
