package main

import (
	"fmt"
	identity "github.com/ic-matcom/model-identity-go/api"
	_ "github.com/ic-matcom/model-identity-go/tools"
)

func main() {
	fmt.Println("probando...", identity.ContractNameIdentity)
}

