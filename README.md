# curl-a-joke
[![Build Status](https://travis-ci.org/dvdmuckle/curl-a-joke.svg?branch=master)](https://travis-ci.org/dvdmuckle/curl-a-joke)

The Woz's dial-a-joke, for the 21st century!

### Design Specs

 * Use HTTP server to serve up jokes

 * Use gorm to access jokes from an sqlite database

 * Use command line arguments for configuring port and location of sqlite db

 * If no database is specified, the default will be used
