package config

import (
	"errors"
	"inviter/hash"
	"log"
	"os"
	"strings"
)

type AppConfig struct {
	OrgName    string //mandatory
	Token      string //mandatory
	GroupName  string //optional
	InviteCode []byte //optional
	HttpPort   string //optional (default 80)
	HttpsPort  string //optional (default 443)
	TlsCert    string //optional
	TlsKey     string //optional
}

var conf AppConfig

type EnabledFeatures struct {
	Code  bool
	Tls   bool
	Group bool
}

var features EnabledFeatures

func Load() {
	// Check for mandatory environment variables
	orgName := strings.Trim(os.Getenv("GITHUB_ORG_NAME"), " ")
	if len(orgName) == 0 {
		log.Fatal("GITHUB_ORG_NAME environment variable must be set")
	}

	token := strings.Trim(os.Getenv("GITHUB_TOKEN"), " ")
	if len(token) == 0 {
		log.Fatal("GITHUB_TOKEN environment variable must be set")
	}

	groupName := strings.Trim(os.Getenv("GITHUB_GROUP_NAME"), " ")
	if len(groupName) != 0 {
		features.Group = true
	}

	inviteCode := strings.Trim(os.Getenv("INVITE_CODE_HASH"), " ")
	if len(inviteCode) != 0 {
		features.Code = true
	}

	// Set the optional environment variables, using defaults if not set
	httpPort := strings.Trim(os.Getenv("HTTP_PORT"), " ")
	if len(httpPort) == 0 {
		httpPort = "80"
	}
	httpsPort := strings.Trim(os.Getenv("HTTPS_PORT"), " ")
	if len(httpsPort) == 0 {
		httpsPort = "443"
	}

	conf = AppConfig{
		OrgName:    orgName,
		Token:      token,
		GroupName:  strings.ToLower(groupName),
		InviteCode: hash.HexToByteArray(inviteCode),
		HttpPort:   httpPort,
		HttpsPort:  httpsPort,
		TlsCert:    strings.Trim(os.Getenv("TLS_CERT"), " "),
		TlsKey:     strings.Trim(os.Getenv("TLS_KEY"), " "),
	}

	if len(conf.TlsCert) > 0 && len(conf.TlsKey) > 0 {
		if _, err := os.Stat(conf.TlsCert); errors.Is(err, os.ErrNotExist) {
			log.Fatalf("Certificate file: %s does not exist", conf.TlsCert)
		}
		if _, err := os.Stat(conf.TlsKey); errors.Is(err, os.ErrNotExist) {
			log.Fatalf("Key file: %s does not exist", conf.TlsKey)
		}

		features.Tls = true
	}
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

func HttpPort() string {
	return conf.HttpPort
}

func HttpsPort() string {
	return conf.HttpsPort
}

func TlsCert() string {
	return conf.TlsCert
}

func TlsKey() string {
	return conf.TlsKey
}

func IsTlsEnable() bool {
	return features.Tls
}

func IsCodeRequired() bool {
	return features.Code
}

func IsGroupEnable() bool {
	return features.Group
}
