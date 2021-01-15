// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../definitions/fields.toml (608B)
// ../definitions/info_object_meta.toml (643B)
// ../definitions/info_storage_meta.toml (107B)
// ../definitions/operations.toml (3.795kB)
// ../definitions/pairs.toml (1.502kB)

package specs

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

var _fieldsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xb1\x4e\xc3\x40\x10\x44\x7b\x7f\x45\x94\x6f\x40\x74\x34\x74\x14\x08\x44\x0a\x8a\x08\xa1\xb5\xbd\x09\x0b\x89\xcf\x9a\x1d\x04\xe1\xeb\xd1\x9d\x13\xdd\x39\xbe\xd2\x6f\x9e\x77\xf6\x76\xdb\xda\x5b\xc3\xd3\xa8\xab\xbb\xd5\xfa\xfe\x10\xba\xaf\x07\x2a\x84\x01\xeb\xa6\xd9\xb6\xd6\xe7\xd4\x09\x1b\xf6\x67\xec\xd7\xfc\x5d\x00\x39\xc5\xb4\x77\xd6\x7e\x52\x20\x63\x05\xa6\x06\x1b\x7a\xfd\xcd\xdc\x06\x46\x7a\x54\x4a\x86\x1b\x06\xc8\x5e\x1f\x95\x12\xc3\x61\xa6\xdf\xde\x24\x26\x47\xad\x95\x86\x0c\x9f\xda\x4f\xed\xd2\xf4\xb0\xdb\xb9\xb2\x32\x25\xd8\xb5\x5e\x1e\x63\x14\x43\xf1\xec\xe7\xf8\x39\x71\x70\xc6\xc1\x33\xe7\x47\x6d\xa7\xd1\xe6\x72\xd9\x51\x9c\xe8\x45\xa5\xd7\x04\xdd\xfe\xb4\xb2\xad\xa3\xab\x8d\x77\x0a\xcd\x69\xdd\xe2\x80\x9b\x4b\x32\x69\xb6\x10\x50\xae\xe2\x0c\xd0\xa5\x12\xa3\x6f\x1c\x6a\xcd\x3f\x19\xbe\xc2\x98\xdc\xff\x00\x00\x00\xff\xff\xec\xb9\x4c\x6e\x60\x02\x00\x00")

func fieldsTomlBytes() ([]byte, error) {
	return bindataRead(
		_fieldsToml,
		"fields.toml",
	)
}

func fieldsToml() (*asset, error) {
	bytes, err := fieldsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fields.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x89, 0x49, 0x23, 0x77, 0x93, 0x92, 0xdc, 0x60, 0xa, 0xe6, 0x7d, 0x32, 0x65, 0xf9, 0xd6, 0x3e, 0x59, 0x1b, 0xac, 0xbc, 0xc4, 0x1c, 0x8e, 0x22, 0x8b, 0xd5, 0x2, 0x19, 0x56, 0x3a, 0x34, 0x45}}
	return a, nil
}

var _info_object_metaToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x31\x4f\xc3\x40\x0c\x85\xf7\xfc\x0a\xab\x4b\xa7\x30\x01\x5b\xb7\x2e\x48\x54\x30\x20\x31\x54\x0c\xd7\x9c\x9b\x98\x24\xe7\xc3\xf6\x01\xe5\xd7\xa3\x4b\x28\x44\x69\x16\xb6\xf8\x3d\xeb\xf3\xb3\x73\xfb\x8a\x83\x61\xb0\xb2\xf7\x37\x2f\x85\x9d\x22\xc2\x06\x56\x6a\x42\xa1\x5e\x15\xc5\xaf\x9d\x9d\x25\x1f\xcd\xd5\x4b\x3a\xf9\x4b\x15\x3f\x23\x8b\xc1\x06\x4c\x12\x16\x1e\xb5\x12\x8a\x46\x1c\x72\xd3\xdd\x16\x48\xc1\x1a\x84\x14\xe8\x2d\x21\xb4\x78\x02\x0a\xa0\xc6\xe2\x6a\xbc\xca\xd0\x8e\x42\x5b\x9a\x93\x1a\xed\x92\x3e\xe3\xdd\x53\x68\x9f\x86\xd6\x33\x57\x4f\x7d\x06\xc0\x08\x80\x23\x0b\x0c\x35\x1f\x5e\xb1\xb2\x61\x40\xcf\x7e\xb2\xe5\xc3\x60\xec\xd8\xe3\x3c\x7b\xb1\xef\x53\x67\x14\x9d\x58\xb9\xb4\xe9\x2c\xcb\xee\xdc\xfc\xb7\x64\xae\x80\x3c\xf0\x71\xfc\x9c\x84\x88\xce\x9a\x7f\x1e\xef\xd1\x59\x93\xc9\x48\xd6\xa0\x0c\x03\xdc\x41\xb9\x4b\x96\x27\x59\x03\x3c\x8a\x82\x9d\x33\x7a\xff\x11\x8d\x3f\x9c\x78\x3d\x9f\x78\xad\xf0\xcc\xd2\x6e\x49\xc0\x63\xc4\xe0\x15\x38\x40\x52\x94\xb5\x02\x85\x98\xc6\x78\x4a\x5f\x93\x1b\x51\xb0\xdb\xeb\x41\x1e\x21\x65\xd5\x39\xd5\xa5\x17\x91\xa2\x77\x86\xbe\x74\x93\x7f\x67\xd4\xe3\xea\x3b\x00\x00\xff\xff\x9f\xc0\x20\xb3\x83\x02\x00\x00")

