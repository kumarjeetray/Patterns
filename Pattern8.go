/* 11111
   1   1
   1   1
   1   1
   11111
   */
   package main
   import "fmt"
   func main(){
	   for i:=1;i<=5;i++{
		   for j:=1;j<=5;j++{
			if (i==1 && j==1) || (i==5 && j==5)|| (i==1) || (i==5){
				   fmt.Print("1")
			   }else if (i>1 && i<5) && (j==1 || j==5){
				   fmt.Print("1")
			   }else{
					 fmt.Print(" ")
					 }
			   
		   }
		   fmt.Println()
	   }
   }