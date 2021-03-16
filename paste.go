package paste

type Service interface {
	UploadPaste(interface{}) (Paste, error)
}

type UploadRequest struct {
	Authorization string
	ContentType   string
	Name          string
	Description   string
	Visibility    string
	Content       []byte
}

type Paste interface {
	GetID() string
	GetURL() string
}
