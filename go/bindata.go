// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../definitions/infos.hcl (1.31kB)
// ../definitions/operations.hcl (5.534kB)
// ../definitions/pairs.hcl (2.243kB)

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

var _infosHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\x3f\x73\xdb\x30\x0c\xc5\x77\x7d\x0a\x1c\x97\x4c\xca\xd4\x74\xcb\x54\x77\xe8\x90\x6b\x06\xdf\x75\xec\x21\x22\x6c\xa1\x16\x09\x95\x7c\x6a\xea\xf6\xfa\xdd\x7b\x94\xec\xf3\xbf\x5a\x69\xe2\xc9\x22\xf1\x7e\xef\x91\x80\x54\x69\x5c\x19\x39\x7b\xfa\x26\x0d\x1c\xb9\x20\x60\x47\xae\xb1\x08\x89\xa8\x83\xbf\x73\xf4\xbb\x22\xc2\xb6\x17\xda\xff\xee\xc9\x65\x24\x8d\x6b\x57\x11\x79\xcd\x7d\xc7\xdb\xaf\x91\x83\x94\x9d\x0f\x93\xf4\x61\x71\xe7\xaa\x3f\x2f\xe0\x0b\xf5\x88\x7f\xc4\xbd\xa6\x14\xf0\xfa\xb5\x89\x3e\x2e\x79\x86\xa8\xfe\x94\x77\x82\x92\x9f\xbd\x25\x94\x45\xa4\x41\x2a\xa2\xc6\x42\x90\x88\x52\xf5\x69\x41\x9a\x09\xad\xd0\x10\xf5\xfb\x20\xb4\x91\x2d\x69\xa4\x0c\x4b\xbc\x96\xdb\xeb\x96\xc1\xbc\x9c\x98\xde\x93\xfb\x3c\xd6\x3c\x94\x9d\x83\xef\xce\xf6\x1a\xa7\xe7\x84\x7a\x36\xff\x51\xdc\x47\x4e\x38\x44\x2e\x52\x52\x4f\xb6\x9a\xfe\x4e\xe8\x99\xcc\x3d\xa3\x7d\xe3\x45\x3d\x32\xda\xe2\x2b\x8a\x56\xd2\x68\xcf\x4f\xd9\xba\x01\x25\x07\x5a\xb2\x69\x31\x49\xc7\xd0\x1f\xbb\x45\xd8\x33\x27\x9f\xf7\xd7\x79\x93\xe9\x8b\xa5\xcd\x42\x13\x79\xe9\x25\xfa\x4c\x16\x69\xc8\x92\x6e\x32\x69\xec\x87\xb9\xf0\x59\x7f\x9d\xcd\x99\x46\xbc\x7f\x37\x23\x98\x4c\xeb\xa6\xe3\x9c\x5f\x37\xa1\xe0\xb4\x16\xfc\x5f\x4f\x96\x63\xed\xbe\x27\x79\x1b\x3a\x8d\x1b\x9a\x08\xb4\xb2\x44\xe3\xf3\x8b\xbd\x19\x7a\xcf\x10\x5f\x33\x4e\xa3\x42\x83\xdc\x2e\x35\xc8\x41\xba\x3b\xd8\x41\xdb\x59\xc3\x50\x8b\xf3\x87\xbc\x90\x95\x77\xeb\x7c\x84\x2f\xa6\xe1\x6c\x7c\x2f\x20\xcf\x96\x36\xb5\xd7\xf4\x76\x50\x06\x43\x33\xb4\x19\x3f\x2a\x43\xc4\x6c\x97\xff\xad\x9b\x19\x8e\xbf\x01\x00\x00\xff\xff\x58\x43\x7e\x1f\x1e\x05\x00\x00")

func infosHclBytes() ([]byte, error) {
	return bindataRead(
		_infosHcl,
		"infos.hcl",
	)
}

func infosHcl() (*asset, error) {
	bytes, err := infosHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "infos.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xad, 0x27, 0xa5, 0x36, 0x2f, 0xa6, 0xb4, 0x25, 0xc7, 0xd4, 0xe4, 0xf, 0x6f, 0xf9, 0x58, 0xcb, 0xa9, 0xf0, 0x6c, 0xf0, 0xc5, 0x40, 0xee, 0x9d, 0x58, 0x16, 0x77, 0x3c, 0x8, 0x9b, 0x82, 0x9e}}
	return a, nil
}

