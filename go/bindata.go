// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../definitions/fields.toml (608B)
// ../definitions/info_object_meta.toml (1.208kB)
// ../definitions/info_storage_meta.toml (107B)
// ../definitions/operations.toml (3.918kB)
// ../definitions/pairs.toml (1.381kB)

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

var _info_object_metaToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x4f\x8f\xd3\x30\x10\xc5\xef\xf9\x14\xa3\x5e\xf6\x14\x4e\xc0\x6d\x0f\x48\x7b\x41\xa2\x5a\x24\x40\x1c\x2a\x84\xdc\x7a\x92\x0c\x8d\xff\x30\x1e\x2f\x1b\x3e\x3d\xb2\xdd\x34\x69\x12\x89\x9e\x5a\xcf\x3c\xbf\x9f\xfd\xc6\x39\x28\xef\xd1\xea\xda\x35\x4d\x40\xf9\x51\xc9\xe0\x11\x1e\x61\x47\x56\xde\xbf\xdd\x55\x1a\xc3\x89\xc9\x0b\x39\x9b\xaa\x1f\xb2\xf8\x39\x6b\x81\x02\x48\x87\x50\x76\x82\x6b\xf2\xaa\xd8\x81\x3b\xfe\xc2\x93\xec\xaa\xea\x70\x72\x56\xd0\x4a\xdd\xa3\x6d\xa5\x5b\x01\x26\x81\xd1\xef\xa6\x6e\x10\x26\xdb\xce\xdb\xa9\xb3\xd5\x47\x51\xed\x56\x9d\xf4\xba\x8a\xaf\xde\xb1\xc0\x23\x08\x47\x5c\xde\xed\xe3\xd3\x78\xa3\x68\xe9\x77\x44\x38\xe3\x00\x64\x21\x88\x63\xd5\xe2\x9b\x64\xda\xab\x20\xb5\x71\x9a\x1a\xc2\x99\xbf\x90\xc1\xdc\x26\x7b\xae\x45\x71\x3b\x4f\x72\x84\x2f\x70\x9f\xc8\x9e\xbf\x66\xe9\x88\x0d\x83\x49\x06\x50\x0c\xa0\x71\x0c\x79\x5d\xb2\xcc\x7c\xe3\xf4\x2c\x84\xe7\xdc\xd8\x3b\x8d\xcb\xab\x55\x07\x13\x7b\x21\xaf\x58\xea\xad\x20\x16\x67\xd9\x8f\xe2\x29\x83\xb4\x02\xd2\x69\xac\xf9\xef\xfc\x10\x57\x6b\x1b\xcd\x11\xb9\x36\xea\x95\x4c\x34\x37\xb3\x5d\x33\x8a\xa8\xb8\x95\x8d\x29\xdd\xab\x19\x38\x8f\xac\x92\xfa\x96\x11\xe8\x2f\x6e\x12\x36\x9e\xe7\x0d\x23\x6d\x04\x8d\x0d\x59\xd4\x70\x1c\xc6\x39\xf2\x96\x3d\xd9\xbb\xec\x8b\xec\xff\xf6\x5e\xcd\x5f\xfa\x5d\xaf\xef\xb3\x92\x2e\x65\x8f\x24\x1d\x72\xf9\x94\x8e\xc1\xf5\x51\xd2\x2c\xa4\x03\x57\x8a\x8c\xbd\x12\x7a\xb9\x14\xc5\xfd\x51\xac\xc3\x08\x7f\x08\xf0\xdd\xf1\xf9\x89\x18\x34\xa6\xef\x30\x80\xb3\x10\x03\xf2\x43\x00\xb2\x3e\x96\x01\x06\xe4\x17\x3a\x61\x6d\x50\x94\x56\xa2\xa6\xa3\x2a\x3b\xac\x6e\xfd\xa5\xa8\xf7\x17\x71\x66\x61\x80\x8b\xc9\x35\x82\xd1\x2c\x01\x12\x71\xc3\xbd\x04\xf1\xf3\xf2\x63\x94\x5f\xb1\xbe\x05\xe4\x25\x28\x99\xad\x29\xff\x02\x00\x00\xff\xff\x25\xb7\x28\xa8\xb8\x04\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x10, 0x57, 0x5, 0xf8, 0x1f, 0x26, 0x7e, 0x5f, 0xe8, 0xc3, 0x88, 0x45, 0x5b, 0x53, 0x36, 0x5a, 0x45, 0xa2, 0x73, 0x78, 0x65, 0xb2, 0x25, 0xd6, 0x2e, 0xd7, 0x6d, 0x5a, 0xef, 0x35, 0xca, 0x74}}
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

