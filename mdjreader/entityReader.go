package mdjreader

import "encoding/json"

type BaseEntity interface {
	GetType() EntityType
}

type ErdEntity struct {
	Id     string     `json:"_id"`
	Name   string     `json:"name"`
	Type   EntityType `json:"_type"`
	Parent Ref        `json:"_parent"`

	OwnedElements ownedElementsList `json:"ownedElements"`
	OwnedViews    ownedViewsList    `json:"ownedViews"`

	// Tags []Tag `json:"tags"`
}

type ownedElementsList []BaseEntity

func (o *ownedElementsList) UnmarshalJSON(data []byte) error {
	var rawJson []json.RawMessage
	if err := json.Unmarshal(data, &rawJson); err != nil {
		return err
	}

	for _, item := range rawJson {
		entityType, err := getEntityType(item)
		if err != nil {
			return err
		}

		var data BaseEntity

		switch entityType {
		case Project:
		case Entity:
			{
				var e ErdEntity
				if err := json.Unmarshal(item, &e); err != nil {
					return err
				}

				data = e
			}
		default:
			{
				var e ErdEntity
				if err := json.Unmarshal(item, &e); err != nil {
					return err
				}

				data = e
			}
		}

		*o = append(*o, data)
	}

	return nil
}

func (e ErdEntity) GetType() EntityType { return e.Type }

func getEntityType(item json.RawMessage) (EntityType, error) {
	typeInfo := struct {
		Type EntityType `json:"_type"`
	}{}

	if err := json.Unmarshal(item, &typeInfo); err != nil {
		return "", err
	}

	return typeInfo.Type, nil
}
