#!/bin/bash

ansible-playbook -i inventory.ini wordpress_playbook.yml $@
