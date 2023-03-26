package redis

import (
	"fmt"
	"testing"

	"github.com/goccy/go-json"
)

func TestSet(t *testing.T) {
	marshal, _ := json.Marshal("ddddd")
	fmt.Println(string(marshal))
	var str string
	json.Unmarshal(marshal, &str)
	fmt.Println(string(str))

}
