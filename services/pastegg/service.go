package pastegg

type PasteGGService struct {
	UrlPaste     string
	ApiUrlCreate string
}

const (
	DefaultPasteGGUrlPaste  = "https://paste.gg"
	DefaultPasteGGUrlCreate = "https://api.paste.gg/v1/pastes"
)

func New() *PasteGGService {
	return &PasteGGService{
		UrlPaste:     DefaultPasteGGUrlPaste,
		ApiUrlCreate: DefaultPasteGGUrlCreate,
	}
}

func (p *PasteGGPaste) GetID() string {
	return p.ID
}

func (p *PasteGGPaste) GetURL() string {
	if p.srv != nil {
		return p.srv.UrlPaste + "/" + p.ID
	}
	return DefaultPasteGGUrlPaste + "/" + p.ID
}
