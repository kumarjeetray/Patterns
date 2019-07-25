/*              *
			   **
			  ***
			 ****
			*****
			 ****
			  ***
			   **
				*
				*/
package main
import "fmt"
func main(){
	a:=4
	s:=4
	for i:=1;i<=5;i++{
		for k:=a;k>=0;k--{
			fmt.Print(" ")
		}
		for j:=1;j<=i;j++{
			fmt.Print("*")
		}
		a--;
		fmt.Println()
	}
		for m:=1;m<=4;m++{
			for n:=1;n<=(m+1);n++{
				fmt.Print(" ")
			}
			for o:=1;o<=s;o++{
				fmt.Print("*")
			}
			s--;
			fmt.Println()
		}
}
	