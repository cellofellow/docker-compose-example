Now we have a bare-basic Flask app and a Golang client for that app.

It has a docker-compose.yml file that specifies how to run the Flask image and
also its dependencies, a Redis container, and the Go client.

Now also with an "aliases" file you can use with "source aliases". Provides
some handy tools to interface with your containers.

* "docker-compose" shortened to "dc" cuz typing.
* "shell" which works like "shell web" to open a bash shell inside a running
  container.
* "container_ip" to find out the IP address of a running container.
