/*
Enter number of rows: 5
Enter number of columns: 5
11011
11011
00000
11011
11011
*/
package main
import "fmt"
func main(){
    var rows, cols int 
    fmt.Print("Enter number of rows: ")
    fmt.Scan(&rows)
    fmt.Print("Enter number of columns: ")
    fmt.Scan(&cols)
	centerRow:=(rows+1)/2
    centerCol:=(cols+1)/2
	for i:=1; i<=rows; i++{
        for j:=1; j<=cols; j++{
            if (centerCol == j || centerRow == i){
                fmt.Print("0")
            }else if ((cols%2 == 0 && centerCol+1 == j) || (rows%2 == 0 && centerRow+1 == i)){
                fmt.Print("0")
            }else{
                fmt.Print("1")
            }
        }

        fmt.Println()
    }
}
