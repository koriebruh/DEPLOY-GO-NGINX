
docker compose up -d

docker compose down --rmi all

docker stop go-sister-node1-1
docker start go-sister-node1-1

docker stop go-sister-node2-1
docker start go-sister-node2-1
