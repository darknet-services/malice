// Code generated by go-bindata.
// sources:
// bindata.go
// config.toml
// plugins.toml
// DO NOT EDIT!

package data

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

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1452971511, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x53\xff\x4b\xe3\x4e\x10\xfd\x3d\x7f\xc5\xd0\xc2\x87\xcf\x81\x8d\x9b\x94\x5a\x5b\x10\x11\xf1\xbe\xa1\x77\xa0\x07\x07\xe7\x89\x6c\xb2\xd3\x64\x71\xb3\x1b\x76\x37\x5a\xfd\xeb\x6f\x66\xd3\x2a\xf7\xe3\x81\x29\x34\x9d\x79\x3b\xef\xcd\xbc\x9d\x4e\xe1\xdc\xf5\xcf\x5e\x37\x6d\x84\xff\xeb\x0f\x50\x8a\x62\x0e\x33\x7e\x2d\xa0\x32\xb2\x7e\x88\xae\x87\xaf\x2e\xb4\x83\x84\x2b\xa9\x2d\x1e\xc0\x99\x31\x70\xcd\x05\x01\xae\x31\xa0\x7f\x44\x95\x67\x53\xb8\x41\x84\xcb\x2f\xe7\x17\xdf\x6e\x2e\x60\xe3\x3c\x18\x5d\xa3\x0d\x08\xda\x52\xd4\xc9\xa8\x9d\xcd\xb3\x6c\xfa\x3e\x0f\xe9\x5d\x9d\xb1\x1a\xb5\x6f\x37\xba\x19\x7c\x12\x80\x7f\xe7\x79\xa7\x7e\xb2\xa8\xa3\x41\x38\x81\xc9\x95\xe4\xc9\xe1\x7a\xb0\x51\x77\xf8\x77\x7f\x93\x2c\xbb\x95\x43\x6c\x9d\xbf\xcb\x00\xac\xec\x52\xc5\xde\xe7\x09\xe5\xa6\xe0\x7c\x23\xad\x7e\x19\xe7\x21\xf4\x93\x8e\x9f\x87\x6a\xc4\x2a\xed\xde\x52\x44\xbd\x71\x83\x55\xe8\xe1\x3f\x38\xbf\xf8\xfe\xdb\x5e\xea\x07\x0c\x10\x65\xa4\x54\x74\x74\x3f\xd2\x2a\xa8\x10\x7d\x3e\x96\x2b\x57\x51\x79\xb1\x5a\xae\x66\x62\x31\x2b\x97\x3f\xc4\x72\x3d\x2f\xd7\x42\xfc\x22\xf0\xa3\xf6\x21\x42\x6d\x64\x08\xa0\x88\x22\x9c\xc2\xcf\xf6\x19\xac\x8b\xa7\xd4\xf5\x13\x56\xdc\xf2\xe0\x0d\x37\x20\xf2\xf4\x59\x1f\x0b\x26\x96\xaa\xd3\xf6\x7e\x07\x15\xe5\x32\x81\xc5\x7a\x4e\x0f\x0f\x8c\x9d\xd4\x86\x8b\x5b\x47\x02\x74\x24\x74\xb1\xcf\x71\x2b\xbb\xde\x60\x5e\xbb\x8e\x39\x7a\xe7\x19\x2b\x17\x2c\x42\x4b\xc5\xe7\xf8\xcd\x1e\x25\x9c\xdb\xa2\x1c\xbf\x9f\x9c\x57\x4c\x4c\x5d\xca\x4a\x06\xbc\x4b\x78\x6c\x19\xef\x92\xfb\xb9\x4a\x86\xa5\xe5\x4c\x54\xc5\xaa\xcc\x8b\xa3\xe3\xbc\xc8\x8b\xbd\x1a\xd3\xdd\xc2\xb1\x10\xc5\xc1\xdb\x77\x09\x4c\x56\x3b\x6b\xb1\xe6\x0b\xb8\xef\xe4\x96\xce\x2d\x84\x10\x94\x47\x2b\x2b\x83\x8a\x12\xd1\x0f\xc8\xa3\xd9\x47\xed\x9d\xed\xd0\x46\xae\xf3\x43\xba\x31\x85\x8f\x68\x5c\xcf\xd9\xd1\xf7\x5d\xbe\xf7\x4e\x0d\xf5\x7e\x0f\x94\xab\x1f\x30\xed\x41\x27\xeb\x96\xfe\x59\xb3\xfd\x3e\xec\xe2\x49\x52\x54\xbd\xd3\x36\xd9\x16\xeb\x7e\x7d\x78\xb8\x9f\x64\xb5\xca\x0b\x21\xd6\xe5\x7c\x79\xc4\x6c\xc6\x35\xcd\xc8\xb6\xd1\x06\xdf\x98\x92\x1b\x04\x4e\x92\xce\x36\xe8\x17\x06\x0a\x31\x86\xb2\xe1\x68\xbe\x8b\x2a\x5a\xc3\xa1\x67\x5f\x96\x94\x30\xae\x96\x26\xad\xf0\x09\x6c\xa4\x09\x3c\x2f\x4d\xb0\x7d\xbe\x7b\x75\xe2\x15\xa1\xbb\x8d\xb1\x67\xc5\xc9\xee\x77\x18\x83\x3f\x01\x00\x00\xff\xff\xcd\xe9\x18\xae\x5f\x04\x00\x00")

