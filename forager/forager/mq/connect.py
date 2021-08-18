import pika
import os

def connect():
    param = os.environ.get("CLOUDAMQP_URL")
    param2 = os.environ.get("CLOUDAMQP")
    if not param:
        param = "localhost"
    
    print("1: %s" % param)
    print("2: %s" % param2)
    connection = pika.BlockingConnection(pika.ConnectionParameters(param))

    return connection.channel(), connection