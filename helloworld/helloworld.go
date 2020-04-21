package helloworld

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
)

func GetHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return home
}
