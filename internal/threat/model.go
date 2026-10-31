package threat

import (
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("угроза не найдена")
	ErrConflict = errors.New("угроза уже была создана")
)

type Threat struct {
	MD5         string    `json:"md5"`
	Name        string    `json:"name"`
	Severity    string    `json:"severity"`
	Type        string    `json:"type"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ThreatCreateRequest struct {
	MD5         string `json:"md5"`
	Name        string `json:"name"`
	Severity    string `json:"severity"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type ThreatUpdateRequest struct {
	Name        string `json:"name"`
	Severity    string `json:"severity"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
