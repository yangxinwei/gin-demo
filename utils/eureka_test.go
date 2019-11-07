package utils

import (
	"testing"
	"fmt"
)

func TestGetEureka(t *testing.T) {
	GetEureka()
	RegisterInstance()
}

func TestEureka(t *testing.T) {
	client:=GetEureka()
	instance,err:=client.GetInstance("10.141.4.154","GATEWAY")
	if err != nil {

	}
	fmt.Println(instance.HostName)
}