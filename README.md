# go-mimetypes

## Description

Provides a large mapping of file extensions to mime types for golang.

Credit to [MimeTypeMap](https://github.com/samuelneff/MimeTypeMap) for the extensive list of mappings.

## Quickstart

``` sh
go get github.com/ao-com/go-mimetypes
```

Getting a mime type from a file extension:

``` go
mime := mimetypes.GetByExtension(".zip")
```