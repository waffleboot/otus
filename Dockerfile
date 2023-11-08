FROM ubuntu:20.04

RUN apt update && apt install -y vim curl wget unzip ansible

RUN wget https://hashicorp-releases.yandexcloud.net/terraform/1.6.3/terraform_1.6.3_linux_arm64.zip && \
    unzip terraform_1.6.3_linux_arm64.zip -d /usr/bin && rm terraform_1.6.3_linux_arm64.zip

RUN useradd -ms /bin/bash otus

USER otus
WORKDIR /home/otus

RUN curl -sSL https://storage.yandexcloud.net/yandexcloud-yc/install.sh | bash

RUN echo 'alias tf=terraform' >> ~/.bashrc
