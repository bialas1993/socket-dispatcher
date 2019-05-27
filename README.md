# Socket Dispatcher [![Build Status](https://travis-ci.org/bialas1993/socket-dispatcher.svg?branch=master)](https://travis-ci.org/bialas1993/socket-dispatcher) [![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md) [![Version](https://img.shields.io/github/tag/bialas1993/socket-dispatcher.svg)](https://github.com/bialas1993/socket-dispatcher/releases)


### RUN
```
env SOCKET_DISPATCHER_PORTS="8000-8004" ./bin/socket-dispatcher --branch master --debug --kill
```

### options: 
  - *SOCKET_DISPATCHER_PORTS* - env define port range to use 
  - *DATABASE_PATH* - env define path for sqlite database
  - *branch* - unique name for resolve port use
  - *debug* - for show debug log
  - *kill* - process which one use selected port should be kill