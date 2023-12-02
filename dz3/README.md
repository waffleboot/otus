# Домашнее задание

Настройка конфигурации веб приложения под высокую нагрузку

# Цель:

Terraform и ansible роль для развертывания серверов веб приложения под высокую нагрузку и отказоустойчивость.

В работе должны применяться:

- keepalived
- nginx,
- uwsgi/unicorn/php-fpm
- некластеризованная бд mysql/mongodb/postgres/redis
- gfs2

Должна быть реализована:

- отказоустойчивость бэкенд и nginx серверов
- отказоустойчивость сессий
- фэйловер без потери статического контента

Должны быть реализованы ansible скрипты с тюнингом:
- параметров sysctl
- лимитов
- настроек nginx
- включением пулов соединений

# Описание/Пошаговая инструкция выполнения домашнего задания:

- Создать несколько инстансов с помощью терраформ (2 nginx, 2 backend, 1 db)
- Развернуть nginx и keepalived на серверах nginx при помощи ansible
- Развернуть бэкенд способный работать по uwsgi/unicorn/php-fpm и базой. (Можно взять что нибудь из Django) при помощи ansible.
- Развернуть gfs2 для бэкенд серверах, для хранения статики
- Развернуть бд для работы бэкенда при помощи ansbile
- Проверить отказоустойчивость системы при выходе из строя серверов backend/nginx

# План

- [ ] посмотреть про nginx
- [ ] посмотреть про keepalived
- [ ] посмотреть про haproxy
- [ ] посмотреть про envoy
- [ ] посмотреть про traefik
- [ ] развернуть nginx и keeplived на серверах nginx при помощи ansible

# Заметки

- VRRP это про отказоустойчивость, HAProxy про балансировку
- файл конфигурации keepalived /etc/keepalived/keepalived.conf
- systemctl enable start keepalived

[Устройство сети в Yandex Cloud](https://cloud.yandex.ru/docs/overview/concepts/network)  
[Deckhouse](https://deckhouse.ru/documentation/v1/modules/450-keepalived/examples.html)  
[Резервирование маршрутизатора с использованием протокола VRRP](https://procloud.ru/blog/cases/rezervirovanie-marshrutizatora-s-ispolzovaniem-protokola-vrrp/)  
[репо с курса](https://github.com/Nickmob/vagrant-ansible-haproxy-keepalived)  

GOOS=linux GOARCH=amd64 go build -o app main.go