package gist

import (
	"github.com/gme-sh/paste-wrapper"
	"strings"
	"time"
)

type GistService struct {
	CreateApiUrl string
}

const (
	DefaultGistCreate = "https://api.github.com/gists"
)

func New() *GistService {
	return &GistService{
		CreateApiUrl: DefaultGistCreate,
	}
}

func WrapPayload(request *paste.UploadRequest) *Create {
	return &Create{
		Authorization: request.Authorization,
		Description:   request.Description,
		Public:        strings.ToLower(request.Visibility) == "public",
		Files: map[string]*GistFile{
			"file": {Content: string(request.Content)},
		},
	}
}

type GistPaste struct {
	URL         string                   `json:"url"`
	ForksURL    string                   `json:"forks_url"`
	CommitsURL  string                   `json:"commits_url"`
	ID          string                   `json:"id"`
	NodeID      string                   `json:"node_id"`
	GitPullURL  string                   `json:"git_pull_url"`
	GitPushURL  string                   `json:"git_push_url"`
	HTMLURL     string                   `json:"html_url"`
	Files       map[string]*responseFile `json:"files"`
	Public      bool                     `json:"public"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
	Description string                   `json:"description"`
	Comments    int                      `json:"comments"`
	CommentsURL string                   `json:"comments_url"`
	Owner       *responseOwner           `json:"owner"`
	Truncated   bool                     `json:"truncated"`
}

func (p *GistPaste) GetID() string {
	return p.ID
}

func (p *GistPaste) GetURL() string {
	return p.HTMLURL
}
