package redash

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// Visualization struct
type Visualization struct {
	ID          int                  `json:"id"`
	Type        string               `json:"type"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Options     VisualizationOptions `json:"options"`
	UpdatedAt   time.Time            `json:"updated_at"`
	CreatedAt   time.Time            `json:"created_at"`
}

// VisualizationOptions struct
type VisualizationOptions struct {
	// CHART TYPE
	// General
	GlobalSeriesType string            `json:"globalSeriesType,omitempty"`
	ColumnMapping    map[string]string `json:"columnMapping,omitempty"`
	// "error_y": {
	// 	"visible": true,
	// 	"type": "data"
	// },
	Legend              VisualizationLegendOptions `json:"legend,omitempty"`
	Series              Series                     `json:"series,omitempty"`
	MissingValuesAsZero bool                       `json:"missingValuesAsZero,omitempty"`
	// X-Axis
	// -scale
	// -name
	XAxis VisualizationAxisOptions `json:"xAxis,omitempty"`
	SortX bool                     `json:"sortX,omitempty"`
	// Y-Axis
	YAxis []VisualizationAxisOptions `json:"yAxis,omitempty"`
	// Series
	SeriesOptions map[string]SeriesOptions `json:"seriesOptions,omitempty"`
	// Colors
	// Data Labels
	ShowDataLabels bool   `json:"showDataLabels,omitempty"`
	NumberFormat   string `json:"numberFormat,omitempty"`
	PercentFormat  string `json:"percentFormat,omitempty"`
	DateTimeFormat string `json:"dateTimeFormat,omitempty"`
	TextFormat     string `json:"textFormat,omitempty"`
	// Unknown
	// "direction": {
	// 	"type": "counterclockwise"
	// },
	// "valuesOptions": {},
	// "customCode": "// Available variables are x, ys, element, and Plotly\n// Type console.log(x, ys); for more info about x and ys\n// To plot your graph call Plotly.plot(element, ...)\n// Plotly examples and docs: https://plot.ly/javascript/",

	// TABLE TYPE
	ItemsPerPage     int                          `json:"itemsPerPage,omitempty"`
	Columns          []VisualizationColumnOptions `json:"columns,omitempty"`
}

type Series struct {
	Stacking string `json:"stacking"`
}

type SeriesOptions struct {
	ZIndex int    `json:"zIndex"`
	Index  int    `json:"index"`
	Type   string `json:"type"`
	YAxis  int    `json:"yAxis"`
}

// VisualizationLegendOptions struct
type VisualizationLegendOptions struct {
	Enabled   bool   `json:"enabled"`
	Placement string `json:"placement"`
}

// VisualizationAxisOptions struct
type VisualizationAxisOptions struct {
	Type     string                    `json:"type"`
	Opposite bool                      `json:"opposite"`
	Labels   VisualizationLabelOptions `json:"labels"`
}

// VisualizationLabelOptions struct
type VisualizationLabelOptions struct {
	Enabled bool `json:"enabled"`
}

// VisualizationColumnOptions struct
type VisualizationColumnOptions struct {
	// Shared
	Visible            bool     `json:"visible"`
	Name               string   `json:"name"`
	Title              string   `json:"title"`
	AlignContent       string   `json:"alignContent"`
	AllowSearch        bool     `json:"allowSearch"`
	Type               string   `json:"type"`
	DisplayAs          string   `json:"displayAs"`

	Order              int      `json:"order"`

	// Text
	AllowHTML          bool     `json:"allowHTML"`
	HighlightLinks     bool     `json:"highlightLinks"`

	// Number
	NumberFormat       string   `json:"numberFormat"`

	// Date/Time
	DateTimeFormat     string   `json:"dateTimeFormat"`

	// Boolean
	BooleanValues      []string `json:"booleanValues"`

	// Link
	LinkUrlTemplate    string   `json:"linkUrlTemplate"`
	LinkTextTemplate   string   `json:"linkTextTemplate"`
	LinkOpenInNewTab   bool     `json:"linkOpenInNewTab"`
	LinkTitleTemplate  string   `json:"linkTitleTemplate"`

	// Image
	ImageUrlTemplate   string   `json:"imageUrlTemplate"`
	ImageWidth         string   `json:"imageWidth"`
	ImageHeight        string   `json:"imageHeight"`
	ImageTitleTemplate string   `json:"imageTitleTemplate"`

	// JSON
}

type VisualizationCreatePayload struct {
	Name        string               `json:"name,omitempty"`
	Type        string               `json:"type,omitempty"`
	QueryId     int                  `json:"query_id,omitempty"`
	Description string               `json:"description,omitempty"`
	Options     VisualizationOptions `json:"options,omitempty"`
}

type VisualizationUpdatePayload struct {
	Name        string               `json:"name,omitempty"`
	Type        string               `json:"type,omitempty"`
	Description string               `json:"description,omitempty"`
	Options     VisualizationOptions `json:"options,omitempty"`
}

// GetVisualization gets a specific visualization
func (c *Client) GetVisualization(queryId, visualizationId int) (*Visualization, error) {
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
func (c *Client) CreateVisualization(visualizationCreatePayload *VisualizationCreatePayload) (*Visualization, error) {
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
	newVisualization := new(Visualization)
	err = json.NewDecoder(response.Body).Decode(newVisualization)
	if err != nil {
		return nil, err
	}

	return newVisualization, nil
}

// UpdateVisualization updates an existing Redash visualization
func (c *Client) UpdateVisualization(id int, visualizationUpdatePayload *VisualizationUpdatePayload) (*Visualization, error) {
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
	newVisualization := new(Visualization)
	json.NewDecoder(response.Body).Decode(newVisualization)
	if err != nil {
		return nil, err
	}

	return newVisualization, nil
}

// DeleteVisualization deletes a visualization
func (c *Client) DeleteVisualization(id int) error {
	path := "/api/visualizations/" + strconv.Itoa(id)

	_, err := c.delete(path, url.Values{})

	return err
}
