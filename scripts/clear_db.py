from gql import gql, Client
from gql.transport.aiohttp import AIOHTTPTransport
import boto3
from botocore.exceptions import ClientError
import logging
import os

# Select your transport with a defined url endpoint
transport = AIOHTTPTransport(url="http://localhost:8081/query")

# Create a GraphQL client using the defined transport
client = Client(transport=transport, fetch_schema_from_transport=True)

try:
    # Register an account
    register = gql("""
        mutation register($email: String!, $password: String!) {
            register(email: $email, password: $password)
        }
        """)

    result = client.execute(register,
                            variable_values={
                                "email": "admin@admin.com",
                                "password": "password"
                            })

    token = result['register']
except:
    # Login to get a token
    login = gql("""
        mutation login($email: String!, $password: String!) {
            login(email: $email, password: $password)
        }
        """)
    result = client.execute(login,
                            variable_values={
                                "email": "admin@admin.com",
                                "password": "password"
                            })
    token = result['login']

# Select your transport with a defined url endpoint
transport = AIOHTTPTransport(url="http://localhost:8081/query",
                             headers={'Authorization': token})

# Create a GraphQL client using the defined transport
client = Client(transport=transport, fetch_schema_from_transport=True)

reset_db = gql("""
    mutation resetDB {
        resetDB
    }
    """)

test = client.execute(reset_db)
