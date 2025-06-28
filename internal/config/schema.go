package config

import (
	"bytes"
	"embed"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

type TestFile struct {
	Endpoint string     `json:"endpoint"`
	Tests    testGroups `json:"tests"`
}

type testGroups struct {
	Functional map[string]testCase `json:"functional"`
}

type testCase struct {
	Method string         `json:"method"`
	Body   map[string]any `json:"body"`
	Expect map[string]any `json:"expect"`
}

//go:embed lavapi.schema.json
var schemaFS embed.FS

var compiledSchema *jsonschema.Schema

func Load() (*jsonschema.Schema, error) {
	if compiledSchema != nil {
		return compiledSchema, nil
	}

	schemaData, err := schemaFS.ReadFile("lavapi.schema.json")
	if err != nil {
		return nil, err
	}

	compiler := jsonschema.NewCompiler()

	if err := compiler.AddResource("lavapi.schema.json", bytes.NewReader(schemaData)); err != nil {
		return nil, err
	}

	compiledSchema, err = compiler.Compile("lavapi.schema.json")
	return compiledSchema, err
}
