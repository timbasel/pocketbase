package jsonschema

import (
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

type Schema struct {
	jsonschema.Schema
}

func CompileSchema(schema string) (*Schema, error) {
	compiler := jsonschema.NewCompiler()
	schemaJSON, err := jsonschema.UnmarshalJSON(strings.NewReader(schema))
	if err != nil {
		return nil, err
	}
	if err := compiler.AddResource("schema.json", schemaJSON); err != nil {
		return nil, err
	}
	s, err := compiler.Compile("schema.json")
	if err != nil {
		return nil, err
	}
	return &Schema{Schema: *s}, nil
}

func CompileSchemaSchema() (*Schema, error) {
	compiler := jsonschema.NewCompiler()
	s, err := compiler.Compile(jsonschema.Draft2020.String())
	if err != nil {
		return nil, err
	}
	return &Schema{Schema: *s}, nil
}

func (schema *Schema) Validate(value string) error {
	valueJSON, err := jsonschema.UnmarshalJSON(strings.NewReader(value))
	if err != nil {
		return err
	}
	if err := schema.Schema.Validate(valueJSON); err != nil {
		verr := err.(*jsonschema.ValidationError)
		return Error{Causes: verr.Causes}
	}
	return nil
}

type Error struct {
	Causes []*jsonschema.ValidationError
}

func (err Error) Error() string {
	var errors []string
	for _, cause := range err.Causes {
		errors = append(errors, cause.Error())
	}
	return strings.Join(errors, " | ")
}
