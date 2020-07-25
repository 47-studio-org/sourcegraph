// Code generated by go-bindata. DO NOT EDIT.
// sources:
// state.html.tmpl (4.802kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _stateHtmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdd\x73\xdb\x36\x12\x7f\xf7\x5f\xb1\xc7\x69\xa7\x49\x5c\x12\xfa\xb4\x63\x1f\xa5\x89\xeb\x8f\x5a\x6e\x6c\xc7\x92\xed\xc6\xb9\xde\x03\x44\xae\x44\x48\x20\xc0\x00\x4b\xd9\xb2\x46\xff\xfb\x0d\x49\x49\xb6\x12\x49\xe7\x5c\x67\xee\x6e\x3a\xc1\x8b\x08\x60\xb1\x9f\xbf\xdd\x15\xe0\xff\x2d\xd4\x01\x8d\x13\x84\x88\x62\xd9\xdc\xf2\xb3\x1f\x90\x5c\xf5\x1b\x0e\x2a\xa7\xb9\xb5\xe5\x47\xc8\xc3\xe6\x16\x00\x80\x4f\x82\x24\x36\x0d\x26\xda\x4d\x93\x90\x13\x1a\xd7\x12\x27\xf4\x59\xb1\x53\x50\x49\xa1\x86\x60\x50\x36\x1c\x4b\x63\x89\x36\x42\x24\x07\x22\x83\xbd\x86\x13\x11\x25\x76\x9f\x31\x4b\x3c\x18\x26\x9c\x22\xaf\xab\x35\x59\x32\x3c\x09\x42\xe5\x05\x3a\x66\x8b\x05\x56\xf3\xea\x5e\x89\x05\xd6\x3e\xad\x79\xb1\x50\x5e\x60\xad\x03\x42\x11\xf6\x8d\xa0\x71\xc3\xb1\x11\xaf\xbe\xad\xb9\x7b\xbc\x45\x15\xd5\x4e\x0e\xcb\x95\x9b\xe1\x5e\xbf\xb3\xd7\xe5\x47\xb2\x56\x2e\x5f\x5c\x1d\x24\x27\xf1\x61\x65\xe7\xf8\xfe\xe0\xf2\xf4\xed\xef\xfd\x4f\xb2\x7e\x7e\x77\xf7\x70\xd2\x0b\xb6\x2f\x82\x0f\xdd\x72\xf8\xdb\xaf\x83\xdd\xce\xd0\x81\xc0\x68\x6b\xb5\x11\x7d\xa1\x1a\x0e\x57\x5a\x8d\x63\x9d\x5a\xe7\x1b\xec\xca\x8c\x18\xd8\x10\xa5\x18\x19\x4f\x21\x31\x95\xc4\xec\x5d\x4f\x1b\xe2\xf7\x68\x75\x8c\xac\xa7\xd5\xfc\xdb\xed\x19\xc4\x77\x75\xaf\x5c\x9b\x99\xc9\xa5\x5c\x18\x38\x13\x9a\x8b\x2a\xbe\xb3\xe1\x75\xad\x4b\x5a\x4b\x12\x89\xdb\xd5\x44\x3a\x06\x6f\x3e\x17\x4a\xa1\x81\xc9\x82\x36\x1b\x31\x7f\x70\xef\x45\x48\xd1\x3e\x94\x4b\xa5\x1f\xff\xbe\xd8\x9c\x16\xec\xd9\x8c\xbf\xcf\x8a\x30\x6f\xf9\x5d\x1d\x8e\x67\xb2\x43\x31\x82\x40\x72\x6b\x1b\x4e\x90\x69\x2d\x14\x1a\xe7\x49\x97\xc9\xe4\x07\x1b\x44\x18\xa6\x12\xcd\x51\x1a\x27\xb0\xdf\x80\x9e\x30\x96\xc0\x9b\x4e\x61\x32\x61\x6f\x20\xe4\xc4\x41\x77\x07\x18\x10\x08\x0b\x1c\xa4\xb0\x04\xba\x07\x21\x76\xd3\x3e\x84\x69\x9c\x58\x78\xc3\xa6\xd3\x05\xd3\xe7\x42\x63\x72\xeb\x70\xef\xee\xb1\x72\xe5\x99\xd8\x9c\x2a\xaa\x2d\x88\xba\x6e\xd5\x69\x76\x66\x8a\xf8\x2c\xaa\x7d\x41\x9a\x2c\xcf\xb3\x71\x1d\x21\xcc\x55\xcf\xb4\xb9\x8f\x50\x41\x06\x6b\x2b\x48\x1b\x81\x16\xfa\x48\x80\xea\x73\x8a\x29\x86\x19\xd4\x34\x50\x84\x70\x93\x83\x1e\xae\xb2\x65\x6f\x59\x0a\xfb\x42\x8c\x4f\xbc\x2b\x71\xae\x64\x31\x21\x7c\x20\x57\x62\x8f\x20\x26\xb7\xe6\x7c\xad\x97\x4f\x59\x14\x16\x87\xb2\x89\x2b\x45\x3f\xa2\x95\xb4\xe6\xeb\xc5\x19\x13\xc8\xa3\xda\x70\x16\x91\xff\xd1\x69\xb6\x8e\x7c\x46\xd1\x4b\x8f\xd4\xb2\x23\x17\x3c\xc6\x8d\x87\x56\x6f\xe4\x9b\x36\xe1\xaa\x39\xf3\x57\x4b\x11\x9a\x11\x97\x3e\xcb\x57\xd7\x1f\x12\x73\xd3\x7b\xdc\x42\x8f\xbb\x42\xf5\xb4\x1b\x08\x13\x48\x84\x78\xec\xf2\x94\x34\xc4\xd2\xad\x3a\x39\xb0\x5c\xd2\xfd\x7e\xa6\xf3\x2c\x01\x9c\xb5\x8c\x01\x20\x2f\x50\x0d\xe7\x90\xcb\x20\x95\x9c\x30\x84\x2e\xb7\x18\x82\x56\x79\x64\x49\xc4\x08\x14\x71\x82\x88\x5b\x40\xc9\x93\x6c\xd3\x0a\x15\x60\xbe\x2f\xb9\x25\x08\x74\x1c\x0b\xfa\x19\x42\x31\x12\x61\xc6\x61\x0c\x1c\x02\xad\x2c\x71\x45\xd0\xe3\x01\x69\x93\xc1\xa9\xe2\xad\x57\x85\xad\x71\xe6\x46\x37\x5f\xe0\x03\xcd\xc0\xb7\x9a\xd0\x67\xab\xd0\x90\xd1\x2e\x8a\xf7\x32\xcf\xa7\x2c\x7f\x3e\x26\x13\xc3\x55\x1f\x61\x39\xaf\xbd\x79\x72\x3d\x4b\xd3\x65\x6e\x6b\x90\x58\x6c\x86\xcd\xc9\xc4\x6b\x63\xa2\xbd\xd6\xd1\x74\xea\x33\x5a\xa1\xd0\x73\xea\x4d\x51\x04\x9f\xcf\x0a\xee\x9c\xe7\x4d\xfb\xfd\x74\xba\x31\xf2\xd9\x58\x09\x97\x62\x35\x91\x3c\xc0\x18\x15\x35\x9c\xa2\xa0\x3a\x73\xac\x2c\x8b\xd8\xac\x17\xe4\xde\x2b\xe8\xb3\xb4\x59\xe3\xaa\x85\x19\x8c\x6f\x70\xc2\x46\x17\xe5\x75\x95\x4c\xaa\x82\x2c\xb5\x62\x21\x85\xc5\x40\xab\xd0\x42\x6a\x85\xea\x83\xc1\x3e\x3e\x2c\xd7\xd4\xaf\x04\xe4\x31\xc9\x09\xdb\x98\xdb\x7f\x20\x25\x38\x7f\xfc\xe1\x79\x6f\xac\x03\xaf\x48\x77\xc8\x64\xbc\xbc\x79\xe6\xbe\x06\xc7\x3a\x2f\x09\xde\x64\xe2\x1d\xa5\xe8\x9d\x68\x13\x73\x02\xe7\x5c\xab\x9f\xa1\x54\x81\x33\xae\xa0\x52\x2a\xed\x40\xb9\xbe\x5f\xaa\xed\x97\xea\x70\xde\xb9\xde\xc8\x70\x35\xa2\x27\x13\x94\xf6\x3f\xc5\x21\x04\x5a\x66\x05\xa8\xe1\x94\x4b\x25\x67\x51\x65\xb3\xa2\x1c\x60\x66\xe8\xbf\x89\xb1\x9f\xcc\xcf\x70\x89\x86\x9c\xe6\x85\x5e\x6e\x1a\x42\x2d\xba\xca\x57\xfd\xe0\x0b\xdb\xbe\xd5\x6a\x15\xae\x30\xda\x67\x2b\xf2\xd8\x67\x79\xb3\x79\x5a\xf4\x59\x28\x46\xcd\x3f\xdb\x60\x9f\xb7\xbe\x97\x35\xd9\x03\x48\x8c\xd0\xd9\x3f\x34\xc8\xfb\x68\x56\x19\x97\xdc\x45\x1a\x8a\x7f\x91\xde\x8a\xb3\xf7\xda\x0c\xd1\x64\xd5\x95\x84\x4a\x75\x6a\xe5\x18\x42\xcc\x19\xd9\xac\x22\xc7\xc0\x55\x08\x16\x73\xe0\xe7\x5c\x72\x86\x7d\x41\x16\xcd\x08\xcd\xf7\xf6\x4c\x51\x11\x34\xa1\xfa\x9b\xa9\x3e\xcc\xc2\xb4\x99\xaa\x93\x39\x5f\x05\xff\x93\x06\x54\x80\x2f\xc7\xde\xf7\x1e\xb4\x71\xfc\x77\x7a\xd0\xcc\xa1\x73\x78\xbd\xb0\x31\xcc\x71\xf6\x42\xf2\x0e\x7e\xfe\x0b\x36\x88\xcf\x45\x01\xfd\x7f\xec\x0e\xcf\x3f\x6d\x60\x44\x42\x60\x4d\xf0\xec\x6e\xab\x43\xf4\x06\x9f\x53\x34\xe3\xfc\xa2\x5e\x7c\xba\x55\xaf\xee\x95\x3d\x2b\x45\x9c\xdf\x5d\x07\x2b\xef\xe6\x47\xbd\x8f\xe1\x63\x25\xa2\x0f\xa7\x25\x69\x3b\x1d\x5b\x57\x87\xd7\x49\x3a\x60\x8f\xe3\xda\xe1\xf6\xe5\xaf\x09\x8f\xf5\xc9\xed\xb8\xfa\xf6\xfc\xf6\x17\x75\xbc\xdd\xea\x76\x6f\xef\x6e\xf0\x7e\xfb\xd2\x1c\x7e\xe4\xed\x61\x6f\xb0\xfe\x6e\xee\xb3\x42\xd7\x4d\x8a\xaf\xba\x94\x27\x3a\x49\xd0\x78\x03\xfb\xae\xec\x95\x77\xbc\x12\x0b\x85\x25\x96\xc6\xe1\x7c\x67\xbd\x31\x57\x3b\xc7\x7b\xed\xd3\x51\xb7\x35\xfe\x74\x72\xa6\x7b\xb4\x5d\x89\xcf\xba\xa7\xfc\xf8\x77\x19\xca\x51\x6b\xaf\x75\x79\x37\xae\xab\xea\xe3\xed\xde\xe3\xe3\x35\xc5\xad\xea\xcd\xd0\x86\x57\xed\xdb\x91\x7e\x38\xef\x69\x7d\xa0\xff\x94\x31\xdf\xf0\x72\x32\xf8\xf2\xe1\x64\xb5\x39\x97\xfd\xdb\xf6\x28\x3d\xb8\xfe\x50\x7e\xdc\x3d\x1b\x9c\xbe\x1f\xa6\x97\x37\xbb\x1f\xef\x77\x4b\xb5\xed\xe8\x6d\xb5\xfe\xde\x6c\xef\x5c\xbd\xdf\xbb\x19\xdd\x0d\x3e\x1d\x57\x5b\x49\xba\x73\x9d\xec\xd6\x07\xbb\xbf\x44\x6c\xd8\x2e\x9d\xfd\xd6\xfa\x46\x73\x9e\xb0\xf7\xc3\xab\x9f\xfe\xb1\xb2\x2a\xfe\xf3\xa7\xd7\xf3\x57\x8d\x57\xaf\xe7\x0f\x15\xb3\xc3\x3e\x2b\x50\xed\xb3\xe2\xc5\xea\x5f\x01\x00\x00\xff\xff\x14\xe0\x09\x78\xc2\x12\x00\x00")

func stateHtmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_stateHtmlTmpl,
		"state.html.tmpl",
	)
}

func stateHtmlTmpl() (*asset, error) {
	bytes, err := stateHtmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "state.html.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9d, 0xda, 0x61, 0xe2, 0xb, 0xa2, 0x14, 0x52, 0x65, 0x24, 0xbe, 0x3e, 0x2e, 0x6, 0x6a, 0x48, 0x64, 0x90, 0x7c, 0x8a, 0x7a, 0x58, 0x13, 0x8a, 0x1d, 0x8c, 0xbd, 0xc1, 0xd2, 0x6, 0x94, 0x59}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"state.html.tmpl": stateHtmlTmpl,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"state.html.tmpl": {stateHtmlTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
