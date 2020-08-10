# Pull latest Golang 1.14
FROM golang

RUN apt-get update && apt-get install vim -y 

# Set the env 
ENV GO111MODULE=auto
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV GOBIN $GOPATH/bin
ENV PATH $GOBIN:$PATH

# Set the go path/bin
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

# Copy the pwd
ADD . /go/src/github.com/vasu81in/simple-grpc

# Get the source from GitHub
RUN go get google.golang.org/grpc

# Install protoc-gen-go
RUN go get github.com/golang/protobuf/protoc-gen-go

# Install APP server
RUN go install /go/src/github.com/vasu81in/simple-grpc/cmd/helloclient

# Set the ENTRY for the APP
#ENTRYPOINT ["/go/bin/helloclient"]