var _operationsHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x4d\x8b\x1c\x37\x10\xbd\xcf\xaf\x28\xfa\x12\x30\xc3\x9c\x92\xdc\x7c\x48\x0c\x81\x1c\x4c\x16\xef\x21\x07\xb3\x2c\x35\xea\xea\x19\xc5\x6a\x49\x48\xea\x99\x8c\xc3\xfe\xf7\x50\x52\xcf\xf4\xf7\x97\x59\x1f\x0c\xbb\x7a\xa5\x7a\xf5\xf4\x54\xa5\xde\x9d\xd4\x81\x5c\x81\x82\x20\x43\x6b\x49\xe7\xe4\x32\xf8\x6f\x07\x90\x93\x17\x4e\xda\x20\x8d\x86\x8f\x90\x49\x0f\xe1\x4c\xd0\xc0\x0b\xe3\xe0\xb7\x18\x01\x8e\x14\x06\xca\xc1\x58\x72\xc8\x01\xfe\x90\xed\x76\x00\xc6\x42\x26\x1c\x61\xa0\xd7\xb4\x77\xda\x79\xb0\xf7\x55\x2a\x05\x09\x08\xa8\x21\x61\xc1\x1c\xff\x21\x11\x0e\x59\x8c\xb0\xe8\xb0\xf4\x10\xff\x7d\x84\xaf\x99\xc5\x70\xce\x5e\xe2\x92\x23\x5f\xa9\xe0\x1f\x4b\x26\xfe\xfe\xad\xce\x7f\x75\x72\x55\xfa\x3a\xa9\x30\x3a\x90\x0e\x10\xcc\x5a\x26\x26\xdb\x43\xe6\xf8\x3f\x2f\xbf\xd3\x04\x27\x5d\x73\x7a\x6b\xeb\x7d\x54\x46\x7c\xdb\x22\xf7\xef\x1c\xb0\x46\xed\xb8\xf3\x1a\xb1\x41\xd3\x15\x22\xfa\x5d\xe5\x5e\xcc\x1f\x61\x1d\xb1\x13\x8b\x75\x1a\xef\x21\x3b\xca\x7c\x41\xea\x5a\x10\x53\x1e\xa5\x5e\xa3\x48\x02\x26\x1a\x9e\x85\x8f\x0e\x58\x3e\xfa\xa3\xcc\x7d\x27\xa5\x92\x3e\x2c\xe7\x63\xd4\x3d\xd9\x91\x94\xd1\x27\x96\x21\x9c\xa5\x5f\xca\x39\x5e\xf6\x51\x8e\x59\x4c\x18\x2b\xb7\x38\xec\x93\xb1\xb7\x96\x9b\x8c\xbd\xcd\x6b\x66\x6f\x2c\xd2\x5f\x91\x30\x18\x07\x65\xa5\x82\xb4\x8a\xea\x1a\x40\xea\x98\xc4\x93\xbb\x48\x41\x13\x25\x79\x27\x58\xc8\xdc\x87\xb1\x12\x0a\x0a\xe2\xbc\xa5\x86\x3f\x38\xa0\x29\x22\xc6\xcf\x55\x11\x01\x50\x38\x53\x02\xc2\x49\x5e\x48\x43\xe5\x14\x9f\x06\xbb\x7e\xee\x42\xec\x21\xab\x9c\x1a\x23\x5d\x9a\xcb\x16\xca\x9f\xcd\x85\x1a\xc6\x1c\x3c\x47\x98\xd7\x1b\x6f\xbe\x87\xc6\xe9\xd8\xd0\x85\x4d\xa4\xef\x41\x6b\x3a\xd2\x23\xc3\xea\xae\xf4\x88\x78\x9f\x96\xb4\x8a\xc0\x48\x5b\x5a\xa2\x31\x6c\x4d\x52\xe7\xf4\xef\xda\xe6\x64\x15\xad\x56\xa7\x46\xb7\x49\x41\x65\x95\xc1\x1c\x30\x8d\x2e\x1f\x5c\x25\x42\x73\x25\xe7\x38\x73\xf8\x48\xe3\x5a\x45\x25\x36\xaf\xb8\x41\xbf\x77\x2d\xcb\x35\x2e\x8c\x1d\x6d\x5f\x16\x4f\x5b\x1c\xf9\x84\x27\x1a\x31\x23\x5c\xcf\x52\x9c\xc1\x57\xd6\x1a\x76\x2b\xea\xdc\x94\xe9\xac\x07\x3e\xe5\x8c\xab\x2d\xca\xe0\x77\x9d\x9b\x4b\xd9\x07\xfe\xf4\x96\x84\x2c\xa4\x00\x53\x14\x9e\xd6\x7b\x34\xc1\xb7\x3c\x56\x1c\xe1\xb6\x36\xfc\x85\x03\x1a\x7d\x63\xfc\x5c\x6d\xd6\x99\x8b\xcc\x59\xda\x2b\xde\xf6\xf5\x99\x09\xd4\x10\x23\x63\x82\x1f\x95\x7a\xa2\x49\xd7\x1d\x73\xbc\x26\xce\x5c\xa2\xd4\x01\xa5\x6e\xdd\x37\x1f\x8c\xe3\x53\xaf\x63\x07\x8d\x6e\xb5\x77\xea\x8d\x1c\x48\xed\x03\xea\xc9\xc6\xad\xb1\x9c\x7a\x54\xf2\x16\xd4\x71\x51\x4e\xdc\x20\xe6\x38\x24\x04\xe0\xf6\xfc\xf7\x1c\x27\x9a\x6d\x0d\x27\x0a\x80\x70\x41\x25\xf3\x61\x8e\x68\x8c\xf9\x41\xb5\xa9\x5e\xee\x43\x8b\x7d\x0a\x95\x1a\x32\xf1\x50\xf1\xa7\x4e\x6a\x5a\x5d\x46\xc3\xb4\xa3\xbd\xc9\x07\x0c\xd2\x07\x29\x24\xea\xf5\xb7\xe2\xf9\x1e\x85\xaa\xb1\x8e\x6f\x7e\x39\x3f\x00\x2a\x1d\xee\x64\x7f\xf2\xf0\x08\xf3\x7b\xf0\x95\x38\x03\x7a\x78\x96\xdf\x69\x0f\x9f\x18\x39\x59\x4e\x1d\x35\x5e\x54\xd2\x69\x7d\x41\xbd\x0b\xd1\x14\xb5\xde\x8b\x8f\x07\x64\x7c\x83\xcd\xdb\xe3\x71\xcb\xef\x16\x28\x29\x60\x8e\x01\xe7\x12\x39\x0a\x95\xd3\x20\x2a\xe7\xb8\x6f\x3e\xcc\x70\x8f\x9d\x50\x8a\x97\x37\xd9\xad\xce\x93\x5c\xd7\xf4\xe6\xa5\x67\xe4\x4b\xbd\x24\x5d\xbd\x12\x97\xd2\x34\x36\xf9\xd4\x6d\x30\xb2\xc3\xcd\x11\xce\x7e\xdf\xf2\x7a\x3c\xbe\x42\x2a\x36\x4f\xab\xee\xa9\xb7\xed\x75\x82\x59\x6f\x88\xec\x53\xf2\x57\x81\x4a\x1d\x51\x7c\x7b\x2d\x2a\x2d\x56\xbd\x7f\xd8\x8a\x73\x9c\x79\x1d\x30\xea\xc7\xc3\x8e\x5b\x8b\xd4\x85\x01\x53\x2c\x7e\x9a\x6d\x1b\xbc\xcb\x33\x97\xe5\xaa\x3f\x0a\xa3\x82\x0b\xd2\xf5\xff\x1e\x30\x90\x30\x59\xf0\x55\x28\xf4\x9e\xb1\xf5\x4c\x7f\x0d\x37\x4b\xed\x9f\xcb\xfc\x97\x1f\x12\xf8\x6d\xb7\x2b\x24\xa9\x9c\xbf\x4f\x53\x75\xbc\x33\x17\xf5\x21\xfe\x15\xe1\xcf\xc0\xaf\x23\xe3\xb2\xdd\x5b\x03\xcc\xbb\x48\x1f\x9c\xd4\xa7\x2e\xc2\x77\x21\x5f\x5f\x06\xa0\xfc\x7e\x47\x26\xb7\x21\xe7\xba\x08\x72\xae\xc3\x24\xbd\xa0\x3b\x10\xa9\x43\x0b\x10\x6f\x66\xb7\xaa\xe7\xa4\xe8\x67\x5e\x69\x80\x7a\xb0\xcb\xaf\x3f\xb7\x97\x79\xd8\xcc\x93\x35\xbd\x3c\xa9\x55\xb5\x01\xe9\x16\xcc\xe6\x31\xfd\x33\x48\xbb\x8c\x1c\x42\x34\x4a\x17\x7c\x38\x1c\x9e\x50\x76\x41\xfc\x74\xef\x1d\xc4\x87\x27\x7e\xb4\xb7\x41\xe1\xbc\x50\x9b\xed\xd3\xe2\x2d\x46\x48\xf5\x8e\x4b\x9a\xc3\x17\xc2\x9c\xda\x10\x4f\xa7\x2e\xe8\x99\x4e\x25\x75\x0e\xcd\xf7\xb3\xd5\x90\x91\x84\xf1\xe6\xcc\x2a\xca\xdf\xb3\xf3\xc5\x35\x73\x6e\xd4\x28\x8f\x41\xdc\x09\x19\x50\xac\x47\xc5\x18\xc7\xf8\x1a\xe9\xd6\x7c\x9f\x9e\x0d\x8a\x9f\x9e\xf3\x3c\xaf\x03\x71\xff\xe6\x8e\x13\x37\xf9\x3f\x00\x00\xff\xff\xe9\xe3\x7f\xbe\x9e\x15\x00\x00")

