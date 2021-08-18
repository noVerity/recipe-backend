import pika
import os

def connect():
    param = os.environ.get("CLOUDAMQP_URL")
    if not param:
        param = "amqp://guest:guest@localhost"
    connection = pika.BlockingConnection(pika.URLParameters(param))

    return connection.channel(), connection