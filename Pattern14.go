/*
	1
	1 2
	3 5 8
	13 21 34 55
	89 144 233 377 610
*/
package main
import "fmt"
func main(){
	t1:= 1
	t2:= 1
	nextTerm:=0
	a:=1
    for i:= 1;i<=5;i++{
		for j:=1;j<=a;j++{
		fmt.Print(t1)
		fmt.Print("  ")
        nextTerm = t1 + t2
        t1 = t2
		t2 = nextTerm
		}
		a++
		fmt.Println()
    }
}
