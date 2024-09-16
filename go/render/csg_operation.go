package render

var _ CSGOperation = (*CSGUnion)(nil)
var _ CSGOperation = (*CSGIntersection)(nil)
var _ CSGOperation = (*CSGDifference)(nil)

var CSG_UNION CSGOperation
var CSG_INTERSECTION CSGOperation
var CSG_DIFFERENCE CSGOperation

func init() {
	CSG_UNION = &CSGUnion{}
	CSG_INTERSECTION = &CSGIntersection{}
	CSG_DIFFERENCE = &CSGDifference{}
}

type CSGOperation interface {
	intersectionAllowed(lHit bool, inL, inR bool) bool
}

type CSGUnion struct{}

type CSGIntersection struct{}

type CSGDifference struct{}

func (op *CSGUnion) intersectionAllowed(lHit bool, inL, inR bool) bool {
	return (lHit && !inR) || (!lHit && !inL)
}

func (op *CSGIntersection) intersectionAllowed(lHit bool, inL, inR bool) bool {
	return (lHit && inR) || (!lHit && inL)
}

func (op *CSGDifference) intersectionAllowed(lHit bool, inL, inR bool) bool {
	return (lHit && !inR) || (!lHit && inL)
}
