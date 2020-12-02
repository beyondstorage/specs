package specs

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -nometadata -pkg "specs" -ignore "\\.go$" -prefix "../definitions/" ../definitions

const (
	pairPath      = "pairs.hcl"
	infoPath      = "infos.hcl"
	operationPath = "operations.hcl"
)

var (
	ParsedPairs      Pairs
	ParsedInfos      Infos
	ParsedOperations Operations
)

func ParseService(filePath string) (*Service, error) {
	srv := &Service{}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("read spec: %w", err)
	}

	err = parseHCL(content, filePath, srv)
	if err != nil {
		return nil, fmt.Errorf("parse spec: %w", err)
	}

	err = FormatHCL(filePath, srv)
	if err != nil {
		return nil, fmt.Errorf("format spec: %w", err)
	}
	return srv, nil
}

func parseHCL(src []byte, filename string, in interface{}) (err error) {
	var diags hcl.Diagnostics
	defer func() {
		if err != nil {
			err = fmt.Errorf("parse hcl: %w", err)
		}
	}()

	file, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return diags
	}

	diags = gohcl.DecodeBody(file.Body, nil, in)
	if diags.HasErrors() {
		return diags
	}

	return nil
}

func FormatHCL(filePath string, in interface{}) error {
	hf := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(in, hf.Body())

	formatBody(hf.Body())

	content := hclwrite.Format(hf.Bytes())
	err := ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		return fmt.Errorf("format spec: %w", err)
	}
	return nil
}

func formatBody(body *hclwrite.Body) {
	for k, v := range body.Attributes() {
		if isAttrEmpty(v) {
			body.RemoveAttribute(k)
			continue
		}
	}
	for _, v := range body.Blocks() {
		if isBlockEmpty(v) {
			body.RemoveBlock(v)
			continue
		}

		formatBody(v.Body())
	}
}

func isBlockEmpty(block *hclwrite.Block) bool {
	if len(block.Body().Blocks()) > 0 {
		return false
	}

	attrs := block.Body().Attributes()
	for _, v := range attrs {
		if !isAttrEmpty(v) {
			return false
		}
	}
	return true
}

func isAttrEmpty(attr *hclwrite.Attribute) bool {
	tokens := attr.Expr().BuildTokens(make([]*hclwrite.Token, 0))

	// xxx = null
	if len(tokens) == 1 && tokens[0].Type == hclsyntax.TokenIdent {
		s := string(tokens[0].Bytes)
		switch s {
		case "true":
			return false
		case "null", "false":
			return true
		default:
			log.Fatalf("not handled token: %s", s)
		}
	}
	// xxx = ""
	if len(tokens) == 2 &&
		tokens[0].Type == hclsyntax.TokenOQuote &&
		tokens[1].Type == hclsyntax.TokenCQuote {
		return true
	}
	// xxx = []
	if len(tokens) == 2 &&
		tokens[0].Type == hclsyntax.TokenOBrack &&
		tokens[1].Type == hclsyntax.TokenCBrack {
		return true
	}

	return false
}

func init() {
	err := parseHCL(MustAsset(pairPath), pairPath, &ParsedPairs)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	err = parseHCL(MustAsset(infoPath), infoPath, &ParsedInfos)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}

	err = parseHCL(MustAsset(operationPath), operationPath, &ParsedOperations)
	if err != nil {
		log.Fatalf("parse: %v", err)
	}
}
