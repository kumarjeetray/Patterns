/*       1234
	 2341
	 3421
	 4321
*/
package main
import "fmt"
func main(){
	c:=1
	d:=0
	j:=0
	for i:=1;i<=4;i++{
		for j=c;j<=4;j++{
			fmt.Print(j)
		}
		if i>1 {
		for k:=d;k>=1;k--{
			fmt.Print(k)
		}
		}
		c++
		d++
		fmt.Println()
	}
}


