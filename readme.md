gRPC Hello World
Follow these setup to run the quick start example:

Get the code:

$ go get google.golang.org/grpc/examples/helloworld/greeter_client
$ go get google.golang.org/grpc/examples/helloworld/greeter_server
Run the server:

$ $(go env GOPATH)/bin/greeter_server &
Run the client:

$ $(go env GOPATH)/bin/greeter_client
Greeting: Hello world
For more details (including instructions for making a small change to the example code) or if you're having trouble running this example, see Quick Start.
