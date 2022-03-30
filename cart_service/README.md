# Cart / Orchestrator Service
## Tech Stack
- PHP, lumen framework
- mysql

## How to Start (PHP server way)
1. Set up virtual environemnt
    ```
    // go to svr-product directory
    $ cd ./product_service/svr-service

    // instal depedency
    $ composer install

    // create .env
    $ cp .env.example .env

	// set mysql database at env
	DB_HOST=[YOUR DB HOST]
	DB_PORT=[YOUR DB PORT]
	DB_DATABASE=[YOUR DB NAME]
	DB_USERNAME=[YOUR DB USERNAME]
	DB_PASSWORD=[YOUR DB PASSWORD]

	// create table
	$ php aritsan migrate
	
	// run php server
	$ php -S localhost:8000 -t public
    ```


## Documentation:

For endpoint, import thunder-collection_svr-product_postman.json to postman or thunder-collection_svr-product.json to thunder client

## Reference:
- https://lumen.laravel.com/docs/7.x
