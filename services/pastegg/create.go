package pastegg

import (
	"fmt"
	"github.com/gme-sh/paste-wrapper"
	"github.com/imroc/req"
)

func (s *pasteGGService) UploadPaste(d interface{}) (p paste.Paste, err error) {
	switch payload := d.(type) {
	case *Create:
		p, err = s.upload(payload)
		return
	case *paste.UploadRequest:
		// wrap UploadRequest
		return s.UploadPaste(WrapPayload(payload))

	case paste.UploadRequest:
	case Create:
		return s.UploadPaste(&payload)
	}
	return nil, fmt.Errorf("invalid parameter: [%T] %v", d, d)
}

func (s *pasteGGService) upload(payload *Create) (paste *PasteGGPaste, err error) {
	var resp *req.Resp
	if resp, err = req.Post(s.ApiUrlCreate, req.BodyJSON(payload)); err != nil {
		return
	}
	res := new(responseCreate)
	if err = resp.ToJSON(res); err != nil {
		return
	}
	paste = s.newPaste(res.Result.ID)
	return
}
