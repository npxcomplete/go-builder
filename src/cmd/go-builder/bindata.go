package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _buildertemplate_txt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\xd1\x4a\xf3\x30\x14\xc7\xaf\x9b\xa7\x38\x8c\x8f\x8f\x76\xb8\x3c\xc0\x60\x17\x2a\x0c\xbc\x50\x84\xf9\x00\xeb\xda\x6c\x4e\x9b\xa6\xa4\x89\x10\xc2\x79\x77\x49\x9b\xce\x95\x36\xb5\x8a\x77\x5b\x4e\xce\xbf\xbf\x5f\x72\x52\xa5\xd9\x7b\x7a\x62\x60\x2d\x7d\x7e\x4a\x39\x43\x24\x44\x99\xaa\x59\x78\x68\x17\xee\xf4\xb9\xc8\x99\x84\x5a\x49\x9d\x29\xb0\x24\xb2\x56\xa6\xe5\x89\xc1\xbf\x8f\xb4\x80\xf5\x06\xe8\xf6\xcc\x8a\xbc\x86\x15\xa2\x2b\xba\x65\xea\xbb\x6e\x95\x92\xe7\x83\x56\x0c\x11\x7c\xe5\xc5\x54\xac\x4d\x86\xbd\x49\x79\xb1\x5e\x04\x5b\x16\xf0\x56\x8b\x72\x72\xc3\x51\x48\x3e\xb5\x61\xef\x88\x56\xc0\xca\x1c\x91\x20\x21\x53\xec\x47\x5d\x66\x10\x73\xe3\x48\xe9\xae\x65\x4c\x3a\xec\x5d\xa3\xff\xc8\xd4\xab\xc8\x11\xe3\x64\x44\xc7\x92\x48\x32\xa5\x65\x09\xdc\xd0\x5e\xdb\x15\x11\x41\xf2\x05\x34\x03\xe7\xe0\x8f\x7f\xe9\xa0\xfa\x57\x72\x61\xf0\xff\x2f\x70\x3f\xb8\x83\x64\x2c\xd8\x99\xf8\xef\xd2\x70\xd6\x06\x82\xb5\xcb\x41\xf8\x94\xbe\xf4\x40\x6c\xe0\xb5\x95\x82\xc7\xb5\xd0\x32\xbb\x9e\xc4\x64\x64\xaf\x43\x5d\x76\x51\x9b\xe1\xd8\x5a\x12\x4d\xcf\x6b\x78\x60\xd7\xd0\x12\xd0\xc0\x04\xdc\x34\xcd\x9d\x55\x34\x26\x3d\x43\xf5\x5e\x54\x26\x0e\x99\xf9\xc0\xff\x7f\xab\xf5\xed\xd5\x0e\xcc\x66\x99\x34\x3f\x9a\x87\xd1\x95\xae\x14\xac\xed\x5e\xd4\x6c\xf4\xc1\xd3\xf9\x1d\xf9\x67\x00\x00\x00\xff\xff\xd6\x45\xc2\x46\xe3\x04\x00\x00")

func buildertemplate_txt() ([]byte, error) {
	return bindata_read(
		_buildertemplate_txt,
		"builderTemplate.txt",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"builderTemplate.txt": buildertemplate_txt,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"builderTemplate.txt": &_bintree_t{buildertemplate_txt, map[string]*_bintree_t{
	}},
}}
