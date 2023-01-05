package main

import (
	"fmt"
	"os"

	_ "github.com/ntnguyencse/cluster-api-sdk/client"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("KUBECONFIG", os.Getenv("KUBECONFIG"))
}

func main() {
	fmt.Println("Main function")
}
