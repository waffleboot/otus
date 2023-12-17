#!/bin/bash

ansible-playbook -i inventory.ini gfs2_playbook.yml $@
