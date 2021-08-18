import os
from forager.mq.connect import connect

sending_queue = os.environ.get("APP_OUT_QUEUE")

def send(message):
    channel, connection = connect()    

    channel.queue_declare(queue=sending_queue)

    channel.basic_publish(exchange='',
                      routing_key=sending_queue,
                      body=message)

    connection.close()