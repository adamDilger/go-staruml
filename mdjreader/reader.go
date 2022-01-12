package mdjreader

import (
	"encoding/json"
	"fmt"
	"io"
)

type EntityType string
type ViewType string

const (
	Project         EntityType = "Project"
	Column          EntityType = "ERDColumn"
	DataModel       EntityType = "ERDDataModel"
	Diagram         EntityType = "ERDDiagram"
	ErdEntityType   EntityType = "ERDEntity"
	Relationship    EntityType = "ERDRelationship"
	RelationshipEnd EntityType = "ERDRelationshipEnd"
	Hyperlink       EntityType = "Hyperlink"
)

const (
	ColumnView            ViewType = "ERDColumnView"
	ColumnCompartmentView ViewType = "ERDColumnCompartmentView"
	EntityView            ViewType = "ERDEntityView"
	RelationshipView      ViewType = "ERDRelationshipView"
	EdgeLabelView         ViewType = "EdgeLabelView"
	HyperlinkView         ViewType = "HyperlinkView"
	LabelView             ViewType = "LabelView"
)

type Ref struct {
	Ref string `json:"$ref"`
}

type Tag struct {
	Type string `json:"_type"`

	Id     string `json:"_id"`
	Parent Ref    `json:"_parent"`
	Name   string `json:"name"`
	Kine   string `json:"kind"`
	Value  string `json:"value"`
}

func ReadMdj(in io.Reader) (*Entity, error) {
	var val json.RawMessage

	decoder := json.NewDecoder(in)
	if err := decoder.Decode(&val); err != nil {
		return nil, err
	}

	t := struct {
		Type EntityType `json:"_type"`
	}{}
	if err := json.Unmarshal(val, &t); err != nil {
		return nil, err
	}

	if t.Type != Project {
		return nil, fmt.Errorf("no valid types found for %s", t.Type)
	}

	var data Entity
	if err := json.Unmarshal(val, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
