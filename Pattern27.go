/*
1
2 3
4 5 6
7 8 9 10
11 12 13 14 15
16 17 18 19 20 21
22 23 24 25 26 27 28
29 30 31 32 33 34 35 36
37 38 39 40 41 42 43 44 45
46 47 48 49 50 51 52 53 54 55

*/
package main
import "fmt"
func main(){
     	k:=1
		r:=10
		
	   fmt.Println("FLOYD'S TRIANGLE")
		for i:=1;i<=r;i++{
			for j:=1;j<=i;j++{
				 fmt.Print(k)
				 fmt.Print(" ")
				 k++
			}
			fmt.Println()
		}
	   
	}
