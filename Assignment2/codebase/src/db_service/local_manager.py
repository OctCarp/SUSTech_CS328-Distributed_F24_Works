import os

from dotenv import load_dotenv
import psycopg2
from psycopg2 import pool  # this line is oddly necessary
from pprint import pprint
"""
Connection pooling: https://www.psycopg.org/docs/pool.html
Connection: https://www.psycopg.org/docs/connection.html
Cursor: https://www.psycopg.org/docs/cursor.html
"""

load_dotenv()

# Create a SimpleConnectionPool for a single-threaded application
simple_pool = psycopg2.pool.SimpleConnectionPool(
    minconn=1,
    maxconn=10,
    user=os.getenv("POSTGRES_USER"),
    password=os.getenv("POSTGRES_PASSWORD"),
    host="localhost",
    port="5432",
    database=os.getenv("POSTGRES_DB"),
)

print(os.getenv("POSTGRES_USER"))

# Single-threaded usage of connections
try:
    # Get a connection from the pool
    conn = simple_pool.getconn()

    # Perform a database operation
    with conn.cursor() as cursor:
        cursor.execute("SELECT * FROM products;")
        results = cursor.fetchall()
        pprint(results)
        # Commit only for INSERT/UPDATE/DELETE operations.
        # cursor.execute("INSERT INTO products (name, stock, price) VALUES (%s, %s, %s) RETURNING id", ("SUSTech Pixel Map", 300, 3.99))
        # results = cursor.fetchone()
        # pprint(results)
        # conn.commit()
        # Rollback when sth is wrong.
        # conn.rollback()
finally:
    # Return the connection to the pool
    simple_pool.putconn(conn)

# Closing all connections in the pool
simple_pool.closeall()
