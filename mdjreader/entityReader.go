package mdjreader

import "encoding/json"

type BaseEntity interface {
	GetType() EntityType
}

type Entity struct {
	Id         string     `json:"_id"`
	Name       string     `json:"name"`
	EntityType EntityType `json:"_type"`
	Parent     Ref        `json:"_parent"`

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
		item := item

		entityType, err := getEntityType(item)
		if err != nil {
			return err
		}

		var data BaseEntity

		switch entityType {
		case Project:
			{
				var e Entity
				if err := json.Unmarshal(item, &e); err != nil {
					return err
				}

				data = e
			}
		case ErdEntityType:
			{
				var e ErdEntity
				if err := json.Unmarshal(item, &e); err != nil {
					return err
				}

				data = e
			}
		default:
			{
				var e Entity
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

func (e Entity) GetType() EntityType { return e.EntityType }

func getEntityType(item json.RawMessage) (EntityType, error) {
	typeInfo := struct {
		Type EntityType `json:"_type"`
	}{}

	if err := json.Unmarshal(item, &typeInfo); err != nil {
		return "", err
	}

	return typeInfo.Type, nil
}
