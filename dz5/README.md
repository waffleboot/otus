# Схема

```
+ Cluster: postgres-cluster (7315484841462623129) -+----+-----------+
| Member | Host         | Role         | State     | TL | Lag in MB |
+--------+--------------+--------------+-----------+----+-----------+
| db-1   | 192.168.0.41 | Replica      | streaming |  7 |         0 |
| db-2   | 192.168.0.42 | Leader       | running   |  7 |           |
| db-3   | 192.168.0.43 | Sync Standby | streaming |  7 |         0 |
+--------+--------------+--------------+-----------+----+-----------+
```

![](arch.png)

# Домашнее задание

Реализация кластера postgreSQL с помощью patroni

# Цель:

Перевести БД веб проекта на кластер postgreSQL с ипользованием patroni, etcd/consul/zookeeper и haproxy/pgbouncer

# Описание/Пошаговая инструкция выполнения домашнего задания

Перевести БД веб проекта на кластер postgreSQL с ипользованием patroni, etcd/consul/zookeeper и haproxy/pgbouncer.

# Порядок запуска

Для установки postgresql-15/haproxy/etcd/patroni был взят https://github.com/vitabaks/postgresql_cluster

- https://github.com/vitabaks/postgresql_cluster
- ./gfs2_install.sh
- ./webapp_install.sh
- ./nginx_install.sh

```bash
cp id_rsa ../../postgresql_cluster/
ansible all -m ping
ansible-playbook deploy_pgcluster.yml
```
скопировать ansible_ssh_common_args из inventory

# Тест

```bash
./run.sh stop nginx all
./run.sh stop back all
./run.sh test

./run.sh start nginx 1
./run.sh start back 1
./run.sh test

./run.sh stop nginx 1
./run.sh stop back 1
./run.sh start nginx 2
./run.sh start back 2
./run.sh test
```

```bash
./run.sh stop db all

./run.sh stop db 1
./run.sh start db 2
./run.sh test

./run.sh stop db 2
./run.sh start db 3
./run.sh test

./run.sh stop db 3
./run.sh start db 1
./run.sh test
```
