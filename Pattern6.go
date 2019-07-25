/*
A
AB
ABC
ABCD
ABCDE
*/
package main
import(
	 "fmt"
)
func main(){
	n:=65
	for i:=1;i<=5;i++{
		for j:=1;j<=i;j++{
			fmt.Print(string(j+64))
		}
		fmt.Println()
		n++
	}
}
