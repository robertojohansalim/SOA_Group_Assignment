import uuid
import json
from flask import (
    Blueprint, jsonify, Flask, request
)

from flask_cors import CORS, cross_origin

from .clients.productClient import ProductClient

from typing import List
from dataclasses import dataclass

print("Product.py Loaded")

@dataclass
class Product:
    id: str = ""
    title: str = ""
    price: int = 0
    stock: int = 0
    status: str = ""

    @classmethod
    def fromClient(cls, json) -> "Product":
        return cls(
            id=json["id"],
            title=json['product_name'],
            price=json['product_price'],
            stock=json['product_stock'],
            status=json['product_status'],
        )

def create_product_bp(productClient: ProductClient):
    print("Creating Product BP")
    bp = Blueprint('product', __name__)

    @cross_origin
    @bp.route("/api/get_product_list", methods=["POST"])
    def getProductList():

        productListClient = productClient.getProductList()
        productList = []
        for clientProduct in productListClient:
            print(clientProduct)
            productList.append(Product.fromClient(clientProduct))

        print(productList)

        return jsonify(productList)

    @cross_origin
    @bp.route("/api/get_product", methods=["POST"])
    def getProduct():
        request_data = request.get_json()
        product_ID = ""
        try:
            product_ID = request_data["ID"]
        except:
            return "Error Bad Request"
        clientProduct = productClient.getProduct(product_ID)
        product = Product.fromClient(clientProduct)

        return jsonify(product)

    return bp