package specs

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -nometadata -pkg "specs" -ignore "\\.go$" -prefix "../definitions/" ../definitions

const (
	featurePath     = "features.toml"
	fieldPath       = "fields.toml"
	infoObjectMeta  = "info_object_meta.toml"
	infoStorageMeta = "info_storage_meta.toml"
	operationPath   = "operations.toml"
	pairPath        = "pairs.toml"
)

var (
	ParsedFeatures   Features
	ParsedPairs      Pairs
	ParsedInfos      Infos
	ParsedOperations Operations
)

func ParseService(filePath string) (Service, error) {
	srv := parseService(filePath)

	// Make sure services has been sorted, so that we can format the specs correctly.
	srv.Sort()

	return srv, nil
}

func init() {
	ParsedFeatures = parseFeatures()
	ParsedFeatures.Sort()

	ParsedPairs = parsePairs()
	ParsedPairs.Sort()

	ParsedInfos = parseInfos()
	ParsedInfos.Sort()

	ParsedOperations = parseOperations()
	ParsedOperations.Sort()
}
