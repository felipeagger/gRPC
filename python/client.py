import grpc

# import the generated classes
import user_pb2
import user_pb2_grpc

# open a gRPC channel
channel = grpc.insecure_channel('0.0.0.0:8300')

# create a stub (client)
stub = user_pb2_grpc.UserServiceStub(channel)

# create a valid request message
requestMsg = user_pb2.UserRequest(username="maria")

# make the call
response = stub.GetUser(requestMsg)

print(type(response))
print(response)
