/*
*                          *
* *                      * *
* * *                  * * *
* * * *              * * * *
* * * * *          * * * * *
* * * * * *      * * * * * *
* * * * * * *  * * * * * * *
* * * * * * *  * * * * * * *
* * * * * *      * * * * * *
* * * * *          * * * * *
* * * *              * * * *
* * *                  * * *
* *                      * *
*                          *

*/
package main
import "fmt"
func main(){
	n:=7
	for i:=1; i<=n; i++{ 
        for j:=1; j<=(2*n); j++{  
            if (i<j){
                fmt.Print(" "); 
			 } else{
                fmt.Print("*"); 
			 }
            if (i<=((2*n)-j)){
                fmt.Print(" ") 
			}else{
                fmt.Print("*"); 
			}
		}	 
        fmt.Println() 
    } 
    for i:=1; i<=n; i++{ 
        for j:=1;j<=(2*n);j++{ 
            if (i>(n-j+1)){ 
                fmt.Print(" ") 
			}else{
                fmt.Print("*") 
			}    
            if ((i+n)>j){ 
                fmt.Print(" ") 
			}else{
				fmt.Print("*")
			}
        } 
        fmt.Println() 
    } 
}