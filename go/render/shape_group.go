package render

var _ Shape = (*Group)(nil)
var _ ConcreteShape = (*Group)(nil)

type Group struct {
	children []Shape
	BaseShape
}

func NewGroup() *Group {
	baseShape := *DefaultBaseShape()
	res := &Group{
		children:  make([]Shape, 0),
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res

	return res
}

func (g *Group) Add(s Shape) {
	s.setParent(g)
	g.children = append(g.children, s)
}

func (g *Group) Children() []Shape {
	return g.children
}

func (s *Group) concreteNormal(p Tuple) Tuple {
	return NewPoint(0, 0, 0)
}

func (s *Group) concreteIntersect(tr *Ray) []*Intersection {
	return nil
}
