package main
import(
	"fmt"
)
func main(){

	// text, err := ReadBook("path")

	// if err != nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println(text)

	food, err := ReadFoods("./Data/Food/food.csv")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(food)

}