package main

/*
* Import function from the package "main"
*/
import(
    "fmt"
    "math/rand"
)

// take argument a and argument b, and an interger of the subtract of a and b
func minus(a,b int) int{
    return a - b
}

// Naked return
func minus_lazy(a,b int) (answer int){
    answer = a - b
    return
}

// Multiple return
func minus_two_numbers(a,b int) (int,int){
    return (a - b),(b - a)
}

//By the way, variable can be ouside of package
var outsider int
func main(){
    fmt.Println("Hello world") 
    fmt.Println("Random number", rand.Intn(3))
    fmt.Println("Random number", rand.Intn(3))
    fmt.Println("3 - 2 = ", minus(3,2))
    fmt.Println("6 - 5 = ", minus_lazy(6,5))
    a,b := minus_two_numbers(6,5)
    fmt.Println("6 - 5 and 5 - 6 = ",a," and ",b)

    //Declare a variable for later used
    var wot,is,dis,bull,shit bool
    fmt.Println("Printing out variable")
    fmt.Println(wot,is,dis,bull,shit) 
    fmt.Println("Have you realized boolean in default is false?")

    //Let's declare a variable with initializer
    //So no need to declare a type for it
    var ignore,me = "blacklisted","me"
    fmt.Println(ignore,me)

    //Here's a tough one
    //Implicit vs explicit
    temp := "Seriously, I am temp"
    //Printing out an implicit variable
    fmt.Println(temp)

    //Let's learn some Computer architecture
    normalNumber0 := 0
    normalNumber1 := 1
    normalNumber2 := 2
    fmt.Println("printing out normal number: ",normalNumber0,normalNumber1,normalNumber2)
    //Oh well, if you can implicit enough then you don't need to have to have var inside the package
    veryBigNumber := uint64(1<<64 - 1)
    var pretendedBigNumber uint8 = 1 << 8 - 1
    fmt.Println("Given value \"veryBigNumber\" ",veryBigNumber)
    fmt.Println("Given value \"pretendedBigNumber\" ",pretendedBigNumber,"\n")
    fmt.Println("After adding one for both of them")
    fmt.Println("\"veryBigNumber\": ",veryBigNumber+1)
    fmt.Println("\"pretendedBigNumber\": ",pretendedBigNumber+1)

    //Let's have some interesting conversation shalle we?
    i := 1
    f := float64(i)
    h := int64(i)
    fmt.Println(i,f+0.5,h << 63)

    //Type interence, aka right hand rule
    //Everything you typed will be depend on the right hand side
    aa := 1 //int
    bb := 2.5 //float64
    cc := 0.69 + 3i //complex128
    fmt.Println("Given three values ",aa,bb,cc)
    fmt.Printf("Their corresponding type would be %T %T %T\n",aa,bb,cc)

    //Constant as usual
    const FUCK = "you"
    //Right! Constant can't be implicit, you know that's a constant instead of just some variable right?
    //FUCK := "you"
    fmt.Println("Printing out Constant variable \"FUCK\":",FUCK)

    //Let's do the constant but for numeric
    //But first let me declare them
    const (
        six = 6
        seven = 7
        eight = six + 2
    )
    //By the way, numberic constant takes the type it needed by its context 
    fmt.Println("Printing out constant six after becoming float: ",six + 2 - 1 - 1 * 0.3)

}