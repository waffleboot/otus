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
[Ansible: Execute task only when a tag is specified](https://serverfault.com/questions/623634/ansible-execute-task-only-when-a-tag-is-specified)

https://www.project-open.com/en/howto-postgresql-port-secure-remote-access
https://info.gosuslugi.ru/articles/%D0%A0%D0%B0%D0%B7%D0%BC%D0%B5%D1%89%D0%B5%D0%BD%D0%B8%D0%B5_%D0%A1%D0%A3%D0%91%D0%94_PostgreSQL_%D0%BD%D0%B0_%D0%BE%D1%82%D0%B4%D0%B5%D0%BB%D1%8C%D0%BD%D0%BE%D0%BC_%D0%BE%D1%82_%D0%B0%D0%B4%D0%B0%D0%BF%D1%82%D0%B5%D1%80%D0%B0_CentOS_%D1%81%D0%B5%D1%80%D0%B2%D0%B5%D1%80%D0%B5/
https://computingforgeeks.com/how-to-install-postgresql-on-centos-rhel-7/
https://hub.docker.com/_/wordpress
https://docs.docker.com/network/network-tutorial-host/
https://docs.docker.com/config/containers/start-containers-automatically/
https://www.techtransit.org/install-mysql-database-server-through-ansible-playbook/
https://copyprogramming.com/howto/install-and-configure-mysql-using-ansible
https://www.digitalocean.com/community/tutorials/how-to-install-wordpress-on-centos-7
https://github.com/geerlingguy/ansible-role-mysql
https://docs.ansible.com/ansible/latest/playbook_guide/playbooks_reuse_roles.html
https://docs.ansible.com/ansible/latest/reference_appendices/config.html
https://askubuntu.com/questions/883404/pip-install-is-not-installing-executables-in-usr-local-bin
https://stackoverflow.com/questions/65225803/ansible-ec2-no-such-file-or-directory-bssh-bssh
https://stackoverflow.com/questions/1559955/host-xxx-xx-xxx-xxx-is-not-allowed-to-connect-to-this-mysql-server
https://docs.docker.com/config/daemon/systemd/
https://docs.docker.com/engine/install/centos/#install-using-the-repository

GOOS=linux GOARCH=amd64 go build -o app main.go
