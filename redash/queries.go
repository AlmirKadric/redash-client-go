package redash

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

// QueryList object structure from Redash's /api/queries endpoint
type QueryList struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Results  []QueryListItem
}

// Query object structure for QueryList items
type QueryListItem struct {
	// Base Data
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	// Query
	DataSourceID int    `json:"data_source_id"`
	Query        string `json:"query"`
	QueryHash    string `json:"query_hash"`

	// Options
	Options QueryOptions `json:"options"`

	// State
	IsDraft    bool `json:"is_draft"`
	IsArchived bool `json:"is_archived"`
	IsSafe     bool `json:"is_safe"`
	Version    int  `json:"version"`

	// User
	User User `json:"user"`

	// Timestamps
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	// Metadata
	APIKey            string        `json:"api_key"`
	Tags              []string      `json:"tags"`
	LatestQueryDataID int           `json:"latest_query_data_id,omitempty"`
	Schedule          QuerySchedule `json:"schedule"`

	// List Item Specific
	LastModifiedByID int       `json:"last_modified_by_id"`
	IsFavorite       bool      `json:"is_favorite"`
	RetrievedAt      time.Time `json:"retrieved_at"`
	Runtime          float32   `json:"runtime"`
}

// Query object structure from Redash's /api/queries/<ID> endpoint
type Query struct {
	// Base Data
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	// Query
	DataSourceID int    `json:"data_source_id"`
	Query        string `json:"query"`
	QueryHash    string `json:"query_hash"`

	// Options
	Options QueryOptions `json:"options"`

	// State
	IsDraft    bool `json:"is_draft"`
	IsArchived bool `json:"is_archived"`
	IsSafe     bool `json:"is_safe"`
	Version    int  `json:"version"`

	// User
	User User `json:"user"`

	// Timestamps
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	// Metadata
	APIKey            string        `json:"api_key"`
	Tags              []string      `json:"tags"`
	LatestQueryDataID int           `json:"latest_query_data_id,omitempty"`
	Schedule          QuerySchedule `json:"schedule"`

	// Query Specific
	LastModifiedBy User                 `json:"last_modified_by"`
	IsFavorite     bool                 `json:"is_favorite"`
	CanEdit        bool                 `json:"can_edit"`
	Visualizations []VisualizationQuery `json:"visualizations"`
}

// Query object structure for Dashboard Widget Visualizations
type QueryDashboard struct {
	// Base Data
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	// Query
	DataSourceID int    `json:"data_source_id"`
	Query        string `json:"query"`
	QueryHash    string `json:"query_hash"`

	// Options
	Options QueryOptions `json:"options"`

	// State
	IsDraft    bool `json:"is_draft"`
	IsArchived bool `json:"is_archived"`
	IsSafe     bool `json:"is_safe"`
	Version    int  `json:"version"`

	// User
	User User `json:"user"`

	// Timestamps
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	// Metadata
	APIKey            string        `json:"api_key"`
	Tags              []string      `json:"tags"`
	LatestQueryDataID int           `json:"latest_query_data_id,omitempty"`
	Schedule          QuerySchedule `json:"schedule"`

	// Dashboard Specific
	LastModifiedBy User `json:"last_modified_by"`
}

type QuerySchedule struct {
	Interval  int         `json:"interval"`
	Time      string      `json:"time"`
	DayOfWeek string      `json:"day_of_week"`
	Until     interface{} `json:"until"`
}

type QueryOptions struct {
	Parameters []QueryOptionsParameter `json:"parameters"`
}

type QueryOptionsParameter struct {
	Name  string `json:"name"`
	Title string `json:"title"`

	ParentQueryId int `json:"parentQueryId"`

	Locals []interface{} `json:"locals"`

	Type        string      `json:"type"`
	Value       interface{} `json:"value"`
	EnumOptions string      `json:"enumOptions,omitempty"`

	Global bool `json:"global"`
}

// QueryCreatePayload defines the schema for creating a new Redash query
type QueryCreatePayload struct {
	// Base Data
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	// Query
	DataSourceID int    `json:"data_source_id"`
	Query        string `json:"query"`
	QueryHash    string `json:"query_hash"`

	// Options
	Options QueryOptions `json:"options"`

	// State
	IsDraft    bool `json:"is_draft"`
	IsArchived bool `json:"is_archived"`
	IsSafe     bool `json:"is_safe"`
	Version    int  `json:"version"`

	// Metadata
	APIKey            string        `json:"api_key"`
	Tags              []string      `json:"tags"`
	LatestQueryDataID int           `json:"latest_query_data_id,omitempty"`
	Schedule          QuerySchedule `json:"schedule"`

	// Query Specific
	IsFavorite bool `json:"is_favorite"`
	CanEdit    bool `json:"can_edit"`
}

// QueryUpdatePayload defines the schema for updating a Redash query
type QueryUpdatePayload struct {
	// Base Data
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	// Query
	DataSourceID int    `json:"data_source_id"`
	Query        string `json:"query"`
	QueryHash    string `json:"query_hash"`

	// Options
	Options QueryOptions `json:"options"`

	// State
	IsDraft    bool `json:"is_draft"`
	IsArchived bool `json:"is_archived"`
	IsSafe     bool `json:"is_safe"`
	Version    int  `json:"version"`

	// Metadata
	APIKey            string        `json:"api_key"`
	Tags              []string      `json:"tags"`
	LatestQueryDataID int           `json:"latest_query_data_id,omitempty"`
	Schedule          QuerySchedule `json:"schedule"`

	// Query Specific
	IsFavorite bool `json:"is_favorite"`
	CanEdit    bool `json:"can_edit"`
}

// GetQueries returns a list of Redash queries
func (c *Client) GetQueries() (*QueryList, error) {
	path := "/api/queries"

	queryParams := url.Values{}
	response, err := c.get(path, queryParams)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	queries := new(QueryList)
	err = json.NewDecoder(response.Body).Decode(queries)
	if err != nil {
		return nil, err
	}

	return queries, nil
}

// GetQuery returns a specific Redash query by its ID
func (c *Client) GetQuery(id int) (*Query, error) {
	path := "/api/queries/" + strconv.Itoa(id)

	queryParams := url.Values{}
	response, err := c.get(path, queryParams)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	query := new(Query)
	err = json.NewDecoder(response.Body).Decode(query)
	if err != nil {
		return nil, err
	}

	return query, nil
}

// CreateQuery creates a new Redash query
func (c *Client) CreateQuery(query *QueryCreatePayload) (*Query, error) {
	path := "/api/queries"

	payload, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	response, err := c.post(path, string(payload), queryParams)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	newQuery := new(Query)
	err = json.NewDecoder(response.Body).Decode(newQuery)
	if err != nil {
		return nil, err
	}

	return newQuery, nil
}

// UpdateQuery updates an existing Redash query
func (c *Client) UpdateQuery(id int, query *QueryUpdatePayload) (*Query, error) {
	path := "/api/queries/" + strconv.Itoa(id)

	payload, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	response, err := c.post(path, string(payload), queryParams)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	newQuery := new(Query)
	err = json.NewDecoder(response.Body).Decode(newQuery)
	if err != nil {
		return nil, err
	}

	return newQuery, nil
}

// ArchiveQuery archives an existing Redash query
func (c *Client) ArchiveQuery(id int) error {
	path := "/api/queries/" + strconv.Itoa(id)

	_, err := c.delete(path, url.Values{})

	return err
}
