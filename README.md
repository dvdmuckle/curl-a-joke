```
    ____ _    ______  __  _____  __________ __ __    ______
   / __ \ |  / / __ \/  |/  / / / / ____/ //_// /   / ____/
  / / / / | / / / / / /|_/ / / / / /   / ,<  / /   / __/   
 / /_/ /| |/ / /_/ / /  / / /_/ / /___/ /| |/ /___/ /___   
/_____/ |___/_____/_/  /_/\____/\____/_/ |_/_____/_____/   
                                                           
+-+-+-+-+-+-+-+-+
|P|R|E|S|E|N|T|S|
+-+-+-+-+-+-+-+-+
                                                                                                            
  _|_|_|  _|    _|  _|_|_|    _|                _|_|                      _|    _|_|    _|    _|  _|_|_|_|  
_|        _|    _|  _|    _|  _|              _|    _|                    _|  _|    _|  _|  _|    _|        
_|        _|    _|  _|_|_|    _|  _|_|_|_|_|  _|_|_|_|  _|_|_|_|_|        _|  _|    _|  _|_|      _|_|_|    
_|        _|    _|  _|    _|  _|              _|    _|              _|    _|  _|    _|  _|  _|    _|        
  _|_|_|    _|_|    _|    _|  _|_|_|_|        _|    _|                _|_|      _|_|    _|    _|  _|_|_|_|  
                                                                                                            
                                                                                                            
```
[![Build Status](https://travis-ci.org/dvdmuckle/curl-a-joke.svg?branch=master)](https://travis-ci.org/dvdmuckle/curl-a-joke)[![Heroku](https://heroku-badge.herokuapp.com/?app=curl-a-joke)](https://curl-a-joke.herokuapp.com)

The Woz's dial-a-joke, for the 21st century!

```bash
curl https://curl-a-joke.herokuapp.com
```

### Install and Run

```bash
go get -u github.com/dvdmuckle/curl-a-joke
```
Run `$GOPATH/bin/curl-a-joke`

By default `curl-a-joke` will expect the jokes database in your current working directory. This behavior can be changed by passing a path to the jokes database with the `--jokesdb`. This database will be populated automatically, and by default will be populated by jokes in the `jokes.json`, which is also expected in the current working directory. This can be changed by passing a path to the json file with `--jokesjsn`. The default port, 8080, can be changed with the `--port` flag. This flag can also take the form of an environment variable `PORT`. This is mostly for Heroku deployment, and will override the `--port` option if both are present.

You can now `curl` the jokes service. If you're running it on your local machine, `curl localhost:8080` will give you a random joke.

### Dockerize

```bash
docker run -d -p 8080:8080 --name curl-a-joke dvdmuckle/curl-a-joke
```

This will use the default port, `jokes.json`, and database location. If you would like to supply your own `jokes.json`, you can mount a volume with the json,  and use the `--jokesjsn` flag to specify its location within the container.

```bash
docker run -d -p 8080:8080 -v /home/dvdmuckle/curl-a-joke:/root/curl-a-joke --name curl-a-joke dvdmuckle/curl-a-joke --jokejson /root/curl-a-joke/jokes.json
```
The port can also be speicified in a similar manner, however this can be more easily achieved using Docker's `-p` option.

### TODO

* Allow for POSTing new jokes given user authentication
