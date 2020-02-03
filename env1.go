package main 

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"strings"
)

func main() {
	flag.Usage = func() {
	}	
	
	iPtr := flag.Bool("i", false, "Ignore current environment")
	var arrayOfCLIArgs []string
	var arrayOfNewENV []string
	var arrayOfUtils []string
	var arrayOfUtils1 []string
	var utilCmds string

	flag.Parse()

	for counter,element := range os.Args[1:] {
			arrayOfCLIArgs = append(arrayOfCLIArgs, element)
			if (strings.ContainsAny(arrayOfCLIArgs[counter], "=") == true) {
				arrayOfNewENV = append(arrayOfNewENV, arrayOfCLIArgs[counter])
			} else {
					arrayOfUtils = append(arrayOfUtils, arrayOfCLIArgs[counter])
					utilCmds = strings.Join(arrayOfUtils, " ")
					utilCmds = strings.Replace(utilCmds, "-i", "", -1)
					arrayOfUtils1 = arrayOfUtils[1:]
			}
	}		
	if *iPtr == true {		
//		fmt.Println("New ENV Vars: ", arrayOfNewENV)
//		fmt.Println("Commands: ", arrayOfUtils1)

		fmt.Printf("No. of Commands: %d \n", len(arrayOfUtils1))
		
		if len(arrayOfNewENV) < 1  {
		
		} else {
			EnvHandler(arrayOfNewENV)	
		}

		if len(arrayOfUtils1) < 1 {
				
		} else {
			Syscall_without_environment(arrayOfUtils1)
		}
		return 
	} else {
		if len(arrayOfNewENV) > 0 {
			EnvHandler(arrayOfNewENV)	
		}
		if len(arrayOfUtils1) > 0  {
 			Syscall_with_environment(arrayOfUtils1)
		}
			PrintCurrentEnvironment()
	}
}

func Syscall_with_environment(utilities []string){
	environ := os.Environ()
	binary,lookupErr := exec.LookPath(utilities[0])
	if lookupErr != nil {
		panic(lookupErr)
	}
	execErr := syscall.Exec(binary,utilities,environ)
	if execErr != nil {
		panic(execErr)
	}
	return
}

func Syscall_without_environment(utilities []string) {
	binary,lookupErr := exec.LookPath(utilities[0])
	if lookupErr != nil {
		panic(lookupErr)
	}
	execErr := syscall.Exec(binary,utilities,nil)
	if execErr != nil {
		panic(execErr)
	}
	return
}

func EnvHandler(argument []string){
	for _,element := range argument{
		tmpEnVar := strings.Split(element, "=")
		name := tmpEnVar[0]
		value := tmpEnVar[1]

		if checkENV(name) == true {
			os.Getenv(name)

		} else {
			SetENV(name, value)
		}
	}
	return
}

func checkENV(name string) (bool){
	_,isset := os.LookupEnv(name)
	if !isset { 
		return false 
	} else { 
		return true 
	}
}

func SetENV(name string, value string){
	err := os.Setenv(name, value)
	if err != nil {
		panic(err)
	} 
}

func PrintCurrentEnvironment(){
	environ := os.Environ()
	for _,i := range environ {
		fmt.Printf("%s \n", i)
	}
}
