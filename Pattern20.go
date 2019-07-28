/*
1 2 3 4 5
      4

    3

  2

1 2 3 4 5
*/
package main
import "fmt"
func main(){
	N:=5
	Top:= 1
	Bottom:= 1
	Diagonal:= N - 1;  
    for index:= 0; index < N; index++{
		fmt.Print(Top)
		Top++
		fmt.Print(" ") 
	}
    fmt.Println()
    for index:= 1; index < N - 1; index++{  
        for side_index:= 0; side_index < 2 * (N - index - 1); 
             side_index++{
            fmt.Print(" ")
			 }
		fmt.Println(Diagonal)
		Diagonal--
		fmt.Println()
    } 
    for index:= 0; index < N; index++{
		fmt.Print(Bottom)
		Bottom++
		fmt.Print(" ") 
	}
} 
