/*
				1
			   123
			  12345
			 1234567
			123456789
			 1234567
			  12345
			   123
				1
*/
package main
import "fmt"
func main(){
	l:=4
	e:=7
	o:=1
	for i:=1;i<=9;i+=2{
		for k:=l;k>0;k--{
			fmt.Print(" ")
		}
		for j:=1;j<=i;j++{
			fmt.Print(j)
		}
		l--
		fmt.Println()
	}
	for a:=1;a<=5;a++{
		for m:=1;m<=o;m++{
			fmt.Print(" ")
		}
		for b:=1;b<=e;b++{
			fmt.Print(b)
		}
		o++
		e=e-2
		fmt.Println()
	}
}