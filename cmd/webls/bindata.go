// Code generated by go-bindata.
// sources:
// listing.tmpl
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

var _listingTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x94\xcf\x6e\x9c\x30\x10\xc6\xef\x7d\x8a\x51\x7a\x8b\xc2\x5a\x55\x73\x48\x88\xbb\x97\x54\x95\x2a\xf5\x9f\x94\xbe\xc0\x80\xc7\xc1\x5a\x8c\x91\xf1\xa6\xcb\x22\xde\xbd\x32\x04\x58\xb3\xb0\x5a\x9f\xf0\xcc\xf7\xcd\x8e\x7f\x33\x5a\x9e\x39\x9d\x6f\x3f\x00\x00\xf0\xca\xd5\x39\xf5\xdf\xfe\x24\x46\xd4\xd0\x8c\x57\x7f\xa4\x29\x5c\x24\x51\xab\xbc\x8e\x41\x9b\xc2\x54\x25\xa6\xf4\x14\x68\x12\x4c\x77\xaf\xd6\xec\x0b\x11\xa5\x26\x37\x36\x86\x24\xc7\x74\x17\x8a\xde\x33\x1f\x3f\x3d\xd0\xc3\xfd\xe3\x94\x6b\xc7\x2f\x8c\xdf\x54\xa5\x1c\x89\x59\x0b\x83\xf3\xfe\xd1\x7b\x17\x9d\x2b\x8e\xf5\xdf\x12\xea\x6d\xe6\x89\xfe\x51\xb2\x53\xce\xbf\x60\xaf\x8b\x28\x35\xfb\xc2\xc5\xf0\xf9\x09\xd8\x2d\x3c\x67\xd6\x68\xba\x83\x17\x94\x68\xd5\x1d\xfc\x2e\xc9\x22\xdc\xb2\xb0\x80\x36\xc7\x45\xf7\x37\x65\x49\x9a\xc3\x5c\x3f\x97\xce\xba\xe4\xec\x64\x38\xdc\x4f\x66\x9a\xd3\xf7\x42\xd0\x01\x8c\x84\xa6\x01\x8b\xc5\x2b\xc1\x46\x28\x5b\x41\xdb\x72\x84\xcc\x92\xfc\x72\xd3\x34\xb0\x79\x36\xba\xf4\x53\x81\xb6\xbd\xd9\xfa\xc0\x57\x65\xbd\x86\xa1\xbf\x51\x97\x00\x9e\x58\x36\x55\xe6\x0e\x93\xd3\x8d\xe8\x63\x19\xa1\x08\x63\x7d\x5c\x6c\xe1\x17\x6a\x02\xce\xdc\x5a\xfe\x07\x56\x0e\x7e\x1a\xa1\xa4\x22\x71\x49\xf8\xb7\x2e\x2f\x16\x7a\x51\xc7\xa5\x3c\x67\x0b\xdd\x71\x17\xf2\x1a\xce\x84\x4b\xaa\x9c\x3c\xaf\x33\x49\x6f\xb7\xe7\xde\xa0\x99\x13\xcc\x9b\x3f\xe8\xb2\x01\x70\x47\xa3\x27\xbc\xf2\x96\xa0\x8c\xb7\x78\x40\x23\x1f\x3f\x8f\xeb\x6c\x1d\xae\xeb\xe5\x1d\xbd\x8b\x72\xce\x96\x5e\x3d\xee\xc9\x1c\x7a\x08\x98\xb3\xd9\xde\x70\x56\xda\x61\x77\x59\xaf\xe5\xac\xfb\xe3\xf9\x1f\x00\x00\xff\xff\x2d\x74\x44\x30\x7f\x04\x00\x00")

func listingTmplBytes() ([]byte, error) {
	return bindataRead(
		_listingTmpl,
		"listing.tmpl",
	)
}

func listingTmpl() (*asset, error) {
	bytes, err := listingTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "listing.tmpl", size: 1151, mode: os.FileMode(436), modTime: time.Unix(1494144717, 0)}
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
	"listing.tmpl": listingTmpl,
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
	"listing.tmpl": &bintree{listingTmpl, map[string]*bintree{}},
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

