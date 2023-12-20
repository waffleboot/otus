#!/bin/bash

ansible-playbook -i inventory.ini pg_playbook.yml $@
