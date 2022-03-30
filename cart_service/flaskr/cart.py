import uuid
from flask import (
    Blueprint, jsonify, Flask, request
)

from .redis import RedisClient

from typing import List
from dataclasses import dataclass

@dataclass
class LineItem:
    title:str
    description:str
    quantity:int
    price:int

@dataclass
class Cart:
    ID: str
    lineItems: List[LineItem]
    totalPrice: int

cart = Cart(
    ID="cart-id",
    lineItems=[
        LineItem(
            title="title-1",
            description="description",
            quantity=1,
            price=5000
        ),
        LineItem(
            title="title-1",
            description="description",
            quantity=1,
            price=15000
        ),
    ],
    totalPrice=20000
   )

def create_cart_bp(redisClient : RedisClient):

    bp = Blueprint('cart', __name__)

    @bp.route("/api/get_cart/<string:cart_id>", methods=["GET"])
    def get_cart(cart_id):
        
        cart = redisClient.get(cart_id)
        return jsonify({"cart":cart})


    @bp.route("/api/upsert_cart", methods=["POST"])
    def upsert_cart():
        request_data = request.get_json()
        cart_ID = ""
        newLineItem = []
        try:
            newLineItem = request_data["lineItems"] or []
        except:
            return "Error Bad Request" 

        #* Generate new ID if ID do not provided
        try:
            cart_ID = request_data["ID"] or ""
        except:
            cart_ID = str(uuid.uuid4())

        #* Get Existing Cart
        cart = redisClient.get(cart_ID)
        
        if cart is None:
            cart = Cart(ID=cart_ID, lineItems=[], totalPrice=0)
        else:
            cart = Cart(**cart)

        #* Update Cart
        cart.lineItems = newLineItem[:]
        cart.totalPrice = sum([item['quantity'] * item['price'] for item in cart.lineItems])

        #* Save Cart to Storage
        redisClient.set(cart.ID, cart.__dict__)

        return jsonify(cart)


    @bp.route("/api/place_order", methods=["POST"])
    def place_order():
        request_data = request.get_json()
        cart_ID = ""
        options = {}
        try:
            cart_ID = request_data["ID"]
            options = request_data["options"]
        except:
            return "Error Bad Request"  
        cart = redisClient.get(cart_ID)

        if cart is None:
            return "Error No Cart to Process" 

        cart = Cart(**cart) 

        redisClient.delete(cart_ID)

        #TODO: Generate Payment
        paymentMethod = options["payment_method"]

        #TODO: Remove Product Quantity

        #TODO: return payment method 
        returnValue = {
            "cart_id": cart_ID,
            "payment_method": paymentMethod,
            "payment_link": "random-payment-link"
        }
        return jsonify(returnValue)

    return bp