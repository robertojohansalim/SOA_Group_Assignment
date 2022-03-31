from itertools import product
import os

from flask import Flask, jsonify
from .redis import RedisClient
from .clients.paymentClient import PaymentClient
from .clients.productClient import ProductClient

from flask_cors import CORS

def create_app(test_config=None):
    # create and configure app
    app = Flask(__name__, instance_relative_config=True)
    cors = CORS(app)

    app.config['CORS_HEADERS'] = 'Content-Type'
    try:
        os.makedirs(app.instance_path)
    except OSError:
        pass

    redisClient = RedisClient(useInMemoryOnly=True)
    paymentClient = PaymentClient()
    productClient = ProductClient()

    from .cart import create_cart_bp 
    cart_bp = create_cart_bp(redisClient, paymentClient,productClient)
    app.register_blueprint(cart_bp, name="cart")

    print("Loading Product")

    from .product import create_product_bp 
    product_bp = create_product_bp(productClient)
    print("Registering Product")
    app.register_blueprint(product_bp, name="product")

    app.add_url_rule('/', endpoint='index')

    return app