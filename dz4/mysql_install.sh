#!/bin/bash

ansible-playbook -i inventory.ini mysql_playbook.yml $@
