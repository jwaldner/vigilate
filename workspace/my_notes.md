
# init swarm
docker swarm init 

# deploy
docker stack deploy -c swarm.yml vigilate_app
# to take down the stack if needed

docker swarm leave --force
docker swarm init 



