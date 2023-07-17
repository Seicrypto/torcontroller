# torcontroller Japanese 日本語説明

[![Badge](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2FSeicrypto%2Ftorcontroller&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://github.com/Seicrypto/torcontroller)

[tor](https://www.torproject.org/)とは？「オニオン・ルーター（The Onion Router）」は、匿名通信を可能にするフリーでオープンソースのソフトウェアです。

torcontrollerはtor、privoxy、systemctlパッケージなどを組み合わせたパッケージです： docker コンテナを含む Debian bullseye 環境でコマンドを実行するだけです。任意のバックエンドプログラムを書いた関数で、アプリケーションを実行し、torルータを制御できるようになります。

## クイック・スタート

使用:

* [Linux Debian / Ubuntu](#linux)
* [Docker コンテナ](#dockerドッカーコンテナで使用する)

### Linux

![Debian](https://img.shields.io/badge/Debian-A81D33?style=for-the-badge&logo=debian&logoColor=white) ![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white)

Now torcontroller suport on Linux Debian / Ubuntu.
LinuxのDebian / Ubuntuでtorcontrollerが使える。

Step1. ダウンロードとインストール

```bash
#!/bin/bash
apt-get update

# Intel / AMD cpu:
wget https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_amd64.deb
apt-get install -y ./torcontroller_1.0-1_amd64.deb

# ARM cpu:
# wget https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_arm64.deb
# apt-get install -y ./torcontroller_1.0-1_arm64.deb

# * ARMまたはIntel / AMDを使用しているマシンを知る方法:
# uname -m
# 応答:
# aarch64 (ARM)
# x86_64 (Intel / AMD)
```

Step2. 認証パスワードを設定する

```bash
#!/bin/bash
torcontoller --version
# torcontroller バージョン X.X

torcontroller --resetpassword
# torcontroller info ...
# Enter old TOR password:
# (デフォルトパスワードはtorcontroller)
```

Step3. torとprivoxyの機能をチェック

```bash
#!/bin/bash
torcontroller --start
# log info...
# Start command succeeded.　(開始コマンド成功)

curl -x 127.0.0.1:8118 http://icanhazip.com/
# 176.10.99.200 (torのIPアドレスの例)
curl http://icanhazip.com/
# 89.196.159.79 (あなたのIPアドレスの例)
```

### Dockerドッカーコンテナで使用する

 ![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)

ドッカーイメージがdebianベースであることを確認してください。
bullseye, bookwormなど:

* golang:bullseye
* golang:bookworm
* python:3.9-bullseye
* node:bullseye
* など...

Step1. パッケージのダウンロードとインストール

```dockerfile
# dockerfile
# bullseye / bookworn その他、Debian・システムに従って作られたものを推奨する。
From golang:bullseye
# もちろん、どんなプログラミングにも対応できる,
# bullseye、bookworn、など。.
WORKDIR /app

# docker イメージで wget が動作していることを確認してください。
RUN apt-get update

# Intel / AMD cpu:
RUN wget https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_amd64.deb
RUN apt-get install -y /app/torcontroller_1.0-1_amd64.deb
# ダウンロードとインストールのパスに注意してください。

# ARM cpu:
# RUN wget https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_arm64.deb
# RUN apt-get install -y /app/torcontroller_1.0-1_arm64.deb
```

Step2. 認証パスワードを設定する

コマンドターミナル（推奨）

```bash
# ターミナル
# このサンプルでは、マシンのsocketPortに9050、controlPortに9051を使用しています。
docker run -it -p 9050:9050 -p 9051:9051 --name <your-container-name> <docker-image>

# torcontrollerが正しく取り付けられていることを確認する。
docker exec <your-container-name> torcontroller --version
# torcontroller version X.X-X

# ターミナルで docker を操作する。
docker exec -i <your-container-name> torcontroller --resetpassword
# torcontroller info ...
# Enter old TOR password:
# (torcontroller set as default password)
```

Step3. torとprivoxyの機能をチェック

あなたのプログラム機能でbashを呼び出す。自分で作りてもよろし、インポートしてもいい。

注意してください！127.0.0.1:8118ポートをプロキシとして使ってください。ドッカーコンテナの tor ソケット 9050 ポートを経由することになる。

Example:

* ![golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) [Golang: gotorcontroller](https://github.com/Seicrypto/gocontroller)
* 今後、他のプログラミング・ファンク・ライブラリが更新されるかもしれない。

## 詳細はこちら

続きを読む torcontrollerコマンド:

[コマンド一覧](./docs/commandListJP.md)

## Reference

[A step-by-step guide how to use Python with Tor and Privoxy](https://gist.github.com/DusanMadar/8d11026b7ce0bce6a67f7dd87b999f6b) :

Which is my basic script content reference.

[tor.service file for systemctl](https://gist.github.com/gtank/f6a8f99c70f682cd8d4acd6a4a9ee696)

[privoxy.service file for systemctl](https://alt.os.linux.mageia.narkive.com/D2i3xOYQ/privoxy-service-file-for-systemd)
