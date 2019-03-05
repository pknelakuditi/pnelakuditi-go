package main

import (
	"fmt"
)

func main() {
	////using package private functions
	//happy()
	//
	////using public methods from different package
	//hello2.Sad();
	//
	////using scanner for reading the string (only reads firt word)
	////var inputString string;
	////fmt.Scanln(&inputString)
	////fmt.Print(inputString)
	//
	////reading multi worded line
	////reader := bufio.NewReader(os.Stdin)
	////var input, _= reader.ReadString('\n')
	////fmt.Print(input)
	//
	////using pointers
	//var pointer *int;
	//var v int = 42;
	//pointer = &v;
	//if pointer != nil {
	//	fmt.Println("pointer : ", pointer, "value = ", *pointer)
	//}

	////using arrays
	//var stringArray [3]string;
	//stringArray[0] = "Hi"
	//stringArray[1] = "world"
	//stringArray[2] = "!"
	//fmt.Println(stringArray, "length ", len(stringArray))
	////using slices similar to Arraylist
	//var sliceList = []string{
	//	"red", "blue"}
	//fmt.Println(sliceList)
	//sliceList = append(sliceList, "violet")
	//fmt.Println(sliceList)
	//
	////sliceList = append(sliceList[:len(sliceList)-1])
	//sliceList = append(sliceList, "green")
	//sort.Strings(sliceList)
	//fmt.Println(sliceList)

	////using map and make
	//nameMap := make(map[int]string)
	//nameMap[1]= "Adam"
	//nameMap[10]="Joe"
	//nameMap[11]="Jim"
	//nameMap[12]="John"
	////delete(nameMap, 10)
	//fmt.Println(nameMap)
	//
	//for k,v := range nameMap{
	//	fmt.Println(k,v)
	//}
	dog := Dog{"breed1", 10}
	fmt.Println(dog.Breed)

}

type Dog struct {
	Breed  string
	Weight int
}
