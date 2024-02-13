package main;
import "fmt"

// Arrays and loops
// Arrays are fixed length,same type,indexable,contiguous in memory
// 	& for address in memory
func main(){
	var intArr [3]int32;
	fmt.Println(intArr[0]);
	fmt.Println(&intArr[0]); 
	var intSlice []int32  = []int32{4,5,6}
	fmt.Println("The len of the slice is %v with capacity %v",len(intSlice),cap(intSlice)) 
	fmt.Println(intSlice) 
	intSlice = append(intSlice,7);
	fmt.Println(intSlice) 
}


// intslice is like a arraylist in java


