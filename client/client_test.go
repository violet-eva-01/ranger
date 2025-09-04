// Package client @author: Violet-Eva @date  : 2025/9/4 @notes :
package client

import (
	"fmt"
	"github.com/violet-eva-01/ranger/policy"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	a := policy.Policy{}
	conv(a)
	b := &policy.Policy{}
	conv(b)
}

func conv(t any) {
	a := &t
	fmt.Println(reflect.ValueOf(a).Kind())
}
