package redash

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
	"time"
)

type QuerySnippet struct {
	Id          int       `json:"id,omitempty"`
	Trigger     string    `json:"trigger,omitempty"`
	Snippet     string    `json:"snippet,omitempty"`
	Description string    `json:"description,omitempty"`
	User        User      `json:"user,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type CreateQuerySnippetPayload struct {
	Trigger     string `json:"trigger,omitempty"`
	Snippet     string `json:"snippet,omitempty"`
	Description string `json:"description,omitempty"`
}

type UpdateQuerySnippetPayload struct {
	Id          int    `json:"id,omitempty"`
	Trigger     string `json:"trigger,omitempty"`
	Snippet     string `json:"snippet,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *Client) GetQuerySnippets() (*[]QuerySnippet, error) {
	path := "/api/query_snippets"
	res, err := c.get(path, url.Values{})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	querySnippets := []QuerySnippet{}
	err = json.Unmarshal(body, &querySnippets)
	if err != nil {
		return nil, err
	}

	return &querySnippets, nil
}

func (c *Client) GetQuerySnippet(id int) (*QuerySnippet, error) {
	path := "/api/query_snippets/" + strconv.Itoa(id)

	res, err := c.get(path, url.Values{})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	querySnippet := QuerySnippet{}
	err = json.Unmarshal(body, &querySnippet)
	if err != nil {
		return nil, err
	}

	return &querySnippet, nil
}

func (c *Client) CreateQuerySnippet(createQuerySnippetPayload CreateQuerySnippetPayload) (*QuerySnippet, error) {
	path := "/api/query_snippets"
	payload, err := json.Marshal(createQuerySnippetPayload)
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

	querySnippet := QuerySnippet{}
	err = json.Unmarshal(body, &querySnippet)
	if err != nil {
		return nil, err
	}

	return &querySnippet, nil
}

func (c *Client) UpdateQuerySnippet(id int, updateQuerySnippetPalyload UpdateQuerySnippetPayload) (*QuerySnippet, error) {
	path := "/api/query_snippets/" + strconv.Itoa(id)
	payload, err := json.Marshal(updateQuerySnippetPalyload)
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

	querySnippet := QuerySnippet{}
	err = json.Unmarshal(body, &querySnippet)
	if err != nil {
		return nil, err
	}

	return &querySnippet, nil
}

func (c *Client) DeleteQuerySnippet(id int) error {
	path := "/api/query_snippets/" + strconv.Itoa(id)
	_, err := c.delete(path, url.Values{})
	if err != nil {
		return err
	}

	return nil
}
