# vault-sible
Demo: Ansible integrates with HashiCorp Vault to securely handle sensitive data in playbooks.

# Step 1:
Start up Vault
```
$ vault server -dev
```
Check the status
```
$ export VAULT_ADDR='http://127.0.0.1:8200'
$ vault status
```
Generate a secret password for Ansible Vault, and push it to HashiCorp Vault
```
$ vault kv put secret/ansible_vault_password password=$(openssl rand -base64 32)
```

# Step 2:
Start up Consul
```
$ consul agent -dev -hcl 'acl { enabled = true default_policy = "deny" tokens { master = "MASTER_TOKEN" }}'
```
store the Consul's master token into **vars.yaml** file
```
$ echo -e '---\ntoken: "MASTER_TOKEN"' > vars.yaml
```

# Step 3:
create a script eg: `vault_password.sh` or `vault_password.go`. 

Set the environment variable `ANSIBLE_VAULT_PASSWORD`
```
$ export ANSIBLE_VAULT_PASSWORD_FILE=/path/to/the/vault_password.sh
```
If using go script, compile it into a binary
```
$ go build vault_password.go
```
Let's encrypt the **vars.yaml**
```
$ ansible-vault encrypt vars.yml
```

# Step 4:
Runs the **playbook.yaml** example playbook, which retrieves the secret password from HashiCorp Vault, decrypts **vars.yaml**, and then completes the remaining tasks.
