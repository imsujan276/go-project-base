package utilities

import (
	"encoding/json"
	"fmt"
	"log"
)

//Pretty display the claims licely in the terminal
func PrettyLog(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
