package controllers

func bbox2polygon(Bbox []float64) [][][2]float64 {
	ptA := [2]float64{Bbox[0], Bbox[1]}
	ptB := [2]float64{Bbox[2], Bbox[1]}
	ptC := [2]float64{Bbox[2], Bbox[3]}
	ptD := [2]float64{Bbox[0], Bbox[3]}
	firstArr := [][2]float64{ptA, ptB, ptC, ptD, ptA}
	geom := [][][2]float64{firstArr}
	return geom
}
