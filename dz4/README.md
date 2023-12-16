# Схема

![](arch.jpg)

# Домашнее задание

Развернуть InnoDB или PXC кластер

# Цель:

Перевести базу веб-проекта на один из вариантов кластера MySQL:
- Percona XtraDB Cluster
- или InnoDB Cluster.

# Описание/Пошаговая инструкция выполнения домашнего задания:

- Разворачиваем отказоустойчивый кластер MySQL (PXC || Innodb) на ВМ или в докере любым способом
- Создаем внутри кластера вашу БД для проекта

# Лекция

- mysql innodb cluster
- percona xtradb cluster (pxc)

# Ссылки

https://stackoverflow.com/questions/38847824/ansible-how-to-get-service-status-by-ansible
https://docs.percona.com/percona-xtradb-cluster/8.0/encrypt-traffic.html
https://docs.percona.com/percona-xtradb-cluster/8.0/encrypt-traffic.html#generate-keys-and-certificates-manually

# Порядок запуска

- mysql_install.sh
- wordpress_install.sh
- nginx_install.sh
- gfs2_install.sh
