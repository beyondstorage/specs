package specs

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -nometadata -pkg "specs" -ignore "\\.go$" -prefix "../definitions/" ../definitions

const (
	pairPath = "pairs.toml"

	infoObjectMeta  = "info_object_meta.toml"
	infoStorageMeta = "info_storage_meta.toml"

	operationPath = "operations.toml"
	fieldPath     = "fields.toml"
)

var (
	ParsedPairs      Pairs
	ParsedInfos      Infos
	ParsedOperations Operations
)

func ParseService(filePath string) (Service, error) {
	srv := Service{}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Service{}, fmt.Errorf("read spec: %w", err)
	}

	err = parseTOML(content, srv)
	if err != nil {
		return Service{}, fmt.Errorf("parse spec: %w", err)
	}

	// Make sure services has been sorted, so that we can format the specs correctly.
	srv.Sort()

	return srv, nil
}

func parseTOML(src []byte, in interface{}) (err error) {
	return toml.Unmarshal(src, in)
}

func init() {
	ParsedPairs = parsePairs()
	ParsedPairs.Sort()

	ParsedInfos = parseInfos()
	ParsedInfos.Sort()

	ParsedOperations = parseOperations()
	ParsedOperations.Sort()
}
