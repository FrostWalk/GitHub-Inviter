package handlers

import (
	"fmt"
	"html/template"
	"inviter/config"
	"inviter/github"
	"inviter/hash"
	"log"
	"net/http"
	"strings"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	// Fetch organization logo
	logoURL, err := github.GetOrgLogo(config.OrgName())
	if err != nil {
		http.Error(w, "Failed to fetch organization logo", http.StatusInternalServerError)
		return
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		OrgName  string
		LogoURL  string
		TeamName string
	}{
		OrgName:  config.OrgName(),
		LogoURL:  logoURL,
		TeamName: config.GroupName(),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}
	password := strings.Trim(r.FormValue("password"), " ")
	if password == "" {
		http.Error(w, "Invitation code is required", http.StatusBadRequest)
		return
	}
	if !hash.Compare(password, config.Password()) {
		http.Error(w, "Invalid username or invitation code", http.StatusUnauthorized)
		log.Printf("User: %s, tried to access with code: %s", username, password)
		return
	}

	err := github.AddUserToGroup(username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add user to group: %v", err), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "User %s successfully added to the group %s in organization %s", username, config.GroupName(), config.OrgName())
	if err != nil {
		return
	}
}
