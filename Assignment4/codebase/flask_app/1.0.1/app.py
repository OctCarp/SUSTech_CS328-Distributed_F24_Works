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


@app.route('/', methods=['GET'])
def hello():
    return f'Hello! This is server in pod "<{POD_NAME}>" (IP=<{POD_IP}>) from node "<{NODE_NAME}>"!'


@app.route('/chat/<username>', methods=['GET'])
def greet_with_info(username):  # retrieve username from URL path
    # retrieve institution from URL query
    institution = request.args.get('institution', None)
    institution_segment = f' from {institution}' if institution else ''
    msg = f'Hello {username}{institution_segment}!'
    return {'message': msg}, 200


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
