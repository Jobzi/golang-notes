OUTPUT
=====

### General Formatting Verbs
The following verbs can be used with all data types:
| VERB             | Description                            |
| ---------------- | ----------------                       |
| %v               |Prints the value in the default format  |
| %#v              |Prints the value in Go-syntax format    |
| %T               |Prints the type of the value            |
| %%               |Prints the % sign                       |

````golang
package main
import ("fmt")

func main() {
  var i = 15.5
  var txt = "Hello World!"

  fmt.Printf("%v\n", i)     //output: 15.5  
  fmt.Printf("%#v\n", i)    //output: 15.5  
  fmt.Printf("%v%%\n", i)   //output: 15.5%  
  fmt.Printf("%T\n", i)     //output: float64
  fmt.Printf("%v\n", txt)   //output: Hello World!  
  fmt.Printf("%#v\n", txt)  //output: "Hello World!"    
  fmt.Printf("%T\n", txt)   //output: string  
}
````
### Integer Formatting Verbs
The following verbs can be used with the integer data type:
| VERB             | Description                  |
| ---------------- | ----------------             |
| %b               |Base 2 "Binario"              |
| %d               |Base 10                       |
| %+d              |Base 10 and always show sign  |
| %o               |Base 8                        |
| %O               |Base 8, with leading 0o       |
| %x               |Base 16,lowercase             |
| %X               |Base 16, uppercase            |
| %#X              |Base 16, with leading 0x      |
| %4d              |Pad with spaces (width 4, right justified)  |
| %-4d             |Pad with spaces (width 4, left justified)   |
| %04d             |Pad with zeroes (width 4      |

````golang
package main
import ("fmt")

func main() {
  var i = 15
 
  fmt.Printf("%b\n", i)   //output 1111  
  fmt.Printf("%d\n", i)   //output 15
  fmt.Printf("%+d\n", i)  //output +15
  fmt.Printf("%o\n", i)   //output 17
  fmt.Printf("%O\n", i)   //output 0o17
  fmt.Printf("%x\n", i)   //output f
  fmt.Printf("%X\n", i)   //output F
  fmt.Printf("%#x\n", i)  //output 0xf
  fmt.Printf("%4d\n", i)  //output   15
  fmt.Printf("%-4d\n", i) //output 15
  fmt.Printf("%04d\n", i) //output 0015
}
````

### String Formatting Verbs
The following verbs can be used with the string data type:
| VERB             | Description                  |
| ---------------- | ----------------             |
| %s               |Prints the value as plain string                            |
| %q               |Prints the value as a double-quoted string                  |
| %8s              |Prints the value as plain string (width 8, right justified) |
| %-8s             |Prints the value as plain string (width 8, left justified)  |
| %x               |Prints the value as hex dump of byte values    |
| % x              |Prints the value as hex dump with spaces       |
````golang
package main
import ("fmt")

func main() {
  var txt = "Hello"
 
  fmt.Printf("%s\n", txt)   //output: Hello
  fmt.Printf("%q\n", txt)   //output: "Hello"
  fmt.Printf("%8s\n", txt)  //output:    Hello
  fmt.Printf("%-8s\n", txt) //output: Hello
  fmt.Printf("%x\n", txt)   //output: 48656c6c6f
  fmt.Printf("% x\n", txt)  //output: 48 65 6c 6c 6f
}
````

### Float Formatting Verbs
The following verbs can be used with the float data type:
| VERB             | Description                  |
| ---------------- | ----------------             |
| %e               |Scientific notation with 'e' as exponent     |
| %f               |Decimal point, no exponent                   |
| %.2f             |Default width, precision 2                   |
| %6.2f            |Width 6, precision 2                         |
| %g               |Exponent as needed, only necessary digits    |

````golang
package main
import ("fmt")

func main() {
  var i = 3.141

  fmt.Printf("%e\n", i)     //output: 3.141000e+00
  fmt.Printf("%f\n", i)     //output: 3.141000
  fmt.Printf("%.2f\n", i)   //output: 3.14
  fmt.Printf("%6.2f\n", i)  //output:   3.14
  fmt.Printf("%g\n", i)     //output: 3.141 
}
````