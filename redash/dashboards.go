package redash

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

// DashboardList object structure from Redash's /api/dashboards endpoint
type DashboardList struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Results  []DashboardListItem
}

// DashboardListItem object structure for DashboardList items
type DashboardListItem struct {
	// Base Data
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`

	// Options
	Layout []interface{} `json:"layout"`

	// State
	IsFavorite              bool `json:"is_favorite"`
	IsArchived              bool `json:"is_archived"`
	IsDraft                 bool `json:"is_draft"`
	DashboardFiltersEnabled bool `json:"dashboard_filters_enabled"`
	Version                 int  `json:"version"`

	// User
	UserID int  `json:"user_id"`
	User   User `json:"user"`

	// Timestamps
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	// Metadata
	Tags []string `json:"tags"`

	// List Item Specific
	// Widgets                 null      `json:"widgets"`
}

// Dashboard object structure from Redash's /api/dashboards/<SLUG> endpoint
type Dashboard struct {
	// Base Data
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`

	// Options
	Layout []interface{} `json:"layout"`

	// State
	IsFavorite              bool `json:"is_favorite"`
	IsArchived              bool `json:"is_archived"`
	IsDraft                 bool `json:"is_draft"`
	DashboardFiltersEnabled bool `json:"dashboard_filters_enabled"`
	Version                 int  `json:"version"`

	// User
	UserID int  `json:"user_id"`
	User   User `json:"user"`

	// Timestamps
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	// Metadata
	Tags []string `json:"tags"`

	// Dashboard Specific
	PublicUrl string            `json:"public_url"`
	CanEdit   bool              `json:"can_edit"`
	Widgets   []WidgetDashboard `json:"widgets"`
	APIKey    string            `json:"api_key"`
}

// DashboardCreatePayload defines the schema for creating a Redash dashboards
type DashboardCreatePayload struct {
	// Base Data
	Name string `json:"name"`
	Slug string `json:"slug"`

	// Options
	// Layout                  []interface{}     `json:"layout"`

	// State
	IsFavorite              bool `json:"is_favorite"`
	IsArchived              bool `json:"is_archived"`
	IsDraft                 bool `json:"is_draft"`
	DashboardFiltersEnabled bool `json:"dashboard_filters_enabled"`

	// Metadata
	Tags []string `json:"tags"`
}

// DashboardUpdatePayload defines the schema for updating a Redash dashboards
type DashboardUpdatePayload struct {
	// Base Data
	Name string `json:"name"`
	Slug string `json:"slug"`

	// Options
	// Layout                  []interface{}     `json:"layout"`

	// State
	IsFavorite              bool `json:"is_favorite"`
	IsArchived              bool `json:"is_archived"`
	IsDraft                 bool `json:"is_draft"`
	DashboardFiltersEnabled bool `json:"dashboard_filters_enabled"`

	// Metadata
	Tags []string `json:"tags"`
}

// GetDashboard gets a specific dashboard by its slug
func (c *Client) GetDashboard(slug string) (*Dashboard, error) {
	path := "/api/dashboards/" + slug

	queryParams := url.Values{}
	response, err := c.get(path, queryParams)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	dashboard := new(Dashboard)
	err = json.NewDecoder(response.Body).Decode(dashboard)
	if err != nil {
		return nil, err
	}

	return dashboard, nil
}

// CreateDashboard creates a new Redash dashboard
func (c *Client) CreateDashboard(dashboard *DashboardCreatePayload) (*Dashboard, error) {
	path := "/api/dashboards"

	payload, err := json.Marshal(dashboard)
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	response, err := c.post(path, string(payload), queryParams)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	newDashboard := new(Dashboard)
	err = json.NewDecoder(response.Body).Decode(newDashboard)
	if err != nil {
		return nil, err
	}

	return newDashboard, nil
}

// UpdateDashboard updates an existing Redash dashboard
func (c *Client) UpdateDashboard(id int, dashboard *DashboardUpdatePayload) (*Dashboard, error) {
	path := "/api/dashboards/" + strconv.Itoa(id)

	payload, err := json.Marshal(dashboard)
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	response, err := c.post(path, string(payload), queryParams)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	newDashboard := new(Dashboard)
	err = json.NewDecoder(response.Body).Decode(newDashboard)
	if err != nil {
		return nil, err
	}

	return newDashboard, nil
}

// ArchiveDashboard archives an existing dashboard
func (c *Client) ArchiveDashboard(slug string) error {
	path := "/api/dashboards/" + slug

	_, err := c.delete(path, url.Values{})

	return err
}
