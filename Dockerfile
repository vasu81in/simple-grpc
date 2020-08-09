FROM golang

ENV GO111MODULE=on
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV GOBIN $GOPATH/bin
ENV PATH $GOBIN:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
RUN git clone https://github.com/vasu81in/simple-grpc.git
ADD . /go/src/github.com/vasu81in/simple-grpc

#RUN cp -rf simple-grpc /go/src/github.com/vasu81in/
