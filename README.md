# paste-wrapper
golang wrapper for some paste services

## Installation
```bash
$ go get -u github.com/gme-sh/paste-wrapper
```

## Services
> to be continued
### GitHub Gist
https://gist.github.com
> **You have to provide an oauth token in order to create GitHub gists**
> * https://docs.github.com/en/developers/apps/scopes-for-oauth-apps
> * https://github.com/settings/tokens

**Simple request**
```go
srv := gist.New()
if p, _ := srv.UploadPaste(&paste.UploadRequest{
    Authorization: "< your auth token >",
    Name:          "Test",
    Description:   "This is a test",
    Content:      []byte("Hello World!"),
}); p != nil {
    log.Println(p.GetURL())
    // https://gist.github.com/....
}
```

**Advanced request**
```go
const cntnt = `import "fmt"
func main() {
    fmt.Println("Hello World!")
}`
srv := gist.New()
if paste, err := srv.UploadPaste(&gist.Create{
    Authorization: "< your auth token >",
    Description:   "This is a test",
    Public:        false,
    Files: map[string]*gist.GistFile{
        "test.go": {cntnt},
    },
}); paste != nil {
    log.Println("URL:", paste.GetURL())
    // https://gist.github.com/449c0d0a635e75bcd172147f2763ec81
}
```

### paste.gg
https://paste.gg

**Simple request**
```go
srv := pastegg.New()
if p, _ := srv.UploadPaste(&paste.UploadRequest{
    Name:          "Test",
    Description:   "This is a test",
    Content:      []byte("Hello World!"),
}); p != nil {
    log.Println(p.GetURL())
    // https://paste.gg/bbd0d8c6747846d19ef6c89d193a111e
}
```

**Advanced request**
```go
srv := pastegg.New()
expires := time.Now().Add(5 * time.Minute)
const contnt = `import "fmt"
func main() {
    fmt.Println("Hello World!")
}`
if p, _ := srv.UploadPaste(&pastegg.Create{
    Name:        "Test",
    Description: "This is a test",
    Visibility:  "unlisted",
    Expires:     &expires,
    Files:       []*pastegg.CreateFile{
        {
            Name:    "hello-world.go",
            Content: &pastegg.CreateFileContent{
                Format:            "text",
                HighlightLanguage: "Go",
                Value:             contnt,
            },
        },
    },
}); p != nil {
    log.Println(p.GetURL())
    // https://paste.gg/3d933dbf87314dde8e7ec24744bb4bdd
}
```