func info_object_metaTomlBytes() ([]byte, error) {
	return bindataRead(
		_info_object_metaToml,
		"info_object_meta.toml",
	)
}

func info_object_metaToml() (*asset, error) {
	bytes, err := info_object_metaTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "info_object_meta.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2f, 0x19, 0xc6, 0xc0, 0x6f, 0x93, 0x27, 0xbf, 0x14, 0x15, 0xb, 0x11, 0xb3, 0xa4, 0xa4, 0x6f, 0x18, 0x2b, 0xd7, 0x6, 0x68, 0xf4, 0xc2, 0x5e, 0xa2, 0xf7, 0xa5, 0x6b, 0x88, 0x53, 0xc5, 0xdb}}
	return a, nil
}

var _info_storage_metaToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xce\xc9\x4f\x4e\x2c\xc9\xcc\xcf\x8b\xe5\x2a\xa9\x2c\x48\x55\xb0\x55\x50\x2a\x2e\x29\xca\xcc\x4b\x57\xe2\xe2\x8a\xce\x4b\xcc\x4d\xc5\x14\x4f\xad\x28\xc8\x2f\x2a\x51\xb0\x55\x28\x29\x2a\x4d\xe5\xe2\x8a\x2e\xcf\x2f\xca\xd6\x4d\xc9\x2c\x22\xa4\x12\x10\x00\x00\xff\xff\x5f\xe0\xd4\x5a\x6b\x00\x00\x00")

func info_storage_metaTomlBytes() ([]byte, error) {
	return bindataRead(
		_info_storage_metaToml,
		"info_storage_meta.toml",
	)
}

func info_storage_metaToml() (*asset, error) {
	bytes, err := info_storage_metaTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "info_storage_meta.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x23, 0xd1, 0x40, 0x2c, 0xba, 0x2b, 0x76, 0x2b, 0x8d, 0x58, 0xb2, 0xa, 0x71, 0xa7, 0xb1, 0xb5, 0x96, 0x12, 0x28, 0xb9, 0x77, 0x6d, 0x13, 0xc5, 0x3a, 0x98, 0x62, 0x3e, 0x73, 0xf2, 0x73, 0x74}}
	return a, nil
}

