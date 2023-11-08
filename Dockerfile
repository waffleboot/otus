FROM ubuntu:20.04

RUN apt update && apt install -y curl ansible

RUN curl -sSL https://storage.yandexcloud.net/yandexcloud-yc/install.sh | bash
