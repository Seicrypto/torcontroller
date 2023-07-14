# torcontroller

Now version: 1.0-1

torcontroller is packages which combines tor, privoxy, systemctl packages, and so on. torcontroller Dev built some scripts let you just command on Linux: Debian bullseye envirements including docker container.

If you are not reading this on github, please go to <https://github.com/Seicrypto/torcontroller-1.0>
Read more

[日本語説明](./READMEJP.md)

## QuickStart

Use in:
* [Linux](#linux)
* [Docker container](#control-in-docker-container)

### Linux

Now torcontroller suport on Linux Debian / Ubuntu else.

step1. download

```bash
# mac cpu:
wget https://github.com/Seicrypto/torcontroller-1.0/release/v1.0/torcontroller_1.0-1_arm64.deb
#
#
```

step2. install

```bash
apt-get update
# mac cpu:
apt-get install -y ./torcontroller_1.0-1_arm64.deb
#
#
```

### Control in docker container:

Controll with Golang

```golang

```

## Detail

Read more torcontroller command:

[torcontroller command list](./docs/commandList.md)

## Reference

[A step-by-step guide how to use Python with Tor and Privoxy](https://gist.github.com/DusanMadar/8d11026b7ce0bce6a67f7dd87b999f6b) :

Which is my basic script content reference.

[tor.service file for systemctl](https://gist.github.com/gtank/f6a8f99c70f682cd8d4acd6a4a9ee696)

[privoxy.service file for systemctl](https://alt.os.linux.mageia.narkive.com/D2i3xOYQ/privoxy-service-file-for-systemd)