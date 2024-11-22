package main

import (
	"net/http"
	"os"
	"log"
)

func main() {
	// Obter a URL de redirecionamento da variável de ambiente
	redirectURL := os.Getenv("REDIRECT_URL")
	if redirectURL == "" {
		log.Fatal("A variável de ambiente REDIRECT_URL não está definida")
	}

	// Configurar o servidor HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, redirectURL, http.StatusFound) // Redireciona com status 302
	})

	// Inicia o servidor na porta 8080
	log.Println("Servidor iniciado na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
