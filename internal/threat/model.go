package threat

import (
	"time"
)

// Threat представляет полную информацию об угрозе, хранящуюся в базе данных.
type Threat struct {
	MD5         string    `json:"md5"`
	Name        string    `json:"name"`
	Severity    string    `json:"severity"`
	Type        string    `json:"type"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ThreatCreateRequest содержит данные, необходимые для создания новой угрозы.
type ThreatCreateRequest struct {
	MD5         string `json:"md5"`
	Name        string `json:"name"`
	Severity    string `json:"severity"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// ThreatUpdateRequest содержит данные для обновления существующей угрозы.
type ThreatUpdateRequest struct {
	Name        string `json:"name"`
	Severity    string `json:"severity"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
