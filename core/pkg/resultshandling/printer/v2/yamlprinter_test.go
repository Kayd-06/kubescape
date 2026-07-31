package printer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewYamlPrinter(t *testing.T) {
	yp := NewYamlPrinter()
	assert.NotNil(t, yp)
	assert.Empty(t, yp)
}

func TestSetWriter_Yaml(t *testing.T) {
	yp := NewYamlPrinter()
	assert.NotNil(t, yp)
	
	// Test without outputFile (should use stdout)
	yp.SetWriter(nil, "")
	assert.NotNil(t, yp.writer)
	yp.CloseWriter()
}
