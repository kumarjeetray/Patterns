package main
import "fmt"
func main(){
     	k:=1
		r:=10
		
	   fmt.Println("FLOYD'S TRIANGLE")
		for i:=1;i<=r;i++{
			for j:=1;j<=i;j++{
				 fmt.Print(k)
				 fmt.Print(" ")
				 k++
			}
			fmt.Println()
		}
	   
	}