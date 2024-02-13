package main
import "fmt"

// data types => int ,int16,int32,int63,float32,float64,string,error,bool
func main(){
	var intNum int = 3
    var floarNum float32 = 4444444.6166;
    
	// printResult(intNum,5);
	if(intNum==32){
		fmt.Println(intNum);
	}else{
		fmt.Println(floarNum);
	}
	printResult(intNum,floarNum)
	
}
func printResult(num1 int,num2 float32){
	fmt.Println(num1,num2);
}