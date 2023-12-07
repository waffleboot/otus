#!/bin/bash

case $1 in
    stop)
        case $2 in
            nginx)
                case $3 in
                    all)
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@nginx-1.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo systemctl stop nginx
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@nginx-2.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo systemctl stop nginx
                        ;;
                    *)
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@nginx-$3.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo systemctl stop nginx
                        ;;
                esac
                ;;
            back)
                case $3 in
                    all)
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@backend-1.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo docker stop wordpress
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@backend-2.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo docker stop wordpress
                        ;;
                    *)
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@backend-$3.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo docker stop wordpress
                        ;;
                esac
                ;;
            bastion)
                ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@bastion.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo systemctl stop nginx
                ;;
        esac
        ;;
    start)
        case $2 in
            nginx)
                case $3 in
                    all)
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@nginx-1.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo systemctl start nginx
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@nginx-2.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo systemctl start nginx
                        ;;
                    *)
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@nginx-$3.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo systemctl start nginx
                        ;;
                esac
                ;;
            back)
                case $3 in
                    all)
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@backend-1.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo docker start wordpress
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@backend-2.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo docker start wordpress
                        ;;
                    *)
                        ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@backend-$3.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo docker start wordpress
                        ;;
                esac
                ;;
            bastion)
                ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@bastion.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa sudo systemctl start nginx
                ;;
        esac
        ;;
    ssh)
        case $2 in
            nginx)
                ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@nginx-$3.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa
                ;;
            back)
                ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@backend-$3.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa
                ;;
            db)
                ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@db.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa
                ;;
            bastion)
                ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no centos@bastion.ru-central1.internal -oProxyCommand="ssh centos@62.84.116.217 -i id_rsa -W %h:%p" -i id_rsa
                ;;
        esac
        ;;
    open)
        open http://62.84.116.217
        ;;
esac
