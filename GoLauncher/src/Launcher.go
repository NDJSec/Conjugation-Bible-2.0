package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"gopkg.in/ini.v1"
)

func main() {
	cfgFile := "../assets/ConjugationBible.ini"
	requiredPackagesSection := "requiredPackages"

	fmt.Println("Welcome to the Conjugation-Bible-2.0 Launcher")
	fmt.Println("Checking if dependencies are installed...")

	cfg, err := ini.Load(cfgFile)
	if err != nil {
		createIniFile(cfgFile)
		fmt.Printf("Creating %v\n", cfgFile)
		log.Fatal(err)
	}

	if cfg.Section(requiredPackagesSection).Key("ConjugationBibleInstalled").MustBool() == false {
		fmt.Printf("Conjugation-Bible-2.0 not installed.\n INSTALLING...\n")

		if cfg.Section(requiredPackagesSection).Key("pythonExecutable").MustBool() == false {
			pythonCmd := exec.Command("python", "--version")
			pyErr := pythonCmd.Run()

			if pyErr != nil {
				log.Fatal("Python not installed. Visit https://www.python.org/ to install")
			}
			cfg.Section(requiredPackagesSection).Key("pythonExecutable").SetValue("true")
		}

		if cfg.Section(requiredPackagesSection).Key("pythonDependencies").MustBool() == false {
			setupScript := "setup.py"
			pyArgs := "install"
			pythonCmd := exec.Command("cmd", "python", setupScript, pyArgs)
			pyErr := pythonCmd.Run()

			if pyErr != nil {
				log.Fatal("Dependencies not installed. Check dependencies by running 'python setup.py install' or contact maintainer")
			}
			cfg.Section(requiredPackagesSection).Key("pythonDependencies").SetValue("true")
		}
		fmt.Println("Conjugation-Bible-2.0 Installed 😊 Enjoy!!")
		cfg.Section(requiredPackagesSection).Key("ConjugationBibleInstalled").SetValue("true")
		cfg.SaveTo(cfgFile)
		cmd := exec.Command("ConjugationBible")
		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Conjugation-Bible-2.0 Installed 😊 Enjoy!!")
		cmd := exec.Command("ConjugationBible")
		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	}
}

func createIniFile(cfgFileName string) {
	requiredPackagesSection := "requiredPackages"
	newCfgFile, err := os.Create(cfgFileName)

	if err != nil {
		log.Fatal(err)
	}

	newCfgFileLoaded, err := ini.Load(newCfgFile)
	if err != nil {
		log.Fatal(err)
	}

	newCfgFileLoaded.NewSection(requiredPackagesSection)
	newCfgFileLoaded.Section(requiredPackagesSection).NewKey("pythonExecutable", "false")
	newCfgFileLoaded.Section(requiredPackagesSection).NewKey("pythonDependencies", "false")
	newCfgFileLoaded.Section(requiredPackagesSection).NewKey("ConjugationBibleInstalled", "false")
	newCfgFileLoaded.Section(requiredPackagesSection).NewKey("version", "2.0")
	newCfgFileLoaded.SaveTo(cfgFileName)
}
