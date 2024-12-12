# Flask Server App - Instruction

Refer to the `app.py` file in the Reverse Proxy - Load Balancing example. Update the root API to return a hello message including the pod name, pod IP, and the node name.

An example template is as follow:
```text
# curl <IP>:<PORT>/
Hello! This is server in pod "<POD_NAME>" (IP=<POD_IP>) from node "<NODE_NAME>"!
```

The Flask server needs to support graceful shutdown. Refer to [this post](https://cloud.google.com/blog/products/containers-kubernetes/kubernetes-best-practices-terminating-with-grace).

Build the Flask server into a light docker image, using `requirements.txt` and `Dockerfile`.

Later, add the `greet-with-info` API from the RESTful API - Hello World example. Build a new image with a new tag (version).

An example:
```text
# curl <IP>:<PORT>/chat/peter?institution=sustech
{"message":"Hello peter from sustech!"}
```
