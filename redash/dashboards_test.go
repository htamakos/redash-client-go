package redash

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetDashboard(t *testing.T) {
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

	dashboard, err := c.GetDashboard("service-slos")
	assert.Nil(err)

	assert.Equal(1, dashboard.ID)
	assert.Equal("service-slos", dashboard.Slug)
	assert.Equal("Service SLOs", dashboard.Name)
	assert.Equal(5388, dashboard.UserID)
	assert.NotNil(dashboard.Layout)
	assert.Equal(false, dashboard.DashboardFiltersEnabled)
	assert.Equal(false, dashboard.IsArchived)
	assert.Equal(false, dashboard.IsDraft)
	assert.Equal([]string{"reliability"}, dashboard.Tags)
	assert.Equal(14, dashboard.Version)
	assert.Equal(true, dashboard.IsFavorite)
	assert.Equal(true, dashboard.CanEdit)

	user := dashboard.User
	assert.NotNil(user)
	assert.Equal(5388, user.ID)
	assert.Equal("Developer", user.Name)
	assert.Equal("developer@example.com", user.Email)
	assert.Equal("https://example.com", user.ProfileImageURL)

	assert.Equal(4, len(dashboard.Widgets))
	widget := dashboard.Widgets[0]
	assert.Equal(65057, widget.ID)
	assert.Equal(1, widget.Width)
	assert.Equal(1, widget.DashboardID)
}

func TestCreateDashboard(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	body, err := ioutil.ReadFile("testdata/create-dashboard.json")
	if err != nil {
		panic(err.Error())
	}
	httpmock.RegisterResponder("POST", "https://com.acme/api/dashboards",
		httpmock.NewStringResponder(200, string(body)))

	dashboard, err := c.CreateDashboard(&DashboardCreatePayload{
		Name: "New Dashboard",
	})
	assert.Nil(err)

	assert.Equal(5, dashboard.ID)
	assert.Equal("New Dashboard", dashboard.Name)
	assert.Equal("new-dashboard", dashboard.Slug)
	assert.Nil(dashboard.Widgets)
}

func TestUpdateDashboard(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	body, err := ioutil.ReadFile("testdata/update-dashboard.json")
	if err != nil {
		panic(err.Error())
	}
	httpmock.RegisterResponder("POST", "https://com.acme/api/dashboards/5",
		httpmock.NewStringResponder(200, string(body)))

	dashboard, err := c.UpdateDashboard(5, &DashboardUpdatePayload{
		Name: "New Name",
	})
	assert.Nil(err)

	assert.Equal(5, dashboard.ID)
	assert.Equal("New Name", dashboard.Name)
	assert.Equal("new-dashboard", dashboard.Slug)
	assert.Nil(dashboard.Widgets)
}

func TestArchiveDashboard(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	c, _ := NewClient(&Config{RedashURI: "https://com.acme/", APIKey: "ApIkEyApIkEyApIkEyApIkEyApIkEy"})

	httpmock.RegisterResponder("DELETE", "https://com.acme/api/dashboards/my-dashboard",
		httpmock.NewStringResponder(200, `{}`))

	err := c.ArchiveDashboard("my-dashboard")
	assert.Nil(err)
}
