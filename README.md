# url-shortener

#### Run the Testing

```bash
$ make test
```

#### Run the Applications
Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/budhip/url-shortener.git

#move to project
$ cd url-shortener

# Build the docker image first
$ make docker

# Run the application
$ make run

# check if the containers are running
$ docker ps

# Execute the call for example: users-by-ids
$ curl localhost:9090/:shortcode

# Stop
$ make stop

# Kill all container and image
$ make kill-all
```