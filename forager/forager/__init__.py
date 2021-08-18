from flask import Flask
from werkzeug.serving import is_running_from_reloader
import threading
from forager.mq.send import send
from forager.mq.receive import receive
__version__ = '0.1.0'

def create_app():
    app = Flask(__name__)

    from forager.views import index
    app.register_blueprint(index)

    if not is_running_from_reloader():
        receivingThread = threading.Thread(target=receive)
        receivingThread.daemon = True
        receivingThread.start()

    return app