import json
import time
from typing import Union
import uuid

from fastapi import FastAPI, Request
import pika
import uvicorn
from schemas import OrderRegister
from config import Config

while True:
    try:
        connection = pika.BlockingConnection(
            pika.URLParameters(Config.BROKER_URL))
        break
    except Exception as error:
        print(f"Error {error}")
        time.sleep(4)

app = FastAPI()

channel = connection.channel()

channel.queue_declare(queue='invetory.queue', durable=True)
channel.queue_declare(queue='invetory.queue', durable=True)
channel.queue_declare(queue='shipping.queue', durable=True)
channel.queue_declare(queue='shipping.queue', durable=True)
channel.queue_declare(queue='events_saga.queue', durable=True)

channel.exchange_declare(exchange='EventsProcess', exchange_type='direct', durable=True)
channel.exchange_declare(exchange='EventsSaga', exchange_type='direct', durable=True)


channel.queue_bind('invetory.queue', exchange='EventsProcess', routing_key='invetory.queue')
channel.queue_bind('payment.queue', exchange='EventsProcess', routing_key='payment.queue')
channel.queue_bind('shipping.queue', exchange='EventsProcess', routing_key='shipping.queue')
channel.queue_bind('events_saga.queue', exchange='EventsSaga')


@app.post("/order")
def read_root(request: Request, order_data: OrderRegister):
    
    channel.basic_publish(exchange='EventsSaga', routing_key='events_saga.queue', body=json.dumps({"event_type": "create_order", "message": order_data.model_dump()}))
    
    return{
        "order_uuid": str(uuid.uuid4())
    }

