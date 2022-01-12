package mdjreader

import (
	"encoding/json"
	"fmt"
)

type ErdEntity struct {
	Entity

	Columns []ErdColumn `json:"columns"`
}

func (e ErdEntity) GetType() EntityType { return e.EntityType }

type ColumnLength string

func (c *ColumnLength) UnmarshalJSON(data []byte) error {
	var stringValue string
	err := json.Unmarshal(data, &stringValue)
	if err == nil {
		*c = ColumnLength(stringValue)
		return nil
	}

	var numberValue int
	err = json.Unmarshal(data, &numberValue)
	if err == nil {
		*c = ColumnLength(fmt.Sprintf("%d", numberValue))
		return nil
	}

	return err
}

type ErdColumn struct {
	Entity

	Type       string       `json:"type"`
	Length     ColumnLength `json:"length"`
	PrimaryKey bool         `json:"primaryKey"`
	ForeignKey bool         `json:"foreignKey"`

	ReferenceTo Ref `json:"referenceTo"`

	Nullable bool `json:"nullable"`
	Unique   bool `json:"unique"`
}

func (e ErdColumn) GetType() EntityType { return e.EntityType }

type ERDRelationshipEnd struct {
	Entity

	Reference   Ref    `json:"reference"`
	Cardinality string `json:"cardinality"`
}

func (e ERDRelationshipEnd) GetType() EntityType { return e.EntityType }

type ERDRelationship struct {
	Entity

	Identifying bool               `json:"identifying"`
	End1        ERDRelationshipEnd `json:"end1"`
	End2        ERDRelationshipEnd `json:"end2"`
}

func (e ERDRelationship) GetType() EntityType { return e.EntityType }
