from gql import gql, Client
from gql.transport.aiohttp import AIOHTTPTransport
import boto3
from botocore.exceptions import ClientError
import logging
import os

S3_BUCKET = 'uploads-76078f4'
UNIT_NAME = "Unit 2"
CLASS_NAME = "Class 1"
ASSIGNMENT_NAME = "Assignment 1"

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

# S3 Client
s3 = boto3.client('s3')


def upload_file(file_name, bucket, object_name=None):
    """Upload a file to an S3 bucket

    :param file_name: File to upload
    :param bucket: Bucket to upload to
    :param object_name: S3 object name. If not specified then file_name is used
    :return: True if file was uploaded, else False
    """

    # If S3 object_name was not specified, use file_name
    if object_name is None:
        object_name = os.path.basename(file_name)

    # Upload the file
    s3_client = boto3.client('s3')
    try:
        response = s3_client.upload_file(file_name, bucket, object_name)
    except ClientError as e:
        logging.error(e)
        return False
    return True


# Upload whole directory to s3, preserving the directory structure
def upload_directory(dir_name, bucket, prefix):
    studentID = os.path.basename(dir_name)
    for subdir, dirs, files in os.walk(dir_name):
        for file in files:
            full_path = os.path.join(subdir, file)
            s3_path = f'{prefix}/{os.path.join(studentID, os.path.relpath(full_path, dir_name))}'
            upload_file(full_path, bucket, s3_path)


# Create a unit
create_unit = gql("""
    mutation createUnit($name: String!) {
        createUnit(input: {name: $name}) {
            id
        }
    }
    """)

# Create a class
create_class = gql("""
    mutation createClass($name: String!, $unitID: ID!) {
        createClass(input: {name: $name, unitID: $unitID}) {
            id
        }
    }
    """)

# Create an assignment
create_assignment = gql("""
    mutation createAssignment($name: String!, $classID: ID!, $dueDate: Int!) {
        createAssignment(input: {name: $name, classID: $classID, dueDate: $dueDate}) {
            id
        }
    }
    """)

# Create a submission
create_submission = gql("""
    mutation createSubmission($assignmentID: ID!, $studentID: String!) {
        createSubmission(input: {assignmentID: $assignmentID, studentID: $studentID}) {
            id
        }
    }
    """)

# Create a test
create_test = gql("""
    mutation createTest($name: String!, $assignmentID: ID!) {
        createTest(input: {name: $name, assignmentID: $assignmentID}) {
            id
        }
    }
    """)

# Create a unit
unit = client.execute(create_unit, variable_values={"name": UNIT_NAME})
unit_id = unit['createUnit']['id']
# Create a class
class_ = client.execute(create_class,
                        variable_values={
                            "name": CLASS_NAME,
                            "unitID": unit_id
                        })
class_id = class_['createClass']['id']
# Create an assignment
assignment = client.execute(create_assignment,
                            variable_values={
                                "name": ASSIGNMENT_NAME,
                                "classID": class_id,
                                "dueDate": 1620000000
                            })
assignment_id = assignment['createAssignment']['id']
# Create a test
test = client.execute(create_test,
                      variable_values={
                          "name": "Test 1",
                          "assignmentID": assignment_id
                      })
test_id = test['createTest']['id']

# Upload the test file
upload_file('./data/Test.java', S3_BUCKET,
            f'{UNIT_NAME}/{ASSIGNMENT_NAME}/Tests/{test_id}/Test.java')

# Upload all project directories
for dir in os.listdir('./data/submissions'):
    # Create a submission
    submission = client.execute(create_submission,
                                variable_values={
                                    "name": "Submission 1",
                                    "assignmentID": assignment_id,
                                    "studentID": dir
                                })
    submission_id = submission['createSubmission']['id']

    upload_directory(f'{os.path.join("./data/submissions", dir)}', S3_BUCKET,
                    f'{UNIT_NAME}/{ASSIGNMENT_NAME}/Projects')

# Execute the query on the transport
# result = client.execute(query)
# print(result)