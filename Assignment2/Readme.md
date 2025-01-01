## Assignment 2

| My Score | Average  | Median |
| -------- | -------- | ------ |
| 100      | 77.02632 | 85.5   |

### Feedback from TA

> 98+20=118 (extra:18)
>
>  
>
> Summary: Almost everything is perfect! Thank you for your hard work, you have made my day!
>
>  
>
> \> API Service
>
> Monitoring:
>
> \1. for the logging service, you configure a go-gin handler which is brilliant, the minor issue here is that the latency calculation does not make sense to me (these two lines are neighbors so how is the latency measured?), but this is not our focus here, so nvm (-0)
>
>  
>
> \> API Service:
>
> YAML:
>
> \1. example is Order is frighteningly mismatched, but it does not matter much (-0)
>
> \2. protected APIs do not necessarily need user_id as input because a user cannot modify other users' info/orders, but your impl ensures this, and actually this is even safer, very good! (-0)
>
>  
>
> \> DB Service
>
> Order:
>
> \1. U R the first ever student that realizes that the product should be updated when order is created/cancelled. Huge applause!
>
>  
>
> \> Compose
>
> Compose file:
>
> \1. submission missing .env, but it's fine (-0)
>
> \2. depends_on cannot solve service dependency, so sometimes db_service will start and fail when postgres is still configuring - need healthcheck (-1)
>
> \3. the inside ad listener addr for kafka should be kafka:9091 but not localhost:9091, and make sure inside listener addr uses the same port=9091 (-1)
>
> DB Init:
>
> \1. In init1.sql, `CREATE USER :"POSTGRES_USER" WITH PASSWORD :'POSTGRES_PASSWORD';` not sure whether these two colons are needed since my postgres container broke on this, but it's fine (-0); also FYI, it seems that a database called "POSTGRES_DB" will be created instead of the value it points to as an env var. I restored it to the initial version and everything works fine.
>
> Dockerfile:
>
> \1. go-impl services has a very nice Dockerfile to squeeze image size
>
> \2. python-impl logging service uses python:3.11 (1.01GB) as the base image, which can be replaced with lighter python:3-slim (120MB) (-0)

### Some Additional Notes from TA

> About Assignment 2, some common questions are answered as follows:
>
> \1 About JWT: JWT (JSON Wen Token) is an authentication approach we can use for RESTful API servers. In general, the user accesses the login API with a password, and receives a token string from the server as the response. Then, when the user wants to access other APIs that require authentication (e.g., the user needs to login first to create an order), he/she needs to manually put this token in the 'Authorization' field of the HTTP request header. 
>
> Check the README file in the JWT demo for more info. Check https://jwt.io/ as well if you want to understand how JWT is structured. 
>
> (FYI: Why use JWT? Because the servers do not need to store JWTs to maintain user "sessions", so we can freely introduce server replicas for load balancing.)
>

> \2. About RESTful API Service: API Service is the only service exposed to the external world. The client/user interacts ONLY with API Service for all operations. 
>
> On the other hand, API Service never DIRECTLY interacts with PostgreSQL DB or Kafka topic. It needs to use gRPC DB Service to play with the database, and use gRPC Logging Service to produce logs to the Kafka topic. Please check the arrows in the architecture figure carefully. 
>
> (For example, when a user wants to retrieve the information of a product, he/she might access the GET /products/{id} RESTful API, then this API might use the DB Service stub to remotely invoke the DBService.GetProduct RPC - this RPC will actually connect to PostgreSQL DB and execute the corresponding SQL query.)

> \3. About Logs: Logging Service provides a client-side streaming RPC for API Service to send log messages occasionally. 
>
> You may design by yourself the log format and the triggering condition. For example, we might generate a new log message upon every API call (e.g., a log message like "GET /products requested by user 12 at 2024-11-27 22:25:00"). We might also record errors like "POST /login: user 17 failed due to password mismatch", and "POST /orders/create: user 21 failed because product 2 is out of stock". 
>
> Remember to briefly describe in the report what your logs look like and when they will be sent to Logging Service from the API Service.

> \4. About Docker Compose: You need to build the docker images for all three implemented services and add them to the Docker Compose file. 
>
> Check the load_balancing codebase in the reverse proxy lab, which gives an example of putting a RESTful API server in the compose file. 
>
> NOTE that if you test the servers on localhost, you find postgres via 'localhost:5432' and kafka via 'localhost:9093'. This is because postgres and kafka containers inside the docker compose network have exposed these ports in the compose file. However, when you put everything in the same docker compose network, you need to change these addresses to the correct '{container_hostname}:{port}'.