func operationsHclBytes() ([]byte, error) {
	return bindataRead(
		_operationsHcl,
		"operations.hcl",
	)
}

func operationsHcl() (*asset, error) {
	bytes, err := operationsHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "operations.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf, 0xd, 0x5f, 0xda, 0xe5, 0x83, 0x39, 0x3d, 0x5f, 0x5d, 0xb7, 0xfd, 0x30, 0x68, 0x1c, 0x58, 0x3, 0x7b, 0x31, 0x57, 0xef, 0xd0, 0x5b, 0x22, 0xff, 0x46, 0x77, 0x0, 0x3d, 0x39, 0x45, 0xcc}}
	return a, nil
}

var _pairsHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\x0c\x7c\x4a\x8a\xd4\xb9\xb4\xbd\xe5\xd2\x02\x05\x02\x34\x68\x80\x76\x4f\xbb\x0b\x81\x26\x47\xf6\xc0\x14\x87\x3b\x1c\xd9\xf1\x2e\xf6\xbf\x2f\x48\x49\xb6\xe2\xc8\x5e\x23\x39\x04\x96\xe6\xf1\xbd\xa7\xf9\xe0\xcc\xa2\x21\x81\xb9\xe5\xa0\x18\xb4\x6a\xdc\xef\x73\xf8\x36\x03\xd0\x7d\x44\x78\x80\x79\x52\xa1\xb0\x9a\xcf\xbe\x9f\x00\x73\xfc\x0a\xe4\x8b\x8e\x40\xfd\xdf\xc3\x21\xb6\xf8\xab\xc7\xcc\x00\x1c\x26\x2b\x14\x95\x38\x8c\x10\x40\x01\x8c\xf7\x20\xf8\xa5\xc5\xd4\x03\x6b\xd3\x7a\x3d\xa5\xfa\xd3\xd8\xcd\x4a\xb8\x0d\xee\xe6\xf6\xb5\x09\x0a\xad\xc9\xb4\x95\xf2\x06\xc3\xa4\x9f\xc1\xfb\x1b\x1b\x29\xa2\xa5\x7a\x0f\xba\x46\x18\x73\x41\xe1\x82\x9a\x05\x3c\x25\xad\x1c\x09\x0c\xbf\xa3\x60\x4d\x2f\x8b\x91\x09\x41\x87\x41\xc9\xf8\x49\xf1\x5f\x8e\xf1\xc5\xb3\xf0\x96\x1c\xca\x05\x2b\x6b\xde\x81\x32\xc4\x0e\x09\xc7\xc3\xc5\x4d\x42\xd9\x92\xc5\x6c\x26\x29\x8b\x59\x61\xa6\x8a\x46\x12\xca\x31\xfd\x23\xc1\x1c\x39\x5a\xc5\xe0\x22\x53\x98\xae\xda\x10\x7c\x87\xcb\xe1\xe8\xf5\x1e\x8f\x62\x27\x0e\x5f\x22\x09\x4e\xfa\xcb\xbe\xcf\x3b\xda\xad\x31\x94\x3a\xb6\x92\xfb\x49\x5b\x09\xe8\x60\xb9\x07\x41\x63\xd7\xb0\x23\xef\xa1\x27\x7f\xeb\xa6\x3c\x3e\x66\xfe\xc1\xc7\x5a\x35\x56\xd6\x53\x1e\x06\x2e\x62\x69\xba\xba\x19\xd8\xe1\x16\xff\xf6\xb8\x09\x93\x18\x0f\x7d\xd6\xb3\x95\x54\xe5\xe7\x4c\x00\x1d\xc3\x51\x9e\x82\xa2\x58\x8c\xca\xf2\x7a\x0c\x1f\x47\x81\x03\xba\xf4\x65\xc3\xee\x64\x64\xff\xa1\xa4\x4f\xf9\xed\x11\xc8\xb6\xf4\xf7\xfb\xa7\x64\x60\x38\x5b\xe9\x41\x2a\x98\x66\xba\x8c\x57\xc9\xf4\x74\x50\x58\x0e\x9c\x5c\xd7\x09\xa7\x9b\x97\x82\xfe\xf1\xdb\x05\xd2\xee\x68\x9f\x74\x4a\xc3\x8d\x73\x77\x50\x2a\x0d\x92\x10\x37\xb9\xaf\x0b\xa4\x3f\xb2\xc4\x9a\x05\x73\x1b\xb9\x0b\x9d\x93\xc5\x07\x9b\xf9\x7f\x15\xd9\x93\xdd\xbf\x2e\xc8\xb3\x21\x79\xee\xde\x8f\xc0\xa2\x15\xb9\xcb\x97\x6d\x16\xaf\xac\xf1\x7e\x69\xec\xa6\xaa\xdb\x60\x27\x93\x90\x03\x37\x1f\x3f\x2f\xf7\x8a\xb7\x17\x47\xc5\x28\x28\x3b\x06\xdc\xa2\xec\x41\xa9\x41\xd8\x75\x9f\x08\xce\xa8\x81\x5a\xb8\x81\xc4\xad\xd8\x51\xf6\x13\x7d\x3d\x3b\x98\x17\x73\x9f\x0f\xfe\x2c\xf3\x1c\xfc\xbe\x33\xe0\xa9\x21\x45\x07\xfd\x2a\x2a\x86\xae\x4c\x7c\xcf\x58\x59\x6f\x52\xba\x9c\xd1\x36\xa1\x54\x66\x85\x67\xee\xc2\xeb\x36\x46\x9b\x94\x1b\xc8\x54\xbf\x16\xaa\x2e\x6f\xa7\xb3\xbc\x63\xd9\xe4\x0d\xf2\x7e\xa1\xcc\x00\x79\x07\x4d\x0f\xdd\x5d\x5f\x47\x8e\x28\xdd\x70\x96\x94\x2e\x73\x45\xbd\x51\xda\xe2\xa1\xa7\x1d\xc9\x02\x06\x43\xf0\xf4\xe1\xbf\xff\x21\xa9\x11\x85\x1d\xe9\x1a\xee\x8b\x40\x47\x36\x54\xa7\x97\x4b\xa3\x63\x03\xfb\xb0\xaa\x95\xe1\x1e\xa8\x86\xc0\x0a\x09\x75\x01\x9f\x02\xfc\xcd\x02\x75\x3a\x65\x81\xe2\x2d\x38\xde\x25\x88\xde\x68\xcd\xd2\xdc\x95\x4f\x5c\xe2\xda\x6c\x89\x05\x28\x41\x1b\x1c\xd6\x14\xd0\x95\x25\xfb\x23\x00\x00\xff\xff\x38\xf9\x8f\x11\xc3\x08\x00\x00")

func pairsHclBytes() ([]byte, error) {
	return bindataRead(
		_pairsHcl,
		"pairs.hcl",
	)
}

func pairsHcl() (*asset, error) {
	bytes, err := pairsHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pairs.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x87, 0xc9, 0xd5, 0x5d, 0x4a, 0x76, 0x33, 0x88, 0xee, 0xcb, 0xf, 0x89, 0x1b, 0xb2, 0x77, 0x39, 0x7a, 0xbc, 0x16, 0x17, 0x5, 0x3e, 0x40, 0xca, 0xe5, 0x52, 0xd2, 0xeb, 0x1b, 0xe1, 0xb3, 0xc2}}
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
	"infos.hcl":      infosHcl,
	"operations.hcl": operationsHcl,
	"pairs.hcl":      pairsHcl,
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
	"infos.hcl":      {infosHcl, map[string]*bintree{}},
	"operations.hcl": {operationsHcl, map[string]*bintree{}},
	"pairs.hcl":      {pairsHcl, map[string]*bintree{}},
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
