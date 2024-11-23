from gc import callbacks

import grpc
from concurrent import futures
from confluent_kafka import Producer
import json

from config import Config
import glog_pb2
import glog_pb2_grpc


class LoggingService(glog_pb2_grpc.LoggingServiceServicer):
    def __init__(self, kafka_config):
        self.producer = Producer(kafka_config)
        self.topic = "log-channel"  # kafka topic

    def publish_to_kafka(self, log_message):
        try:
            message_dict = {
                'level': glog_pb2.LogLevel.Name(log_message.level),
                'service_name': log_message.service_name,
                'message': log_message.message,
                'timestamp': log_message.timestamp,
                'trace_id': log_message.trace_id
            }

            print(json.dumps(message_dict))
            self.producer.produce(
                self.topic,
                value=json.dumps(message_dict).encode('utf-8'),
            )
            self.producer.poll(0)

        except Exception as e:
            print(f"Error publishing to Kafka: {str(e)}")
            return False
        return True

    def delivery_report(self, err, msg):
        if err is not None:
            print(f'Message delivery failed: {err}')
        else:
            print(f'Message delivered to {msg.topic()} [{msg.partition()}]')

    def StreamLogs(self, request_iterator, context):
        try:
            # Iterate through each log message sent by the client
            for log_message in request_iterator:
                # Publish each log to Kafka
                success = self.publish_to_kafka(log_message)

                # If publishing fails, return error response immediately
                if not success:
                    return glog_pb2.LogResponse(
                        success=False,
                        message="Failed to publish logs to Kafka"
                    )

            # After processing all messages, ensure they are sent to Kafka
            self.producer.flush()

            # Return success response
            return glog_pb2.LogResponse(
                success=True,
                message="Successfully processed all logs"
            )

        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f'An error occurred: {str(e)}')
            return glog_pb2.LogResponse(
                success=False,
                message=f"Error processing logs: {str(e)}"
            )


def serve():
    kafka_config = {
        'bootstrap.servers': Config.KAFKA_SEVER,
        'client.id': 'logging-service'
    }

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    glog_pb2_grpc.add_LoggingServiceServicer_to_server(
        LoggingService(kafka_config), server
    )

    server.add_insecure_port(f"[::]:{Config.GRPC_PORT}")
    server.start()
    print(f"Logging Service started on port {Config.GRPC_PORT}")

    try:
        server.wait_for_termination()
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
