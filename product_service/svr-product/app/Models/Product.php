<?php

namespace App\Models;

use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Log;
use PDO;

class Product {

    const PRD_ACTIVE = 1001;
    const PRD_NONACTIVE = 1002;

    public static function addProduct($product){
        return DB::table('product')->insertGetId($product);
    }

    public static function getProduct($id){
        $product = DB::table('product')->where('id', $id)->first();
        $product = (array)$product;
        if($product['product_status'] == self::PRD_ACTIVE){
            $product['product_status'] = 'Active';
        }
        else{
            $product['product_status'] = 'Inactive';
        }
        return $product;
    }

    public static function updateProduct($id, $update){
        return DB::table('product')->where('id', $id)->update($update);
    }

    public static function enableProduct($id){
        $update = [
            'product_status' => self::PRD_ACTIVE
        ];
        return self::updateProduct($id, $update);
    }

    public static function disableProduct($id){
        $update = [
            'product_status' => self::PRD_NONACTIVE
        ];
        return self::updateProduct($id, $update);
    }

    public static function getAllProduct($enabled = false){
        if($enabled){
            $productList = DB::table('product')->where('product_status','=', self::PRD_ACTIVE)->get();
        }
        else{
            $productList = DB::table('product')->get();
        }

        $productList = $productList->toArray();
        $productList = array_map(function($value){return (array)$value;},$productList);

        foreach ($productList as $key => $value) {
            if($value['product_status'] == self::PRD_ACTIVE){
                $productList[$key]['product_status'] = 'Active';
            }
            else{
                $productList[$key]['product_status'] = 'Inactive';
            }
        }
        return $productList;
    }
}
