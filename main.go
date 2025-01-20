package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func initializeConfig() {
	viper.SetDefault("app.port", 8000)
	viper.AutomaticEnv()
}

func main() {
	initializeConfig()

	port := viper.GetInt("app.port")
	address := fmt.Sprintf(":%d", port)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode("ok")
	}).Methods("GET")

	fmt.Printf("🚀 Server is running on port %d\n", port)
	if err := http.ListenAndServe(address, r); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

}
