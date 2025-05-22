#!/bin/bash
export VAULT_ADDR='http://127.0.0.1:8200'
vault kv get -field=password secret/ansible_vault_password
