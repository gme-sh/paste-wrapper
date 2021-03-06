package pastegg

import (
	"fmt"
	"github.com/gme-sh/paste-wrapper"
	"github.com/imroc/req"
)

func (s *PasteGGService) UploadPaste(d interface{}) (p paste.Paste, err error) {
	switch payload := d.(type) {
	// wrap UploadRequest
	case *paste.UploadRequest:
		return s.UploadPaste(WrapPayload(payload))

	// non-pointers
	case paste.UploadRequest:
	case Create:
		return s.UploadPaste(&payload)

	// correct type
	case *Create:
		p, err = s.doCreateRequest(payload)
		return
	}
	return nil, fmt.Errorf("invalid parameter: [%T] %v", d, d)
}

func (s *PasteGGService) doCreateRequest(payload *Create) (paste *PasteGGPaste, err error) {
	// Authorization Header
	header := req.Header{}
	if payload.Authorization != "" {
		header["Authorization"] = "Key " + payload.Authorization
	}

	var resp *req.Resp
	if resp, err = req.Post(s.ApiUrlCreate, req.BodyJSON(payload), header); err != nil {
		return
	}
	res := new(responseCreate)
	if err = resp.ToJSON(res); err != nil {
		return
	}
	paste = res.Result
	paste.srv = s
	return
}
