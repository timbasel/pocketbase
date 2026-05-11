package jsonschema

import (
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

func ValidateSchema(schema string) error {
	compiler := jsonschema.NewCompiler()
	s, err := compiler.Compile(jsonschema.Draft2020.String())
	if err != nil {
		return err
	}
	schemaJSON, err := jsonschema.UnmarshalJSON(strings.NewReader(schema))
	if err != nil {
		return err
	}
	if err := s.Validate(schemaJSON); err != nil {
		return err
	}
	return nil
}

func Validate(value string, schema string) error {
	c := jsonschema.NewCompiler()
	schemaJSON, err := jsonschema.UnmarshalJSON(strings.NewReader(schema))
	if err != nil {
		return err
	}
	if err := c.AddResource("schema.json", schemaJSON); err != nil {
		return err
	}
	s, err := c.Compile("schema.json")
	if err != nil {
		return err
	}
	valueJSON, err := jsonschema.UnmarshalJSON(strings.NewReader(value))
	if err != nil {
		return err
	}
	if err := s.Validate(valueJSON); err != nil {
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
