/*                          1
			   2 2
			  3 3 3
			 4 4 4 4
			5 5 5 5 5
		*/
		package main
		import "fmt"
		func main(){
			a:=4
			for i:=1;i<=5;i++{
				for k:=a;k>0;k--{
					fmt.Print(" ")
				}
				for j:=1;j<=i;j++{
					fmt.Print(i)
					fmt.Print(" ")
				}
				a--
				fmt.Println()
			}
		}
