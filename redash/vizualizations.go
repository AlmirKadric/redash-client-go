package redash

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// Visualization object structure for Queries
type VisualizationQuery struct {
	// Base Data
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	// Options
	Type    string      `json:"type"`
	Options interface{} `json:"options"`

	// Timestamps
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// Visualization object structure for Dashboard Widget
type VisualizationDashboard struct {
	// Base Data
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	// Options
	Type    string      `json:"type"`
	Options interface{} `json:"options"`

	// Timestamps
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	// Dashboard Specific
	Query QueryDashboard `json:"query"`
}

// TABLE Options
type TableOptions struct {
	ItemsPerPage int           `json:"itemsPerPage"`
	Columns      []TableColumn `json:"columns"`
}

type TableColumn struct {
	// Shared
	Visible bool   `json:"visible"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	// Type
	Type         string `json:"type"`
	DisplayAs    string `json:"displayAs"`
	AlignContent string `json:"alignContent"`
	AllowSearch  bool   `json:"allowSearch"`
	Order        int    `json:"order"`
	// Text
	AllowHTML      bool `json:"allowHTML"`
	HighlightLinks bool `json:"highlightLinks"`
	// Number
	NumberFormat string `json:"numberFormat,omitempty"`
	// Date/Time
	DateTimeFormat string `json:"dateTimeFormat,omitempty"`
	// Boolean
	BooleanValues []string `json:"booleanValues"`
	// Link
	LinkUrlTemplate   string `json:"linkUrlTemplate"`
	LinkTitleTemplate string `json:"linkTitleTemplate"`
	LinkTextTemplate  string `json:"linkTextTemplate"`
	LinkOpenInNewTab  bool   `json:"linkOpenInNewTab"`
	// Image
	ImageUrlTemplate   string `json:"imageUrlTemplate"`
	ImageTitleTemplate string `json:"imageTitleTemplate"`
	ImageWidth         string `json:"imageWidth"`
	ImageHeight        string `json:"imageHeight"`
	// JSON
}

// CHART Options
type ChartOptions struct {
	// General
	GlobalSeriesType    string              `json:"globalSeriesType"`
	ColumnMapping       ChartColumnsMapping `json:"columnMapping"`
	ErrorY              ChartErrorY         `json:"error_y"`
	Legend              ChartLegend         `json:"legend"`
	Series              ChartSeries         `json:"series"`
	MissingValuesAsZero bool                `json:"missingValuesAsZero"`
	// X-Axis
	// -scale
	// -name
	XAxis ChartXAxis `json:"xAxis"`
	SortX bool       `json:"sortX"`
	// Y-Axis
	YAxis []ChartYAxis `json:"yAxis"`
	// Series
	SeriesOptions ChartSeriesOptions `json:"seriesOptions"`
	// Colors
	// Data Labels
	ShowDataLabels bool   `json:"showDataLabels"`
	NumberFormat   string `json:"numberFormat"`
	PercentFormat  string `json:"percentFormat"`
	DateTimeFormat string `json:"dateTimeFormat"`
	TextFormat     string `json:"textFormat"`
	// Unknown, not mapped yet
	// "direction": {
	// 	"type": "counterclockwise"
	// },
	// "valuesOptions": {},
	// "customCode": "// Available variables are x, ys, element, and Plotly\n// Type console.log(x, ys); for more info about x and ys\n// To plot your graph call Plotly.plot(element, ...)\n// Plotly examples and docs: https://plot.ly/javascript/",
}

type ChartColumnsMapping map[string]string

type ChartLegend struct {
	Enabled bool `json:"enabled"`
	// Placement string `json:"placement"`
}

type ChartSeries struct {
	Stacking string      `json:"stacking,omitempty"`
	ErrorY   ChartErrorY `json:"error_y"`
}

type ChartXAxis struct {
	Type   string `json:"type"`
	Labels struct {
		Enabled bool `json:"enabled"`
	} `json:"labels"`
}

type ChartYAxis struct {
	Type     string `json:"type"`
	Opposite bool   `json:"opposite"`
}

type ChartErrorY struct {
	Visible bool   `json:"visible"`
	Type    string `json:"type"`
}

type ChartSeriesOptions map[string]ChartSeriesOption

type ChartSeriesOption struct {
	ZIndex int    `json:"zIndex"`
	Index  int    `json:"index"`
	Type   string `json:"type"`
	YAxis  int    `json:"yAxis"`
}

// VisualizationCreatePayload defines the schema for creating a Redash visualizations
type VisualizationCreatePayload struct {
	// Base Data
	Name        string `json:"name"`
	Description string `json:"description"`

	// Options
	Type    string      `json:"type"`
	Options interface{} `json:"options"`

	// References
	QueryId int `json:"query_id"`
}

// VisualizationUpdatePayload defines the schema for updating a Redash visualizations
type VisualizationUpdatePayload struct {
	// Base Data
	Name        string `json:"name"`
	Description string `json:"description"`

	// Options
	Type    string      `json:"type"`
	Options interface{} `json:"options"`
}

// GetVisualization gets a specific Redash visualization by its query and visualization ID
func (c *Client) GetVisualization(queryId, visualizationId int) (*VisualizationQuery, error) {
	query, err := c.GetQuery(queryId)
	if err != nil {
		return nil, err
	}

	for _, v := range query.Visualizations {
		if v.ID == visualizationId {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("visualization %d not found in query %d", visualizationId, queryId)
}

// CreateVisualization creates a new Redash visualization
func (c *Client) CreateVisualization(visualizationCreatePayload *VisualizationCreatePayload) (*VisualizationQuery, error) {
	path := "/api/visualizations"

	payload, err := json.Marshal(visualizationCreatePayload)
	if err != nil {
		return nil, err
	}

	response, err := c.post(path, string(payload), url.Values{})
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	newVisualization := new(VisualizationQuery)
	err = json.NewDecoder(response.Body).Decode(newVisualization)
	if err != nil {
		return nil, err
	}

	return newVisualization, nil
}

// UpdateVisualization updates an existing Redash visualization
func (c *Client) UpdateVisualization(id int, visualizationUpdatePayload *VisualizationUpdatePayload) (*VisualizationQuery, error) {
	path := "/api/visualizations/" + strconv.Itoa(id)

	payload, err := json.Marshal(visualizationUpdatePayload)
	if err != nil {
		return nil, err
	}

	response, err := c.post(path, string(payload), url.Values{})
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	newVisualization := new(VisualizationQuery)
	json.NewDecoder(response.Body).Decode(newVisualization)
	if err != nil {
		return nil, err
	}

	return newVisualization, nil
}

// DeleteVisualization deletes a Redash visualization
func (c *Client) DeleteVisualization(id int) error {
	path := "/api/visualizations/" + strconv.Itoa(id)

	_, err := c.delete(path, url.Values{})

	return err
}
