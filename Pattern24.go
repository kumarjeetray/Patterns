/*
 
 ******
 *     *
 *     *
 *     *
 *     *
 *     *
 *     *
 *     *
 ******
 
 */
 package main
import "fmt"
func main(){
	n:=9
	for i:= 0; i < n; i++{  
        for j:= 0; j <= n; j++{ 
            if (j == 1 || ((i == 0 || i == n-1) && 
               (j > 1 && j < n-2)) || (j == n-2 && 
                i != 0 && i != n-1)){
                fmt.Print("*")  
                }else{
                fmt.Print(" ") 
                }
        } 
          
    fmt.Println()
    } 
}
