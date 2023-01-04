package redash

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Destination struct {
	Name    string                 `json:"name,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
	Type    string                 `json:"type,omitempty"`
	Icon    string                 `json:"icon,omitempty"`
}

type DestinationType struct {
	Name                string                             `json:"name,omitempty"`
	Type                string                             `json:"type,omitempty"`
	Icon                string                             `json:"icon,omitempty"`
	ConfigurationSchema DestinationTypeConfigurationSchema `json:"configuration_schema,omitempty"`
}

type DestinationTypeConfigurationSchema struct {
	Secret     interface{}                             `json:"secret,omitempty"`
	Required   []string                                `json:"required,omitempty"`
	Type       string                                  `json:"type,omitempty"`
	Order      []string                                `json:"order,omitempty"`
	Properties map[string]DestinationTypePropertyField `json:"properties,omitempty"`
}

type DestinationTypePropertyField struct {
	Type    string
	Title   string
	Default interface{}
}

type CreateDestinationPayload struct {
	Name    string                 `json:"name,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
	Type    string                 `json:"type,omitempty"`
}

func (c *Client) GetDestinations() (*[]Destination, error) {
	path := "/api/destinations"
	res, err := c.get(path, url.Values{})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	destinations := []Destination{}
	err = json.Unmarshal(body, &destinations)

	if err != nil {
		return nil, err
	}

	return &destinations, nil
}

func (c *Client) GetDestination(id int) (*Destination, error) {
	path := "/api/destinations/" + strconv.Itoa(id)
	res, err := c.get(path, url.Values{})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	destination := Destination{}
	err = json.Unmarshal(body, &destination)
	if err != nil {
		return nil, err
	}

	return &destination, nil
}

func (c *Client) SanitizeDestinationOptions(destination *Destination) (*Destination, error) {
	destinationTypes, err := c.GetDestinationTypes()
	if err != nil {
		return nil, err
	}

	for _, dst := range destinationTypes {
		if dst.Type == destination.Type {
			for _, required := range dst.ConfigurationSchema.Required {
				_, exists := destination.Options[required]
				if !exists {
					return nil, fmt.Errorf("Required Field missiong: %s", required)
				}
			}

			for propName, propVal := range destination.Options {
				_, exists := dst.ConfigurationSchema.Properties[propName]
				if !exists {
					if c.IsStrict() {
						return nil, fmt.Errorf("Invalid field (%s) for type: %s", propName, destination.Type)
					}

					log.Warn(fmt.Sprintf("[WARN] Ignoring invalid field (%s) for type: %s", propName, destination.Type))
					delete((*destination).Options, propName)
					continue
				}

				switch propVal.(type) {
				case int:
					if dst.ConfigurationSchema.Properties[propName].Type != "number" {
						return nil, fmt.Errorf("Invalid value type for %s", propName)
					}
				case string:
					if dst.ConfigurationSchema.Properties[propName].Type != "string" {

						return nil, fmt.Errorf("Invalid value type for %s", propName)
					}

				case bool:

					if dst.ConfigurationSchema.Properties[propName].Type != "boolean" {
						return nil, fmt.Errorf("Invalid value type for %s", propName)
					}
				default:
					return nil, fmt.Errorf("Invalid value type for %s", propName)
				}
			}
		}
	}

	return destination, nil
}

func (c *Client) CreateDestination(destinationPayload *Destination) (*Destination, error) {
	path := "/api/destinations"

	destinationPayload, err := c.SanitizeDestinationOptions(destinationPayload)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(destinationPayload)
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

	destination := Destination{}
	err = json.Unmarshal(body, &destination)
	if err != nil {
		return nil, err
	}

	return &destination, nil
}

func (c *Client) UpdateDestination(id int, destinationPayload *Destination) (*Destination, error) {
	path := "/api/destinations/" + strconv.Itoa(id)

	destinationPayload, err := c.SanitizeDestinationOptions(destinationPayload)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(destinationPayload)
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

	destination := Destination{}
	err = json.Unmarshal(body, &destination)
	if err != nil {
		return nil, err
	}

	return &destination, nil
}

func (c *Client) DeleteDestination(id int) error {
	path := "/api/destination/" + strconv.Itoa(id)

	_, err := c.delete(path, url.Values{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetDestinationTypes() ([]DestinationType, error) {
	path := "/api/destinations/types"
	res, err := c.get(path, url.Values{})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	destinationTypes := []DestinationType{}
	err = json.Unmarshal(body, &destinationTypes)
	if err != nil {
		return nil, err
	}

	return destinationTypes, nil
}
