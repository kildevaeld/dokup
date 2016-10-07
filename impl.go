// Code generated by go-bindata.
// sources:
// index.js
// index2.js
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _indexJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x56\x5f\x6f\xbb\x36\x14\x7d\xe7\x53\x38\xbc\xc4\xd6\x0f\x79\xdb\x6b\xb2\xad\xda\xa6\x3d\xec\xa5\x9b\x34\xf5\x29\x8a\x2a\x0a\x97\xd4\x2d\xd8\xcc\x36\x69\xd7\x25\xfb\xec\xbb\x36\x36\x01\x12\xd6\x1f\x95\x42\xf0\xb9\x3e\x3e\xe7\xfe\x09\x3d\xe6\x9a\xfc\xdc\x89\xba\x04\x4d\x7e\x20\x1a\xfe\xea\x84\x06\x9a\x96\xaa\x78\x05\xcd\x9f\x7a\x24\x65\xdb\xc4\x05\x76\x62\x1c\xd3\x89\xb8\x2e\x71\x39\x90\xf0\x7b\x65\x45\x25\x8a\xdc\x0a\x25\x7b\x34\xd7\x87\xe3\x78\x5f\x23\xa4\x68\x84\xb1\x29\xa3\xad\x56\x05\x18\xc3\x5d\x08\x37\xb5\x28\x80\x7e\xc7\x02\x69\x65\xc6\x9b\x2a\x13\x0f\xab\x5a\x5c\x8f\x1b\x8b\xb7\x92\x32\xf2\x85\xa4\xdf\x94\xea\xb5\x6b\xf9\x8b\x49\xb7\x49\x92\x68\xb0\x9d\x96\x48\xc1\x8d\xcd\x2d\xad\x5a\x96\x10\xbc\xb8\x7d\x06\x49\xab\x4e\x16\x4e\x1d\xa1\x0e\x64\xe4\x1f\x8f\xf9\x0f\x77\xb9\x33\x3c\xd9\xe8\x78\x24\x40\xda\x71\x04\x48\xe7\xe9\x00\xf6\x11\xbf\x51\x36\x05\xd1\xce\xf0\x2c\x2a\x42\x9d\xbd\x1d\xc6\xed\x4f\xa7\xf8\x75\xf7\xed\x7e\xcf\x5c\x20\xb2\x2c\xc0\x03\x05\xd4\x06\x46\xa1\xfc\x91\xd7\x20\x0f\xf6\x99\xdc\x85\x67\x8c\x26\x1b\x92\xa2\x1d\x6d\xd3\x61\xdb\x44\xc2\x6a\xb5\xfa\x77\xb7\xf6\x11\xeb\x0c\xef\xaa\xc5\x9b\x86\x46\x1d\x61\x9d\x91\xb5\x2f\xf4\x7a\xcf\x85\x2c\xe1\xfd\xf7\xca\x29\x66\x31\x33\xf1\x2a\x94\x34\xaa\x06\x5e\xab\x03\x4d\x1f\x4c\x7e\x80\x4d\x48\xd4\xf7\x86\x3a\x62\x76\x72\xbc\x27\x4d\x3d\x2d\x3b\x3d\x51\xc7\xca\x7e\x4c\xd9\x84\x27\xd6\x0e\xde\x85\xc5\x72\x6f\x07\xf0\x7c\xc9\x62\x28\xa0\xee\x24\xf5\x47\x64\xce\x69\x76\xc9\x77\x2c\xc7\x99\x71\x6c\xb5\xe2\x79\x54\x54\x18\xeb\x9e\x68\xbe\x57\x24\x76\x09\xc1\xc3\x8d\x35\xab\x34\x23\x10\x14\x9c\x1d\x67\x32\xf0\x0c\x47\x05\xb6\x58\x47\x8e\xf2\xcb\xae\x8f\x39\x9d\xfa\x02\xb4\x2c\x0a\x4e\x2f\x68\xba\x9d\xee\xc3\xdc\x1f\x84\x3c\xf4\x35\xe6\xe6\xb2\x25\x00\x21\x3e\xac\xae\x4b\x38\x42\xad\xda\x06\xa4\x5d\x6f\x93\xf3\x48\x99\x4b\x4a\xa3\xca\x8c\x14\x0d\x7e\xa0\xc6\x28\x31\x8e\x60\xa1\x21\xb7\x10\x9e\xfa\x50\x17\x35\x24\x65\x3e\x06\x61\xca\x1d\xcd\xa4\x50\x38\xa7\xf6\xe1\xb7\x01\xde\x4e\x51\xf3\x26\x30\xf1\x84\xa2\x8a\x79\xa7\x4c\xba\x6f\x28\x45\x8e\x5d\x1c\x7a\x74\x13\x6d\x06\x6e\xee\x97\xa9\xd5\x1d\x8c\xfa\xe1\x13\xaa\xbe\x7b\xaf\xb9\xfa\x75\x4f\x96\x91\x05\xca\x28\x46\xb5\xb7\xb4\xa8\x96\x7e\xbd\x0c\xbf\xed\x9a\xc5\xdf\x97\x1c\x95\x50\xe5\x5d\x6d\x37\x57\x80\xa7\xfd\x9f\x39\x73\x69\x0a\x53\xe6\x5d\xce\x67\x2b\x5e\x8b\x33\xe6\x2e\x6c\xa6\xcb\xcc\x7d\x32\x40\x73\x41\x63\x37\x67\x36\xe9\xcb\x59\xbb\x04\x12\xf7\x7b\xd8\xf6\x7b\x62\x6a\x94\xa4\xb7\xde\x16\xbf\x1e\xb1\xd7\x33\x32\x12\x92\x91\x66\xac\xc5\x51\x19\xab\x17\x5e\x36\x3b\xd8\xbb\x17\x01\xfe\x7d\x21\xf4\x27\xad\xf3\xbf\xb9\x30\xfe\x4e\x91\xe5\x8e\x34\xbc\xc9\xdb\x91\xcd\x8f\xb9\xcd\x50\xc0\x0f\x2e\xf3\x06\xc6\x09\x7a\x51\x42\x52\x24\x66\xf8\x23\xdb\x78\xf4\x92\xf4\x38\x05\xd7\x49\x73\xcd\x21\xb9\x57\x8a\xd3\xbd\xb9\x05\xfe\xe9\xca\xb9\x04\xfe\xe2\xa6\x78\x79\xa7\x6a\xdb\x2b\xd0\x57\x1e\xd3\xd3\x09\xfe\x47\xdf\x00\xd4\xe5\xcb\x65\x85\x73\x9e\xde\x68\xc4\xb6\xd7\x40\xaf\xbb\xe8\x09\x8f\x7f\xdd\x2e\x5a\x5a\xf6\x03\xb7\x31\x6f\x67\x01\xf3\x6e\xe6\x58\x90\xd7\x15\xde\x07\xfe\x47\x22\xe1\xa6\x03\xf4\x2b\xbb\xba\xbe\x46\x66\x0e\xce\xf1\x8d\x91\x9c\xff\x0b\x00\x00\xff\xff\x88\x8b\x18\xb7\xf5\x08\x00\x00")

