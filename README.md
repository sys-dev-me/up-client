# up-client

Hardcoded host & port, should be refactored configuration part

- added gocron, timer in 1 minutes between messaging
- added Dockerfile, sysdevme\go:up-client

Installation

- create subdir /var/spool/up-client
- create settings file ".settings" inside created folder
- create update file for docker container (created by me update.sh) and added content
`
echo "stop container" && docker stop up-client
echo "remove container" && docker rm up-client
docker container prune
docker pull sysdevme/go:up-client
docker run -d -h=`hostname` --restart unless-stopped --volume=/var/spool/up-client:/var/spool/up-client --name=up-client sysdevme/go:up-client

` 
