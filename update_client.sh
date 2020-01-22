#linux

echo "stop container" && docker stop up-client
echo "remove container" && docker rm up-client
docker container prune
docker pull sysdevme/go:up-client
docker run -d -e "GATEWAY=$(curl -s ifconfig.me)" -e "DOCKER_HOST=$(ifconfig $1 | grep -Po -i -w 'inet (.*)(\d.{1,4})+$' | awk '{print $2}')" -h=`hostname` --restart unless-stopped --volume=/var/spool/up-client:/var/spool/up-client --name=up-client sysdevme/go:up-client


#mac
#echo "stop container" && docker stop up-client
#echo "remove container" && docker rm up-client
#docker container prune
#docker pull sysdevme/go:up-client
#docker run -d -e "GATEWAY=$(curl -s ifconfig.me)" -e "DOCKER_HOST=$(ifconfig $1 | grep -Eo -i -w 'inet (.*)(\d.{1,4})+$' | awk '{print $2}')" -h=`hostname` --restart unless-stopped --volume=/var/spool/up-client:/var/spool/up-client --name=up-client sysdevme/go:up-client
