package main 
import "sort"
import "fmt"

type earthMas float64 
type au float64 

type Planet struct{
	name string 
	mass earthMas 
	distance au 
}

type planetSorter struct{
	planets []Planet 
	by  func(p1, p2 *Planet) bool
}

func(s *planetSorter) Len()int{
	return len(s.planets)
}

func(s *planetSorter) Swap(i, j int){
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

func(s *planetSorter) Less(i, j int) bool{
	return s.by(&s.planets[i], &s.planets[j])
}


type By func(p1, p2 *Planet) bool 

func(by By)Sort(planets []Planet){
	ps := &planetSorter{
		planets: planets,
		by: by,
	}
	sort.Sort(ps)
}

func main(){
	planets := []Planet{
		{"Mercury", 0.055, 0.4},
		{"Venus", 0.815, 0.7},
		{"Earth", 1.0, 1.0},
		{"Mars", 0.107, 1.5},
	}

	name := func(p1, p2 *Planet) bool{
		return p1.name < p2.name
	}
	mass := func(p1, p2 *Planet) bool{
		return p1.mass < p2.mass
	}
	distance := func(p1, p2 *Planet) bool{
		return p1.distance < p2.distance
	}	
	decreasingDistance := func(p1, p2 *Planet) bool{
		return !distance(p1, p2)
	}
	//类型转换成By类型，然后调用Sort方法
	By(name).Sort(planets)
	fmt.Println("By name:", planets)

	By(mass).Sort(planets)
	fmt.Println("By mass:", planets)

	By(distance).Sort(planets)
	fmt.Println("By distance:", planets)

	By(decreasingDistance).Sort(planets)

	
	
}