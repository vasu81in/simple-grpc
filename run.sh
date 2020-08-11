#!/bin/sh -x

source ./dockerCleanupScript.sh

cleanDocker() {
		cleanNoneImages
		cleanDeadExitedContainers
		cleanDanglingImages
		cleanupUnusedVolume
}

removeSimplenet() {
	net="simplenet"
	echo "\n ======= remove network namespace ${net} ======== \n"
	simplenet=$($DOCKER network ls | $GREP ${net} | $AWK '{print $1}')
	if [[ -z ${simplenet} ]]; then
		echo -n "\n ==== namespace ${net} NOT FOUND ===== \n" >> runLog
	else
		${DOCKER} network rm ${simplenet} >> runLog
		echo -n "\n ==== namespace ${net}:${simplenet} DELETED =====  \n" >> runLog
	fi
}

createSimplenet() {
	net="simplenet"
	echo "\n ======= create network namespace ${net} ======== \n"
	simplenet=$($DOCKER network ls | $GREP ${net} | $AWK '{print $1}')
	if [[ -z ${simplenet} ]]; then  
		${DOCKER} network create -d bridge ${net} >> runLog
		echo -n "\n ==== namespace ${net} CREATED ===== \n" >> runLog
	else
		echo -n "\n ==== namespace ${net} NOT FOUND ====== \n" >> runLog
	fi
}

removeSimpleGrpcServer() {
	name="simple-grpc-server"
	echo "\n ======= remove image/container called ${name} ======== \n"
	match=$($DOCKER images | $GREP ${name} | $AWK '{print $1}')
	if [[ -z ${match} ]]; then  
		echo -n "\ ==== container image ${name} NOT FOUND ===== \n" >> runLog
	else
		${DOCKER} rm -f ${match} >> runLog
		echo -n "\ ==== container ${match} DELETED ===== \n" >> runLog
		${DOCKER} image rm -f ${match} >> runLog
		echo -n "\ ==== container image ${match} DELETED ===== \n" >> runLog
	fi
}

createSimpleGrpcServer() {
	name="simple-grpc-server"
	net="simplenet"
	echo "\n ======= create image/container called ${name} ======== \n"
	match=$($DOCKER images | $GREP ${name} | $AWK '{print $1}')
	if [[ -z ${match} ]]; then  
		echo -n "\ ==== image ${name} NOT FOUND ====== \n" >> runLog
		echo -n "\ ==== creating container image with name ${name} ===== \n" >> runLog
		${DOCKER} build -t ${name} -f server/Dockerfile . >> runLog
		# this may fail if the port is already found to be BOUND!
		${DOCKER} run -d -p 50051:50051 --network=${net} --name ${name} ${name} >> runLog
	else
		echo -n "\ ==== container image ${match} FOUND (already) ===== " >> runLog
		echo -n "\ ==== Continuing to use the image ${match} ===== " >> runLog
	fi
}

createSimpleGrpcClient1() {
	name="simple-grpc-client1"
	net="simplenet"
	echo "\n ======= create image/container called ${name} ======== \n"
	match=$($DOCKER images | $GREP ${name} | $AWK '{print $1}')
	if [[ -z ${match} ]]; then  
		echo -n "\ ==== image ${name} NOT FOUND ====== \n" >> runLog
		echo -n "\ ==== creating container image with name ${name} ===== \n" >> runLog
		${DOCKER} build -t ${name} -f client/Dockerfile . >> runLog
	else
		echo -n "\ ==== container image ${match} FOUND (already) ===== " >> runLog
		echo -n "\ ==== Continuing to use the image ${match} ===== " >> runLog
	fi
}

removeSimpleGrpcClient1() {
	name="simple-grpc-client1"
	echo "\n ======= remove image/container called ${name} ======== \n"
	match=$($DOCKER images | $GREP ${name} | $AWK '{print $1}')
	if [[ -z ${match} ]]; then  
		echo -n "\ ==== container image ${name} NOT FOUND ===== \n" >> runLog
	else
		${DOCKER} image rm -f ${match} >> runLog
		echo -n "\ ==== container image ${match} DELETED ===== \n" >> runLog
	fi
}

createSimpleGrpcClient2() {
	name="simple-grpc-client2"
	net="simplenet"
	echo "\n ======= create image/container called ${name} ======== \n"
	match=$($DOCKER images | $GREP ${name} | $AWK '{print $1}')
	if [[ -z ${match} ]]; then  
		echo -n "\ ==== image ${name} NOT FOUND ====== \n" >> runLog
		echo -n "\ ==== creating container image with name ${name} ===== \n" >> runLog
		${DOCKER} build -t ${name} -f client/Dockerfile . >> runLog
	else
		echo -n "\ ==== container image ${match} FOUND (already) ===== " >> runLog
		echo -n "\ ==== Continuing to use the image ${match} ===== " >> runLog
	fi
}

removeSimpleGrpcClient2() {
	name="simple-grpc-client2"
	echo "\n ======= remove image/container called ${name} ======== \n"
	match=$($DOCKER images | $GREP ${name} | $AWK '{print $1}')
	if [[ -z ${match} ]]; then  
		echo -n "\ ==== container image ${name} NOT FOUND ===== \n" >> runLog
	else
		${DOCKER} image rm -f ${match} >> runLog
		echo -n "\ ==== container image ${match} DELETED ===== \n" >> runLog
	fi
}

clean() {
	cleanDocker
	echo "\n ======= clean up simple-grpc application ======= \n"
	echo -n "\n ======= clean up simple-grpc application ======= \n" >> runLog
	removeSimplenet
	removeSimpleGrpcServer
	removeSimpleGrpcClient1
	removeSimpleGrpcClient2
	echo -n "\n\n ============================ END OF APP CLEANUP ============================= \n\n"
}

setup() {
	clean
	createSimplenet
	createSimpleGrpcServer
	createSimpleGrpcClient1
	createSimpleGrpcClient2
	
	echo -n "\n\n ============================ SERVICES DEPLOYED ============================= \n\n"
}


help() {
  echo "-------------------------------------------------------------------------"
  echo "                      Available commands                                -"
  echo "-------------------------------------------------------------------------"
  echo "   > ./run.sh clean          clean the app docker                        "
  echo "   > ./run.sh setup          clean and sets up app docker                "
  echo "   > help                    Display this help                           "
  echo "-------------------------------------------------------------------------"
}

if [ $# -eq 0 ] ; then
    help
fi

$*
