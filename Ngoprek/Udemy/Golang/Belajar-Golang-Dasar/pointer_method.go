package main
import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr " + man.Name
}

func main(){
	kinur := Man{"Kinur"}
	kinur.Married()

	fmt.Println(kinur.Name)
}
