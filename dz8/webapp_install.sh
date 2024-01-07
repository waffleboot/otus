#!/bin/bash

ansible-playbook -i inventory.ini webapp_playbook.yml $@
