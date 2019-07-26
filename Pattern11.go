/* 11111
   10001
   10001
   10001
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
					 fmt.Print("0")
					 }
			   
		   }
		   fmt.Println()
	   }
   }