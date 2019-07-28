/*
1       1
 2     2
  3   3
   4 4
    5
   4 4
  3   3
 2     2
1       1
*/
package main

import "fmt"
func main() {
	{
		var N int
	
		fmt.Print("Enter N: ")
		fmt.Scan(&N);
		for i:=1; i<=N; i++{
		
			for j:=1; j<i; j++{
				fmt.Print(" ")
			}
	
			fmt.Print(i)
	
			for j:=1; j<=((N - i) * 2 - 1); j++{
				fmt.Print(" ");
			}
	
		
			if (i != N){
				fmt.Print(i)
			}
	
				fmt.Println()
		}
	
		for i:=N-1; i>=1; i--{
			
			for j:=1; j<i; j++{
				fmt.Print(" ")
			}
	
			fmt.Print(i)
			m:=(N - i ) * 2 - 1
	
			for j:=1; j<=m; j++{
				fmt.Print(" ");
			}
	
			fmt.Print(i)
	
			fmt.Println()
		}

	}
}