package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// const baseURL = "https://app.asana.com/api/1.0"
const workspaceID = "1207408218614750"

type Client struct {
	client           *http.Client
	baseUrl          string
	limitRateSeconds int
}

func NewAstanaClient(baseUrl string) *Client {
	return &Client{
		client:  &http.Client{},
		baseUrl: baseUrl,
	}
}

func (c *Client) GetAllProjects(limit string) ([]Project, error) {
	var data ProjectData
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/projects", c.baseUrl), nil)
	if err != nil {
		return nil, err
	}
	addHeaders(req, limit)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request did not respond with 200, but with %d instead", resp.StatusCode)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}

	if len(data.Projects) == 0 {
		return nil, fmt.Errorf("did not get any projects")
	}

	return data.Projects, err
}

func (c *Client) GetAllUsers(limit string) ([]User, error) {
	var data UsersData
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/users", c.baseUrl), nil)
	if err != nil {
		return nil, err
	}
	addHeaders(req, limit)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request did not respond with 200, but with %d instead", resp.StatusCode)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}

	if len(data.Users) == 0 {
		return nil, fmt.Errorf("did not get any users")
	}

	return data.Users, err
}

func addHeaders(r *http.Request, limit string) {
	// Bearer to be added as .env
	r.Header.Add("authorization", "Bearer 2/1207408404715783/1207408322162260:faebb040c25defdd82b1f7de5dc9a7ef")
	r.Header.Add("content-type", "application/json")
	q := r.URL.Query()
	q.Add("limit", limit)
	q.Add("workspace", workspaceID)
	r.URL.RawQuery = q.Encode()
}
