import pika

def connect():
    connection = pika.BlockingConnection(pika.ConnectionParameters("localhost"))

    return connection.channel(), connection