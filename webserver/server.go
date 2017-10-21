package webserver

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"context"
	"../model"
	"time"
)

type WebserverData struct {
	Stacks model.Stacks
}

func (wd *WebserverData) UpdateStacks(stacks model.Stacks) {
	wd.Stacks = stacks
}

func (wd *WebserverData) HandleGetStacks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(wd.Stacks)
}

type Server struct {
	logger *log.Logger
	mux    *http.ServeMux
}


func StartServer(port string, data *WebserverData, c chan bool) {
	router := mux.NewRouter()
	router.HandleFunc("/", data.rootHandler).Methods("GET")
	listenAddress := fmt.Sprintf(":%s", port)
	server := &http.Server{Addr: listenAddress, Handler: router}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	if b := <-c; b {
		fmt.Printf("We got told to quit\n")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		server.Shutdown(ctx)
	}
	c <- true
}


func (d *WebserverData) rootHandler(w http.ResponseWriter, r *http.Request) {

	pageHeader := "" +
		"<!DOCTYPE html>" +
		"<html>" +
		"<head>" +
		"<meta charset=\"utf-8\">" +
		"<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">" +
		"<title>Docker Flow Proxy Index Site</title>" +
		"<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css\">" +
		"<link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.0/css/bulma.min.css\">" +
		"</head>" +
		"<body>" +
		"<section class=\"hero is-dark is-bold\">" +
		"<div class=\"hero-body\">" +
		"<div class=\"container\">" +
		"<h1 class=\"title\">Docker Stacks</h1>" +
		"</div>" +
		"</div>" +
		"</section>"

	tableHeader := "" +
		"<div class=\"table is-hoverable is-fullwidth\">" +
		"<table class=\"table is-hoverable\">" +
		"<thead>" +
		"<tr>" +
		"<th>Service Name</th>" +
		"<th>Service Domain</th>" +
		"</tr>" +
		"</thead>" +
		"<tbody>"

	tableFooter := "" +
		"</tbody>" +
		"</table>" +
		"</div>"

	pageFooter := "" +
		"</div>" +
		"</div>" +
		"</section>" +
		"</body>" +
		"</html>"

	userAgent := r.Header.Get("User-Agent")
	fmt.Printf("  > [URI: %s, Method: %s, User-Agent: %s]\n", r.RequestURI, r.Method, userAgent)
	fmt.Fprint(w, pageHeader)
	for _,stack := range d.Stacks {
		fmt.Fprint(w, "<section class=\"hero is-light\">")
		fmt.Fprint(w, "<div class=\"hero-body\">")
		fmt.Fprint(w, "<div class=\"container\">")
		tableTitle := fmt.Sprintf("<h2 class=\"title\">%s</h2>", stack.Name)

		fmt.Fprint(w, tableTitle)
		fmt.Fprint(w, tableHeader)
		for _,service := range stack.Services {

			if len(service.ProxyConfigurations) > 0 {
				serviceName := service.Name
				if service.Alias != "" {
					serviceName = service.Alias
				}
				serverListItem := fmt.Sprintf("<td><a href=\"%s\">%s</td><td>%s</td>", service.ProxyConfigurations[0].ServicePath, serviceName, service.ProxyConfigurations[0].ServiceDomain)
				fmt.Fprint(w, "<tr>")
				fmt.Fprint(w, serverListItem)
				fmt.Fprint(w, "</tr>")
			}
		}
		fmt.Fprint(w, tableFooter)
		fmt.Fprint(w,"</div>")
		fmt.Fprint(w,"</div>")
		fmt.Fprint(w,"</section>")
	}
	fmt.Fprint(w, pageFooter)
}