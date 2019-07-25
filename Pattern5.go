/*
ABCDE
BCDE
CDE
DE
E
*/
package main
import(
	 "fmt"
)
func main(){
	n:=65
	a:=5
	for i:=1;i<=5;i++{
		for j:=n;j<=69;j++{
			fmt.Print(string(j))
		}
		fmt.Println()
		n++
		a--
	}
}