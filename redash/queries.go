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
		ID                int           `json:"id,omitempty"`
		IsArchived        bool          `json:"is_archived,omitempty"`
		CreatedAt         time.Time     `json:"created_at,omitempty"`
		RetrievedAt       time.Time     `json:"retrieved_at,omitempty"`
		UpdatedAt         time.Time     `json:"updated_at,omitempty"`
		Name              string        `json:"name,omitempty"`
		Description       string        `json:"description,omitempty"`
		Query             string        `json:"query,omitempty"`
		QueryHash         string        `json:"query_hash,omitempty"`
		Version           int           `json:"version,omitempty"`
		LastModifiedByID  int           `json:"last_modified_by_id,omitempty"`
		Tags              []string      `json:"tags,omitempty"`
		APIKey            string        `json:"api_key,omitempty"`
		DataSourceID      int           `json:"data_source_id,omitempty"`
		LatestQueryDataID int           `json:"latest_query_data_id,omitempty"`
		Schedule          QuerySchedule `json:"schedule,omitempty"`
		User              User          `json:"user,omitempty"`
		IsFavorite        bool          `json:"is_favorite,omitempty"`
		IsDraft           bool          `json:"is_draft,omitempty"`
		IsSafe            bool          `json:"is_safe,omitempty"`
		Runtime           float32       `json:"runtime,omitempty"`
		Options           QueryOptions  `json:"options,omitempty"`
	}
}

// Query models the response from Redash's /api/queries endpoint
type Query struct {
	ID                int                  `json:"id,omitempty"`
	Name              string               `json:"name,omitempty"`
	Description       string               `json:"description,omitempty"`
	Query             string               `json:"query,omitempty"`
	QueryHash         string               `json:"query_hash,omitempty"`
	Version           int                  `json:"version,omitempty"`
	Schedule          QuerySchedule        `json:"schedule,omitempty"`
	APIKey            string               `json:"api_key,omitempty"`
	IsArchived        bool                 `json:"is_archived,omitempty"`
	IsDraft           bool                 `json:"is_draft,omitempty"`
	UpdatedAt         time.Time            `json:"updated_at,omitempty"`
	CreatedAt         time.Time            `json:"created_at,omitempty"`
	DataSourceID      int                  `json:"data_source_id,omitempty"`
	LatestQueryDataID int                  `json:"latest_query_data_id,omitempty"`
	Tags              []string             `json:"tags,omitempty"`
	IsSafe            bool                 `json:"is_safe,omitempty"`
	User              User                 `json:"user,omitempty"`
	LastModifiedBy    User                 `json:"last_modified_by,omitempty"`
	IsFavorite        bool                 `json:"is_favorite,omitempty"`
	CanEdit           bool                 `json:"can_edit,omitempty"`
	Options           QueryOptions         `json:"options,omitempty"`
	Visualizations    []QueryVisualization `json:"visualizations,omitempty"`
}

// QuerySchedule struct
type QuerySchedule struct {
	Interval  int         `json:"interval,omitempty"`
	Time      string      `json:"time,omitempty"`
	DayOfWeek string      `json:"day_of_week,omitempty"`
	Until     interface{} `json:"until,omitempty"`
}

// QueryOptions struct
type QueryOptions struct {
	Parameters []QueryOptionsParameter `json:"parameters,omitempty"`
}

// QueryOptionsParameter struct
type QueryOptionsParameter struct {
	Title       string        `json:"title,omitempty"`
	Name        string        `json:"name,omitempty"`
	Type        string        `json:"type,omitempty"`
	EnumOptions string        `json:"enum_options,omitempty"`
	Locals      []interface{} `json:"locals,omitempty"`
	Value       string        `json:"value,omitempty"`
}

// QueryVisualization struct
type QueryVisualization struct {
	ID          int                       `json:"id,omitempty"`
	Type        string                    `json:"type,omitempty"`
	Name        string                    `json:"name,omitempty"`
	Description string                    `json:"description,omitempty"`
	Options     QueryVisualizationOptions `json:"options,omitempty"`
}

// QueryVisualizationOptions struct
type QueryVisualizationOptions struct {
	GlobalSeriesType string                          `json:"global_series_type,omitempty"`
	SortX            bool                            `json:"sort_x,omitempty"`
	Legend           QueryVisualizationLegendOptions `json:"legend,omitempty"`
	YAxis            []QueryAxisOptions              `json:"y_axis,omitempty"`
	XAxis            QueryAxisOptions                `json:"x_axis,omitempty"`
}

// QueryVisualizationLegendOptions struct
type QueryVisualizationLegendOptions struct {
	Enabled   bool   `json:"enabled"`
	Placement string `json:"placement"`
}

// QueryAxisOptions struct
type QueryAxisOptions struct {
	Type     string                 `json:"type,omitempty"`
	Opposite bool                   `json:"opposite,omitempty"`
	Labels   QueryAxisLabelsOptions `json:"labels,omitempty"`
}

// QueryAxisLabelsOptions struct
type QueryAxisLabelsOptions struct {
	Enabled bool `json:"enabled"`
}

// QueryCreatePayload defines the schema for creating a new Redash query
type QueryCreatePayload struct {
	Name         string `json:"name"`
	Query        string `json:"query"`
	DataSourceID int    `json:"data_source_id"`
	Description  string `json:"description,omitempty"`
}

// QueryUpdatePayload defines the schema for updating a Redash query
type QueryUpdatePayload struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Query        string `json:"query,omitempty"`
	DataSourceID int    `json:"data_source_id,omitempty"`
	IsDraft      bool   `json:"is_draft,omitempty"`
	Options      bool   `json:"options,omitempty"`
	Version      bool   `json:"version,omitempty"`
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
