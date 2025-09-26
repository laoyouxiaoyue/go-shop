package config

import (
	"fmt"
	"testing"
)

func Test_configLoaderViper_Load(t *testing.T) {
	clv := ConfigLoaderViper{}
	load, err := clv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", load)
}
