package CustomErrorHandling

import (
	"fmt"
	"log"
	"math"
)

// creating own error method struct
type destError struct {
	lat  string
	long string
	err  error
}

func (n destError) Error() string {
	return fmt.Sprintf("a error occured in lat %v long %v with this %v", n.lat, n.long, n.err)
}

func CustomErrorHandling() {
	_, err := Sqrt(-100)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("error :: Sqrt of negative number")
		// formating errors
		nme := fmt.Errorf("error :: Sqrt of negative number : %v", f)
		return 0, destError{"122.4818 N", "32.232 S", nme}
	}
	return math.Sqrt(f), nil
}
