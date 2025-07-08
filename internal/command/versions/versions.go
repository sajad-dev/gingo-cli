package versions

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

const VERSION = "v0.1.0"
const REQUEREMENTS = "requirements.txt"

func Versions() error {

	fmt.Println("#VERSION")
	color.Green(VERSION)
	
	data,err := os.ReadFile(REQUEREMENTS)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("#REQUIREMENTS")
	color.Green(string(data))
	return nil
}
