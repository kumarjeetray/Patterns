/*
AAAAA
BBBB
CCC
DD
E
*/
package main
import(
	 "fmt"
)
func main(){
	n:=65
	a:=5
	for i:=n;i<=69;i++{
		for j:=1;j<=a;j++{
			fmt.Print(string(i))
		}
		fmt.Println()
		n++
		a--
	}
}