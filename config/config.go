package config

import (
	"inviter/hash"
	"log"
	"os"
	"strings"
)

type AppConfig struct {
	OrgName    string //mandatory
	Token      string //mandatory
	GroupName  string //mandatory
	InviteCode []byte //mandatory
	Port       string //optional (default 8080)
	TlsCert    string //optional
	TlsKey     string //optional
}

var conf AppConfig

func Load() bool {
	// Check for mandatory environment variables
	orgName := strings.Trim(os.Getenv("GITHUB_ORG_NAME"), " ")
	if len(orgName) == 0 {
		log.Fatal("GITHUB_ORG_NAME environment variable must be set")
	}

	token := strings.Trim(os.Getenv("GITHUB_TOKEN"), " ")
	if len(token) == 0 {
		log.Fatal("GITHUB_TOKEN environment variable must be set")
	}

	inviteCode := strings.Trim(os.Getenv("INVITE_CODE"), " ")
	if len(inviteCode) == 0 {
		log.Fatal("INVITE_CODE environment variable must be set")
	}

	groupName := strings.Trim(os.Getenv("GITHUB_GROUP_NAME"), " ")
	if len(groupName) == 0 {
		log.Fatal("GROUP_NAME environment variable must be set")
	}

	// Set the optional environment variables, using defaults if not set
	port := strings.Trim(os.Getenv("PORT"), " ")
	if len(port) == 0 {
		port = "8080"
	}

	conf = AppConfig{
		OrgName:    orgName,
		Token:      token,
		GroupName:  strings.ToLower(groupName),
		InviteCode: hash.CalculateHash(inviteCode),
		Port:       port,
		TlsCert:    strings.Trim(os.Getenv("TLS_CERT"), " "),
		TlsKey:     strings.Trim(os.Getenv("TLS_KEY"), " "),
	}
	return len(os.Getenv("TLS_CERT")) > 0 && len(os.Getenv("TLS_KEY")) > 0
}

func OrgName() string {
	return conf.OrgName
}

func Token() string {
	return conf.Token
}

func GroupName() string {
	return conf.GroupName
}

func InviteCode() []byte {
	return conf.InviteCode
}

func Port() string {
	return conf.Port
}

func TlsCert() string {
	return conf.TlsCert
}

func TlsKey() string {
	return conf.TlsKey
}
