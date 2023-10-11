package main

/*В рассматриваемом линейном пространстве существует 3 типа точек: обычные, совместные, удаленные.
  Каждый тип точек обладает своим уникальным набором свойств - функционал главных из которых нам и нужно реализовать.
  У обычных точек: число Пифагора, т. е. x^2 + y^2
  У совместных точек: угол радиус-вектора, например, через обратные тригоном. функции
  У удаленных точек: разложение вида (x + y) + 1/(x + y)

  Входные данные: 1 < N < 10^6 точек с указанием типа (standard, joint, distant) и координат {x,y}.
  Известно, что точки могут оказаться и совместными, и удаленными одновременно.
  Вывод: для каждой из точек вывести результат подсчета. Для точек двойного типа - оба результата. */

import "fmt"
import "math"
import "strconv"

type dots struct {
	x    int
	y    int
	name string
}

func work() {
	//const power float64 = math.Pow( 10,6)
	var newdots [3]dots
	var stringjoint string = " "
	var stringdistant string = " "
	for i := 0; i < 3; i++ {
		fmt.Scanln("%s %d %d", &newdots[i].name, &newdots[i].x, &newdots[i].y)
	}
	//i := 0
	//newdots := [const a  = (math.Pow( 10,6))]dots
	//for i := 0; i <= int(math.Pow(10, 6)); i++ {
	for i := 0; i < 3; i++ {
		switch newdots[i].name {
		case "standard":
			fmt.Println(math.Pow(float64(newdots[i].x), 2) + math.Pow(float64(newdots[i].y), 2))

		case "joint":
			stringjoint = strconv.Itoa(int(math.Atan2(float64(newdots[i].y), float64(newdots[i].x))))
			//fmt.Println(math.Atan2(float64 (newdots[i].y), float64 (newdots[i].x)))
			fallthrough

		case "distant":

			stringdistant = strconv.Itoa((newdots[i].x + newdots[i].y) + 1/(newdots[i].x+newdots[i].y))
			fmt.Println(stringdistant + stringjoint)
		}
	}
}

func main() {
	work()
}
