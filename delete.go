package main
 
import (
"fmt" 
)
 

 
func main() {
 
	a := []int{0, 1, 2, 3, 4}
	fmt.Println(a)

	//删除第i个元素
	i := 4
	fmt.Println(i)

	

	if i == len(a) {
		a = append(a[:i])
		fmt.Println(a)
	}else{
		a = append(a[:i], a[i+1:]...)
		fmt.Println(a)
	}



}
