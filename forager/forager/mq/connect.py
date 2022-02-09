import pika
import os
import time
import logging

RETRY_TIMEOUT = 60 * 2
logger = logging.getLogger(__name__)

def connect():
    param = os.environ.get("CLOUDAMQP_URL")
    if not param:
        param = "amqp://guest:guest@localhost"

    retryIn = 5
    while(True):
        try:
            connection = pika.BlockingConnection(pika.URLParameters(param))
            return connection.channel(), connection
        except pika.exceptions.ConnectionClosedByBroker:
            logger.error("Connection was closed by broker, retrying...")
            time.sleep(retryIn)
            retryIn = 2 * retryIn
            if(retryIn> RETRY_TIMEOUT):
                break
            continue
        except pika.exceptions.AMQPChannelError as err:
            logger.error("Caught a channel error: {}, stopping...".format(err))
            break
        except pika.exceptions.AMQPConnectionError:
            logger.error("Connection was closed, retrying...")
            time.sleep(retryIn)
            retryIn = 2 * retryIn
            if(retryIn> RETRY_TIMEOUT):
                break
            continue
