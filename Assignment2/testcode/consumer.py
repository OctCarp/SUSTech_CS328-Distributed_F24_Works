import os
from datetime import datetime

from confluent_kafka import Consumer
import json

from config import Config


def save_log_to_file(log_data):
    # Create logs directory if not exists
    log_dir = "logs"
    if not os.path.exists(log_dir):
        os.makedirs(log_dir)

    # Generate filename with current date (format: YYYYMMDD_HH)
    filename = os.path.join(log_dir, f"log_{datetime.now().strftime('%Y%m%d_%H')}.txt")

    # Append log entry with timestamp to file
    with open(filename, 'a', encoding='utf-8') as f:
        f.write(f"{datetime.now().isoformat()}: {json.dumps(log_data, indent=2, ensure_ascii=False)}\n")


def main():
    # Kafka consumer configuration
    consumer_port = Config.CONSUMER_PORT
    config = {
        'bootstrap.servers': f'localhost:{consumer_port}',
        'group.id': 'log-consumer-group',
        'auto.offset.reset': 'earliest'
    }

    # Initialize Kafka consumer and subscribe to log channel
    consumer = Consumer(config)
    consumer.subscribe(['log-channel'])

    try:
        while True:
            # Poll for messages with 1 second timeout
            msg = consumer.poll(1.0)
            if msg is None:
                continue
            if msg.error():
                print(f"Consumer error: {msg.error()}")
                continue

            try:
                # Parse and save received log message
                log_data = json.loads(msg.value().decode('utf-8'))
                save_log_to_file(log_data)
                print(f"Received log: {log_data}")
            except Exception as e:
                print(f"Error processing message: {e}")

    except KeyboardInterrupt:
        pass
    finally:
        # Ensure proper cleanup of consumer
        consumer.close()


if __name__ == '__main__':
    main()
