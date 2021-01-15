package specs

import "log"

type tomlField struct {
	Type string
}

type tomlFields map[string]tomlField

type tomlInfo struct {
	Type        string `toml:"type"`
	Export      bool   `toml:"export"`
	Description string `toml:"description"`
}

type tomlInfos map[string]tomlInfo

type tomlPair struct {
	Type        string `toml:"type"`
	Description string `toml:"description,optional"`
}

type tomlPairs map[string]tomlPair

type tomlOperation struct {
	Description string   `toml:"description"`
	Params      []string `toml:"params"`
	Pairs       []string `toml:"pairs"`
	Results     []string `toml:"results"`
}
type tomlInterface struct {
	Description string                   `toml:"description"`
	Ops         map[string]tomlOperation `toml:"op"`
}

type tomlInterfaces map[string]tomlInterface

func parsePairs() Pairs {
	var tp tomlPairs

	err := parseTOML(MustAsset(pairPath), &tp)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	var ps Pairs

	for k, v := range tp {
		p := Pair{
			Name:        k,
			Type:        v.Type,
			Description: v.Description,
		}

		ps = append(ps, p)
	}

	return ps
}

func parseInfos() Infos {
	var ti tomlInfos
	var ps Infos

	// Parse object meta
	err := parseTOML(MustAsset(infoObjectMeta), &ti)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	for k, v := range ti {
		p := Info{
			Scope:       "object",
			Category:    "meta",
			Name:        k,
			Type:        v.Type,
			Export:      v.Export,
			Description: v.Description,
		}

		ps = append(ps, p)
	}

	// Parse storage meta
	err = parseTOML(MustAsset(infoStorageMeta), &ti)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	for k, v := range ti {
		p := Info{
			Scope:       "storage",
			Category:    "meta",
			Name:        k,
			Type:        v.Type,
			Export:      v.Export,
			Description: v.Description,
		}

		ps = append(ps, p)
	}

	return ps
}

func parseOperations() Operations {
	var (
		ti tomlInterfaces
		tf tomlFields
		ps Operations
	)

	// Parse operations.
	err := parseTOML(MustAsset(operationPath), &ti)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	// Parse fields
	err = parseTOML(MustAsset(fieldPath), &tf)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	for k, v := range tf {
		ps.Fields = append(ps.Fields, Field{
			Name: k,
			Type: v.Type,
		})
	}

	for k, v := range ti {
		it := Interface{
			Name:        k,
			Description: v.Description,
		}

		for k, v := range v.Ops {
			it.Ops = append(it.Ops, Operation{
				Name:        k,
				Description: v.Description,
				Params:      v.Params,
				Pairs:       v.Pairs,
				Results:     v.Results,
			})
		}

		ps.Interfaces = append(ps.Interfaces, it)
	}

	return ps
}
