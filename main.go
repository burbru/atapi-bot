package main

import (
	"encoding/json"
	"fmt"
	"os"

	atapi_client "github.com/burbru/atapi-client"
)

func main() {
	token := os.Getenv("AT_AUTH")
	atapiWrapper := atapi_client.NewAtapiWrapper(token)
	bav, err := atapiWrapper.GetMyBankAccount()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		jsonData, _ := json.MarshalIndent(bav, "", "  ")
		fmt.Println(string(jsonData))
	}
}
