package gist

import (
	"errors"
	"fmt"
	"github.com/gme-sh/paste-wrapper"
	"github.com/imroc/req"
)

func (s *GistService) UploadPaste(d interface{}) (p paste.Paste, err error) {
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

func (s *GistService) doCreateRequest(payload *Create) (paste *GistPaste, err error) {
	if payload.Authorization == "" {
		return nil, errors.New("gist requires a authentication")
	}
	var resp *req.Resp
	if resp, err = req.Post(s.CreateApiUrl, req.BodyJSON(payload), req.Header{
		"Authorization": "token " + payload.Authorization,
	}); err != nil {
		return
	}
	res := new(GistPaste)
	if err = resp.ToJSON(res); err != nil {
		return
	}
	paste = res
	return
}
