import json
import os
import requests


class ProductClient:

    __getProductList = "/getproductlist"
    __getProductDetail = "/productdetail"

    def __init__(self, host:str = "", secret:str = ""):
        if host == "" and secret == "":
            host = os.getenv("PRODUCT_HOST")
            secret = os.getenv("PRODUCT_SECRET")
            
        self.host = host
        self.secret = secret
            
    
    def getProductList(self):
        url = f"{self.host}{self.__getProductList}"

        request = {
           "enableOnly":True
        }
        
        # response = requests.get(url,json=request)
        # response_json = response.json()


        response_json = """[{"id": 1, "product_name": "Aqua Botol", "product_desc": "TEST qweqwe", "product_price": 10000, "product_stock": 999, "product_status": "Active", "created_at": "2022-03-31 03:02:50", "updated_at": "2022-03-31 03:02:50"}, {"id": 2, "product_name": "Aqua Galon", "product_desc": "TEST qweqwe", "product_price": 10000, "product_stock": 999, "product_status": "Active", "created_at": "2022-03-31 03:38:00", "updated_at": "2022-03-31 03:38:00"}]"""
        print(response_json)
        response_json = json.loads(response_json)
        print(response_json)
        # print(response_json["paymentLink"])
        # print(response.json())
        return response_json

    def getProduct(self, productID):
        url = f"{self.host}{self.__getProductDetail}"

        request = {
           "enableOnly":True
        }
        
        # response = requests.get(url,json=request)
        # response_json = response.json()
        response_json = """{"id": 1, "product_name": "Aqua Botol", "product_desc": "MOCK DATA", "product_price": 10000, "product_stock": 999, "product_status": "Active", "created_at": "2022-03-31 03:02:50", "updated_at": "2022-03-31 03:02:50"}"""
        if productID == 2:
            response_json = """{"id": 2, "product_name": "Aqua Galon", "product_desc": "MOCK DATA", "product_price": 10000, "product_stock": 999, "product_status": "Active", "created_at": "2022-03-31 03:02:50", "updated_at": "2022-03-31 03:02:50"}"""
        print(response_json)
        response_json = json.loads(response_json)
        print(response_json)
        # print(response_json["paymentLink"])
        # print(response.json())
        return response_json

        


    