var _operationsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x57\x3d\xaf\xeb\x36\x0c\xdd\xf3\x2b\x08\x2f\x5d\x82\xbb\x75\x7c\x43\x5b\xa0\x5b\xd1\xa2\xeb\xc3\x45\xc0\xc8\x74\xa2\x3e\x59\x12\x24\x39\x69\xfa\xeb\x0b\x52\x4e\xe2\xd8\xb2\xa3\x7b\x87\x0b\x47\x22\x79\x8e\x0e\x3f\x2c\x7f\x47\xef\xc9\xb6\x14\x3e\x77\x2d\x45\x15\xb4\x4f\xda\x59\xf8\x06\x8d\x8e\x90\xce\x04\xda\x26\x0a\x1d\x2a\x82\xce\x05\xf8\x45\xac\x21\x90\xc1\x44\x2d\x38\x4f\x01\xd9\x21\x7e\x34\xbb\xdd\x23\xd6\x87\xf3\x1f\x2a\x10\x26\x3a\xe4\xa5\x45\xf0\xab\x36\x06\xb2\x09\xa0\x85\x6c\x05\xee\xf8\x0f\xa9\xf4\xd1\xec\x3c\x06\xec\x23\x7c\x83\xef\x8d\xc7\x74\x6e\x3e\x77\x81\xe2\x60\x52\x5e\x72\xcd\xe7\x0c\xec\x1a\xf4\x1b\xac\x11\x41\x39\x9b\xc8\x26\x48\xee\x0d\xac\x6b\xf6\xd0\x04\xfe\x17\xf5\x7f\x34\x23\x60\x85\xc0\xd1\x38\xf5\xa3\x56\xb8\x5f\xd9\x78\x4d\xb7\x31\xd2\x44\x36\x59\xd9\x56\x0d\x2c\x5d\x41\xec\xbe\xa6\xdb\x04\x2c\xcb\xb6\x81\x25\x06\x2f\xa2\x65\xc4\x4d\xad\xf6\xd0\x1c\x75\xbb\x25\x99\x1c\xd4\xf5\x47\x6d\xb7\x4f\x9a\x4d\x32\x64\x64\x41\x25\x6b\xab\xe9\x3a\xea\x36\xce\x61\x8c\x8e\x69\x0b\x83\xf7\xef\x00\x47\x32\xce\x9e\xf8\x98\xe9\xac\xe3\x0a\xce\xeb\xb1\x8e\x5a\x00\x95\xf3\xba\xb6\x12\x7e\x73\xfe\x26\x59\xcf\x4e\x59\x0b\x7f\x5b\x93\xc0\xdf\xf8\xcc\x7f\x0a\x17\x70\x01\xfa\xc1\x24\xed\x0d\x8d\xf4\x40\x5b\xc1\x88\x14\x2e\x5a\xd1\x2b\xdb\x18\x14\xeb\xd2\xc6\x24\x2c\x3b\x4a\xea\x5c\x4b\xf3\x77\x36\x16\x9e\xa3\x1b\x13\x95\xc7\x32\x53\xd9\x82\x2e\xb8\x1e\x10\x4e\xfa\x42\x16\x86\x60\x58\x4c\xae\xc4\x42\x6d\xee\xa1\x19\x82\x11\x62\xbd\xbb\xd4\xd2\xfa\xc3\x5d\x48\x58\x89\x0f\x73\xe2\x87\x32\x25\xde\x79\x16\xcc\x57\x94\xca\x22\x63\x48\xd5\xb4\xee\x0e\x6b\x2d\x3e\x89\x38\x69\xf3\xc7\x6a\x45\xab\x3f\x6c\x2b\xfb\x7c\x86\x98\x7b\xfd\x0d\x60\xa1\xdf\x57\x60\x97\x3d\xaf\x6d\x4b\xff\x16\xbb\x7e\x7e\x76\xd7\x7b\x43\x15\xa7\x1f\xed\xa6\x24\x60\xf0\xc6\x61\x0b\x98\x67\x79\x4c\x61\x50\xe9\xd9\x20\x05\x8e\xec\x15\x4b\x34\x64\x32\xbc\xa1\x20\xd3\x41\x22\xcc\x87\xc3\xaa\x2c\xaf\x02\xf8\x3c\x1f\x3c\x9e\x6a\x2b\xe9\x2f\x3c\x51\xa1\x88\xe0\x7a\xd6\xea\x0c\x71\xf0\xde\x71\x95\xa1\x6d\x5d\x9f\xf3\x25\xf5\x25\x08\x93\xca\xe2\xdf\x15\x45\xc5\x66\x5f\x7b\x7d\x3c\x80\x72\x41\xad\xe3\x2c\x6a\x29\x7a\x52\xba\xd3\x0a\x5c\xd7\x45\x7a\x5b\x4f\xd9\xaa\x58\x50\x81\xb0\x7e\x90\xfd\xcd\xc6\xa2\xd1\xe8\xc6\xe4\xe5\xb1\xcc\xdb\x07\x77\xd1\x2d\x0b\x74\xc5\xdb\x7e\xd4\x5d\xa1\x05\xf1\x11\x80\x4a\xc1\xee\xe3\x6d\x9c\x38\x4b\xbe\x1c\xb5\x47\x6d\x13\x6a\x3b\xa9\xf1\x98\x5c\xe0\xbc\x8c\x7e\x79\x80\xdc\x83\x3c\x73\x5c\x91\xde\x31\x52\x00\x6d\x63\x42\x3b\x9f\x7a\x16\xfb\xf9\xd5\x86\x3d\xe8\x85\x35\x03\xb6\xc4\x8d\x58\x06\xcc\x7b\x80\x95\x60\xaf\x81\x4f\xb4\xd2\x79\x27\x4a\x80\x70\x41\xa3\xdb\x65\x60\xc9\x6b\x71\x8e\x57\x9f\x88\x3b\x7b\xa3\xe7\xd1\x98\x25\x6c\x84\x81\x2f\x9c\x79\x00\x3c\xe1\x5f\xc1\x72\xc3\xdf\x7d\xeb\x4a\x74\x96\xf0\x9c\xef\x31\x42\xad\xfc\x8f\x3b\x82\xbc\x82\x17\xe2\xc8\xdf\xa4\x4e\x5f\x01\x7a\x4a\xd8\x62\xc2\x32\x44\xa0\x34\x04\x0b\x6a\x08\x81\x5b\xf9\xa1\xcb\xdd\x6b\xa2\xc1\x1d\x85\xb7\x16\x28\xeb\x9a\x8f\x08\x59\xfa\xe7\xa0\x78\xbd\x39\xcc\x8f\xe0\x51\x87\x71\x55\x96\xf3\x40\x77\xed\xb4\x00\xee\x7b\x4e\x2f\xd8\x04\xc2\x95\x6f\x05\xde\x91\x2c\x75\xda\xd0\x4f\x11\xc6\x33\x16\x69\xec\xa1\xb9\x16\xb8\xcc\x86\x18\xcf\x36\xc2\xf6\xa0\xd0\x98\x23\xaa\x1f\x87\x6e\xb0\xaa\x40\xd3\x2e\x58\xc6\x84\x2b\x9a\xf1\x0e\xa0\x68\xc4\xd3\x95\x1b\x46\xdb\xce\x81\xeb\x4a\x97\xe4\xb9\x78\x0b\x81\x16\xc8\x32\xbf\xb7\x46\x3b\xcb\x32\x5e\xca\x45\xa9\x0d\x89\xa6\x1f\x52\x0b\xa9\x32\xe6\x41\x19\x8c\x91\xed\xc6\x57\xc6\x21\xdd\x3c\x4d\x7f\xf7\xed\xcf\x72\xcf\x70\x0f\x19\xcb\x0a\xfe\x1f\x00\x00\xff\xff\xea\x39\x74\xe4\xd3\x0e\x00\x00")

