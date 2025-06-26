package models

import "time"

type Package struct {
    Name         string    `json:"name"`
    Version      string    `json:"version"`
    Source       string    `json:"source"`        // "apt" or "github"
    InstallPath  string    `json:"install_path"`
    Installed    bool      `json:"installed"`
    InstalledDate time.Time `json:"installed_date"`
}