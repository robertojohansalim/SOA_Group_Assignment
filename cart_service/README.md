# Cart / Orchestrator Service
## Tech Stack
- Python with Flask
- Redis DB https://github.com/redis/redis-py

## How to Setup
1. Set up virtual environemnt
    ```
    // Create a venv
    $ python -m venv ./venv

    // source the venv (in ubuntu)
    $ source ./venv/bin/activate

    // Instal the required modules
    $ pip install -r requirements.txt
    ```
2. Set up Redis 
    ```
    // Example of setting up redis using docker
    docker run --name my-redis -p 6379:6379 -d redis
    ```

3. Set up the .env
    - Leave everything as is if no configuration required (will be run as developemnt)

## How to run
1. Start server using:
    ```
    python -m flask run -p 5000
    ```


## Documentation:
#### Cart Structure:
```
{
  "cart": {
    "ID": "cart-id",
    "lineItems": [
      {
        "title": "Item 1",
        "description": "Item 1 Very Long Description",
        "quantity": 1,
        "price": 15000
      },
      {
        "title": "Item 2",
        "description": "Item 2 Very Long Description",
        "quantity": 1,
        "price": 5000
      }
    ],
    "totalPrice": 20000
  }
}
```

### Get Cart
`GET` `http://127.0.0.1:5000/api/get_cart/{cart-id}`

Request Body:
```
// Left empty
```


### Upsert Cart
`POST` `http://127.0.0.1:5000/api/upsert_cart`

Request body:
```
{
  "ID": "9458b918-0afc-4163-aa20-219f8c0bb38f",
  "lineItems": [
    {
      "description": "Item 1 Very Long Description",
      "price": 15000,
      "quantity": 1,
      "title": "Item 1"
    }
  ]
}
```
### Place Order (Checkout)
`POST` `http://127.0.0.1:5000/api/place_order`

Request body:
```
{
  "ID": "9458b918-0afc-4163-aa20-219f8c0bb38f",
  "action":"CHECKOUT"
}
```

## Reference:
- https://towardsdatascience.com/the-right-way-to-build-an-api-with-python-cd08ab285f8f
- https://flask.palletsprojects.com/en/2.1.x/tutorial/layout/