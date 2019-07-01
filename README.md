# Crawler
This is a simple crawler which takes URL and returns all the links from the given site

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/crawler)
[![Build Status](https://travis-ci.org/shamsher31/crawler.svg)](https://travis-ci.org/shamsher31/crawler)
[![GitHub release](http://img.shields.io/github/release/shamsher31/crawler.svg?style=flat-square)](release)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)


## Getting Started

### How to use
If you want to try crawler, you can use crawler docker image to play around with it. Make sure you have [Docker](https://www.docker.com/) installed and running on your system.

How to pull crawler docker image
```docker
docker pull shamsher31/crawler
```

How to run crawler locally
```docker
docker run --rm -p 8080:8080 shamsher31/crawler
```

Use [POSTMAN](https://www.getpostman.com/) to hit the API.
```
localhost:8080/?url=https://www.redhat.com
```

If you don't have postman installed you can try curl command.

```
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET localhost:8080/?url=https://www.redhat.com
```

### Local Development
If you wish to work on Crawler you'll first need [Go](https://www.golang.org/) installed on your machine (version 1.11+ is required).

For local dev first make sure [Go](https://www.golang.org/) is properly installed, including setting up a [GOPATH](https://golang.org/doc/code.html#GOPATH). Ensure that $GOPATH/bin is in your path. Crawler uses [Go Modules](https://github.com/golang/go/wiki/Modules), so it is recommended that you clone the repository outside of the GOPATH.

 Next, clone this repository.
 ```git
 git clone https://github.com/shamsher31/crawler.git
 ```

Once cloned, move into the current directory ie: `cd crawler`.

How to build
```make
make build
```

How to run
```
make run
```

How to run test
```
make test
```

Once you are done testing press `CTR + C` to gracefully shutdown service.

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
