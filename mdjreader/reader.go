package mdjreader

import (
	"encoding/json"
	"fmt"
	"io"
)

type ErdView struct {
}

type EntityType string

const (
	Project   = "Project"
	Entity    = "ErdEntity"
	DataModel = "ErdDataModel"
)

type Tag struct {
	Type string `json:"_type"`
	Id   string `json:"_id"`
	// Parent struct { Ref string `json: "$ref"` } `json: "_parent"`
	Name  string `json:"name"`
	Kine  string `json:"kind"`
	Value string `json:"value"`
}

type ErdEntity struct {
	Id   string     `json:"_id"`
	Name string     `json:"name"`
	Type EntityType `json:"_type"`

	// Parent struct { Ref string `json: "$ref"` } `json: "_parent"`

	OwnedElements []ErdEntity `json:"ownedElements"`
	// OwnedViews    []BaseView   `json:"ownedViews"`

	// Tags []Tag `json:"tags"`
}

func (e *ErdEntity) GetId() string   { return e.Id }
func (e *ErdEntity) GetName() string { return e.Id }
func (e *ErdEntity) GetType() string { return e.Id }

type BaseEntity interface {
	GetId() string
	GetName() string
	GetType() EntityType
}

type BaseView interface {
	Id() string
	Name() string
	Type() string
}

func ReadMdj(in io.Reader) (*ErdEntity, error) {
	var val json.RawMessage

	decoder := json.NewDecoder(in)
	if err := decoder.Decode(&val); err != nil {
		return nil, err
	}

	type tp struct {
		Type  EntityType `json:"_type"`
		Hello string     `json:"hello"`
	}

	var t tp
	if err := json.Unmarshal(val, &t); err != nil {
		return nil, err
	}

	switch t.Type {
	case Project:
		{
			var data ErdEntity
			if err := json.Unmarshal(val, &data); err != nil {
				return nil, err
			}

			return &data, nil
		}
	}

	return nil, fmt.Errorf("no valid types found for %s", t.Type)
}
