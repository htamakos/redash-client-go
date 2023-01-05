package redash

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
	"time"
)

type Alert struct {
	ID              int         `json:"id,omitempty"`
	Name            string      `json:"name,omitempty"`
	Options         AlertOption `json:"options,omitempty"`
	State           string      `json:"state,omitempty"`
	LastTriggeredAt *time.Time  `json:"last_triggered_at,omitempty"`
	UpdatedAt       time.Time   `json:"updated_at,omitempty"`
	CreatedAt       time.Time   `json:"created_at,omitempty"`
	Rearm           *int        `json:"rearm,omitempty"`
	Query           Query       `json:"query,omitempty"`
	User            User        `json:"user,omitempty"`
}

type AlertOption struct {
	Op            string      `json:"op,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	Muted         bool        `json:"muted,omitempty"`
	Column        string      `json:"column,omitempty"`
	CustomBody    *string     `json:"custom_body,omitempty"`
	CustomSubject *string     `json:"custom_subject,omitempty"`
}

type CreateAlertPayload struct {
	Name    string      `json:"name,omitempty"`
	QueryId int         `json:"query_id,omitempty"`
	Options AlertOption `json:"options,omitempty"`
	Rearm   *int        `json:"rearm,omitempty"`
}

type UpdateAlertPayload struct {
	Name    string      `json:"name,omitempty"`
	QueryId int         `json:"query_id,omitempty"`
	Options AlertOption `json:"options,omitempty"`
	Rearm   *int        `json:"rearm,omitempty"`
}

type CreateAlertSubscriptionPayload struct {
	AlertId       int `json:"alert_id,omitempty"`
	DestinationId int `json:"destination_id,omitempty"`
}

type AlertSubscription struct {
	Id          int         `json:"id,omitempty"`
	AlertId     int         `json:"alert_id,omitempty"`
	User        User        `json:"user,omitempty"`
	Destination Destination `json:"destination,omitempty"`
}

func (c *Client) GetAlerts() (*[]Alert, error) {
	path := "/api/alerts"
	query := url.Values{}
	response, err := c.get(path, query)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	alerts := []Alert{}
	err = json.Unmarshal(body, &alerts)
	if err != nil {
		return nil, err
	}

	return &alerts, nil
}

func (c *Client) GetAlert(id int) (*Alert, error) {
	path := "/api/alerts/" + strconv.Itoa(id)
	query := url.Values{}
	response, err := c.get(path, query)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	alert := Alert{}
	err = json.Unmarshal(body, &alert)
	if err != nil {
		return nil, err
	}

	return &alert, nil
}

func (c *Client) CreateAlert(createAlert CreateAlertPayload) (*Alert, error) {
	path := "/api/alerts"

	payload, err := json.Marshal(createAlert)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	res, err := c.post(path, string(payload), query)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	alert := Alert{}
	err = json.Unmarshal(body, &alert)
	if err != nil {
		return nil, err
	}

	return &alert, nil
}

func (c *Client) UpdateAlert(id int, updateAlertPayload *UpdateAlertPayload) (*Alert, error) {
	path := "/api/alerts/" + strconv.Itoa(id)

	payload, err := json.Marshal(updateAlertPayload)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	res, err := c.post(path, string(payload), query)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	alert := Alert{}

	err = json.Unmarshal(body, &alert)
	if err != nil {
		return nil, err
	}
	return &alert, nil
}

func (c *Client) DeleteAlert(id int) error {
	path := "/api/alerts/" + strconv.Itoa(id)

	_, err := c.delete(path, url.Values{})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetAlertSubscriptions(id int) (*[]AlertSubscription, error) {
	path := "/api/alerts/" + strconv.Itoa(id) + "/subscriptions"

	res, err := c.get(path, url.Values{})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var subscriptions []AlertSubscription
	err = json.Unmarshal(body, &subscriptions)
	if err != nil {
		return nil, err
	}

	return &subscriptions, nil
}

func (c *Client) CreateAlertSubscription(createAlertSubsciptionPayload CreateAlertSubscriptionPayload) (*AlertSubscription, error) {
	path := "/api/alerts/" + strconv.Itoa(createAlertSubsciptionPayload.AlertId) + "/subscriptions"

	payload, err := json.Marshal(createAlertSubsciptionPayload)
	if err != nil {
		return nil, err
	}

	res, err := c.post(path, string(payload), url.Values{})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	alertSubscription := AlertSubscription{}
	err = json.Unmarshal(body, &alertSubscription)
	if err != nil {
		return nil, err
	}

	return &alertSubscription, nil
}

func (c *Client) DeleteAlertSubscription(alertId int, subscriptionId int) error {
	path := "/api/alerts/" + strconv.Itoa(alertId) + "/subscriptions" + strconv.Itoa(subscriptionId)

	_, err := c.delete(path, url.Values{})
	if err != nil {
		return err
	}
	return nil
}
