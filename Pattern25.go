/*
1     4
2     3
3     2
4 3 2 1
3     2
2     3
1     4
*/
package main
import "fmt"
func main(){
	left:= 0
	N:=4
	middle:= N - 1
	right:= N + 1
    for row:= 0; row < 2 * N - 1; row++ {  
        if row < N{ 
			left++
            fmt.Print(left)
		}else{
			left--
			fmt.Print(left)
		}
        for col:= 1; col < N - 1; col++{ 
            if (row != N - 1){
				fmt.Print("  ")
			}else{
			   fmt.Print(" ")
			   fmt.Print(middle)
			   middle--
			}
        } 
        if (row < N){
		   fmt.Print(" ")
		   right--
		   fmt.Print(right) 
		}else{
		   fmt.Print(" ")
		   right++
		   fmt.Print(right)
		} 
        fmt.Println()
    } 
} 
