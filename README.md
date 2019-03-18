# SETTING UP DOCKER

Steps :

1) Install Docker in machine ("https://docs.docker.com/docker-for-windows/install/")
2) Create a folder in you machine
3) CD into the folder
4) Create file with name "Dockerfile"(D should be caps).
5) Copy the any of the Dockerfile from the repository.
6) Run the following command in Docker
$ docker build -t myTestBox .
$ docker run myTestBox
$ docker attach myTestBox

# DOCKER

Enterprise development and teams that run business-critical containerized applications and services in production at large scales.

Provides a support mechanism for issues with the underlying engine as well as deployment challenges using it.

When using Docker EE on certified platform, organizations are assured through Docker certification that their applications will work as expected and have support available if they do not.

# Enterprise edition comes in
- Basic
   - Platform for certified infrastructure, containers and plugins with support from docker.
- Standard
   - Adds advanced image and container management, LDAP and RBAC  
- Advanced tiers
   - Adds security scanning and vulnerability monitoring

# Docker EE includes
- Docker engine with support
- Docker trusted registry
- Universal Control Plane
# Compatibility
- Docker Engine 17.06+
- DTR 2.3+

# Docker installation : Centos/Red Hat
  $ yum install -y yum-utils device-mapper-persistent-data lvm2 (device mapper for storage sub system
  Community edition ce-repo)
  $ yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce-repo

  Enterprise Edition EE -repo    
    $ yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ee-repo    
    $ sudo yum update      
    $ sudo yum install docker-ce
    $ systemctl enable docker && systemctl start docker && systemctl status docker

# Allow docker command as non root
  $ usermod -aG docker user  
  $ exit (exit terminal and login back)  
  $ ssh user@tcox1.mylabserver.com

# Docker install: Debain/ubuntu
  $ apt-get install apt-transport-https ca-certificates curl software-properties-common
  $ curl -fsSL https://download.docker.com/linux/ubuntu/gpg  | sudo apt-key add -  
  $ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu
  $(lsb_release -cs) stable"  
  $ apt-get update  
  $ apt-get install docker-ce  
  $ systemctl status  
  $ cd /var/run  
  $ usermod -aG docker user

# Selecting storage driver:
  $ docker info | grep storage
  $ cd /etc/docker  
  key.json -> demon.json  
  {
    "storage driver" : "devicemapper"  
  }

# View the location where the device is
  $ cd /car/lib/docker
  $ cd devicemapper

# Configuring Logging Drivers :
  $ docker info | grep logging
  $ docker container inspect | grep IPAddr
  $ docker logs <container-name>
  $ cd /etc
  $ vim rsyslog.conf
  $ systemctl start rsyslog

  Setting logging to all the docker
  daemon.json
  {
    "log-driver" : "syslog"
    "log-opts" :{
      "syslog-address" : "udp://172.31.125.216:514"   
    }
  }

# Setting logging for particular container
  $ docker container run -d --name testjson --log-driver json-file httpd

# DOCKER SWARM
  Check the ip by typing ifconfig
  $ ifconfig
  $ docker swarm init --advertise-addr 172.31.4.218
Once swarm manager is initialized copy the Token created. Token is important for additional nodes you can retrieve the token by using command
  $ docker swarm join-token worker
Adding additional manager, this will get new token  
  $ docker swarm join-token manager
Display the state of our manager
  $ docker node ls
  $ docker system info | more

# Swarm back up and restore
# Swarm cluster has

  $ docker service create --name bkupweb --publish 80:80 --replicas 2 httpd
  - three nodes 2 running all the time
  - Backup the manager
      - stop the docker first
      - go to cd /var/lib/docker --> cd swarm
      - json files, states of swarm/certificates

# create tarfile
  $ tar cvf  swarm.tar swarm/

# Docker pull commands
  $ docker pull hellow-world - will pull the latest image
  $ docker pull -a hellow-world - will pull all the tagged images present
  $ docker pull --disable-content-trust hellow-world - pull un - trusted  contents
  $ docker images - list out the images
                  - Repository name
                  - When it was created
                  - Size
                  - Image ID

  $ docker run hellow-world
  $ docker images --digest

# Digest and Image ID are two different ways of addressing the docker.
  $ docker images --filter "before=centos"
  $ docker images --filter "since=centos"
  $ docker images --no-trunc - will show full image id without truncation.
  $ docker images --q - only shows the image id's

# Searching an image repository
  $ docker search apache
  $ docker search apache | wc -l  gives the count
  $ docker search --filter stars=50 apache
  $ docker search --filter stars=50 --filter is-official=true apache

# Tagging an image
  $ docker tag centos:6 mycentos:v1

