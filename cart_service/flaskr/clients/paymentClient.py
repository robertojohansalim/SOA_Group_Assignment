import json
import os
import requests


class PaymentClient:

    __makePaymentRequestPATH = "/api/make_payment"

    def __init__(self, host:str = "", secret:str = ""):
        if host == "" and secret == "":
            host = os.getenv("PAYMENT_HOST")
            secret = os.getenv("PAYMENT_SECRET")
            
        self.host = host
        self.secret = secret
            
    def makePayment(self, externalID: str, amount:int, method:str ):
        Headers = {"Authorization":self.secret}
        url = f"{self.host}{self.__makePaymentRequestPATH}"

        request = {
            "external_id": externalID,
            "method": method,
            "amount": amount,
            "callback_url": "",
        }

        response = requests.post(url,json=request,  headers=Headers)
        response_json = response.json()
        print(response_json["paymentLink"])
        print(response.json())
        return response_json["paymentLink"]


        


    
