/*
1
2*2
3*3*3
4*4*4*4
4*4*4*4
3*3*3
2*2
1

*/
package main
import "fmt"
func main() {
	for i:=1;i<=4;i++{
		for j:=1;j<=i;j++ {
			if(j<i){
					fmt.Print(i)
					fmt.Print("*")
			}else{ 
					fmt.Print(i)
			}
		}
		fmt.Println()
	}
	for i:=4;i>=1;i--{
		for j:=1;j<=i;j++{
			if(j<i){
					
					fmt.Print(i)
					fmt.Print("*")
			}else {
				fmt.Print(i)
				}
		}
		fmt.Println()
	}
}
