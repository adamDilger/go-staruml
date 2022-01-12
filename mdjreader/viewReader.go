package mdjreader

import "encoding/json"

type BaseView interface {
	GetType() ViewType
}

type ErdView struct {
	Id     string   `json:"_id"`
	Name   string   `json:"name"`
	Type   ViewType `json:"_type"`
	Parent Ref      `json:"_parent"`
	Model  Ref      `json:"model"`

	SubViews ownedViewsList `json:"subViews"`
}

func (e ErdView) GetType() ViewType { return e.Type }

type ownedViewsList []BaseView

func (o *ownedViewsList) UnmarshalJSON(data []byte) error {
	var rawJson []json.RawMessage
	if err := json.Unmarshal(data, &rawJson); err != nil {
		return err
	}

	for _, item := range rawJson {
		item := item

		entityType, err := getViewType(item)
		if err != nil {
			return err
		}

		var data BaseView

		switch entityType {
		default:
			{
				var e ErdView
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

func getViewType(item json.RawMessage) (ViewType, error) {
	typeInfo := struct {
		Type ViewType `json:"_type"`
	}{}

	if err := json.Unmarshal(item, &typeInfo); err != nil {
		return "", err
	}

	return typeInfo.Type, nil
}
