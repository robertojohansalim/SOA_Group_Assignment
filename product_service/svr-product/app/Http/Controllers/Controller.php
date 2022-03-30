<?php

namespace App\Http\Controllers;

use App\Models\Product;
use Illuminate\Http\Request;
use Laravel\Lumen\Routing\Controller as BaseController;

class Controller extends BaseController
{
    public function createProduct(Request $request){
        $id = Product::addProduct([
            'product_name' => $request['product_name'],
            'product_price' => $request['product_price'],
            'product_stock' => $request['product_stock'],
            'product_status' => Product::PRD_ACTIVE,
            'created_at' => gmdate("Y-m-d H:i:s", time()),
            'updated_at' => gmdate("Y-m-d H:i:s", time())
        ]);

        $product = Product::getProduct($id);
        return json_encode($product);
    }

    public function getProductDetail(Request $request){
        $product = Product::getProduct($request['id']);
        return json_encode($product);
    }

    public function updateProduct(Request $request){
        // $update = [
        //     'product_name' => $request['product_name'],
        //     'product_price' => $request['product_price'],
        //     'product_stock' => $request['product_stock']
        // ];

        $update = [];
        if(isset($request['product_name'])){
            $update['product_name'] = $request['product_name'];
        }
        if(isset($request['product_price'])){
            $update['product_price'] = $request['product_price'];
        }
        if(isset($request['product_stock'])){
            $update['product_stock'] = $request['product_stock'];
        }

        $before = Product::getProduct($request['id']);

        Product::updateProduct($request['id'], $update);

        $after = Product::getProduct($request['id']);

        return json_encode(['before' => $before, 'after' => $after]);
    }

    public function addStock(Request $request){
        $product = Product::getProduct($request['id']);

        $stock = $product->product_stock + $request['stock'];

        Product::updateProduct($request['id'], ['product_stock' => $stock]);

        return json_encode(Product::getProduct($request['id']));
    }

    public function reduceStock(Request $request){
        $product = Product::getProduct($request['id']);

        $stock = $product->product_stock - $request['stock'];

        if($stock >= 0){
            Product::updateProduct($request['id'], ['product_stock' => $stock]);

            return json_encode(Product::getProduct($request['id']));
        }
        else{
            return json_encode(['error' => 'updated stock less than 0. Requested: '.strval($request['stock']). ' Current stock: '.strval($product->product_stock)]);
        }
    }

    public function enableProduct(Request $request){
        Product::enableProduct($request['id']);
        return json_encode(Product::getProduct($request['id']));
    }

    public function disableProduct(Request $request){
        Product::disableProduct($request['id']);
        return json_encode(Product::getProduct($request['id']));
    }

    public function getAllProduct(Request $request){
        if(isset($request['enabledOnly']) && $request['enabledOnly']){
            $productList = Product::getAllProduct(true);
        }
        else{
            $productList = Product::getAllProduct();
        }
        return json_encode($productList->toArray());
    }
}