func operationsTomlBytes() ([]byte, error) {
	return bindataRead(
		_operationsToml,
		"operations.toml",
	)
}

func operationsToml() (*asset, error) {
	bytes, err := operationsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "operations.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf9, 0xf9, 0x90, 0x5d, 0x54, 0x43, 0x69, 0x7, 0xc9, 0x23, 0x26, 0xb7, 0x3d, 0x8, 0x8b, 0x80, 0xf9, 0xac, 0x21, 0x8f, 0xe7, 0x3d, 0xa8, 0x50, 0x6e, 0x81, 0x36, 0x12, 0xdb, 0xc7, 0xb, 0x66}}
	return a, nil
}

var _pairsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\x41\x8f\xd3\x30\x10\x85\xef\xf9\x15\xa3\x9c\x97\xee\x05\xb8\x71\x42\x42\x42\x62\xa5\x95\x80\x13\x42\x91\x6b\x4f\x9a\x51\x1d\x8f\x19\x4f\x1a\xc2\xaf\x47\x76\xd2\x36\x40\xdb\xdd\xe5\x16\xd9\xf3\xbe\x37\x7e\xe3\xf8\x9b\xe5\xa0\x18\xb4\xe9\xdd\x9b\xef\x95\x4e\x11\xe1\x1d\xd4\x49\x85\xc2\xae\xae\xaa\xd3\x76\xde\xb9\xb6\x4f\x61\x30\x4a\x1c\x1a\xe5\x3d\x86\x7f\xab\x1c\x26\x2b\x14\x73\x49\x59\x8e\x68\xa9\x9d\x40\x3b\x84\xb5\x1c\x8a\x1c\x5a\x16\xf0\x94\xb4\xd0\x05\x1d\x06\x25\xe3\x9f\x4d\xed\x78\x04\x65\x88\xc2\x07\x72\x08\x67\x42\x01\x27\x94\x03\x59\x84\xfc\xa9\x2c\x66\x87\xd9\x06\x83\x8b\x4c\x41\xff\xd7\xe4\xa8\xbf\x65\xf1\x33\x92\xac\x12\xa4\xa0\x57\xe9\x63\x87\xa1\xa4\x33\x88\x07\x41\x1d\x24\xa0\x83\xed\x04\x82\xc6\x76\x30\x92\xf7\x30\xf3\x32\x99\x82\xa2\x58\x8c\xca\x72\xc6\x7f\x3c\x2f\xe6\x9a\x9c\x67\xd3\xb3\x5b\x35\xf0\x89\x92\x3e\xb0\x2b\x08\xcf\xb6\x4c\xe0\x45\xa3\x3b\x8a\x6e\x1c\x3a\x98\xfe\xc2\xa5\xb9\xc1\x5c\xb4\x90\x85\x19\xc0\x6d\x9b\x50\xff\x48\xed\xed\xeb\xab\x84\xb9\xba\xf4\xa3\x1d\x25\x10\xfc\x31\x60\xd2\xbb\x13\xb6\x24\x97\x10\xf7\x79\x78\xa5\x64\x91\x6c\xb1\x65\xc1\x9c\xaf\xcb\xb6\xd1\x90\x34\x91\x3d\xd9\xe9\xec\xfd\x68\x48\x1e\xcb\x5a\x2e\xe9\x07\xaf\x14\x8d\x68\x43\xee\xd2\x7f\x41\xdc\x58\xe3\xfd\xd6\xd8\xfd\x6a\x2a\xfc\x7e\x59\xbb\x31\x7b\xa3\xa0\xec\x18\xf0\x80\x32\x81\x52\x8f\x30\xce\xad\x81\x33\x6a\xa0\x15\xee\x21\xf1\x20\xb6\x44\x94\xe8\x17\x3e\x37\xa0\x5c\xfb\x54\x3c\x1c\xfc\x34\xbb\x79\xea\x49\xd1\xc1\xf2\x04\x14\xf7\xe2\x38\x97\x37\xd6\x9b\x94\x2e\x9d\x7d\x48\x28\x8d\xd9\xe1\x0b\x7e\xa8\xf2\x16\x0c\x49\xb9\x87\xac\x7e\x55\xd4\xf3\x51\xad\x27\x0c\xe5\x35\x18\x59\xf6\x8d\x23\x79\x92\x5a\xd7\xd5\x1a\x9c\x75\xe0\x48\xae\xdc\xd4\xbb\x25\x6a\x8e\x28\xf3\x8d\x2e\x41\x6c\x73\xe8\xde\x28\x1d\xf0\x74\x5d\x1c\xc9\xa6\x3a\xb6\x01\x0f\x5f\x3f\x7f\x81\xa4\x46\x14\x46\xd2\x0e\xee\x8b\xc1\x0c\x3b\x66\xba\xd8\xa5\x95\xec\x48\x77\xd8\x9a\xc1\xe7\x69\xc3\x3d\x50\x0b\x81\x15\x12\xea\xa6\xfa\xc0\x02\x6d\xfa\x1b\x01\xa5\xb1\xe0\x78\x4c\x10\xbd\xd1\x96\xa5\xbf\x2b\xe7\xdb\x62\x67\x0e\xc4\x02\x94\x60\x08\x0e\x5b\x0a\xe8\x36\x55\x5d\xd7\xbf\x03\x00\x00\xff\xff\xd9\x63\x8c\x54\xde\x05\x00\x00")

