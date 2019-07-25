/*               0
				101
			   21012
			  3210123
			 432101234
			54321012345
			*/
			package main
			import "fmt"
			func main(){
				a:=5
				b:=0
				for i:=0;i<=5;i++{
					for k:=a;k>=1;k--{
						fmt.Print(" ")
					}
						for j:=b;j>=0;j--{
							fmt.Print(j)
						}
					a--;
					b++;
					if i>=1{
						for l:=1;l<=i;l++{
							fmt.Print(l)
						}
					}
					fmt.Println()
				}
			}