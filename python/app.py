import grpc
import time
import requests
from concurrent import futures

# import the generated classes
import user_pb2
import user_pb2_grpc


# create a class to define the server functions, derived from
class UserServicer(user_pb2_grpc.UserServiceServicer):

    def GetUser(self, request, context):
        print(f"Receive message from client: {request.username}")

        req = requests.get(f"https://api.github.com/users/{request.username}")
        response = user_pb2.UserResponse()

        if req.ok:
            body = req.json()
            response.id = body['id']
            response.name = body['name']
            response.username = body['login']
            response.avatarurl = body['avatar_url']
            response.location = body['location']
            response.repos = body['public_repos']

        return response


# create a gRPC server
server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

# to add the defined class to the server
user_pb2_grpc.add_UserServiceServicer_to_server(UserServicer(), server)

print('Starting server. Listening on port 8300.')
server.add_insecure_port('[::]:8300')
server.start()

# since server.start() will not block,
# a sleep-loop is added to keep alive
try:
    while True:
        time.sleep(86400)
except KeyboardInterrupt:
    server.stop(0)
