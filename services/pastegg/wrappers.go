package pastegg

import (
	"github.com/gme-sh/paste-wrapper"
)

func WrapPayload(request *paste.UploadRequest) *Create {
	return &Create{
		Authorization: request.Authorization,
		Name:          request.Name,
		Description:   request.Description,
		Visibility:    request.Visibility,
		Files: []*CreateFile{
			{
				Name: "",
				Content: &CreateFileContent{
					Format:            "text",
					HighlightLanguage: "",
					Value:             string(request.Content),
				},
			},
		},
	}
}
