package main 

import ("fmt";"math")


/*
* Just learn how to do Sqrt with newton method
*/
func Sqrt(x float64) float64 {
	z := float64(1)
	for fuck := 0; fuck < 100; fuck++ {
		z = z - ((math.Pow(z,2) - x) / (2 * z))
	}
	return z
}

func plusOne(x int)(int){
    x++
    return x
}
func main(){
    
    //This is the end of the story
    

    fmt.Println("Fuck, it worked!")

    //For loop, til for is the new while
    for i := 0; i < math.MaxInt8 >> 6;i++{
        fmt.Println("Print me")
    }
    //For loop, where init and post are optional
    //Dude this is basically a while loop...
    //With the condition of while(wot < 3)...
    wot := 1
    for ; wot < 3; {
        fmt.Println(wot)
        wot += 1
    }
    
    //Okay here's an actual way to write while loop in Go
    for wot < 5{
        fmt.Println(wot)
        wot += 1
    }

    //Even thought, it still won't deal with infinite loop despite how safe Go is
    // for{}

    //Swift time, basically the syntax is the same
    if wot < 6{
        fmt.Println(wot)
    }

    //Lazy plus one evaluation
    if dis := plusOne(wot); dis < 7{
        fmt.Println(dis)
    }else{
        for{}
    }

    //No while loop, but still have switch case
    //Just like the opposite of python...
    fmt.Println("Sqrt(100) is 10")
    switch os := Sqrt(100); os{
        case 100:
            fmt.Println("Impossible")
        case 10:
            fmt.Println("Yep")
        default:
            fmt.Println("U wot m8")
    }

    //Second switch, that has no condition
    switch{
        case 1 + 1 == 3:
            fmt.Println("This case is 1 + 1 = 3, do you think it's possible to get in here?")
        case Sqrt(4) < 2:
            fmt.Println("This case is Sqrt(4) < 2, do you think it's possible to get in here?")
        default:
            fmt.Println("Good evening")
    }

    //NEW STUFF: Defer
    //Defer holds the expression until the rest is done
    fmt.Println("Mission start")
    fmt.Println("I no scope Jose")
    fmt.Println("I no scope Nilo")
    fmt.Println("I no scope myself")
    fmt.Println("ah! I died bitch!")

    //So what defer does, is to put the on the stack
    for i := 0; i < 10; i++{
        defer fmt.Println("Counting people who had been noscoped: ",i)
    }
    defer fmt.Println("Credit roll")
    defer fmt.Println("Mission complete")
}
