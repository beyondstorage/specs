package main

import (
	"log"

	. "github.com/aos-dev/specs/go"
)

const (
	pairPath      = "../definitions/pairs.hcl"
	infoPath      = "../definitions/infos.hcl"
	operationPath = "../definitions/operations.hcl"
)

func main() {
	err := FormatHCL(pairPath, ParsedPairs)
	if err != nil {
		log.Fatalf("format: %v", err)
	}
	err = FormatHCL(infoPath, ParsedInfos)
	if err != nil {
		log.Fatalf("format: %v", err)
	}
	err = FormatHCL(operationPath, ParsedOperations)
	if err != nil {
		log.Fatalf("format: %v", err)
	}
}
