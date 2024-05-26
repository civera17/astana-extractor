package client

type Project struct {
	Gid          string `json:"gid"`
	ResourceType string `json:"resource_type"`
	Name         string `json:"name"`
	Archieved     string
	Color         string `json:"created_at"`
	CreatedAt     string
	CurrentStatus ProjectStatus `json:"current_status"`
}

type ProjectStatus struct {
	Gid         string
	ResourcType string
	Color       string
	HtmlText    string
	Text        string
	Title       string
}

type User struct {
	Gid          string
	ResourceType string `json:"resource_type"`
	Email        string
	Name         string
	Workspaces   string
}

type Workspace struct {
	Gid            string
	ResourceType   string `json:"resource_type"`
	Name           string
	EmailDomains   string `json:"email_domains"`
	IsOrganisation bool   `json:"is_organisation"`
}

type IAstanaClient interface {
	GetAllProjects() ([]Project, error)
	GetAllUsers() ([]User, error)
}

type ProjectData struct {
	Projects []Project `json:"data"`
}

type UsersData struct {
	Users []User `json:"data"`
}
