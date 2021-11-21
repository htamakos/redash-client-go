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
	ID          int                  `json:"id,omitempty"`
	Type        string               `json:"type,omitempty"`
	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	Options     VisualizationOptions `json:"options,omitempty"`
	UpdatedAt   time.Time            `json:"updated_at,omitempty"`
	CreatedAt   time.Time            `json:"created_at,omitempty"`
}

// VisualizationOptions struct
type VisualizationOptions struct {
	XAxis            VisualizationAxisOptions   `json:"xAxis,omitempty"`
	YAxis            []VisualizationAxisOptions `json:"yAxis,omitempty"`
	Series           map[string]interface{}     `json:"series,omitempty"`
	GlobalSeriesType string                     `json:"globalSeriesType,omitempty"`
	SortX            bool                       `json:"sortX,omitempty"`
	SeriesOptions    map[string]SeriesOptions   `json:"seriesOptions,omitempty"`
	ColumnMapping    map[string]string          `json:"columnMapping,omitempty"`
	Legend           VisualizationLegendOptions `json:"legend,omitempty"`
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
	Placement string `json:"placement,omitempty"`
}

// VisualizationAxisOptions struct
type VisualizationAxisOptions struct {
	Type     string                    `json:"type"`
	Opposite bool                      `json:"opposite,omitempty"`
	Labels   VisualizationLabelOptions `json:"labels,omitempty"`
}

// VisualizationLabelOptions struct
type VisualizationLabelOptions struct {
	Enabled bool `json:"enabled"`
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
