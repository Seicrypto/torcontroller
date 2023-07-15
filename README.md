# torcontroller

Now version: 1.0-1

torcontroller is a Debian package which combines tor, privoxy, systemctl packages, and so on. torcontroller Dev built some scripts to let you just command on Linux: Debian bullseye environments including docker container. You will be able to run your application and control tor router by a function that writes any back-end program.

If you are not reading this on github, please go to <https://github.com/Seicrypto/torcontroller>
Read more

torcontrollerはtor、privoxy、systemctlパッケージなどを組み合わせたパッケージです： docker コンテナを含む Debian bullseye 環境でコマンドを実行するだけです。任意のバックエンドプログラムを書いた関数で、アプリケーションを実行し、torルータを制御できるようになります。

githubでこれを読んでいない場合は、<https://github.com/Seicrypto/torcontroller>にアクセスしてください。

[日本語説明こちら](./READMEJP.md)

## QuickStart

Use in:

* [Linux Debian / Ubuntu](#linux)
* [Docker container](#use-in-docker-container)

### Linux

Now torcontroller suport on Linux Debian / Ubuntu else.

Step1. Download and install

```bash
#!/bin/bash
apt-get update

# Intel / AMD cpu:
wget https://github.com/Seicrypto/torcontroller/release/v1.0/torcontroller_1.0-1_amd64.deb
apt-get install -y ./torcontroller_1.0-1_amd64.deb

# ARM cpu:
# wget https://github.com/Seicrypto/torcontroller/release/v1.0/torcontroller_1.0-1_arm64.deb
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

Please make sure your docker image base on debina.
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

# Make sure there's wget in your docker image.
RUN apt-get update

# Intel / AMD cpu:
RUN wget https://github.com/Seicrypto/torcontroller/release/v1.0/torcontroller_1.0-1_amd64.deb
RUN apt-get install -y /app/torcontroller_1.0-1_amd64.deb

# ARM cpu:
# RUN wget https://github.com/Seicrypto/torcontroller/release/v1.0/torcontroller_1.0-1_arm64.deb
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

* [Golang: gotorcontroller](https://github.com/Seicrypto/gocontroller)
* Other programming func lib might update in future.

## Detail

Read more torcontroller command:

[torcontroller command list](./docs/commandList.md)

## Reference

[A step-by-step guide how to use Python with Tor and Privoxy](https://gist.github.com/DusanMadar/8d11026b7ce0bce6a67f7dd87b999f6b) :

Which is my basic script content reference.

[tor.service file for systemctl](https://gist.github.com/gtank/f6a8f99c70f682cd8d4acd6a4a9ee696)

[privoxy.service file for systemctl](https://alt.os.linux.mageia.narkive.com/D2i3xOYQ/privoxy-service-file-for-systemd)
