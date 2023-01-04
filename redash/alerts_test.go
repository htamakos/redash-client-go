package redash

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAlerts(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})
	httpmock.RegisterResponder("GET", "https://com.acme/api/alerts",
		httpmock.NewStringResponder(200, `[{
  "id": 2,
  "name": "New Query: fv > 1",
  "options": { "op": ">", "value": 1, "muted": false, "column": "col1" },
  "state": "unknown",
  "last_triggered_at": null,
  "updated_at": "2023-01-03T02:57:24.435Z",
  "created_at": "2023-01-03T02:57:24.435Z",
  "rearm": null,
  "query": {
    "id": 5,
    "latest_query_data_id": 18,
    "name": "New Query",
    "description": null,
    "query": "select *\nfrom dual",
    "query_hash": "xxxxxxxxxxxxxxxxxxxxxxx",
    "schedule": null,
    "api_key": "xxxxxxxxxxxxxxxxxxxxxxx",
    "is_archived": false,
    "is_draft": false,
    "updated_at": "2023-01-03T02:57:24.435Z",
    "created_at": "2023-01-03T02:49:53.949Z",
    "data_source_id": 16,
    "options": { "apply_auto_limit": true, "parameters": [] },
    "version": 1,
    "tags": [],
    "is_safe": true,
    "user": {
      "id": 1,
      "name": "admin",
      "email": "xxxxxx@example.com",
      "profile_image_url": "https://example.com/image",
      "groups": [1, 2],
      "updated_at": "2023-01-03T03:04:44.481Z",
      "created_at": "2022-12-27T05:39:09.998Z",
      "disabled_at": null,
      "is_disabled": false,
      "active_at": "2023-01-03T03:04:08Z",
      "is_invitation_pending": false,
      "is_email_verified": true,
      "auth_type": "password"
    },
    "last_modified_by": {
      "id": 1,
      "name": "admin",
      "email": "xxxxx@examplecom",
      "profile_image_url": "https://example.com",
      "groups": [1, 2],
      "updated_at": "2023-01-03T03:04:44.481Z",
      "created_at": "2022-12-27T05:39:09.998Z",
      "disabled_at": null,
      "is_disabled": false,
      "active_at": "2023-01-03T03:04:08Z",
      "is_invitation_pending": false,
      "is_email_verified": true,
      "auth_type": "password"
    }
  },
  "user": {
    "id": 1,
    "name": "admin",
    "email": "xxxxxx@example.com",
    "profile_image_url": "https://example.com/images",
    "groups": [1, 2],
    "updated_at": "2023-01-03T03:04:44.481Z",
    "created_at": "2022-12-27T05:39:09.998Z",
    "disabled_at": null,
    "is_disabled": false,
    "active_at": "2023-01-03T03:04:08Z",
    "is_invitation_pending": false,
    "is_email_verified": true,
    "auth_type": "password"
  }
}]`))

	palerts, _ := c.GetAlerts()
	assert.NotNil(palerts)

	alerts := *palerts
	assert.Equal(1, len(alerts))

	alert := alerts[0]
	assert.Nil(alert.LastTriggeredAt)
	assert.Equal(AlertOption{Op: ">", Value: 1.0, Muted: false, Column: "col1"}, alert.Options)
}

func TestGetAlert(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})
	httpmock.RegisterResponder("GET", "https://com.acme/api/alerts/2",
		httpmock.NewStringResponder(200, `{
  "id": 2,
  "name": "New Query: fv > 1",
  "options": { "op": ">", "value": 1, "muted": false, "column": "col1" },
  "state": "unknown",
  "last_triggered_at": null,
  "updated_at": "2023-01-03T02:57:24.435Z",
  "created_at": "2023-01-03T02:57:24.435Z",
  "rearm": null,
  "query": {
    "id": 5,
    "latest_query_data_id": 18,
    "name": "New Query",
    "description": null,
    "query": "select *\nfrom dual",
    "query_hash": "xxxxxxxxxxxxxxxxxxxxxxx",
    "schedule": null,
    "api_key": "xxxxxxxxxxxxxxxxxxxxxxx",
    "is_archived": false,
    "is_draft": false,
    "updated_at": "2023-01-03T02:57:24.435Z",
    "created_at": "2023-01-03T02:49:53.949Z",
    "data_source_id": 16,
    "options": { "apply_auto_limit": true, "parameters": [] },
    "version": 1,
    "tags": [],
    "is_safe": true,
    "user": {
      "id": 1,
      "name": "admin",
      "email": "xxxxxx@example.com",
      "profile_image_url": "https://example.com/image",
      "groups": [1, 2],
      "updated_at": "2023-01-03T03:04:44.481Z",
      "created_at": "2022-12-27T05:39:09.998Z",
      "disabled_at": null,
      "is_disabled": false,
      "active_at": "2023-01-03T03:04:08Z",
      "is_invitation_pending": false,
      "is_email_verified": true,
      "auth_type": "password"
    },
    "last_modified_by": {
      "id": 1,
      "name": "admin",
      "email": "xxxxx@examplecom",
      "profile_image_url": "https://example.com",
      "groups": [1, 2],
      "updated_at": "2023-01-03T03:04:44.481Z",
      "created_at": "2022-12-27T05:39:09.998Z",
      "disabled_at": null,
      "is_disabled": false,
      "active_at": "2023-01-03T03:04:08Z",
      "is_invitation_pending": false,
      "is_email_verified": true,
      "auth_type": "password"
    }
  },
  "user": {
    "id": 1,
    "name": "admin",
    "email": "xxxxxx@example.com",
    "profile_image_url": "https://example.com/images",
    "groups": [1, 2],
    "updated_at": "2023-01-03T03:04:44.481Z",
    "created_at": "2022-12-27T05:39:09.998Z",
    "disabled_at": null,
    "is_disabled": false,
    "active_at": "2023-01-03T03:04:08Z",
    "is_invitation_pending": false,
    "is_email_verified": true,
    "auth_type": "password"
  }
}`))

	palert, _ := c.GetAlert(2)
	assert.NotNil(palert)

	alert := *palert
	assert.Nil(alert.LastTriggeredAt)
	assert.Equal(AlertOption{Op: ">", Value: 1.0, Muted: false, Column: "col1"}, alert.Options)
}

