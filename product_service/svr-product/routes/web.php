<?php

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It is a breeze. Simply tell Lumen the URIs it should respond to
| and give it the Closure to call when that URI is requested.
|
*/

$router->get('/', function () use ($router) {
    return $router->app->version();
});

$router->post('/addnewproduct', [
    'uses' => 'Controller@createProduct'
]);

$router->get('/productdetail', [
    'uses' => 'Controller@getProductDetail'
]);

$router->post('/updateproduct', [
    'uses' => 'Controller@updateProduct'
]);

$router->post('/addstock', [
    'uses' => 'Controller@addStock'
]);

$router->post('/reducestock', [
    'uses' => 'Controller@reduceStock'
]);

$router->post('/enableproduct', [
    'uses' => 'Controller@enableProduct'
]);

$router->post('/disableproduct', [
    'uses' => 'Controller@disableProduct'
]);

$router->get('/getproductlist', [
    'uses' => 'Controller@getAllProduct'
]);

