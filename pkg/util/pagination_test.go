package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGetPage(t *testing.T) {

	c := gin.Context{
		Request:  nil,
		Writer:   nil,
		Params:   gin.Params{{"dd", "dd"}},
		Keys:     nil,
		Errors:   nil,
		Accepted: nil,
	}

	a := GetPage(&c)
	fmt.Println(a)
	if a < 0 {
		t.Errorf("%d", a)
	}
}
