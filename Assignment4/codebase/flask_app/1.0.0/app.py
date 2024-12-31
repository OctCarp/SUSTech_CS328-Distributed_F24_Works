from flask import Flask, request
import os
import signal
import sys

app = Flask(__name__)

POD_NAME = os.getenv('HOSTNAME', 'unknown')
NODE_NAME = os.getenv('NODE_NAME', 'unknown')
POD_IP = os.getenv('POD_IP', 'unknown')


def graceful_shutdown(signum, frame):
    print("Received shutdown signal, exiting...")
    sys.exit(0)


signal.signal(signal.SIGTERM, graceful_shutdown)


@app.route('/')
def hello():
    return f'Hello! This is server in pod "<{POD_NAME}>" (IP=<{POD_IP}>) from node "<{NODE_NAME}>"!'


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
