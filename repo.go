package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Repo struct {
	Name       string   `json:"name"`
	IconURL    string   `json:"iconURL"`
	Identifier string   `json:"identifier"`
	Apps       []App    `json:"apps"`
	News       []string `json:"news"`
}

type App struct {
	Name                 string   `json:"name"`
	BundleIdentifier     string   `json:"bundleIdentifier"`
	DeveloperName        string   `json:"developerName"`
	Version              string   `json:"version"`
	VersionDate          string   `json:"versionDate"`
	DownloadURL          string   `json:"downloadURL"`
	LocalizedDescription string   `json:"localizedDescription"`
	IconURL              string   `json:"iconURL"`
	TintColor            string   `json:"tintColor"`
	Size                 int      `json:"size"`
	ScreenshotURLs       []string `json:"screenshotURLs"`
}

func getReleases(repo string) (*Repo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var releases []struct {
		Name    string `json:"name"`
		Version string `json:"tag_name"`
		Assets  []struct {
			BrowserDownloadURL string `json:"browser_download_url"`
			VersionDate        string `json:"created_at"`
			Size               int    `json:"size"`
			Uploader           struct {
				AvatarURL string `json:"avatar_url"`
			} `json:"uploader"`
		} `json:"assets"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	r := Repo{
		Name:       repo,
		Identifier: repo,
		IconURL:    "https://cdn.simpleicons.org/github",
		Apps:       []App{},
		News:       []string{},
	}
	for _, release := range releases {
		for _, asset := range release.Assets {
			if asset.BrowserDownloadURL == "" || !strings.HasSuffix(asset.BrowserDownloadURL, ".ipa") {
				continue
			}
			r.Apps = append(r.Apps, App{
				Name:                 release.Name,
				BundleIdentifier:     repo,
				DeveloperName:        repo,
				Version:              release.Version,
				VersionDate:          asset.VersionDate,
				DownloadURL:          asset.BrowserDownloadURL,
				LocalizedDescription: release.Name,
				IconURL:              asset.Uploader.AvatarURL,
				Size:                 asset.Size,
			})
		}
	}

	return &r, nil
}
