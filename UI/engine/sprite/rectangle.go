package sprite

import "image"

type Rectangle struct {
	pointHG image.Point
	pointBD image.Point
}

func NewRectangleFromBounds(postition image.Point, bounds image.Rectangle) *Rectangle {
	pointBD := image.Point{postition.X + bounds.Dx(), postition.Y + bounds.Dy()}
	return &Rectangle{postition, pointBD}
}

func NewRectangle(hg, bd image.Point) *Rectangle {
	return &Rectangle{hg, bd}
}

func (r Rectangle) PointHG() image.Point {
	return r.pointHG
}

func (r *Rectangle) SetPointHG(p image.Point) {
	r.pointHG = p
	nouv_p := image.Point{p.X, p.Y}
	nouv_p = nouv_p.Add(r.Transformation())
	r.pointBD = nouv_p
}

func (r Rectangle) PointBD() image.Point {
	return r.pointBD
}

func (r *Rectangle) SetPointBD(p image.Point) {
	r.pointBD = p
	nouv_p := image.Point{p.X, p.Y}
	nouv_p = nouv_p.Add(r.Transformation())
	nouv_p = nouv_p.Mul(-1)
	r.pointBD = nouv_p

}

func (r Rectangle) Transformation() image.Point {
	x := r.pointBD.X - r.pointHG.X
	y := r.pointHG.Y - r.pointBD.Y

	return image.Point{x, y}
}