func TestCreateAlert(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})
	httpmock.RegisterResponder("POST", "https://com.acme/api/alerts",
		httpmock.NewStringResponder(200, `{
  "id": 2,
  "name": "NewQuery",
  "options": { "op": ">", "value": 1, "muted": false, "column": "col1" },
  "state": "unknown",
  "last_triggered_at": null,
  "updated_at": "2023-01-03T02:57:24.435Z",
  "created_at": "2023-01-03T02:57:24.435Z",
  "rearm": null,
  "query": {
    "id": 5,
    "latest_query_data_id": 18,
    "name": "New Query",
    "description": null,
    "query": "select *\nfrom dual",
    "query_hash": "xxxxxxxxxxxxxxxxxxxxxxx",
    "schedule": null,
    "api_key": "xxxxxxxxxxxxxxxxxxxxxxx",
    "is_archived": false,
    "is_draft": false,
    "updated_at": "2023-01-03T02:57:24.435Z",
    "created_at": "2023-01-03T02:49:53.949Z",
    "data_source_id": 16,
    "options": { "apply_auto_limit": true, "parameters": [] },
    "version": 1,
    "tags": [],
    "is_safe": true,
    "user": {
      "id": 1,
      "name": "admin",
      "email": "xxxxxx@example.com",
      "profile_image_url": "https://example.com/image",
      "groups": [1, 2],
      "updated_at": "2023-01-03T03:04:44.481Z",
      "created_at": "2022-12-27T05:39:09.998Z",
      "disabled_at": null,
      "is_disabled": false,
      "active_at": "2023-01-03T03:04:08Z",
      "is_invitation_pending": false,
      "is_email_verified": true,
      "auth_type": "password"
    },
    "last_modified_by": {
      "id": 1,
      "name": "admin",
      "email": "xxxxx@examplecom",
      "profile_image_url": "https://example.com",
      "groups": [1, 2],
      "updated_at": "2023-01-03T03:04:44.481Z",
      "created_at": "2022-12-27T05:39:09.998Z",
      "disabled_at": null,
      "is_disabled": false,
      "active_at": "2023-01-03T03:04:08Z",
      "is_invitation_pending": false,
      "is_email_verified": true,
      "auth_type": "password"
    }
  },
  "user": {
    "id": 1,
    "name": "admin",
    "email": "xxxxxx@example.com",
    "profile_image_url": "https://example.com/images",
    "groups": [1, 2],
    "updated_at": "2023-01-03T03:04:44.481Z",
    "created_at": "2022-12-27T05:39:09.998Z",
    "disabled_at": null,
    "is_disabled": false,
    "active_at": "2023-01-03T03:04:08Z",
    "is_invitation_pending": false,
    "is_email_verified": true,
    "auth_type": "password"
  }
}`))

	palert, _ := c.CreateAlert(CreateAlertPayload{
		Name:    "NewQuery",
		QueryId: 5,
		Options: AlertOption{
			Op:     ">",
			Value:  1.0,
			Muted:  false,
			Column: "col1",
		},
	})
	assert.NotNil(palert)

	alert := *palert
	assert.Nil(alert.LastTriggeredAt)
	assert.Equal(AlertOption{Op: ">", Value: 1.0, Muted: false, Column: "col1"}, alert.Options)

}
