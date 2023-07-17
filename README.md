# torcontroller

Now version: 1.0-1

[![Badge](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2FSeicrypto%2Ftorcontroller&count_bg=%236DAC3D&title_bg=%23555555&icon=grafana.svg&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://github.com/Seicrypto/gotorcontroller)

What is [tor](https://www.torproject.org/)? "The Onion Router," is free and open-source software for enabling anonymous communication.

torcontroller is a Debian package which combines tor, privoxy, systemctl packages, and so on. torcontroller Dev built some scripts to let you just command on Linux: Debian bullseye environments including docker container. You will be able to run your application and control tor router by a function that writes any back-end program.

If you are not reading this on github, please go to <https://github.com/Seicrypto/torcontroller>
Read more

[tor](https://www.torproject.org/)とは？「オニオン・ルーター（The Onion Router）」は、匿名通信を可能にするフリーでオープンソースのソフトウェアです。

torcontrollerはtor、privoxy、systemctlパッケージなどを組み合わせたパッケージです： docker コンテナを含む Debian bullseye 環境でコマンドを実行するだけです。任意のバックエンドプログラムを書いた関数で、アプリケーションを実行し、torルータを制御できるようになります。

githubでこれを読んでいない場合は、<https://github.com/Seicrypto/torcontroller>にアクセスしてください。

Japanese README:
[日本語説明こちら](./READMEJP.md)

## QuickStart

Use in:

* [Linux Debian / Ubuntu](#linux)
* [Docker container](#use-in-docker-container)

### Linux

![Debian](https://img.shields.io/badge/Debian-A81D33?style=for-the-badge&logo=debian&logoColor=white) ![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white)

Now torcontroller suport on Linux Debian / Ubuntu else.

Step1. Download and install

```bash
#!/bin/bash
apt-get update

# Intel / AMD cpu:
wget https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_amd64.deb
apt-get install -y ./torcontroller_1.0-1_amd64.deb

# ARM cpu:
# wget https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_arm64.deb
# apt-get install -y ./torcontroller_1.0-1_arm64.deb

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
# torcontroller version X.X

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

### Use in docker container

 ![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)

Please make sure your docker image base on debian.
Such as bullseye, bookworm:

* golang:bullseye
* golang:bookworm
* python:3.9-bullseye
* node:bullseye
* so on...

Step1. Download and Install package

```dockerfile
# dockerfile
# Recommend bullseye / bookworn ... else built according degina system.
From golang:bullseye
# Of course work for any programming you want,
# just remmember use image: bullseye, bookworm, and so on.
WORKDIR /app

# Make sure there's wget in your docker image.
RUN apt-get update

# Intel / AMD cpu:
RUN wget https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_amd64.deb
RUN apt-get install -y /app/torcontroller_1.0-1_amd64.deb
# Be careful about your download and install path.

# ARM cpu:
# RUN wget https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_arm64.deb
# RUN apt-get install -y /app/torcontroller_1.0-1_arm64.deb
```

Step2. Set up your AUTHENTICATE password

Terminal (Recommend)

```bash
# Terminal
# This sample using 9050 as socketPort, 9051 as controlPort on machine.
docker run -it -p 9050:9050 -p 9051:9051 --name <your-container-name> <docker-image>

# Check torcontroller installed properly.
docker exec <your-container-name> torcontroller --version
# torcontroller version X.X-X

# On terminal to control docker.
docker exec -i <your-container-name> torcontroller --resetpassword
# torcontroller info ...
# Enter old TOR password:
# (torcontroller set as default password)
```

Step3. Check tor and privoxy feature

Call bash in your programing function. You could build it by yourself or import it.

Be careful! Use 127.0.0.1:8118 Port as proxy in your application. It would go through tor socket 9050 Port in your docker container.

Example:

* ![golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) [Golang: gotorcontroller](https://github.com/Seicrypto/gocontroller)
* Other programming func lib might update in future.

## Detail

Read more torcontroller command:

[torcontroller command list](./docs/commandList.md)

## Reference

[A step-by-step guide how to use Python with Tor and Privoxy](https://gist.github.com/DusanMadar/8d11026b7ce0bce6a67f7dd87b999f6b) :

Which is my basic script content reference.

[tor.service file for systemctl](https://gist.github.com/gtank/f6a8f99c70f682cd8d4acd6a4a9ee696)

[privoxy.service file for systemctl](https://alt.os.linux.mageia.narkive.com/D2i3xOYQ/privoxy-service-file-for-systemd)
