package redash

import (
	"github.com/jarcoal/httpmock"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestGetVisualization(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	body, err := ioutil.ReadFile("testdata/get-query.json")
	if err != nil {
		panic(err.Error())
	}
	httpmock.RegisterResponder("GET", "https://com.acme/api/queries/1",
		httpmock.NewStringResponder(200, string(body)))

	visualization, err := c.GetVisualization(1, 2)
	assert.Nil(err)

	assert.Equal(2, visualization.ID)
	assert.Equal("DAU", visualization.Name)
	assert.Equal("CHART", visualization.Type)
	assert.Equal("", visualization.Description)

	options := &ChartOptions{}
	mapstructure.Decode(visualization.Options, &options)
	assert.NotNil(options)
	assert.Equal("line", options.GlobalSeriesType)
	assert.Equal(true, options.SortX)
	assert.Equal(true, options.Legend.Enabled)

	xAxis := options.XAxis
	assert.NotNil(xAxis)
	assert.Equal("datetime", xAxis.Type)
	assert.Equal(true, xAxis.Labels.Enabled)
}

func TestCreateVisualization(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	body, err := ioutil.ReadFile("testdata/create-visualization.json")
	if err != nil {
		panic(err.Error())
	}
	httpmock.RegisterResponder("POST", "https://com.acme/api/visualizations",
		httpmock.NewStringResponder(200, string(body)))

	visualization, err := c.CreateVisualization(&VisualizationCreatePayload{
		Name:        "Results",
		Type:        "TABLE",
		QueryId:     1,
		Description: "Query results",
		Options:     TableOptions{},
	})
	assert.Nil(err)

	assert.Equal(8, visualization.ID)
	assert.Equal("Results", visualization.Name)
	assert.Equal("TABLE", visualization.Type)
	assert.Equal("Query results", visualization.Description)
}

func TestUpdateVisualization(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	body, err := ioutil.ReadFile("testdata/update-visualization.json")
	if err != nil {
		panic(err.Error())
	}
	httpmock.RegisterResponder("POST", "https://com.acme/api/visualizations/8",
		httpmock.NewStringResponder(200, string(body)))

	visualization, err := c.UpdateVisualization(8, &VisualizationUpdatePayload{
		Type:        "CHART",
		Description: "Pie",
	})
	assert.Nil(err)

	assert.Equal(8, visualization.ID)
	assert.Equal("Results", visualization.Name)
	assert.Equal("CHART", visualization.Type)
	assert.Equal("Pie", visualization.Description)
}

func TestDeleteVisualization(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	httpmock.RegisterResponder("DELETE", "https://com.acme/api/visualizations/9",
		httpmock.NewStringResponder(200, "{}"))

	err := c.DeleteVisualization(9)
	assert.Nil(err)
}
