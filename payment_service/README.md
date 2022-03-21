# Payment Service (Golang)

## Technology Stack:
- [Language] Golang
- [Database] Postgresql


## How to run
1. Install Go
2. ```go build -o bin && HTTP_SERVER_PORT=8080 ./bin  ```


## Service Documentation
### Authorization
All Request must provide Authorization Header that contains the username
```
"authorization": "registered_account_id"
```

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


### Get Payment Record
`GET:` `/api/get_payment/{id}`
```
// No Body Parameter
```
### Complete Payment Record
`POST:` `/api/manage_payment`
```
{
    "external_id": "your-unique-id",
    "action": "pay"
}
```

List of action:
- "pay"
- "cancel"
- "denied"


### Status Callbacks
#### Setting up Callbacks
`POST:` `/api/setup_callback`
```
{
    "callback_type": "update_status_payment",
    "callback_url": "http://your.domain/recieve_callback/url"
}
```

#### Update Status Callbacks
```
{
    "id": "payment_id",
    "external_id": "your-unique-id",
    "status": "PAID"
    "method": "BCA_VA",
    "amount": 20000,
    "expiry_date" : "2006-01-02 15:04:05"
}
```
