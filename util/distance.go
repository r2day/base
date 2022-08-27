package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:::                                                                         :::
//:::  This routine calculates the distance between two points (given the     :::
//:::  latitude/longitude of those points). It is based on free code used to  :::
//:::  calculate the distance between two locations using GeoDataSource(TM)   :::
//:::  products.                                                              :::
//:::                                                                         :::
//:::  Definitions:                                                           :::
//:::    South latitudes are negative, east longitudes are positive           :::
//:::                                                                         :::
//:::  Passed to function:                                                    :::
//:::    lat1, lon1 = Latitude and Longitude of point 1 (in decimal degrees)  :::
//:::    lat2, lon2 = Latitude and Longitude of point 2 (in decimal degrees)  :::
//:::    optional: unit = the unit you desire for results                     :::
//:::           where: 'M' is statute miles (default, or omitted)             :::
//:::                  'K' is kilometers                                      :::
//:::                  'N' is nautical miles                                  :::
//:::                                                                         :::
//:::  Worldwide cities and other features databases with latitude longitude  :::
//:::  are available at https://www.geodatasource.com                         :::
//:::                                                                         :::
//:::  For enquiries, please contact sales@geodatasource.com                  :::
//:::                                                                         :::
//:::  Official Web site: https://www.geodatasource.com                       :::
//:::                                                                         :::
//:::          Golang code James Robert Perih (c) All Rights Reserved 2018    :::
//:::                                                                         :::
//:::           GeoDataSource.com (C) All Rights Reserved 2017                :::
//:::                                                                         :::
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
func Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

func parserPosition(pos string) (float64, float64) {
	xy := strings.Split(pos, ",")
	if len(xy) != 2 {
		panic("pos is invalid, should be `12.34,45.67`")
	}
	x, err := strconv.ParseFloat(xy[0], 64)
	if err != nil {
		panic("x is invalid, should be a float string")
	}
	y, err := strconv.ParseFloat(xy[1], 64)
	if err != nil {
		panic("y is invalid, should be a float string")
	}
	return x, y
}

func DistancePosition(a, b string) float64 {
	ax, ay := parserPosition(a)
	bx, by := parserPosition(b)
	return Distance(ax, ay, bx, by, "K")
}

func DistanceShow(dis float64) string {
	// print("dis-->", dis)
	val2 := math.Dim(dis, 1)
	val := ""
	if val2 < 1 {
		m := dis * 1000
		val = fmt.Sprintf("%dm", int(m))
	} else {
		val = fmt.Sprintf("%.2fkm", dis)
	}
	return val
}
