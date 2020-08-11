# SIMPLE-GRPC

This is a simple project to demonstrate the working of gRPC and how the protobufs are defined to faciliate the Client/Server communication. I have used golang tools and some third party libraries to provide CRUD methods for the client to create a Graph over gRPC and query for its properties exposed via command line tools. The server supports multiple clients.

# Layout
```
|
├── README.md
├── cleanUpLog
├── client
│   └── Dockerfile
├── cmd
│   ├── helloclient
│   │   └── main.go
│   └── helloserver
│       └── main.go
├── dockerCleanupScript.sh
├── go.mod
├── go.sum
├── hellopb
│   ├── hello.pb.go
│   └── hello.proto
├── internal
│   └── graphlib
│       ├── graphlib.go
│       ├── graphlib_test.go
│       └── util.go
├── run.sh
├── runLog
└── server
    └── Dockerfile
```

# How to install
```
Points
1. git clone the repository to your machine. 
2. cd ~/go/src/github.com/vasu81in/simple-grpc
3. Read the run.sh (Note: It cleans all the None/Exited/Dead/Dangling docker images/containers)
4. ./run.sh setup

help() {
  echo "-------------------------------------------------------------------------"
  echo "                      Available commands                                -"
  echo "-------------------------------------------------------------------------"
  echo "   > ./run.sh clean          clean the app docker                        "
  echo "   > ./run.sh setup          clean and sets up app docker                "
  echo "   > help                    Display this help                           "
  echo "-------------------------------------------------------------------------"
}

Once the script runs to completion, just verify three simple-grpc images have been created.
The server should come up and listen on :50051. This is hardcoded.

```
Install logs:
```
:~/go/src/github.com/vasu81in/simple-grpc
runLog
cleanUpLog
```

