package redash

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Widget struct {
	ID            int                    `json:"id"`
	Width         int                    `json:"width"`
	Options       WidgetOptions          `json:"options"`
	DashboardID   int                    `json:"dashboard_id"`
	Text          string                 `json:"text"`
	UpdatedAt     time.Time              `json:"updated_at"`
	CreatedAt     time.Time              `json:"created_at"`
	Visualization DashboardVisualization `json:"visualization"`
}

type WidgetOptions struct {
	IsHidden          bool                              `json:"is_hidden"`
	Position          WidgetPosition                    `json:"position"`
	ParameterMappings map[string]WidgetParameterMapping `json:"parameterMappings"`
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

type WidgetParameterMapping struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	MapTo string `json:"mapTo"`
	Value string `json:"value"`
	Title string `json:"title"`
}

type WidgetCreatePayload struct {
	DashboardID     int           `json:"dashboard_id"`
	Text            string        `json:"text"`
	VisualizationID int           `json:"visualization_id"`
	Width           int           `json:"width"`
	WidgetOptions   WidgetOptions `json:"options"`
}

type WidgetUpdatePayload struct {
	Text          string        `json:"text"`
	Width         int           `json:"width"`
	WidgetOptions WidgetOptions `json:"options"`
}

// GetWidget returns a specific Widget
func (c *Client) GetWidget(dashboardSlug string, widgetId int) (*Widget, error) {
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

func (c *Client) CreateWidget(widgetCreatePayload *WidgetCreatePayload) (*Widget, error) {
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
	newWidget := new(Widget)
	err = json.NewDecoder(response.Body).Decode(newWidget)
	if err != nil {
		return nil, err
	}

	return newWidget, nil
}

func (c *Client) UpdateWidget(id int, widgetUpdatePayload *WidgetUpdatePayload) (*Widget, error) {
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
	newWidget := new(Widget)
	err = json.NewDecoder(response.Body).Decode(newWidget)
	if err != nil {
		return nil, err
	}

	return newWidget, nil
}

func (c *Client) DeleteWidget(id int) error {
	path := "/api/widgets/" + strconv.Itoa(id)

	_, err := c.delete(path, url.Values{})

	return err
}
