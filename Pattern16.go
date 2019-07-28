/*
Enter number of rows: 5
Enter number of columns: 5
55555
54444
54333
54322
54321
*/
package main
import "fmt"
func main(){
    var rows, cols int 
    fmt.Print("Enter number of rows: ")
    fmt.Scan(&rows)
    fmt.Print("Enter number of columns: ")
    fmt.Scan(&cols)
	for i:=1; i<=rows; i++{
        for j:=cols; j>cols-i; j--{
            fmt.Print(j)
        }

        for j:=1; j<=cols-i; j++{
            fmt.Print(rows - i + 1)
        }

      fmt.Println()
    }
}