package main

import (
	"flag"
	"fmt"
	"os"

	dkd_crypto "github.com/bloxapp/ssv-dkg/pkgs/crypto"
)

func main() {
	privKeyPath := flag.String("key", "initiator_encrypted_key.json", "a string")
	privKeyPassPath := flag.String("pass", "initiator_password", "a string")
	flag.Parse()
	if _, err := os.Stat(*privKeyPath); os.IsNotExist(err) {
		fmt.Print(fmt.Errorf("🔑 private key file: %s", err))
		return
	}
	if _, err := os.Stat(*privKeyPassPath); os.IsNotExist(err) {
		fmt.Print(fmt.Errorf("🔑 password file: %s", err))
		return
	}
	key, err := dkd_crypto.ReadEncryptedRSAKey(*privKeyPath, *privKeyPassPath)
	if err != nil {
		fmt.Print(fmt.Errorf("🔑 error decrypting the key: %s", err))
		return
	}
	pkBytes, err := dkd_crypto.EncodePublicKey(&key.PublicKey)
	if err != nil {
		fmt.Print(fmt.Errorf("🔑 error encoding the key: %s", err))
		return
	}
	fmt.Printf("Pub key base64 encoded: %s", string(pkBytes))
}
