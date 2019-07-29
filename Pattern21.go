/*
   *****
  **     **
 **       **
 *         *
*           *
*           *
*           *
*           *
*           *
 *         *
 **       **
  **     **
	*****
	
*/
package main
import(
	 "fmt"
	 "math"
)

func main(){
	radius:=6.0
	for  i:= 0; i <=12; i++ { 
	  for j:= 0; j <=12; j++ { 
		  a:=(float64(i)-radius)*(float64(i)-radius)
		  b:=(float64(j)-radius)*(float64(j)-radius)
		dist:= math.Sqrt(float64(a+b))
		if (dist > float64(radius-0.5) && dist <float64(radius+0.5)){ 
		  fmt.Print("*")}else{
		  fmt.Print(" ")}       
	  } 
	
	  fmt.Println() 
	} 
  } 
	

