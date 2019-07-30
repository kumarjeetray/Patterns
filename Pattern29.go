/*
1        1
12      21
123    321
1234  4321
1234554321
*/
package main
import "fmt"
func main(){
    	for i:=1;i<=5;i++{
    		for j:=1;j<=5;j++{
    			if j<=i{
						fmt.Print(j)
				}else{
						fmt.Print(" ")
				}
    		}
    		for j:=5;j>=1;j--{
    			if j<=i{
						fmt.Print(j)
				}else{
						fmt.Print(" ")
				}
    		}
    		fmt.Println()
    	}
    }