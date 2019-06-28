package main
import "fmt"
type catfish int
var x catfish
var y int
func main(){
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T",y)
}