func pairsTomlBytes() ([]byte, error) {
	return bindataRead(
		_pairsToml,
		"pairs.toml",
	)
}

func pairsToml() (*asset, error) {
	bytes, err := pairsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pairs.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdd, 0x3c, 0x36, 0x9, 0x41, 0x12, 0x95, 0x5a, 0x8d, 0x11, 0xf4, 0xf, 0xd1, 0xfa, 0xcb, 0xe4, 0xc1, 0x9e, 0x62, 0x3d, 0xd6, 0xd3, 0x2b, 0x3e, 0x37, 0x29, 0xd8, 0x95, 0x2, 0xef, 0x45, 0x49}}
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
	"fields.toml":            fieldsToml,
	"info_object_meta.toml":  info_object_metaToml,
	"info_storage_meta.toml": info_storage_metaToml,
	"operations.toml":        operationsToml,
	"pairs.toml":             pairsToml,
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
	"fields.toml":            {fieldsToml, map[string]*bintree{}},
	"info_object_meta.toml":  {info_object_metaToml, map[string]*bintree{}},
	"info_storage_meta.toml": {info_storage_metaToml, map[string]*bintree{}},
	"operations.toml":        {operationsToml, map[string]*bintree{}},
	"pairs.toml":             {pairsToml, map[string]*bintree{}},
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
