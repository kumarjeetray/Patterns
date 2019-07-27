/* 11111
   0000
   111
   00
   1
   */
   package main
   import "fmt"
   func main(){
	   a:=5
	   for i:=1;i<=5;i++{
		   for j:=1;j<=a;j++{
			   if i%2==0{
				   fmt.Print("0")
			   }else{
				   fmt.Print("1")
			   }
		   }
		   a--
		   fmt.Println()
	   }
   }