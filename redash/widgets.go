package redash

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// Widget object structure for Dashboards
type WidgetDashboard struct {
	// Base Data
	ID          int `json:"id"`
	DashboardID int `json:"dashboard_id"`

	//
	Text  string `json:"text"`
	Width int    `json:"width"`

	// References
	Visualization VisualizationDashboard `json:"visualization,omitempty"`

	// Options
	Options WidgetOptions `json:"options"`

	// Timestamps
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WidgetOptions struct {
	IsHidden          bool                              `json:"isHidden"`
	ParameterMappings map[string]WidgetParameterMapping `json:"parameterMappings"`
	Position          WidgetPosition                    `json:"position"`
}

type WidgetParameterMapping struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	MapTo string `json:"mapTo"`
	Value string `json:"value"`
	Title string `json:"title"`
}

type WidgetPosition struct {
	AutoHeight bool `json:"autoHeight"`
	SizeX      int  `json:"sizeX"`
	SizeY      int  `json:"sizeY"`
	MaxSizeY   int  `json:"maxSizeY"`
	MaxSizeX   int  `json:"maxSizeX"`
	MinSizeY   int  `json:"minSizeY"`
	MinSizeX   int  `json:"minSizeX"`
	Col        int  `json:"col"`
	Row        int  `json:"row"`
}

// WidgetCreatePayload defines the schema for creating a Redash widget
type WidgetCreatePayload struct {
	// Base Data
	DashboardID int `json:"dashboard_id"`

	//
	Text  string `json:"text"`
	Width int    `json:"width"`

	// References
	VisualizationID *int `json:"visualization_id"`

	// Options
	Options WidgetOptions `json:"options"`
}

// WidgetUpdatePayload defines the schema for updating a Redash widget
type WidgetUpdatePayload struct {
	//
	Text  string `json:"text"`
	Width int    `json:"width"`

	// References
	VisualizationID *int `json:"visualization_id"`

	// Options
	Options WidgetOptions `json:"options"`
}

// GetWidget returns a specific Widget by its dashboard slug and widget ID
func (c *Client) GetWidget(dashboardSlug string, widgetId int) (*WidgetDashboard, error) {
	dashboard, err := c.GetDashboard(dashboardSlug)
	if err != nil {
		return nil, err
	}

	for _, w := range dashboard.Widgets {
		if w.ID == widgetId {
			return &w, nil
		}
	}

	return nil, fmt.Errorf("widget %d not found in dashboard %s", widgetId, dashboardSlug)
}

// CreateWidget creates a new Redash widget
func (c *Client) CreateWidget(widgetCreatePayload *WidgetCreatePayload) (*WidgetDashboard, error) {
	path := "/api/widgets"

	payload, err := json.Marshal(widgetCreatePayload)
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	response, err := c.post(path, string(payload), queryParams)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	newWidget := new(WidgetDashboard)
	err = json.NewDecoder(response.Body).Decode(newWidget)
	if err != nil {
		return nil, err
	}

	return newWidget, nil
}

// UpdateWidget updates an existing Redash widget
func (c *Client) UpdateWidget(id int, widgetUpdatePayload *WidgetUpdatePayload) (*WidgetDashboard, error) {
	path := "/api/widgets/" + strconv.Itoa(id)

	payload, err := json.Marshal(widgetUpdatePayload)
	if err != nil {
		return nil, err
	}

	queryParams := url.Values{}
	response, err := c.post(path, string(payload), queryParams)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	newWidget := new(WidgetDashboard)
	err = json.NewDecoder(response.Body).Decode(newWidget)
	if err != nil {
		return nil, err
	}

	return newWidget, nil
}

// DeleteWidget deletes a Redash widget
func (c *Client) DeleteWidget(id int) error {
	path := "/api/widgets/" + strconv.Itoa(id)

	_, err := c.delete(path, url.Values{})

	return err
}
