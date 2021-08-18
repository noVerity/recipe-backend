import pika
import os

def connect():
    param = os.environ.get("CLOUDAMQP_URL")
    if not param:
        param = "localhost"
    connection = pika.BlockingConnection(pika.ConnectionParameters(param))

    return connection.channel(), connection