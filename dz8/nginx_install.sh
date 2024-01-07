#!/bin/bash

ansible-playbook -i inventory.ini nginx_playbook.yml $@
