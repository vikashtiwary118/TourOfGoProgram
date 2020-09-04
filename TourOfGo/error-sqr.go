package main 

import ("fmt"
        "log"
        "errors")
type error interface {
    Error() string
}

// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}


// New returns an error that formats as the given text.
func New(text string) error {
    return &errorString{text}
}



func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    
    return f
    // implementation
}



func main() {
	f, err := Sqrt(-1)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(f)

}