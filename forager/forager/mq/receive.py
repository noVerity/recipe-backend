import os
import json
import logging
from forager.mq.connect import connect
from forager.mq.send import send
from forager.fooddata.lookup import lookup

receiving_queue = os.environ.get("APP_IN_QUEUE")
logger = logging.getLogger(__name__)

def receive():
    channel, connection = connect()

    channel.queue_declare(queue=receiving_queue)

    channel.basic_consume(queue=receiving_queue,
                      auto_ack=True,
                      on_message_callback=callback)

    channel.start_consuming()

    connection.close()

def callback(ch, method, properties, body):
    try:
        message = json.loads(body)
        logger.info("received message %s" % message)
        ingredient_result = lookup(message.get("searchTerm"))
        ingredient_result.update({"recipeId": message.get("recipeId")})
        send(json.dumps(ingredient_result), message.shard)
    except Exception as e:
        print("error handling mq message %s" % e)
        pass