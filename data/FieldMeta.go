package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command: slipscheme -pkg jiradata -overwrite ../schemas/TransitionsMeta.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

type FieldMeta struct {
	AllowedValues   []interface{} `json:"allowedValues,omitempty" yaml:"allowedValues,omitempty"`
	AutoCompleteURL string        `json:"autoCompleteUrl,omitempty" yaml:"autoCompleteUrl,omitempty"`
	HasDefaultValue bool          `json:"hasDefaultValue,omitempty" yaml:"hasDefaultValue,omitempty"`
	Key             string        `json:"key,omitempty" yaml:"key,omitempty"`
	Name            string        `json:"name,omitempty" yaml:"name,omitempty"`
	Operations      []string      `json:"operations,omitempty" yaml:"operations,omitempty"`
	Required        bool          `json:"required,omitempty" yaml:"required,omitempty"`
	Schema          *JsonType     `json:"schema,omitempty" yaml:"schema,omitempty"`
}