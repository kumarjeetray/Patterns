/*
 A1
     AB12
    ABC123
   ABCD1234
  ABCDE12345
 ABCDEF123456
ABCDEFG1234567
*/
package main
import "fmt"
func main(){
		n:=7
		for i:= 1; i <= n; i++{ 
			k:= 'A'
			m:= 1
			for j:= 1; j <= (2 * n); j++ { 
				if (j >= n + 1 - i && (j <= n + i)) { 
					if j <= n{ 
						fmt.Print(string(k))
						k++
					} else { 
	  				fmt.Print(m)
						m++; 
					} 
				} else{
					fmt.Print(" ")
				}
					 
			} 
				fmt.Println()
		} 
	} 
