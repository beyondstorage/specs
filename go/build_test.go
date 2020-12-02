package specs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	assert.NotNil(t, ParsedPairs)
	assert.NotNil(t, ParsedOperations)
	assert.NotNil(t, ParsedInfos)
}
