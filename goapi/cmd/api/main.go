package main 

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println(`
	___           ___                    ___           ___                 
	/\  \         /\  \                  /\  \         /\  \          ___   
   /::\  \       /::\  \                /::\  \       /::\  \        /\  \  
  /:/\:\  \     /:/\:\  \              /:/\:\  \     /:/\:\  \       \:\  \ 
 /:/  \:\  \   /:/  \:\  \            /::\~\:\  \   /::\~\:\  \      /::\__\
/:/__/_\:\__\ /:/__/ \:\__\          /:/\:\ \:\__\ /:/\:\ \:\__\  __/:/\/__/
\:\  /\ \/__/ \:\  \ /:/  /          \/__\:\/:/  / \/__\:\/:/  / /\/:/  /   
 \:\ \:\__\    \:\  /:/  /                \::/  /       \::/  /  \::/__/    
  \:\/:/  /     \:\/:/  /                 /:/  /         \/__/    \:\__\    
   \::/  /       \::/  /                 /:/  /                    \/__/    
	\/__/         \/__/                  \/__/                              
	`)

	fmt.Println("Listening on localhost:8080")

	err:= http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Error(err)
	}


}

