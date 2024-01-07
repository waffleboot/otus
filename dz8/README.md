# Домашнее задание

Consul cluster для service discovery и DNS

# Цель:

Реализовать consul cluster который выдает доменное имя для веб портала с прошлой ДЗ.
Плавающий IP заменить на балансировку через DNS.
В случае умирание одного из веб серверов IP должен убираться из DNS.

# Описание/Пошаговая инструкция выполнения домашнего задания:

- Реализовать consul cluster который выдает доменное имя для веб портала с прошлой ДЗ.
- Плавающий IP заменить на балансировку через DNS.
- В случае умирание одного из веб серверов IP должен убираться из DNS.

# Install

- tf validate
- tf apply
- ./gfs2_install.sh
- ./webapp_install.sh
- ./nginx_install.sh

скопировать ansible_ssh_common_args из inventory

```bash
cp /home/otus/otus/dz8/id_rsa /home/otus/postgresql_cluster
cd /home/otus/postgresql_cluster
vim /home/otus/postgresql_cluster/inventory
ansible all -m ping
ansible-playbook deploy_pgcluster.yml
```