var _operationsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x57\xbb\x8e\x2b\x37\x0c\xed\xfd\x15\xc4\x34\x69\x8c\xed\x52\xde\x22\x09\x90\x2e\x48\x90\xf6\xc2\x58\xd0\x1a\x8e\xad\x5c\x8d\x24\x48\x1a\x3b\xce\xd7\x07\xa4\x66\xec\x79\x5b\xbb\xc5\xc2\x96\x48\x9d\xc3\xc3\x87\xe4\xef\xe8\x3d\xd9\x9a\xc2\xe9\x50\x53\x54\x41\xfb\xa4\x9d\x85\x6f\x50\xe9\x08\xe9\x4a\xa0\x6d\xa2\xd0\xa0\x22\x68\x5c\x80\x5f\xc4\x1a\x02\x19\x4c\x54\x83\xf3\x14\x90\x1d\xe2\x47\x75\x38\x3c\xcf\xfa\x70\xfe\x43\x05\xc2\x44\x9f\x79\x69\x71\xf8\x5d\x1b\x03\xd9\x04\xd0\x42\xb6\x02\x77\xfe\x87\x54\xfa\xa8\x0e\x1e\x03\xb6\x11\xbe\xc1\xf7\xca\x63\xba\x56\xa7\x43\xa0\xd8\x99\x94\x97\x5c\x75\x9a\x81\xdd\x83\x7e\x83\xd5\x23\x28\x67\x13\xd9\x04\xc9\xbd\x81\x75\xd5\x11\xaa\xc0\xff\xa2\xfe\x8f\x66\x04\xac\x10\x38\x1b\xa7\x7e\x94\x0a\xf7\x2b\x1b\x6f\xe9\xd6\x9f\x34\x92\x4d\x56\xf6\x55\x03\x4b\x77\x10\xbb\xaf\xe9\x36\x02\xcb\xb2\xed\x60\x89\xc1\x44\xb4\x8c\xb8\xab\xd5\x11\xaa\xb3\xae\xf7\x24\x93\x40\x5d\x7b\xd6\x76\x3f\xd2\x6c\x92\x21\x23\x0b\x2a\x59\xdb\x4c\xd7\x59\xd7\x71\x0e\x63\x74\x4c\x7b\x18\xbc\x3f\x00\x9c\xc9\x38\x7b\xe1\x30\xd3\x55\xc7\x0d\x9c\x69\x58\x67\x2d\x80\xca\x79\x5d\x5a\x09\xbf\x39\xff\x90\xac\x67\xa7\xac\x85\x7f\x6c\x49\xe0\x1f\x1c\xf3\x9f\xc2\x05\x5c\x80\xb6\x33\x49\x7b\x43\x3d\x3d\xd0\x56\x30\x22\x85\x9b\x56\x34\x65\x1b\x83\x62\x5d\xea\x98\x84\x65\x43\x49\x5d\x4b\x69\xfe\xce\xc6\xc2\xb3\x77\x63\xa2\xf2\x71\x9d\xa9\x6c\x41\x13\x5c\x0b\x08\x17\x7d\x23\x0b\x5d\x30\x2c\x26\x57\xe2\x4a\x6d\x1e\xa1\xea\x82\x11\x62\xad\xbb\x95\xd2\xfa\xc3\xdd\x48\x58\x89\x0f\x73\xe2\x0f\xeb\x94\x78\xe7\x55\x30\x5f\x51\x2a\x8b\x8c\x21\x15\xd3\x1a\x1c\xb6\x5a\x7c\x74\xe2\xa8\xcd\x9f\xab\x05\xad\xfe\xb4\x2d\xec\xf3\x19\x62\xee\xf5\x37\x80\x2b\xfd\xbe\x01\xbb\xec\x79\x6d\x6b\xfa\x77\xb5\xeb\xe7\xb1\xbb\xd6\x1b\x2a\x88\xbe\xb7\x1b\x93\x80\xce\x1b\x87\x35\x60\x9e\xe5\x31\x85\x4e\xa5\x57\x83\xac\x70\x64\xaf\xb8\x46\x43\x26\xc3\x1b\x0a\x32\x1d\xe4\x84\xf9\x70\xd8\x94\x65\x2a\x80\xcf\xf3\xc1\xe3\xa5\xb4\x92\xfe\xc2\x0b\xad\x14\x11\xdc\xaf\x5a\x5d\x21\x76\xde\x3b\xae\x32\xb4\xb5\x6b\x73\xbe\xa4\xbe\x04\x61\x54\x59\xfc\xbd\xa0\xa8\xd8\xec\x6b\xd7\xc7\x13\x28\x17\xd4\x36\xce\xa2\x96\xa2\x27\xa5\x1b\xad\xc0\x35\x4d\xa4\xb7\xf5\x94\xad\x56\x0b\x2a\x10\x96\x0f\xb2\xbf\xd9\x58\x34\xea\xdd\x98\xbc\x7c\x5c\xe7\xed\x83\xbb\xe9\x9a\x05\xba\xe3\xe3\xd8\xeb\xae\xd0\x82\xf8\x08\x40\xa1\x60\xc3\x78\xeb\x27\xce\x92\x2f\x9f\xda\xa2\xb6\x09\xb5\x1d\xd5\x78\x4c\x2e\x70\x5e\x7a\xbf\x3c\x40\x86\x43\x5e\x39\x2e\x48\x6f\x7f\x52\x00\x6d\x63\x42\x3b\x9f\x7a\x16\xdb\xf9\xd3\x86\x3d\x68\xc2\x9a\x01\x6b\xe2\x46\x5c\x07\xcc\x7b\x80\x85\x60\xd3\x83\x2f\xb4\xd1\x79\x17\x4a\x80\x70\x43\xa3\xeb\xe5\xc1\x92\xd7\xd5\x39\x5e\x1c\x11\x77\xf6\x4e\xcf\xa3\x31\x4b\xd8\x08\x1d\x3f\x38\xf3\x00\x78\xc1\x4f\xc1\x72\xc3\x0f\xbe\x65\x25\x3a\x4b\x78\xce\x77\x7f\x42\xa9\xfc\xcf\x37\x82\x5c\xc1\x0b\x71\xe4\x6f\x54\xa7\x53\x80\x96\x12\xd6\x98\x70\x1d\x22\x50\xea\x82\x05\xd5\x85\xc0\xad\xfc\xd4\x65\xf0\x1a\x69\x30\xa0\xf0\xd6\x02\x65\x5b\xf3\x1e\x21\x4b\xff\x1a\x14\xd3\x97\xc3\x3c\x04\x8f\x3a\xf4\xab\xb2\x9c\x07\xba\xab\xc7\x05\x30\xec\x39\xbd\x60\x13\x08\x37\x7e\x2b\xf0\x8e\x64\xa9\xd1\x86\x7e\x8a\xd0\xc7\xb8\x4a\xe3\x08\xd5\x7d\x85\xcb\x6c\x88\xf1\xf5\xe8\x3e\x15\x1a\x73\x46\xf5\x63\x85\x9f\x5d\xd0\x2b\xee\xf1\xfe\x75\x73\xd7\xe9\xea\x3a\xbe\x0b\x1f\x80\x5e\x03\x83\x6d\x6b\xb7\xd0\xa7\x3a\x1d\x8c\x53\x68\x9e\x31\xa4\xd0\xd1\x8c\x52\x4c\xb8\x91\x3f\xde\x01\x94\x7c\xf1\xa4\xe7\xe6\xd5\xb6\x71\xe0\x9a\xb5\x07\x7b\x09\x99\x29\xb2\xdc\x25\x7b\xd7\x0c\xa7\xa8\xff\x81\x20\x59\xdb\x49\xd7\xf8\x47\xdd\x3c\x6d\xfd\x6d\xf5\x99\x1e\x5e\xd2\x37\x7c\x6f\xeb\x9f\x8b\x72\xf8\x7f\x00\x00\x00\xff\xff\x96\x9e\x58\x9f\x4e\x0f\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x53, 0x5a, 0xd1, 0xa2, 0xf1, 0x10, 0xa3, 0xd5, 0x19, 0xd7, 0xa5, 0x42, 0x88, 0xf, 0xcc, 0x5b, 0xf1, 0xb9, 0xe0, 0x16, 0x37, 0x11, 0x4c, 0x27, 0x5d, 0x99, 0x77, 0x71, 0xf1, 0x19, 0x9a, 0xe9}}
	return a, nil
}

