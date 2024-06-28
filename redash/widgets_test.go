package redash

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestGetWidget(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	body, err := ioutil.ReadFile("testdata/get-dashboard.json")
	if err != nil {
		panic(err.Error())
	}
	httpmock.RegisterResponder("GET", "https://com.acme/api/dashboards/service-slos",
		httpmock.NewStringResponder(200, string(body)))

	widget, err := c.GetWidget("service-slos", 64399)
	assert.Nil(err)

	assert.Equal(64399, widget.ID)
	assert.Equal(1, widget.DashboardID)
	assert.Equal(1, widget.Width)
	assert.Equal("", widget.Text)

	assert.NotNil(widget.Options)
	assert.Equal(234610, widget.Visualization.ID)
}

func TestCreateWidget(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	body, err := ioutil.ReadFile("testdata/create-widget.json")
	if err != nil {
		panic(err.Error())
	}
	httpmock.RegisterResponder("POST", "https://com.acme/api/widgets",
		httpmock.NewStringResponder(200, string(body)))

	widget, err := c.CreateWidget(&WidgetCreatePayload{
		DashboardID:     5,
		Text:            "Widget text",
		VisualizationID: 1,
		Width:           1,
		Options: WidgetOptions{
			IsHidden: false,
			Position: WidgetPosition{
				AutoHeight: false,
				SizeX:      3,
				SizeY:      5,
				MaxSizeY:   12,
				MaxSizeX:   12,
				MinSizeY:   0,
				MinSizeX:   0,
				Col:        0,
				Row:        10,
			},
			ParameterMappings: nil,
		},
	})
	assert.Nil(err)

	assert.Equal(5, widget.DashboardID)
	assert.Equal("Widget text", widget.Text)
	assert.Equal(1, widget.Visualization.ID)
	assert.Equal(1, widget.Width)

	options := widget.Options
	assert.NotNil(options)
	assert.Equal(false, options.IsHidden)
	assert.Nil(options.ParameterMappings)

	position := options.Position
	assert.NotNil(position)
	assert.Equal(false, position.AutoHeight)
	assert.Equal(3, position.SizeX)
	assert.Equal(5, position.SizeY)
	assert.Equal(12, position.MaxSizeX)
	assert.Equal(12, position.MaxSizeY)
	assert.Equal(12, position.MinSizeX)
	assert.Equal(12, position.MinSizeY)
	assert.Equal(0, position.Col)
	assert.Equal(10, position.Row)
}

func TestUpdateWidget(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	body, err := ioutil.ReadFile("testdata/update-widget.json")
	if err != nil {
		panic(err.Error())
	}
	httpmock.RegisterResponder("POST", "https://com.acme/api/widgets/112",
		httpmock.NewStringResponder(200, string(body)))

	widget, err := c.UpdateWidget(112, &WidgetUpdatePayload{
		Text:  "text",
		Width: 1,
		Options: WidgetOptions{
			IsHidden: false,
			Position: WidgetPosition{
				AutoHeight: false,
				SizeX:      4,
				SizeY:      5,
				Col:        0,
				Row:        2,
			},
			ParameterMappings: nil,
		},
	})
	assert.Nil(err)

	assert.Equal("text", widget.Text)
	assert.Equal(1, widget.Width)

	options := widget.Options
	assert.NotNil(options)
	assert.Equal(false, options.IsHidden)
	assert.Nil(options.ParameterMappings)

	position := options.Position
	assert.NotNil(position)
	assert.Equal(false, position.AutoHeight)
	assert.Equal(4, position.SizeX)
	assert.Equal(5, position.SizeY)
	assert.Equal(0, position.Col)
	assert.Equal(2, position.Row)
}

func TestDeleteWidget(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	httpmock.RegisterResponder("DELETE", "https://com.acme/api/widgets/112",
		httpmock.NewStringResponder(200, "{}"))

	err := c.DeleteWidget(112)
	assert.Nil(err)
}
