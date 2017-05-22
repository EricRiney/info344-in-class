#!/usr/bin/env bash

#TODO: 
# - create private network,
# - run hellosvc in network with no published ports
# - run gateway in network publishing port 443
#   and using volumes to give it access to your
#   cert and key files in `gateway/tls`
#   and setting environment variables
#    - CERTPATH = path to cert file in container
#    - KEYPATH = path to private key file in container
#    - HELLOADDR = net address of hellosvc container
if [ -z "$(docker network ls -q -f name=msprivnet)" ]
then
    docker network create msprivnet
fi 

docker run -d \
--name hellosvc1 \
--network msprivnet \
drstearns/hellosvc

docker run -d \
--name hellosvc2 \
--network msprivnet \
drstearns/hellosvc

docker run -d \
--name gateway \
--network msprivnet \
-p 443:443 \
-v $(pwd)/gateway/tls:/tls:ro \
-e CERTPATH=/tls/fullchain.pem \
-e KEYPATH=/tls/privkey.pem \
-e HELLOSVCADDR=hellosvc1,hellosvc2 \
drstearns/gateway
