package utils

import "github.com/RugiSerl/smallEditor/app/math"

type AnchorType int

const (
	ANCHOR_LEFT AnchorType = iota
	ANCHOR_RIGHT
	ANCHOR_TOP
	ANCHOR_BOTTOM
	ANCHOR_MIDDLE
)

type RelativePosition struct {
	HorizontalAnchor, VerticalAnchor AnchorType
	math.Vec2
}

type RelativeRect struct {
	Position RelativePosition
	Size     math.Vec2
}

func (r RelativePosition) GetAbsolutePosition(referential math.Rect, size math.Vec2) math.Vec2 {
	var x, y float64

	switch r.HorizontalAnchor {
	case ANCHOR_LEFT:
		x = referential.Position.X + r.X
	case ANCHOR_RIGHT:
		x = referential.Position.X + referential.Size.X - r.X - size.X
	case ANCHOR_MIDDLE:
		x = referential.Position.X + referential.Size.X/2 - r.X - size.X/2
	default:
		x = r.X
	}

	switch r.VerticalAnchor {
	case ANCHOR_TOP:
		y = referential.Position.Y + r.Y
	case ANCHOR_BOTTOM:
		y = referential.Position.Y + referential.Size.Y - r.Y - size.Y
	case ANCHOR_MIDDLE:
		y = referential.Position.Y + referential.Size.Y/2 - r.Y - size.Y/2
	default:
		y = r.Y
	}

	return math.NewVec2(x, y)
}

func (r RelativeRect) GetAbsolutePosition(referential math.Rect) math.Vec2 {
	return r.Position.GetAbsolutePosition(referential, r.Size)
}

func (r RelativeRect) GetAbsoluteRect(referential math.Rect) math.Rect {
	return math.Rect{Position: r.Position.GetAbsolutePosition(referential, r.Size), Size: r.Size}
}
