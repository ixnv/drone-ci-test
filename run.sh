#!/bin/bash
docker build -t ll-go .
docker run -it --rm --name ll-go-app ll-go
IP=`docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' container_name_or_id`
echo "open http://$IP:8080"
