/* 
  *    **********    *
  **   ****  ****   **
  ***  ***    ***  ***
  **** **      ** ****
  ******        ******
	   
  */ 
	  package main
	  import "fmt"
	  func main(){
		  a:=1
		  b:=1
		  c:=1
		  d:=1
		  e:=1
		  f:=1
		  g:=1
		  h:=1
		  for i:=1;i<=5;i++{
		  for j:=1;j<=a;j++{
			  fmt.Print("*")
		  }
		  a++
		  for k:=4;k>=b;k--{
			  fmt.Print(" ")
		  }
		  b++
		  for l:=5;l>=c;l--{
			  fmt.Print("*")
		  }
		  c++
		  if c>2{
		  for m:=d;m>1;m--{
			  fmt.Print(" ")
		  }
		  }
		  d++
		  if d>2{
		  for n:=1;n<e;n++{
			  fmt.Print(" ")
		  }
		  }
		  e++
		  for o:=4;o>=f-1;o--{
			fmt.Print("*")
		  }
		  f++
		if f>1{
		for q:=4;q>=h;q--{
			fmt.Print(" ")
		}
		}
		h++
		for p:=g;p>=1;p--{
			fmt.Print("*")
		}
	    g++
		fmt.Println()
	}
}


		 