# Client CLI
```
After the install, client containers are not UP. Get inside them using interactive mode. 

tty1:
docker run --network="simplenet" -it simple-grpc-client1 /bin/bash

tty2:
docker run --network="simplenet" -it simple-grpc-client2 /bin/bash

Just make sure,
"ping simple-grpc-server" between client and server.

root@ee19cd9a58ef:/go# ping simple-grpc-server
PING simple-grpc-server (192.168.208.2) 56(84) bytes of data.
64 bytes from simple-grpc-server.simplenet (192.168.208.2): icmp_seq=1 ttl=64 time=0.092 ms

```
Sample Commands:
```
root@ee19cd9a58ef:/go# helloclient

available commands:
	helloclient create
	helloclient add_edge
	helloclient delete
	helloclient spf
```
create COMMAND:
```
root@ee19cd9a58ef:/go# helloclient create
2020/08/11 02:35:18 Hi Client: ee19cd9a58ef, Creating Graph: Size:10
2020/08/11 02:35:18 Receive response => [GraphID: 52fdfc07-2182-654f-163f-5f0f9a621d72]

```
add_edge COMMAND:
```
root@ee19cd9a58ef:/go# helloclient add_edge

usage: helloclient add_edge <graphID> <vertexA> <vertexB>

Add a Edge to a graph.
       eg, helloclient add_edge --graphID=abcd --vertexA=1 --vertexB=5
Arguments:
	<clientID> ... hostname by default (Optional).
	<graphID>  ... uuid of the graphID.
	<vertexA>  ... vertexA (Mandatory.
	<vertexB>  ... vertexB (Mandatory.


Create the following Graph:
0--1--2
|  |  |
3--4--5

root@ee19cd9a58ef:/go# helloclient add_edge  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=0 --vertexB=1
2020/08/11 02:46:01 Hi Client: ee19cd9a58ef, Graph 52fdfc07-2182-654f-163f-5f0f9a621d72, Adding Edge: [0] -- [1]
2020/08/11 02:46:01 Receive response => [Added: true]
root@ee19cd9a58ef:/go# helloclient add_edge  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=1 --vertexB=2
2020/08/11 02:46:07 Hi Client: ee19cd9a58ef, Graph 52fdfc07-2182-654f-163f-5f0f9a621d72, Adding Edge: [1] -- [2]
2020/08/11 02:46:07 Receive response => [Added: true]
root@ee19cd9a58ef:/go# helloclient add_edge  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=3 --vertexB=0
2020/08/11 02:46:15 Hi Client: ee19cd9a58ef, Graph 52fdfc07-2182-654f-163f-5f0f9a621d72, Adding Edge: [3] -- [0]
2020/08/11 02:46:15 Receive response => [Added: true]
root@ee19cd9a58ef:/go# helloclient add_edge  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=3 --vertexB=4
2020/08/11 02:46:22 Hi Client: ee19cd9a58ef, Graph 52fdfc07-2182-654f-163f-5f0f9a621d72, Adding Edge: [3] -- [4]
2020/08/11 02:46:22 Receive response => [Added: true]
root@ee19cd9a58ef:/go# helloclient add_edge  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=4 --vertexB=1
2020/08/11 02:46:28 Hi Client: ee19cd9a58ef, Graph 52fdfc07-2182-654f-163f-5f0f9a621d72, Adding Edge: [4] -- [1]
2020/08/11 02:46:28 Receive response => [Added: true]
root@ee19cd9a58ef:/go# helloclient add_edge  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=4 --vertexB=5
2020/08/11 02:46:32 Hi Client: ee19cd9a58ef, Graph 52fdfc07-2182-654f-163f-5f0f9a621d72, Adding Edge: [4] -- [5]
2020/08/11 02:46:32 Receive response => [Added: true]
root@ee19cd9a58ef:/go# helloclient add_edge  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=5 --vertexB=2
2020/08/11 02:46:41 Hi Client: ee19cd9a58ef, Graph 52fdfc07-2182-654f-163f-5f0f9a621d72, Adding Edge: [5] -- [2]
2020/08/11 02:46:41 Receive response => [Added: true]

```
spf COMMAND:
```
root@ee19cd9a58ef:/go# helloclient spf

usage: helloclient spf <graphID> <vertexA> <vertexB>

Add a Edge to a graph.
       eg, helloclient spf --graphID=abcd --vertexA=1 --vertexB=5
Arguments:
	<clientID> ... hostname by default (Optional).
	<graphID>  ... uuid of the graphID (Required).
	<vertexA>  ... vertexA (Required).
	<vertexB>  ... vertexB (Required).
  
root@ee19cd9a58ef:/go# helloclient spf  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=0 --vertexB=1
2020/08/11 02:48:59 Hi Client: ee19cd9a58ef, Using GraphID: 52fdfc07-2182-654f-163f-5f0f9a621d72, get SPF between Edges: [0] -- [1]
2020/08/11 02:48:59 Receive response => [Distance: 1]
root@ee19cd9a58ef:/go# helloclient spf  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=0 --vertexB=5
2020/08/11 02:49:03 Hi Client: ee19cd9a58ef, Using GraphID: 52fdfc07-2182-654f-163f-5f0f9a621d72, get SPF between Edges: [0] -- [5]
2020/08/11 02:49:03 Receive response => [Distance: 3]
root@ee19cd9a58ef:/go# helloclient spf  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=0 --vertexB=2
2020/08/11 02:49:06 Hi Client: ee19cd9a58ef, Using GraphID: 52fdfc07-2182-654f-163f-5f0f9a621d72, get SPF between Edges: [0] -- [2]
2020/08/11 02:49:06 Receive response => [Distance: 2]
root@ee19cd9a58ef:/go# helloclient spf  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=0 --vertexB=4
2020/08/11 02:49:08 Hi Client: ee19cd9a58ef, Using GraphID: 52fdfc07-2182-654f-163f-5f0f9a621d72, get SPF between Edges: [0] -- [4]
2020/08/11 02:49:08 Receive response => [Distance: 2]
root@ee19cd9a58ef:/go# helloclient spf  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=4 --vertexB=0
2020/08/11 02:49:20 Hi Client: ee19cd9a58ef, Using GraphID: 52fdfc07-2182-654f-163f-5f0f9a621d72, get SPF between Edges: [4] -- [0]
2020/08/11 02:49:20 Receive response => [Distance: 2]
root@ee19cd9a58ef:/go#

```
delete COMMAND:
```

usage: helloclient delete <graphID> <clientID>

Delete graph.
       eg, helloclient delete --graphID=abcd
Arguments:
	<graphID>  ... uuid of the graph (Required).
	<clientID> ... hostname by default (Required).
  
root@ee19cd9a58ef:/go# helloclient delete --graphID=52fdfc07-2182-654f-163f-5f0f9a621d73
2020/08/11 02:50:55 Hi Client: ee19cd9a58ef, Deleting Graph: 52fdfc07-2182-654f-163f-5f0f9a621d73
2020/08/11 02:50:55 Receive response => [Deleted: false] <========================== INCORRECT graphID given, so the request is not successful

root@ee19cd9a58ef:/go# helloclient delete --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72
2020/08/11 02:51:38 Hi Client: ee19cd9a58ef, Deleting Graph: 52fdfc07-2182-654f-163f-5f0f9a621d72
2020/08/11 02:51:38 Receive response => [Deleted: true]

root@ee19cd9a58ef:/go# helloclient spf  --graphID=52fdfc07-2182-654f-163f-5f0f9a621d72 --vertexA=0 --vertexB=1
2020/08/11 02:51:48 Hi Client: ee19cd9a58ef, Using GraphID: 52fdfc07-2182-654f-163f-5f0f9a621d72, get SPF between Edges: [0] -- [1]
2020/08/11 02:51:48 Receive response => [Distance: 0] <============================== Distance is 0. The graphID is already deleted.
```

# Server Docker Commands
```
for interactive attach,
docker run --network="simplenet" -it simple-grpc-server /bin/bash

for detached, 
docker run -d -p 50051:50051 --network="simplenet" --name hello-grpc-server  hello-grpc-server

```
# Client Docker Commands
```
for interactive attach, 
docker run --network="simplenet" -it simple-grpc-client1 /bin/bash
docker run --network="simplenet" -it simple-grpc-client2 /bin/bash
```

# Known Issues:

- If a server crashes, it doesn't get automatically RESTARTED and LISTEN. As a workaround, 
  just clean up Docker containers or use the run.sh. 
- Client CLI is rudimentary and "Help" may not be so intuitive. 

# Tested Environment:
- Docker version 19.03.12. Macintosh requires grep and awk for scripting.
