package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"inviter/config"
	"io"
	"net/http"
)

const baseUrl = "https://api.github.com"

func AddUserToGroup(username string) error {
	url := fmt.Sprintf("%s/orgs/%s/teams/%s/memberships/%s", baseUrl, config.OrgName(), config.GroupName(), username)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(`{"role":"member"}`)))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+config.Token())
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("the token does not have the required permissions")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to add user to group: %s", string(body))
	}

	return nil
}

func AddUserToOrg(username string) error {
	url := fmt.Sprintf("%s/orgs/%s/memberships/%s", baseUrl, config.OrgName(), username)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(`{"role":"member"}`)))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+config.Token())
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("the token does not have the required permissions")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to add user to group: %s", string(body))
	}

	return nil
}

func GetOrgLogoUrl(orgName string) (string, error) {
	url := fmt.Sprintf("%s/orgs/%s", baseUrl, orgName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+config.Token())
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return "", fmt.Errorf("provided token is invalid")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	logoURL, ok := result["avatar_url"].(string)
	if !ok {
		return "", fmt.Errorf("logo URL not found")
	}

	return logoURL, nil
}
