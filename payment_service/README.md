# Payment Service (Golang)

## Technology Stack:
- [Language] Golang
- [Database] Postgresql


## How to run
1. Install Go
2. ```go build -o bin && HTTP_SERVER_PORT=8080 ./bin  ```


## Service Documentation
### Make Payment Record
`POST:` ```/api/make_payment```

Request Body:

```
{
    "external_id": "your-unique-id",
    "method": "BCA_VA",
    "amount": 20000,
    "active_duration": 3600 // InSeconds
}
```
Response:
```
{
    "payment_id": "payment_01"
}
```


### Get Payment Record
`GET:` `/api/get_payment/{id}`
```
// No Body Parameter
```
Response:
```
{
	"external_id"` : "your-external-id",
	"method"` : "BCA_VA",
	"status"` : "UNPAID",
	"amount"` : 20000
}
```
### Complete Payment Record
`POST:` `/api/manage_payment`
```
{
    "ID" :"Generated ID",
    "action": "PAY"
}
```

List of action:
- "PAY"
- "CANCEL"


