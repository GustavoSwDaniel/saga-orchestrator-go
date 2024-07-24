import os


class Config:
    BROKER_URL = os.getenv('BROKER_URL', 'amqp://admin:admin@localhost:5672')
