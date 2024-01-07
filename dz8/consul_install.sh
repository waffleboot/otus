#!/bin/bash

ansible-playbook -i inventory.ini consul_playbook.yml $@
