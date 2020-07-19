# sample-gorm-gin

# gin-gonic
Gin is a web framework written in Go (Golang). It features a martini-like API with much better performance, up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
In this section we will walk through what Gin is, what problems it solves, and how it can help your project.
Or, if you are ready to use Gin in to your project, visit the Quickstart.

## Features

* Fast Radix tree based routing, small memory foot print. No reflection. Predictable API performance.

* Middleware support An incoming HTTP request can be handled by a chain of middlewares and the final action. For example: Logger, Authorization, GZIP and finally post a message in the DB.

* Crash-free Gin can catch a panic occurred during a HTTP request and recover it. This way, your server will be always available. As an example - it’s also possible to report this panic to Sentry!

* JSON validation Gin can parse and validate the JSON of a request - for example,checking the existence of required values.

* Routes grouping Organize your routes better. Authorization required vs non required, different API versions… In addition, the groups can be nested unlimitedly without degrading performance.

* Error management Gin provides a convenient way to collect all the errors occurred during a HTTP request. Eventually, a middleware can write them to a log file, to a database and send them through the network.

* Rendering built-in Gin provides an easy to use API for JSON, XML and HTML rendering.

* Extendable Creating a new middleware is so easy, just check out the sample codes.

## Install gin
```sh
 go get -u github.com/gin-gonic/gin
```

# GORM Guides
The fantastic ORM library for Golang, aims to be developer friendly.

## Overview
* Full-Featured ORM (almost)
* Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism)
* Hooks (Before/After Create/Save/Update/Delete/Find)
* Preloading (eager loading)
* Transactions
* Composite Primary Key
* SQL Builder
* Auto Migrations
* Logger
* Extendable, write Plugins based on GORM callbacks
* Every feature comes with tests
* Developer Friendly

## Install GORM
```sh
go get -u github.com/jinzhu/gorm
```

## insomnia test
 import file Insomnia_2020-07-15.json in insomnia for test rest 