package main

import (
    "html/template"
    "net/http"
)

// Função para servir a página inicial
func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(
        "views/profile_views/home.html",
        "views/components/header.html",
        "views/components/professional_card.html",
        "views/profile_views/registration.html",
    ))
    tmpl.ExecuteTemplate(w, "home.html", nil)
}

// Função para servir a página de cadastro
func registrationHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(
        "views/profile_views/registration.html",
        "views/components/header.html",
    ))
    tmpl.ExecuteTemplate(w, "registration.html", nil)
}

func main() {
    mux := http.NewServeMux()

    // Servir arquivos estáticos
    mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

    // Rota para a página inicial
    mux.HandleFunc("/", homeHandler)

    // Rota para a pagina de resitrar
    mux.HandleFunc("/registration", registrationHandler)
    
    // Iniciar o servidor
    http.ListenAndServe(":8080", mux)
}
