<?php

namespace App\Models;

use Illuminate\Support\Facades\DB;
use PDO;

class Product {

    const PRD_ACTIVE = 1001;
    const PRD_NONACTIVE = 1002;

    public static function addProduct($product){
        return DB::table('product')->insertGetId($product);
    }

    public static function getProduct($id){
        return DB::table('product')->where('id', $id)->first();
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
            return DB::table('product')->where('product_status','=', self::PRD_ACTIVE)->get();
        }
        else{
            return DB::table('product')->get();
        }
    }
}
