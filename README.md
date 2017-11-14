# curl-a-joke
[![Build Status](https://travis-ci.org/dvdmuckle/curl-a-joke.svg?branch=master)](https://travis-ci.org/dvdmuckle/curl-a-joke)![Heroku](https://heroku-badge.herokuapp.com/?app=curl-a-joke)

The Woz's dial-a-joke, for the 21st century!

```bash
curl https://curl-a-joke.herokuapp.com
```

### Install and Run

```bash
go get -u github.com/dvdmuckle/curl-a-joke
```
Run `$GOPATH/bin/curl-a-joke`

By default `curl-a-joke` will expect the jokes database in your current working directory. This behavior can be changed by passing a path to the jokes database with the `--jokesdb` flag. If you create a new jokes database, be sure to follow the same schema as the example jokes database. The default port, 8080, can be changed with the `--port` flag. This flag can also take the form of an environment variable `PORT`. This is mostly for Heroku deployment, and will override the `--port` option if both are present.

You can now `curl` the jokes service. If you're running it on your local machine, `curl localhost:8080` will give you a random joke.

### Dockerize

```bash
docker run -d -p 8080:8080 --name curl-a-joke dvdmuckle/curl-a-joke
```

This will use the default port and jokes database. If you would like to supply your own jokes database, you can mount a volume with the database,  and use the `--jokesdb` flag to specify its location within the container.

```bash
docker run -d -p 8080:8080 -v /home/dvdmuckle/curl-a-joke:/root/curl-a-joke --name curl-a-joke dvdmuckle/curl-a-joke --jokesdb /root/curl-a-joke/jokes.db
```
The port can also be speicified in a similar manner, however this can be more easily achieved using Docker's `-p` option.

### TODO

* Parse jokes from json and automigrate to jokes database

* Allow for easy addition of jokes via a command line option

* Allow for POSTing new jokes given user authentication
