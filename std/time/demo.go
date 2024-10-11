package main

import "time"
import "fmt"


func main(){
	// d := time.Date(2022, 7, 1, 13, 56, 24, 0, time.UTC)
	// fmt.Println(d)
	// fmt.Println(d.Local())
	// fmt.Println(d.UTC())

	// loc, _ := time.LoadLocation("Asia/Shanghai")
	// t := d.In(loc)
	// fmt.Println(t.Location().String())
	// fmt.Println(t.Zone())

	// fmt.Println(d== t)
	// fmt.Println(d.Equal(t))

	// t = t.AddDate(0,0, 1)
	// fmt.Println(d.Before(t))
	// fmt.Println(d.After(t))

	// d1, _:= time.ParseDuration("1h2m3s100us")
	// fmt.Println(d1)

	// d1, _ = time.ParseDuration("1us1ns")
	// fmt.Println(d1.Nanoseconds())

	// var t2 time.Time
	// fmt.Println(t2.IsZero())
	// t2 = time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	// fmt.Println(t2.IsZero())
	// fmt.Println(t2.Unix())
	// fmt.Println(t2.UnixNano())


	// fmt.Println(t2.Format("2006-01-02 15:04:05 -0700 MST"))
	// fmt.Println(t2.Format("2006-01-02 15:04:05"))
	// fmt.Println(t2.Format("06-1-2 3:4:5PM"))

	testZone()
}


func testZone(){
		// China doesn't have daylight saving. It uses a fixed 8 hour offset from UTC.
		secondsEastOfUTC := int((8 * time.Hour).Seconds())
		beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	
		// If the system has a timezone database present, it's possible to load a location
		// from that, e.g.:
		//    newYork, err := time.LoadLocation("America/New_York")
	
		// Creating a time requires a location. Common locations are time.Local and time.UTC.
		timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
		sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	
		// Although the UTC clock time is 1200 and the Beijing clock time is 2000, Beijing is
		// 8 hours ahead so the two dates actually represent the same instant.
		timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
		fmt.Println(timesAreEqual)
	
	
}
