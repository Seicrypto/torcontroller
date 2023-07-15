# torcontroller

Now version: 1.0-1

torcontroller is packages which combines tor, privoxy, systemctl packages, and so on. torcontroller Dev built some scripts let you just command on Linux: Debian bullseye environments including docker container. You will be able to run your application and control tor router by a function wrote any back-end program.

If you are not reading this on github, please go to <https://github.com/Seicrypto/torcontroller-1.0>
Read more

torcontrollerはtor、privoxy、systemctlパッケージなどを組み合わせたパッケージです： docker コンテナを含む Debian bullseye 環境でコマンドを実行するだけです。任意のバックエンドプログラムを書いた関数で、アプリケーションを実行し、torルータを制御できるようになります。

githubでこれを読んでいない場合は、<https://github.com/Seicrypto/torcontroller-1.0>にアクセスしてください。

[日本語説明こちら](./READMEJP.md)

## QuickStart

Use in:

* [Linux Debian / Ubuntu](#linux)
* [Docker container](#control-in-docker-container)

### Linux

Now torcontroller suport on Linux Debian / Ubuntu else.

Step1. Download and install

```bash
#!/bin/bash
apt-get update
# ARM cpu:
wget https://github.com/Seicrypto/torcontroller-1.0/release/v1.0/torcontroller_1.0-1_arm64.deb
apt-get install -y ./torcontroller_1.0-1_arm64.deb
# Intel / AMD cpu:
# wget https://github.com/Seicrypto/torcontroller-1.0/release/v1.0/torcontroller_1.0-1_amd64.deb
# apt-get install -y ./torcontroller_1.0-1_amd64.deb

# * How to know your machine using ARM or Intel / AMD
# uname -m
# Response :
# aarch64 (Means ARM)
# x86_64 (Means Intel / AMD)
```

Step2. Set up your AUTHENTICATE password

```bash
#!/bin/bash
torcontoller --version
# torcontroller version 1.0
torcontroller --resetpassword
# torcontroller info ...
# Enter old TOR password:
# (torcontroller set as default password)
```

Step3. Check tor and privoxy feature

```bash
#!/bin/bash
torcontroller --start
# log info...
# Start command succeeded.
curl -x 127.0.0.1:8118 http://icanhazip.com/
# 176.10.99.200 (a example tor ip address)
curl http://icanhazip.com/
# 89.196.159.79 (a example your ture ip address)
```

### Control in docker container:

* [Controll with Golang](#golang-sample)
* [Others program control sample](#other-program-sample)

#### Golang sample

Step1. Download and Install package

Please make sure your docker image base on debina.
Such as bullseye, bookworm:

* golang:bullseye
* golang:bookworm
* python:3.9-bullseye
* node:bullseye
* so on...

```dockerfile
# dockerfile
# Recommend bullseye / bookworn ... else built according degina system.
From golang:bullseye
RUN apt-get update
# Make sure there's wget in your docker image.
RUN wget https://github.com/Seicrypto/torcontroller-1.0/release/v1.0/torcontroller_1.0-1_arm64.deb
RUN apt-get install -y /app/torcontroller_1.0-1_arm64.deb
```

Step2. Set up your AUTHENTICATE password

Terminal (Recommend)

```bash
# Terminal
docker run -p 2043:9050 -p 2044:9051 <docker-image>
# On terminal to control docker.
# Of course you could control in golang func,
# if you want.
docker exec <your-container>

```

Step3. Check tor and privoxy feature

Call bash in your programing funtcion. You could build it by yourself or import it.
Example:

* [Gotorcontroller]()
* Other programing might update in future.

#### Other program sample

Coming soon.

## Detail

Read more torcontroller command:

[torcontroller command list](./docs/commandList.md)

## Reference

[A step-by-step guide how to use Python with Tor and Privoxy](https://gist.github.com/DusanMadar/8d11026b7ce0bce6a67f7dd87b999f6b) :

Which is my basic script content reference.

[tor.service file for systemctl](https://gist.github.com/gtank/f6a8f99c70f682cd8d4acd6a4a9ee696)

[privoxy.service file for systemctl](https://alt.os.linux.mageia.narkive.com/D2i3xOYQ/privoxy-service-file-for-systemd)