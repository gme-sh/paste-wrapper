# paste-wrapper
golang wrapper for some paste services

## Installation
```bash
$ go get -u github.com/gme-sh/paste-wrapper
```

## Services
> to be continued
### paste.gg
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