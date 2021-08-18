from flask import Flask
import asyncio
from forager.mq.send import send
from forager.mq.receive import receive
__version__ = '0.1.0'

def create_app():
    app = Flask(__name__)

    from forager.views import index
    app.register_blueprint(index)

    asyncio.run(receive())

    return app