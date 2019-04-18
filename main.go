package main

import (
	sdk "github.com/gaia-pipeline/gosdk"
	"log"
	"os/exec"
	"time"
)

func thisiscmddemo() {
	log.Println("this is cmd demo : docker -v")
	cmd := exec.Command("docker -v")
	data, err := cmd.CombinedOutput()
	result := string(data)
	if len(result) > 0 {
		log.Println(result)
	}
	log.Println(result)
	log.Println(err)
}

func thisisgitdemo() {
	log.Println("this is cmd demo : git --version")
	cmd := exec.Command("git --version")
	data, err := cmd.CombinedOutput()
	result := string(data)
	if len(result) > 0 {
		log.Println(result)
	}
	log.Println(result)
	log.Println(err)
}

func CreateUser(args sdk.Arguments) error {
	log.Println("CreateUser is my defined has been started!")
	thisiscmddemo()
	thisisgitdemo()
	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateUser has been finished!")
	return nil
}

func MigrateDB(args sdk.Arguments) error {
	log.Println("MigrateDB has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("MigrateDB has been finished!")
	return nil
}

func CreateNamespace(args sdk.Arguments) error {
	log.Println("CreateNamespace has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateNamespace has been finished!")
	return nil
}

func CreateDeployment(args sdk.Arguments) error {
	log.Println("CreateDeployment has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateDeployment has been finished!")
	return nil
}

func CreateService(args sdk.Arguments) error {
	log.Println("CreateService has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateService has been finished!")
	return nil
}

func CreateIngress(args sdk.Arguments) error {
	log.Println("CreateIngress has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("CreateIngress has been finished!")
	return nil
}

func Cleanup(args sdk.Arguments) error {
	log.Println("Cleanup has been started!")

	// lets sleep to simulate that we do something
	time.Sleep(5 * time.Second)
	log.Println("Cleanup has been finished!")
	return nil
}

func main() {
	// Serve
	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}
