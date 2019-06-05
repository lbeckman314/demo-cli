// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// cert.pem
package cert

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _certPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x55\xc9\xce\xb3\xb8\xda\xdc\x73\x15\xff\x3e\xfa\x15\x20\xf3\xa2\x17\x8f\x8d\x19\x92\x00\x31\x10\x12\xd8\x31\x24\x86\x0c\xcc\x60\xe0\xea\x8f\xde\xf7\xd3\x69\xb5\xd4\x8b\xe3\x65\xd9\x55\x8f\xab\x6c\xa9\xfe\xff\x67\x21\xa2\x19\xd6\xff\x61\xe2\x78\x86\x6a\x60\xf0\xc8\x2f\x2a\x98\x86\xa1\xfa\x1e\xc6\x48\x91\x19\x70\x03\x01\x33\x5c\x90\xf1\x9b\xf7\x76\x22\xbd\x7b\xb5\x5e\x03\x59\xa8\xa1\x19\x4a\x4e\x69\x82\xa8\x61\xb7\xd6\x5c\x23\x5e\x29\x94\x20\xcc\xaf\x20\x98\xa4\x1c\xf1\x0c\x47\xc4\x2c\x1f\x41\xe0\xc1\xc7\xf7\x4c\x27\xe0\x2a\x04\x8a\x4f\xe9\x89\x70\xc9\x0c\xef\xb4\x48\x26\xa4\xc6\x5f\x6b\x7a\xdc\x91\x68\xba\x26\x37\xe8\xef\xbe\x22\x90\xb1\xfa\xd7\x01\x83\xa8\x52\xaa\x65\x43\xf2\xfd\x88\x0f\x17\x05\xe6\x0c\x0f\x95\x8b\xa3\xed\x81\x6c\x2a\x74\x32\x3d\x3a\x5a\x4a\x1a\xfd\x60\x82\xed\xc1\xc6\x54\x8c\xbf\x41\x13\x95\xa3\xa6\xc0\xed\xcf\x8d\x4c\x4f\x91\xc7\x2a\xb8\x49\x79\x78\xb3\x9a\xf8\xa6\xf6\xe7\xaf\x35\xc4\x1e\xc6\xe0\x1a\x5c\xa1\xc1\x51\x38\x95\xa1\x91\x0d\x89\x05\x94\x20\x44\x41\x61\x8c\x5c\x40\xc1\x18\x68\x89\x19\x23\x08\xcc\xe1\xf5\x21\xf5\x69\x48\x3a\xef\xf2\xe8\x96\xfd\x0d\x2b\x7b\x6f\xba\x44\x87\xb2\x2a\x4f\x57\x2a\xf8\xcf\x2c\xbf\x5b\xa7\x9d\x74\xef\x88\xf9\xb4\x82\xbe\x0b\xc4\xcb\xfa\xad\xbb\x57\x34\xdd\x9a\x6f\x24\x47\x73\x50\x7c\x31\xdd\xae\xe0\x1b\x8b\xbd\xe3\x4b\x5a\x87\xf5\x5d\xb5\x9e\x87\x2a\xb1\x37\xaa\xf0\xbc\x9f\xd9\x4e\xae\xa4\x6e\x1e\x1e\x4f\x45\x0f\x44\x3c\x78\xc6\x4a\xcd\x58\xce\x1b\xc3\x92\xe9\x59\x0c\x97\x45\x3d\x54\xc4\x0f\xde\x87\x75\xb0\x2f\xdc\x2e\x9a\x02\xec\xb4\x4a\x28\x8b\x2b\xec\x0b\x0b\x45\xbf\x9c\x4e\xd2\xf4\xf0\x96\x1a\xd5\x22\x2d\x47\x67\xf7\x13\xf5\xef\xe3\xae\x79\xb0\x83\x87\x8e\xeb\x5d\xc4\x8d\x95\x7b\x28\x9f\x71\x76\xbc\xec\x8a\x49\xff\x2c\x25\xeb\x73\xcc\xd6\x24\x75\xab\x44\x90\xa4\x6b\xeb\xc9\x41\x2c\x66\x1c\xee\xab\x57\xdf\x8f\x7c\xb9\x0a\x93\x76\x4c\x7a\xbd\xdc\x2f\x9d\xcf\xe5\xe1\x3f\x88\x2d\xde\x8c\xc5\x25\xe7\x9d\x32\xb4\xea\xa6\x8d\x6d\xe0\xd7\xfa\xbd\x87\x53\x29\x6c\xdd\x30\xac\xdd\xb3\xa5\xc5\xe3\x5d\xe5\xca\xdc\x27\x0d\xf2\xaf\xdd\xa2\xb2\x6a\x6b\xf7\x66\x2d\x0a\x57\x90\x84\x13\xbd\x21\x63\xd2\x3e\x77\xb6\x0d\x9e\xf6\x73\x21\x62\xe0\x04\x20\xb2\xb1\x00\x5f\x93\x33\x76\x7c\x9a\xb0\xd6\x40\xba\xa6\x0a\x27\x68\xc9\x29\x01\x6e\xa8\xa5\x02\xe9\xcf\x73\xea\xee\x95\xa8\x2f\xb8\x22\xc6\x1a\xc4\x88\x8a\x68\xa2\x00\x0d\x8c\x13\x0f\x90\x80\xe8\x55\x07\x6e\x70\xe5\xe7\xe7\x39\xa2\x07\x54\x5f\x22\x30\x38\xfc\x97\x4c\xd7\x44\x65\xf4\xfa\x59\xe7\xd8\xad\x82\xb5\x77\xff\x78\xf1\x2e\xf4\x15\x6c\xbb\x81\x14\xbb\xc2\xbe\x34\xb8\xce\x7f\xc9\x2f\x84\x18\x57\x4b\xb8\xd6\xa4\xaa\x03\xee\x1c\x56\x71\xf3\x15\xbd\xef\xaa\x7a\xfb\x9f\xdd\xd2\x2e\x77\x27\xc2\x63\xfe\x67\xf2\xef\x60\x4a\x88\x10\xcc\x28\x33\xf1\x5a\xc3\xb8\xd5\x80\x5e\x55\x34\x03\xca\xf2\x63\x99\xea\x0e\xb7\xf3\xfd\x10\xcb\xd6\x9c\xe0\x4d\x15\x17\xb4\x7b\x28\x66\x1f\x6b\xbe\x98\xc8\x7e\x1f\xac\x8e\x9b\x44\xa7\xbd\x10\xaf\x8e\x85\x89\xf7\xff\x14\xc0\x59\x6e\xfd\x2d\x10\xc8\xfe\x94\xfe\x2f\x81\xf3\x0c\xd1\xaf\x5f\x87\x10\x32\x83\xc3\xd8\xa1\x8d\x6e\x6a\x17\x7c\xfd\x57\x24\x4b\x59\x9c\x6f\x5e\xb1\x2c\x72\xef\x4f\x4e\x0c\x91\x2b\xa7\x33\x18\x88\x85\x85\xc0\x28\x47\xc0\x08\xb7\x78\x70\xfe\xf5\x06\x1a\x5e\x8d\x26\x02\x4a\xf8\x49\x81\xef\x3f\x92\xc7\xe0\x04\x51\xa4\x3b\x62\xa2\x94\xc3\x59\xb6\x78\x32\x6d\x5a\x21\xbc\x3b\x73\x78\xdb\xbc\x92\xe2\xc3\x53\xbc\x19\x92\x6f\xc2\x19\x23\x0a\x62\xe5\x0f\x91\x40\x1a\x6e\x80\x51\x8c\x0c\x5d\x44\x86\x3e\x82\xbd\x87\x94\x22\xf1\xd9\xd5\x8a\x19\x89\xb3\x40\x8e\x24\x2b\xcc\xf5\xc7\x13\x8f\x9c\x37\xcb\xc3\xfd\x6d\x30\x6c\x06\x2b\xeb\x5e\x7c\x89\x3e\x98\x3e\x03\x80\x5b\xfb\xae\x0e\xc9\x0d\x00\x08\x70\xa4\x99\x84\x62\x03\xe9\x43\x5e\xd9\xb5\xf0\x2a\xdd\xf1\x28\x16\xf9\xeb\xa3\x35\x4e\x18\xec\x7d\x4f\xd7\x77\x5b\x55\xf3\x1e\x43\x89\x16\x79\xdf\x24\xb0\x86\x1c\xdd\x90\x3d\xd7\x59\xbb\xcf\xb7\x6c\x44\x43\xa7\x2c\xfc\x17\x90\x57\x69\xc4\x02\xdd\x76\xfd\xd0\xb4\xf5\xab\xb0\xb0\x39\x4b\x8b\x03\x45\x32\x68\x97\x49\xde\xaf\x4a\x7b\xcf\xdb\x99\x77\x53\x86\xd3\x7b\x34\xdb\xef\xd7\x4d\x5d\xbd\x76\x92\x54\xbd\xf2\x71\x94\xb3\xab\x7b\xc8\x2d\x01\x00\x50\x34\xb9\x85\x34\x9a\x00\x40\x15\x20\x09\x77\xa8\xc1\xcc\xef\xa4\x71\xa3\xc7\xa7\xb7\xf9\x1c\xaf\x55\x40\x9b\x78\xb9\xa8\xb4\x16\x3d\x1b\x52\x6d\x57\x8f\xf8\x6c\x99\xb6\x9f\xa6\x42\xd0\x03\x36\xa8\x92\xad\xfb\x53\x62\xd5\x97\x61\xbc\x78\x15\x4d\xa8\x56\x28\xb7\x36\x69\xfb\x01\x5d\xf6\x83\x6f\x3d\xbd\xa5\x7a\x31\xab\xc7\xfd\xc6\x66\xb0\x10\x7b\xd7\xd9\x3b\xd7\x0e\x5c\x44\x02\xd0\x56\x05\xb0\x31\x50\x02\xec\xb4\x9d\xab\x07\x43\xa5\xba\x35\x34\xd9\xd9\xb9\x85\xd3\x9f\xc4\x6b\xff\x3c\x01\x09\xc2\xd8\x3e\xb4\xea\xf2\x9c\x76\x5d\xd4\x49\xdf\x67\x28\x6a\xe3\xfb\xf6\x10\x2a\x36\xd9\x09\x7e\x18\xf5\x26\xec\x2e\x4d\x3a\x05\x5e\xaa\x4d\xf3\xc7\xf0\xfc\x59\x57\xae\xab\xe1\x60\x6f\xc3\xf1\x74\x74\x70\xbe\xe5\xb5\x78\x5e\xed\xc7\x68\xd7\x6f\xbd\xde\xbe\xaf\xfd\xc3\xb5\xd0\x05\x29\xfb\x70\x75\x5e\x46\x19\x1b\x20\x30\xfc\x63\xc5\xbb\x6f\xda\x41\x78\xdc\xb2\xc3\xb6\x51\xf3\xd0\x9b\x8a\x57\xd5\x93\xa6\xe9\xbd\xfe\x20\x5f\xd3\x65\x3a\x17\xdc\x3f\x2f\x23\x71\x76\xad\xd4\x13\xfc\xfd\x23\xbf\x0c\x59\x16\xce\xe5\x29\x92\xe5\xc2\x74\x48\x93\xea\x84\xdc\x45\x8d\x38\x1b\x86\xbe\xb5\xbc\x79\x6a\x1f\xfa\xf4\x0e\x54\xa9\xfd\xc1\x79\xcd\x99\x71\x1e\x95\xef\x6e\xaf\xb1\xf7\x26\x14\x4a\xb4\xed\x76\xe4\xa4\x04\xcd\x49\x0f\xc5\x7d\xee\x6e\x64\x65\x1e\xf1\x79\x29\xa9\x69\xc3\x5e\x56\x4d\x56\xfd\xc2\xcb\x9a\x77\xfd\x4e\x8b\xcf\x9b\x14\xd1\xa7\x15\xdb\xf1\xf9\x2e\xbc\xdb\xd0\x1e\xce\xc2\xbe\xfe\xf2\xa7\x36\xc7\x76\xd1\x92\xa2\x58\xaa\x97\xc1\x58\x0e\xfd\x71\xbb\x37\xeb\x4f\x97\x84\x63\xcc\xff\xfa\x4b\xf8\xad\x51\x62\x29\xff\xae\xd6\xff\x04\x00\x00\xff\xff\xa4\x77\x42\xbd\x77\x07\x00\x00")

func certPemBytes() ([]byte, error) {
	return bindataRead(
		_certPem,
		"cert.pem",
	)
}

func certPem() (*asset, error) {
	bytes, err := certPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cert.pem", size: 1911, mode: os.FileMode(420), modTime: time.Unix(1559688895, 0)}
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
	"cert.pem": certPem,
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
	"cert.pem": &bintree{certPem, map[string]*bintree{}},
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