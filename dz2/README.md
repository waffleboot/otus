- виртуалка с iscsi
- 3 виртуальных машины с разделяемой файловой системой gfs2 поверх clvm

```bash
ssh centos@node1.ru-central1.internal -oProxyCommand="ssh centos@51.250.87.225 -i id_rsa -W %h:%p" -i id_rsa
```

[Яндекс документация](https://terraform-provider.yandexcloud.net//Resources/compute_instance)  
[Gateways](https://cloud.yandex.com/en/docs/vpc/concepts/gateways)  
[Routing](https://cloud.yandex.com/en-ru/docs/tutorials/routing/nat-instance)  
[NAT Gateways](https://cloud.yandex.com/en-ru/docs/vpc/operations/create-nat-gateway)  
[Locals](https://developer.hashicorp.com/terraform/language/values/locals)  
[VM Metadata](https://cloud.yandex.com/en-ru/docs/compute/concepts/vm-metadata)  
[Templatefile](https://developer.hashicorp.com/terraform/language/functions/templatefile)  
