import os

from flask import Flask, jsonify
from .redis import RedisClient

def create_app(test_config=None):
    # create and configure app
    app = Flask(__name__, instance_relative_config=True)
 
    try:
        os.makedirs(app.instance_path)
    except OSError:
        pass

    # redisClient = RedisClient()
    redisClient = RedisClient(useInMemoryOnly=True)

    from . import cart
    cart_bp = cart.create_cart_bp(redisClient)
    app.register_blueprint(cart_bp)

    app.add_url_rule('/', endpoint='index')

    return app