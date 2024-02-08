package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func get2(n string) []string {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var init []string
	for _, v := range names {
		init = append(init, v[:3])
	}
	return init
}
func sayGreeting(n string) {
	fmt.Printf("good morning %v \n", n)
}
func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}
func sayFuck() {
	menu := map[string]float64{
		"soup":       4.99,
		"sea":        4.99,
		"arab juice": 4.99,
	}
	for k, v := range menu {
		fmt.Println(k, " ", v)
	}
}

func main() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run()
}

/*
math.Pi
sort.Ints(ages);
Sort Integer how can we loop of slice of numbers
slice of same kind of loop
slice of 2 items greeting items
greeting function every item
common for this right now
let me save
do this in the console function
len(variable )
func (circleArea)

standard library about that pacakge greeting simple string
greeting := hello
standared library values
fmt printf string method
and store them in a new sli
fixed length
slices arrays change
sprintF dh 3bara 3n eh
any b3ml
sprintf commit save and printf
fmt.spreintF()
var fmt
var agetwo=30
int type
so 20 30
int 8
int16
var numtwo =-128 int8
so lets try to into unnassinged int
this is entire functoin of the application one main function
go standard libary
packages formating  string and printing message to the console
inside the pacakge there is PrintLin
PrintLin("helloWorld");
program terminal   string and numbers
plus 25
minus 25 you int
doesn't specify zero to 255
signed vs unsinged integers
sgined -218
unsgined postive
//string
varaibles var nameOne string
var 2 we just set equal to a string
go look wi explicity type it
third way name3  namethree string
empty string
i can't change the type  mmkn a3mlo modify
nameThree="browser"
fmt.PrintLin(nameOne,nameTwo,namethree);
:= short hand versioned for string
most of the varaibles  you can't use this outside of functoin
var someName="hello";
unsgined all of the time
don't speicfy the no of bits how many bits
var scroreOne float :=
float64 large numebr realy doesn't matter
"fmt"
printLi
*/
