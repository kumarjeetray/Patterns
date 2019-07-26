/*           1
		   4 9 16
		25 36 49 64 81
	100 121 144 169 196 225 256
289 324 361 400 441 484 529 576 625
*/
package main
import "fmt"
func main(){
	l:=20
	c:=1
	a:=1
	t:=0
	d:=3
	m:=2
	for i:=1;i<=5;i++{
		for k:=l;k>0;k--{
			fmt.Print(" ")
		}
		for j:=a;j<=c;j++{
			fmt.Print(j*j)
			fmt.Print(" ")
			t++
		}
		c+=d
		l-=m
		a=t+1
		d+=2
		m++
		fmt.Println()
	}
}