func configTomlBytes() ([]byte, error) {
	return bindataRead(
		_configToml,
		"config.toml",
	)
}

func configToml() (*asset, error) {
	bytes, err := configTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.toml", size: 1119, mode: os.FileMode(420), modTime: time.Unix(1452396945, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x93\x4b\x6f\x13\x31\x10\xc7\xef\xfb\x29\x46\xc9\x05\x10\x6d\x78\x88\x23\x87\x28\x69\xa1\x28\x84\x68\x53\x7a\xa9\x7a\x98\xd8\xb3\x9b\x11\x7e\xc9\xf6\x56\xe4\xdb\x33\x4e\xb5\x85\x86\xa2\x2c\x52\xf6\xb0\x96\xc7\x33\xf3\xff\x69\x1e\x63\x98\xf9\xb0\x8b\xdc\x6e\x33\xbc\x50\x2f\xe1\xdd\x9b\xb7\xef\xe1\xac\x1c\x1f\x60\x63\x50\xfd\xc8\x3e\xc0\x17\x9f\xb6\x1d\xc2\x57\x64\x47\xaf\x61\x6a\x0c\xd4\x25\x20\x41\x4d\x89\xe2\x3d\xe9\xf3\x6a\x0c\x6b\x22\x58\x5c\xcd\x2e\x96\xeb\x0b\x68\x7c\x04\xc3\x8a\x5c\x22\x60\x27\x37\x8b\x99\xbd\x3b\xaf\xaa\xf1\x69\x3e\xd1\x5b\x2d\xbe\x7f\xba\x5a\x0a\xbe\x6b\xb8\xed\xe2\x5e\x00\xfe\x3f\xcf\x89\x78\xaa\xdb\xdb\x60\xba\x96\xdd\xdd\x5d\x05\xe0\xd0\x12\x7c\x84\x91\x4b\xd1\x8c\xe4\x4e\x0e\x37\x86\xb4\x98\x1a\x34\x89\xc4\xa2\x29\xa9\xc8\x61\x0f\x2d\x8e\xcb\x75\xbd\x80\x39\x66\xdc\xa0\x94\xec\x33\xa6\xad\xd4\x13\xa3\xda\x96\x68\x85\x99\x5a\x1f\x77\xc5\x71\xc3\x0e\xe3\xae\x58\xd9\x62\xbb\x17\xe9\xdb\x34\xe9\xd5\x2c\x3f\xa8\xbf\x1a\xfd\x03\xab\x61\x43\xa5\x2d\x4f\xd1\x72\xec\x9e\x21\x4b\x49\x13\x85\xc9\x75\xcd\xf3\x09\xfd\xe4\x26\x7b\x6f\x0e\xa1\x2c\x65\xd4\x02\xff\x04\xcb\x62\x99\x80\xc9\x9f\x5a\x47\xc1\x94\x41\x8b\xf7\x43\xb0\x66\xe2\x39\xbd\x39\x04\x79\x88\x3d\x44\xf8\x9d\xf5\x78\x65\x42\xf4\x79\x88\xfe\xe5\xd9\xaa\xfe\x76\x0d\x53\x97\xf9\x86\x63\x97\x86\x91\x3c\xa6\x3f\x0a\x12\x74\xc3\x7a\xd0\xe8\xac\xe6\x97\x3c\x97\x95\x2d\x8d\x91\x9f\xf0\x72\x11\x15\x33\x68\xaf\x3a\x4b\x2e\xff\x45\xd7\x3f\x3c\x3f\x47\x8f\xda\x3d\x25\x86\x20\xf8\xfb\x0d\x2b\x8f\xa3\xea\x57\x00\x00\x00\xff\xff\x2d\x0a\x9b\xfb\x36\x04\x00\x00")

func pluginsTomlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsToml,
		"plugins.toml",
	)
}

func pluginsToml() (*asset, error) {
	bytes, err := pluginsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins.toml", size: 1078, mode: os.FileMode(420), modTime: time.Unix(1452755045, 0)}
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
	"bindata.go":   bindataGo,
	"config.toml":  configToml,
	"plugins.toml": pluginsToml,
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
	"bindata.go":   &bintree{bindataGo, map[string]*bintree{}},
	"config.toml":  &bintree{configToml, map[string]*bintree{}},
	"plugins.toml": &bintree{pluginsToml, map[string]*bintree{}},
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
