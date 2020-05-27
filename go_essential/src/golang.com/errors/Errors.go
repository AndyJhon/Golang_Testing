package main

import "fmt"
import "os"
import "github.com/pkg/errors"

func main(){
	callReadConfig()
}
/*********************************************/
// Read config
type Config struct{}

func readConfig(path string) (*Config, error) {
	file, error := os.Open(path)
	if error != nil {
		return nil, errors.Wrap(error, "can not open configuration file")
	}

	defer file.Close()

	cfg := &Config{}

	return cfg, nil

}

func callReadConfig(){
	config, error := readConfig("/path/jarvis/file.txt")
	if error != nil {
		fmt.Fprintf(os.Stderr,"error %+s\n", error)
		os.Exit(1)
	}
	fmt.Println(config)
}