func indexJsBytes() ([]byte, error) {
	return bindataRead(
		_indexJs,
		"index.js",
	)
}

func indexJs() (*asset, error) {
	bytes, err := indexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.js", size: 2293, mode: os.FileMode(420), modTime: time.Unix(1475173241, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _index2Js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe2\xe2\x4a\xad\x28\xc8\x2f\x2a\x29\xd6\x4b\x2a\xcd\xcc\x49\x51\xb0\x55\x48\x2b\xcd\x4b\x2e\xc9\xcc\xcf\x53\xd0\xd0\x54\xa8\xe6\xe2\x4c\xce\xcf\x2b\xce\xcf\x49\xd5\xcb\xc9\x4f\xd7\x50\x77\x0a\xf5\xf4\x71\x51\xd7\xe4\xaa\x05\x04\x00\x00\xff\xff\x10\xfd\xc0\x46\x37\x00\x00\x00")

func index2JsBytes() ([]byte, error) {
	return bindataRead(
		_index2Js,
		"index2.js",
	)
}

func index2Js() (*asset, error) {
	bytes, err := index2JsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index2.js", size: 55, mode: os.FileMode(420), modTime: time.Unix(1475617235, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"index.js": indexJs,
	"index2.js": index2Js,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"index.js": &bintree{indexJs, map[string]*bintree{}},
	"index2.js": &bintree{index2Js, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

