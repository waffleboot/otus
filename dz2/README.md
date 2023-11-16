- виртуалка с iscsi
- 3 виртуальных машины с разделяемой файловой системой gfs2 поверх clvm

```bash
ssh centos@node1.ru-central1.internal -oProxyCommand="ssh centos@51.250.87.225 -i id_rsa -W %h:%p" -i id_rsa
```

```bash
$ ansible-playbook -i inventory.ini playbook.yml

PLAY [server] ********************************************************************************************************************

TASK [Create file] ***************************************************************************************************************
changed: [server.ru-central1.internal]

TASK [Install targetcli] *********************************************************************************************************
changed: [server.ru-central1.internal]

TASK [command] *******************************************************************************************************************
fatal: [server.ru-central1.internal]: FAILED! => {"changed": true, "cmd": ["targetcli", "/iscsi", "delete", "iqn.2018-09.ru.otus:storage.target00"], "delta": "0:00:00.306727", "end": "2023-11-16 23:48:14.778008", "msg": "non-zero return code", "rc": 1, "start": "2023-11-16 23:48:14.471281", "stderr": "No such Target in configfs: /sys/kernel/config/target/iscsi/iqn.2018-09.ru.otus:storage.target00", "stderr_lines": ["No such Target in configfs: /sys/kernel/config/target/iscsi/iqn.2018-09.ru.otus:storage.target00"], "stdout": "Warning: Could not load preferences file /root/.targetcli/prefs.bin.", "stdout_lines": ["Warning: Could not load preferences file /root/.targetcli/prefs.bin."]}
...ignoring

TASK [command] *******************************************************************************************************************
fatal: [server.ru-central1.internal]: FAILED! => {"changed": true, "cmd": ["targetcli", "/backstores/fileio", "delete", "disk01"], "delta": "0:00:00.145013", "end": "2023-11-16 23:48:15.654172", "msg": "non-zero return code", "rc": 1, "start": "2023-11-16 23:48:15.509159", "stderr": "No storage object named disk01.", "stderr_lines": ["No storage object named disk01."], "stdout": "", "stdout_lines": []}
...ignoring

TASK [Map disk] ******************************************************************************************************************
changed: [server.ru-central1.internal]

TASK [Create target] *************************************************************************************************************
changed: [server.ru-central1.internal]

TASK [Create LUN] ****************************************************************************************************************
changed: [server.ru-central1.internal]

PLAY [clients] *******************************************************************************************************************

TASK [Install pacemaker] *********************************************************************************************************
changed: [node2.ru-central1.internal]
changed: [node3.ru-central1.internal]
changed: [node1.ru-central1.internal]

TASK [Start pacemaker] ***********************************************************************************************************
changed: [node2.ru-central1.internal]
changed: [node1.ru-central1.internal]
changed: [node3.ru-central1.internal]

TASK [Set hacluster password] ****************************************************************************************************
changed: [node3.ru-central1.internal]
changed: [node2.ru-central1.internal]
changed: [node1.ru-central1.internal]

PLAY [admin] *********************************************************************************************************************

TASK [Auth to pcs] ***************************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [Setup pcs] *****************************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [Enable pcs] ****************************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [Start pcs] *****************************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [Stop stonith] **************************************************************************************************************
changed: [node1.ru-central1.internal]

PLAY [clients] *******************************************************************************************************************

TASK [Discover target] ***********************************************************************************************************
changed: [node3.ru-central1.internal]
changed: [node2.ru-central1.internal]
changed: [node1.ru-central1.internal]

TASK [command] *******************************************************************************************************************
changed: [node3.ru-central1.internal]
changed: [node2.ru-central1.internal]
changed: [node1.ru-central1.internal]

TASK [Grant access] **************************************************************************************************************
changed: [node1.ru-central1.internal -> server.ru-central1.internal]
changed: [node3.ru-central1.internal -> server.ru-central1.internal]
changed: [node2.ru-central1.internal -> server.ru-central1.internal]

TASK [Start iscsi] ***************************************************************************************************************
changed: [node3.ru-central1.internal]
changed: [node2.ru-central1.internal]
changed: [node1.ru-central1.internal]

TASK [Login to target] ***********************************************************************************************************
changed: [node3.ru-central1.internal]
changed: [node2.ru-central1.internal]
changed: [node1.ru-central1.internal]

PLAY [server] ********************************************************************************************************************

TASK [command] *******************************************************************************************************************
changed: [server.ru-central1.internal]

TASK [ansible.builtin.copy] ******************************************************************************************************
changed: [server.ru-central1.internal -> localhost]

PLAY [admin] *********************************************************************************************************************

TASK [Create dlm] ****************************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [Create clvmd] **************************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [Create volume group] *******************************************************************************************************
FAILED - RETRYING: Create volume group (3 retries left).
changed: [node1.ru-central1.internal]

TASK [Create logical volume] *****************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [Format filesystem] *********************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [Mount filesystem] **********************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [Add order constraint] ******************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [Add order constraint] ******************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [Add colocation constraint] *************************************************************************************************
changed: [node1.ru-central1.internal]

PLAY [admin] *********************************************************************************************************************

TASK [command] *******************************************************************************************************************
changed: [node1.ru-central1.internal]

TASK [ansible.builtin.copy] ******************************************************************************************************
changed: [node1.ru-central1.internal -> localhost]

PLAY [src] ***********************************************************************************************************************

TASK [ansible.builtin.copy] ******************************************************************************************************
changed: [node1.ru-central1.internal]

PLAY [dst] ***********************************************************************************************************************

TASK [command] *******************************************************************************************************************
changed: [node2.ru-central1.internal]

TASK [debug] *********************************************************************************************************************
ok: [node2.ru-central1.internal] => {
    "msg": "123456"
}

PLAY RECAP ***********************************************************************************************************************
node1.ru-central1.internal : ok=25   changed=25   unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
node2.ru-central1.internal : ok=10   changed=9    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
node3.ru-central1.internal : ok=8    changed=8    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
server.ru-central1.internal : ok=9    changed=9    unreachable=0    failed=0    skipped=0    rescued=0    ignored=2
```

[Яндекс документация](https://terraform-provider.yandexcloud.net//Resources/compute_instance)  
[Gateways](https://cloud.yandex.com/en/docs/vpc/concepts/gateways)  
[Routing](https://cloud.yandex.com/en-ru/docs/tutorials/routing/nat-instance)  
[NAT Gateways](https://cloud.yandex.com/en-ru/docs/vpc/operations/create-nat-gateway)  
[Locals](https://developer.hashicorp.com/terraform/language/values/locals)  
[VM Metadata](https://cloud.yandex.com/en-ru/docs/compute/concepts/vm-metadata)  
[Templatefile](https://developer.hashicorp.com/terraform/language/functions/templatefile)  
[Ansible Inventory](https://docs.ansible.com/ansible/latest/inventory_guide/intro_inventory.html)  
[Ansible Bastion](https://www.jeffgeerling.com/blog/2022/using-ansible-playbook-ssh-bastion-jump-host)  
[ssh port forwarding](https://habr.com/ru/articles/331348/)
[Настройка mkfs.gfs2](https://www.altlinux.org/GFS2_%D0%BD%D0%B0_iSCSI_%D1%81_Multipath)  
