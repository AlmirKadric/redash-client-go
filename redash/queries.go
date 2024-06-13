package redash

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

// QueriesList models the response from Redash's /api/queries endpoint
type QueriesList struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Results  []struct {
		ID                int           `json:"id"`
		IsArchived        bool          `json:"is_archived"`
		CreatedAt         time.Time     `json:"created_at"`
		RetrievedAt       time.Time     `json:"retrieved_at"`
		UpdatedAt         time.Time     `json:"updated_at"`
		Name              string        `json:"name"`
		Description       string        `json:"description"`
		Query             string        `json:"query"`
		QueryHash         string        `json:"query_hash"`
		Version           int           `json:"version"`
		LastModifiedByID  int           `json:"last_modified_by_id"`
		Tags              []string      `json:"tags"`
		APIKey            string        `json:"api_key"`
		DataSourceID      int           `json:"data_source_id"`
		LatestQueryDataID int           `json:"latest_query_data_id"`
		Schedule          QuerySchedule `json:"schedule"`
		User              User          `json:"user"`
		IsFavorite        bool          `json:"is_favorite"`
		IsDraft           bool          `json:"is_draft"`
		IsSafe            bool          `json:"is_safe"`
		Runtime           float32       `json:"runtime"`
		Options           QueryOptions  `json:"options"`
	}
}

// Query models the response from Redash's /api/queries endpoint
type Query struct {
	ID                int             `json:"id"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	Query             string          `json:"query"`
	QueryHash         string          `json:"query_hash"`
	Version           int             `json:"version"`
	Schedule          QuerySchedule   `json:"schedule"`
	APIKey            string          `json:"api_key"`
	IsArchived        bool            `json:"is_archived"`
	IsDraft           bool            `json:"is_draft"`
	UpdatedAt         time.Time       `json:"updated_at"`
	CreatedAt         time.Time       `json:"created_at"`
	DataSourceID      int             `json:"data_source_id"`
	LatestQueryDataID int             `json:"latest_query_data_id"`
	Tags              []string        `json:"tags"`
	IsSafe            bool            `json:"is_safe"`
	User              User            `json:"user"`
	LastModifiedBy    User            `json:"last_modified_by"`
	IsFavorite        bool            `json:"is_favorite"`
	CanEdit           bool            `json:"can_edit"`
	Options           QueryOptions    `json:"options"`
	Visualizations    []Visualization `json:"visualizations"`
}

// QuerySchedule struct
type QuerySchedule struct {
	Interval  int         `json:"interval"`
	Time      string      `json:"time"`
	DayOfWeek string      `json:"day_of_week"`
	Until     interface{} `json:"until"`
}

// QueryOptions struct
type QueryOptions struct {
	Parameters []QueryOptionsParameter `json:"parameters"`
}

// QueryOptionsParameter struct
type QueryOptionsParameter struct {
	Title       string        `json:"title"`
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	EnumOptions string        `json:"enum_options"`
	Locals      []interface{} `json:"locals"`
	Value       interface{}   `json:"value"`
}

// QueryCreatePayload defines the schema for creating a new Redash query
type QueryCreatePayload struct {
	Name         string       `json:"name,omitempty"`
	Description  string       `json:"description,omitempty"`
	Query        string       `json:"query,omitempty"`
	DataSourceID int          `json:"data_source_id,omitempty"`
	IsDraft      bool         `json:"is_draft,omitempty"`
	Options      QueryOptions `json:"options,omitempty"`
	Version      int          `json:"version,omitempty"`
}

// QueryUpdatePayload defines the schema for updating a Redash query
type QueryUpdatePayload struct {
	Name         string       `json:"name,omitempty"`
	Description  string       `json:"description,omitempty"`
	Query        string       `json:"query,omitempty"`
	DataSourceID int          `json:"data_source_id,omitempty"`
	IsDraft      bool         `json:"is_draft,omitempty"`
	Options      QueryOptions `json:"options,omitempty"`
	Version      int          `json:"version,omitempty"`
}

// GetQueries returns a paginated list of queries
func (c *Client) GetQueries() (*QueriesList, error) {
	path := "/api/queries"

	queryParams := url.Values{}
	response, err := c.get(path, queryParams)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	queries := new(QueriesList)
	err = json.NewDecoder(response.Body).Decode(queries)
	if err != nil {
		return nil, err
	}

	return queries, nil
}

// GetQuery gets a specific query
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
	json.NewDecoder(response.Body).Decode(newQuery)
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
