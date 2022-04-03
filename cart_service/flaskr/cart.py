import uuid
import json
from flask import (
    Blueprint, jsonify, Flask, request
)

from flask_cors import CORS, cross_origin

from .redis import RedisClient
from .clients.paymentClient import PaymentClient
from .clients.productClient import ProductClient

from typing import List
from dataclasses import dataclass


@dataclass
class LineItem:
    product_id: str = ""
    title: str = ""
    description: str = ""
    quantity: int = ""
    price: int = ""


@dataclass
class Cart:
    ID: str = ""
    lineItems: List[LineItem] = None
    paymentMethod: str = ""
    totalPrice: int = 0


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
    # paymentMethod=
    totalPrice=20000
)


def create_cart_bp(redisClient: RedisClient, paymentClient: PaymentClient, productClient: ProductClient):

    bp = Blueprint('cart', __name__)

    @cross_origin
    @bp.route("/api/get_cart/<string:cart_id>", methods=["GET"])
    def get_cart(cart_id):
        cart = redisClient.get(cart_id)
        return jsonify({"cart":cart})

    @cross_origin
    @bp.route("/api/upsert_cart", methods=["POST"])
    def upsert_cart():
        print("Upsert Request Recieved", request)
        print(str(request.get_data()))
        request_data = json.loads(request.data, strict=False)
        cart_ID, paymentMethod = "", ""
        newLineItem = []
        try:
            newLineItem = request_data["lineItems"] or []
        except:
            return "Error Bad Request"

        try:
            paymentMethod = request_data["paymentMethod"]
        except:
            pass

        # * Generate new ID if ID do not provided
        try:
            cart_ID = request_data["ID"] or ""
        except:
            pass

        if cart_ID == "":
            cart_ID = str(uuid.uuid4())

        # * Get Existing Cart
        cart = redisClient.get(cart_ID)

        if cart is None:
            cart = Cart(ID=cart_ID, lineItems=[], paymentMethod=paymentMethod)
        else:
            cart = Cart(**cart)

        # * Update Cart
        cart.lineItems = newLineItem[:]

        if paymentMethod != "":
            cart.paymentMethod = paymentMethod

        cart.totalPrice = sum([item['quantity'] * item['price']
                              for item in cart.lineItems])

        # * Save Cart to Storage
        redisClient.set(cart.ID, cart.__dict__)
        print(cart)
        response = jsonify(cart)
        return response

    @cross_origin
    @bp.route("/api/place_order", methods=["POST"])
    def place_order():
        print("PLACE ORDER")
        # request_data = request.get_json()
        print(str(request.get_data()))
        request_data = json.loads(request.data, strict=False)
        print(request_data)
        cart_ID, paymentMethod = "",  ""
        options = {}
        try:
            print(request_data)
            cart_ID = request_data["ID"]
        except:
            return "Error Bad Request"
        cart = redisClient.get(cart_ID)

        if cart is None:
            return "Error No Cart to Process"

        cart = Cart(**cart)

        # Empty Checkouted Cart
        redisClient.delete(cart_ID)

        paymentMethod = cart.paymentMethod
        externalID = str(uuid.uuid1())
        paymentID = paymentClient.makePayment(
            externalID=externalID, amount=cart.totalPrice, method=paymentMethod)


        print()
        # TODO: Remove Product Quantity

        # TODO: return payment method
        # paymentMet
        returnValue = {
            "cart_id": cart_ID,
            "payment_method": paymentMethod or "",
            "payment_id": paymentID,
            # "payment_link": paymentLink
        }
        return jsonify(returnValue)

    return bp
