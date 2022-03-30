import os

from flask import Flask, jsonify
from .redis import RedisClient
from .clients.paymentClient import PaymentClient

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

    from . import cart
    cart_bp = cart.create_cart_bp(redisClient, paymentClient)
    app.register_blueprint(cart_bp)

    app.add_url_rule('/', endpoint='index')

    return app