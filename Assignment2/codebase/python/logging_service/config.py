import os


class Config:
    GRPC_PORT = os.getenv('GRPC_PORT', 50152)

    KAFKA_SEVER = os.getenv("KAFKA_BROKER", "127.0.0.1:19092")
