# Procd

System resource and profile information emission. A simple method of exposing system resources statistics through standard interfaces. Current interfaces are stdout, JSON over HTTP, and as a client to [Mozilla Hekad](https://github.com/mozilla-services/heka).

While the Heka client will accept transport of TCP or UDP, though UDP is recommended for system metric data.

# Installation

You can utilize the scripts/build.sh to generate a binary assuming Go is installed.

Manual build & installation steps:

- Install Go v1.2+
- Install Git
- Install Mercurial (Needed to download the UUID package from Google)
- git clone https://github.com/bellycard/procd.git
- go get github.com/bellycard/toml
- go get code.google.com/p/go-uuid/uuid
- go get github.com/mozilla-services/heka/client
- go get github.com/mozilla-services/heka/message
- go build
- ./procd -config=procd.toml


# Configuration

Configuration of Procd is based on a specified [TOML](https://github.com/mojombo/toml) configuration file. A reasonable default configuration is in /conf/procd.toml. Defined outputs will have metrics and profile information emitted if specified.

### Example Configuration

This configuration will emit metric information in JSON format to STDOUT, localhost:5596, and to a local Hekad agent.

```
ticker_interval = 5 # Time (seconds) to poll resources and collect metrics

[output.stdout]

[output.http]
bind_address = "0.0.0.0:5596"

[output.heka]
server = "127.0.0.1:5565"
sender = "udp"
payload = false
hostname = "super.coolhost.com" # Optional. Overwrites os.Hostname() for Heka messages.
```


# Usage

### HTTP

To view all resources: ```GET http://hostname:5596/v1/resources.json```
To view CPU resources: ```GET http://hostname:5596/v1/resources/cpu.json```
To view memory resources: ```GET http://hostname:5596/v1/resources/memory.json```
To view disk resources: ```GET http://hostname:5596/v1/resources/disk.json```
To ensure the application is operational: ```GET http://hostname:5596/ping```

### Benchmark Information

To view benchmark information: ```go test -bench=".*"```


# Roadmap
- Statsd output
- Currently only Linux /proc stats are collected. Profile information will be next.
- Only the Linux platform is supported. Other platforms, Darwin notibly, will be included.


# Versioning

Versioning adheres to [Semantic Versioning 2.0.0](http://semver.org/spec/v2.0.0.html).


# Author(s) & Credit

**Christian Vozar**

+ [http://twitter.com/christianvozar](http://twitter.com/christianvozar)
+ [http://github.com/christianvozar](http://github.com/christianvozar)

Credit given to Mozilla's [PushGo](https://github.com/jrconlin/pushgo/blob/master/src/mozilla.org/util/heka_log.go) for Heka client code for which the Heka code is heavily based.

## Copyright and License

Copyright 2013 Belly, Inc. under [the Apache 2.0 license](LICENSE.md).
