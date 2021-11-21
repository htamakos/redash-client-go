package redash

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

type Dashboard struct {
	ID                      int           `json:"id"`
	Slug                    string        `json:"slug"`
	Name                    string        `json:"name"`
	UserID                  int           `json:"user_id"`
	User                    User          `json:"user"`
	Layout                  []interface{} `json:"layout"`
	DashboardFiltersEnabled bool          `json:"dashboard_filters_enabled"`
	Widgets                 []Widget      `json:"widgets"`
	IsArchived              bool          `json:"is_archived"`
	IsDraft                 bool          `json:"is_draft"`
	Tags                    []string      `json:"tags"`
	UpdatedAt               time.Time     `json:"updated_at"`
	CreatedAt               time.Time     `json:"created_at"`
	Version                 int           `json:"version"`
	IsFavorite              bool          `json:"is_favorite"`
	CanEdit                 bool          `json:"can_edit"`
}

type DashboardVisualization struct {
	ID          int                       `json:"id"`
	Type        string                    `json:"type"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Options     QueryVisualizationOptions `json:"options"`
	Query       Query                     `json:"query"`
}

type DashboardCreatePayload struct {
	Name string `json:"name"`
}

type DashboardUpdatePayload struct {
	Name string `json:"name"`
}

// GetDashboard gets a specific dashboard
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

func (c *Client) ArchiveDashboard(slug string) error {
	path := "/api/dashboards/" + slug

	_, err := c.delete(path, url.Values{})

	return err
}
