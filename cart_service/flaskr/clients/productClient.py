from itertools import product
import json
from math import prod
import os
import requests


class ProductClient:

    __getProductList = "/getproductlist"
    __getProductDetail = "/productdetail"
    __reduceStock = "/reducestock"

    def __init__(self, host: str = "", secret: str = ""):
        if host == "" and secret == "":
            host = os.getenv("PRODUCT_HOST")
            secret = os.getenv("PRODUCT_SECRET")

        self.host = host
        self.secret = secret

    def reduceStock(self, product_id, stock):
        url = f"{self.host}{self.__reduceStock}"

        request = {
            "id": product_id,
            "stock": stock
        }

        response = requests.post(url, json=request)
        response_json = response.json()

        response_json = json.loads(response_json)
        return response_json
