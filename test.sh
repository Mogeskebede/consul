#!/bin/bash

set -euo pipefail

unset CDPATH

cd "$(dirname "$0")"

readonly COMMAND="${1:-}"

old_version='hashicorp/consul:1.10.8'
new_version='hashicorp/consul:1.14.8'

docker rm -f server1 || true
docker rm -f server2 || true
docker rm -f server3 || true

docker run --name server1 -p 8501:8500 -d $new_version agent -hcl '
server = true
bootstrap = true
client_addr = "0.0.0.0"
'

ip_server_1="$(docker inspect server1 | jq -r '.[0].NetworkSettings.IPAddress')"
echo "ip of server1 is: $ip_server_1"

docker run --name server2 -p 8502:8500 -d $new_version agent -hcl '
server = true
client_addr = "0.0.0.0"
retry_join = ["'$ip_server_1'"]
'

docker run --name server3 -p 8503:8500 -d $old_version agent -hcl '
server = true
retry_join = ["'$ip_server_1'"]
client_addr = "0.0.0.0"
'

docker exec -it server1 consul members -detailed