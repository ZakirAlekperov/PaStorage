package main

import (
	"fmt"

	"github.com/ZakirAlekperov/PaStorage/internal/vault"
)

// main запускает приложение PaStorage.
func main() {
	storage := vault.New()

	fmt.Printf("PaStorage: записей %d\n", storage.Count())
}
