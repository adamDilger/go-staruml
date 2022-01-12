package mdjreader

type LabelView {
	underline = false;
	text = "";
	horizontalAlignment = Graphics.AL_CENTER;
	verticalAlignment = Graphics.AL_MIDDLE;
	direction = DK_HORZ;
	wordWrap = false;
}

type ERDColumnView struct {
	ErdView

	Identifying bool               `json:"identifying"`
	End1        ERDRelationshipEnd `json:"end1"`
	End2        ERDRelationshipEnd `json:"end2"`
}

func (e ERDRelationship) GetType() EntityType { return e.EntityType }
