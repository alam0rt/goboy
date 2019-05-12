package rom

import (
	"fmt"
	"io/ioutil"
)

func OpenROM(f string) []byte {

	r, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("error reading ROM")
	}
	return r

}
