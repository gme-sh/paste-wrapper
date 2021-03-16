package pastegg

type pasteGGService struct {
	UrlPaste     string
	ApiUrlCreate string
}

const (
	DefaultPasteGGUrlPaste  = "https://paste.gg"
	DefaultPasteGGUrlCreate = "https://api.paste.gg/v1/pastes"
)

func NewWithURLs(paste, apiCreate string) *pasteGGService {
	return &pasteGGService{
		UrlPaste:     paste,
		ApiUrlCreate: apiCreate,
	}
}

func New() *pasteGGService {
	return NewWithURLs(DefaultPasteGGUrlPaste, DefaultPasteGGUrlCreate)
}

///

type PasteGGPaste struct {
	ID     string
	prefix string
}

func (s *pasteGGService) newPaste(id string) *PasteGGPaste {
	return &PasteGGPaste{
		prefix: s.UrlPaste + "/",
		ID:     id,
	}
}

func (p *PasteGGPaste) GetID() string {
	return p.ID
}

func (p *PasteGGPaste) GetURL() string {
	return p.prefix + p.ID
}