# Image layers:
$ docker image history <image name>
union file system - Union file system allows files and directories of separate files system or branches to be overlaid so that they form a single file
$ docker image history <image name> --no-trunc
To see image layer file
$ cd /var/lib/docker

# Docker CLI commands
$ docker image  -  will give the list of options
- build - history - import - inspect - load - ls - prune - pull - push - rm - save - tag
To package the image into a Tar file
$ docker image save myrepo/mycentos:ver2 > mycentos.custom.tar
$ tar tvf mycentos.custom.tar | more
$ docker import mycentos.custom.tar  localimport:centos6

# Docker inspect
  $ docker image inspect centos:6
  $ docker image inspect centos:6 > centos.output
  $ docker image inspect centos:6 | grep ContainerConfig
  $ docker image inspect centos:6 --format '{{ .ContainerConfig }}'
  $ docker image inspect centos:6 --format '{{ json.ContainerConfig }}'
  $ docker image inspect centos:6 --format '{{ .ContainerConfig }}'
  $ docker image inspect centos:6 --format '{{ .RepoTags }}'

# Dockerfile   
  - create directory  
  - create Dockerfile <- D should be caps
  - FROM/LABEL/RUN
  - cmd -->  $ docker build -t customimage:v1 .
  - Each instruction in docker file create one layer
  - use cache while rebuilding unless its stated not to do so
  - specify specify file to act as docker file --> $ docker build -t customimage:v1 -f Dockerfile2 .
  - "." represent the context
  - always pull from the repository and merge the layer into single layer -> $ docker build --pull --no-cache --squash -t optimized:v1 .
  - "squash" will help in merging of the layers.

# Docker Life cycle restart policies :
  $ docker run -d --name testweb httpd
  $ docker stop testweb
  $ docker start testweb
  $ sudo systemctl restart docker
  $ docker container run -d --name testweb --restart none
  $ docker container run -d --name testweb --restart always
  - automatically restart
  $ docker container run -d --name testweb --restart unless-stopped httpd
  -
  $ docker container run -d --name testweb --restart on-failure
  - restart

# Docker - running / attaching to and executing commands
  $ docker images
  $ docker ps - list the container up and running
  $ docker ps -a  -list all the container
  $ docker run centos:6
      '-d' detached mode
      '-i' interactive mode
      '-t' current terminal root
      '--name' name the container
      '--rm' removes the container once exited
      '--privilege'
      '--env MYVAR=whatever' pass parameter

# Dockerfile Options
- Comment using #
- FROM
  - Available in public or Private repositories
- MAINTAINER (obsolete)
- LABEL  
     - LABEL maintainer= "manoj.chandran@oouve.com"
- RUN
     - Shell form and json format
     - json array
     - RUN ["yum", "install"]
     - RUN yum update -y && install httpd net-tools
- CMD
     - CMD echo "remember to check container IP address"
     - when multiple CMD commands available only last one will run
- ENV ENVIRONMENT="PRODUCTION"
- EXPOSE
   - EXPOSE 80
- ENTRYPOINT
   - Configure the container
   - Eg . ENTRYPOINT apachectl ".DFOREGROUND"
- ARG

# Non configurable
    Raft consensus between manager nodes
    Gossip protocol for overlay networking
    etcd
    RethinkDB
    Stand-alone swarm
# SWARM Cluster management:
- Discovery
- Scheduling
      - Filters
          - Ports
          - Health Check
          - Constraints
          - Affinities
      -  Strategy
          - Binpack - Fill and move to next engine
          - Random - Create randomly in different engine available
          - Spread  - Create evenly across the Docker engine
- High Availability
      - Among Swarm manager check the storage (Console/etcd/zookeeper) emerge a leader swarm
      - Other manager will be notified  and starts communicating with the leader
      - When leader manager node fail, all manager starts contacting the storage. Will select a new leader
- Networking
  - Docker network command
  - Network of type overlay
  -
-
- Volumes

$ docker volume create my-mount
$ docker volume ls
$ cd /var/lib/docker/volumes/my-mount/

$ docker network ls - will give the network details
$ docker network inspect bridge | more
$ docker container inspect testweb | grep IPAdd
$ docker container inspect --format="{{.NetworkSettings.Netwroks.bridge.IPAddress}}" testweb

# Docker DNS
  $ cat /etc/resolve.conf



# Quorum
The minimum number of members of an assembly or society that must be present at any of its meetings to make the proceedings of that meeting valid.

  $ docker rm kind_boyd
  $ docker rm $(docker ps -a -q)
  UCP and DTR

Signing an Image:
Signing an image is turning on Docker Content Trust.