var _pairsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\x41\x8f\xd3\x30\x10\x85\xef\xf9\x15\xa3\x9c\x57\xbb\x17\xe0\xc6\x09\x09\x09\x89\x95\x56\x02\x4e\x08\x45\xae\x3d\xde\x8c\xea\x78\xcc\x78\xd2\x10\x7e\x3d\xb2\x93\xb6\x01\xb6\xdd\x85\x5b\x65\xcf\xfb\xde\xcc\x1b\x37\x5f\x2d\x47\xc5\xa8\xdd\xe0\x5e\x7f\x6b\x74\x4e\x08\x6f\xa1\xcd\x2a\x14\x1f\xdb\xa6\x39\x5d\x97\x9b\x4b\xf7\x14\x47\xa3\xc4\xb1\x53\xde\x63\xfc\xbb\xca\x61\xb6\x42\xa9\x94\xd4\xe3\x84\x96\xfc\x0c\xda\x23\x6c\xe5\x50\xe5\xe0\x59\x20\x50\xd6\x4a\x17\x74\x18\x95\x4c\x78\x31\xb5\xe7\x09\x94\x21\x09\x1f\xc8\x21\x9c\x09\x15\x9c\x51\x0e\x64\x11\xca\x4f\x65\x31\x8f\x58\x6c\x30\xba\xc4\x14\xf5\x7f\x4d\x8e\xfa\x6b\x16\x3f\x12\xc9\x26\x41\x8a\x7a\x91\x3e\xf5\x18\x6b\x3a\xa3\x04\x10\xd4\x51\x22\x3a\xd8\xcd\x20\x68\x6c\x0f\x13\x85\x00\x0b\xaf\x90\x29\x2a\x8a\xc5\xa4\x2c\x67\xfc\x87\xf3\x61\xa9\x29\x79\x76\x03\xbb\x4d\x03\x1f\x29\xeb\x3d\xbb\x8a\x08\x6c\xeb\x06\xfe\x69\x75\x47\xd1\x95\xa1\xa3\x19\x9e\x78\x34\x57\x98\xab\x16\x8a\xb0\x00\xd8\xfb\x8c\xfa\x5b\x6a\x6f\x5e\x5d\x24\x2c\xd5\xb5\x1f\xed\x29\x83\xe0\xf7\x11\xb3\xde\x9c\xb0\x35\xb9\x8c\xb8\x2f\xcb\xab\x25\xab\x64\x87\x9e\x05\x4b\xbe\xae\xd8\x26\x43\xd2\x25\x0e\x64\xe7\xb3\xf7\x83\x21\x79\xa8\x67\xa5\x64\x18\x83\x52\x32\xa2\x1d\xb9\xa7\xfe\x17\xc4\x9d\x35\x21\xec\x8c\xdd\x6f\xb6\xc2\xef\xd6\xb3\x2b\xbb\x37\x0a\xca\x8e\x01\x0f\x28\x33\x28\x0d\x08\xd3\xd2\x1a\x38\xa3\x06\xbc\xf0\x00\x99\x47\xb1\x35\xa2\x4c\x3f\xf1\xa5\x01\x95\xda\xe7\xe2\xe1\x18\xe6\xc5\x2d\xd0\x40\x8a\x0e\xd6\x4f\x40\x75\x2f\x8e\x13\xcb\xbe\x73\x24\xcf\x6e\xb6\x6d\x9b\xed\x72\x8b\x0e\x1c\xc9\x85\x07\x73\xb3\x4e\xcc\x09\x65\x79\x58\xb5\x9f\x5d\x99\x3d\x18\xa5\x03\x9e\xb6\xe6\x48\x6e\x9b\x63\x1b\x70\xff\xe5\xd3\x67\xc8\x6a\x44\x61\x22\xed\xe1\xae\x1a\x2c\xb0\xe3\x68\xab\x5d\xde\xc8\x8e\x74\x87\xde\x8c\xa1\x84\x0e\x77\x40\x1e\x22\x2b\x64\xd4\xdb\xe6\x3d\x0b\xf8\xfc\x27\x02\x6a\x63\xd1\xf1\x94\x21\x05\xa3\x9e\x65\xb8\xa9\xf3\xed\xb0\x37\x07\x62\x01\xca\x30\x46\x87\x9e\x22\xba\xdb\xa6\x6d\xdb\x5f\x01\x00\x00\xff\xff\xe3\x4d\xb1\x49\x65\x05\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x41, 0xbb, 0xd3, 0xf9, 0x76, 0x45, 0xb9, 0x84, 0x6d, 0xe1, 0x7a, 0x3e, 0xef, 0x80, 0x27, 0x12, 0xbe, 0x28, 0x7d, 0x34, 0x2d, 0x3b, 0xf8, 0x7f, 0xef, 0x67, 0x2, 0x51, 0x7a, 0xe3, 0xb0, 0x6b}}
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
