#Argyle technical challenge

A fibonacci calculation library has been created in fibonacci/
Tests can be run in the fibonacci directory with `go test -v `.

There is also a benchmark test in the same directory that can be executed with `go test -bench=.`

# Execution
An environment variable (PORT_NUM) is expected by the code, which defines the
port that the HTTP service is listening on.
The code can be executed directly with `PORT_NUM=8000 go run cmd/main.go cmd/routes.go`

# Dockerfile
A Dockerfile has been created that will create a container, the PORT_NUM env var
will still need to be set when the container is run.
`docker build -t fibonacci .`
To run the container locally use `PORT_NUM=8000 docker run fibonacci`

# Kubernetes
fib.yaml has been created, it builds a pod, and deploys it. The PORT_NUM
variable has been set in this file and therefore does not require to be set in
the environment. The file expects the container that it is creating the pod from
to be tagged `fibonacci`,

# Makefile
A simple makefile has been created, it only has a single default target that
assumes that the system is being built locally with minikube as the container
registry.
