# Socket Dispatcher [![Build Status](https://travis-ci.org/bialas1993/socket-dispatcher.svg?branch=master)](https://travis-ci.org/bialas1993/socket-dispatcher)


### RUN
```
env SOCKET_DISPATCHER_PORTS="8000-8004" ./bin/socket-dispatcher --branch master --debug 1 --kill 1
```

### options: 
  - SOCKET_DISPATCHER_PORTS - env define port range to use 
  - branch - unique name for resolve port use
  - debug - for show debug log
  - kill - process which one use selected port should be kill