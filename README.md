# lmqtt-ping-test

This repository contains an MQTT client and an MQTT server originally created for debugging connectivity problems.  I published it since it is a rather minimal example of how to make
an MQTT client and server pair using [lmqtt](https://github.com/lab5e/lmqtt), which is a fork of the [gmqtt](https://github.com/DrmagicE/gmqtt) library that we've tried to clean up a bit.

## Build

```shell
make
```

This produces six binaries in the bin directory. A `client` and `server` binary for the platform you are on, a amd64/linux binary and an ARM/linux binary.