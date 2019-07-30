/*
*           *           *
#           #           #
##         ###         ##
###       #####       ###
####     #######     ####
#####   #########   #####
###### ########### ######
#########################
#########################
#########################
#########################
*************************

*/

package main
import "fmt"
func main(){
	length:=25
	height:=12
    for i:= 0; i < height; i++ { 
        for j:= 0; j < length; j++{ 
            if i == 0{ 
                if (j == 0 || j == height || j == length - 1){ 
                   fmt.Print("*")
                }else{ 
                    fmt.Print(" ")
			}

			}else if i == height- 1{ 
                fmt.Print("*") 
            }else if ((j < i || j > height - i) &&(j < height + i || j >= length - i)){ 
				fmt.Print("#")
			}else{
				fmt.Print(" ") 
			}
        } 
       fmt.Println() 
    } 
} 
