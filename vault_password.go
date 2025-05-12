package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("vault", "kv", "get", "-field=password", "secret/ansible_vault_password")
	cmd.Env = append(os.Environ(), "VAULT_ADDR=http://127.0.0.1:8200")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Output:", string(output))
		return
	}
	fmt.Print(string(output))
}
