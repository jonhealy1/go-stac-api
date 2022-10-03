package controllers

import "strings"

func bbox2polygon(Bbox []float64) [][][2]float64 {
	ptA := [2]float64{Bbox[0], Bbox[1]}
	ptB := [2]float64{Bbox[2], Bbox[1]}
	ptC := [2]float64{Bbox[2], Bbox[3]}
	ptD := [2]float64{Bbox[0], Bbox[3]}
	firstArr := [][2]float64{ptA, ptB, ptC, ptD, ptA}
	geom := [][][2]float64{firstArr}
	return geom
}

func returnDatetime(datetime string) []string {
	result := strings.Split(datetime, "/")
	start := ""
	end := ""
	if result[0] != ".." {
		start = result[0][0:19] + "Z"
	} else {
		start = "1900-01-01T01:01:01Z"
	}
	if result[1] != ".." {
		end = result[1][0:19] + "Z"
	} else {
		end = "3000-01-01T01:01:01Z"
	}
	parsed := make([]string, 2)
	parsed[0] = start
	parsed[1] = end
	return parsed
}
