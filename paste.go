package paste

type Service interface {
	UploadPaste(interface{}) (Paste, error)
}

type UploadRequest struct {
	Name          string
	Description   string
	Visibility    string
	Content       []byte
	Authorization string
}

type Paste interface {
	GetID() string
	GetURL() string
}
