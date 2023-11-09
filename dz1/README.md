## Сгенерить ключи

```bash
ssh-keygen -t rsa -f id_rsa
```

## Создать ВМ

```bash
terraform apply
```

## Поставить nginx через Ansible

```bash
ansible-playbook -i inventory.ini playbook.yml
```

## Вытащить public IP

Был в output Terraform

```
Outputs:

public_ip = "51.250.77.84"
```

## Проверить что nginx работает

```bash
curl http://51.250.77.84
```

## Удалить ВМ

```bash
terraform destroy
```
