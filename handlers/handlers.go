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
	"sync"
)

var (
	templateCache *template.Template
	logoCache     string
	cacheMutex    sync.RWMutex
)

func InitCache() error {
	cacheMutex.RLock()
	var err error
	templateCache, err = template.ParseFiles("templates/index.html")
	if err != nil {
		cacheMutex.RUnlock()
		return err
	}

	logoCache, err = github.GetOrgLogoUrl(config.OrgName())
	if err != nil {
		cacheMutex.RUnlock()
		return err
	}

	cacheMutex.RUnlock()
	return nil
}

func MainPage(w http.ResponseWriter, _ *http.Request) {
	cacheMutex.RLock()
	cachedTemplate := templateCache
	cachedLogo := logoCache
	cacheMutex.RUnlock()

	data := struct {
		OrgName  string
		LogoURL  string
		TeamName string
	}{
		OrgName:  config.OrgName(),
		LogoURL:  cachedLogo,
		TeamName: config.GroupName(),
	}

	err := cachedTemplate.Execute(w, data)
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
	inviteCode := strings.Trim(r.FormValue("inviteCode"), " ")
	if inviteCode == "" {
		http.Error(w, "Invite code is required", http.StatusBadRequest)
		return
	}
	if !hash.Compare(inviteCode, config.InviteCode()) {
		http.Error(w, "Invalid username or invitation code", http.StatusUnauthorized)
		log.Printf("User: %s, tried to access with code: %s", username, inviteCode)
		return
	}

	err := github.AddUserToGroup(username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add user to group: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/success", http.StatusSeeOther)
}
