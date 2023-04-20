package main

import (
	"fmt"
	"path"
	"time"

	"github.com/repoowner/sdk-for-go/appwrite"
)

func main() {
	stringInArray := []interface{}{"string in array"}

	client := appwrite.NewClient()
	err := client.SetTimeout(10 * time.Second)
	if err != nil {
		panic(err)
	}
	client.AddHeader("Origin", "http://localhost")
	fmt.Print("\n\nTest Started\n")
	testFooService(client, stringInArray)
	testBarService(client, stringInArray)
	testGeneralService(client, stringInArray)
}

func testFooService(client appwrite.Client, stringInArray []interface{}) {
	foo := appwrite.NewFoo(client)
	// Foo Service
	response, err := foo.Get("string", 123, stringInArray)
	if err != nil {
		fmt.Printf("foo.Get => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])

	response, err = foo.Post("string", 123, stringInArray)
	if err != nil {
		fmt.Printf("foo.Post => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])

	response, err = foo.Put("string", 123, stringInArray)
	if err != nil {
		fmt.Printf("foo.Put => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])

	response, err = foo.Patch("string", 123, stringInArray)
	if err != nil {
		fmt.Printf("foo.Patch => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])

	response, err = foo.Delete("string", 123, stringInArray)
	if err != nil {
		fmt.Printf("foo.Delete => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])
}

func testBarService(client appwrite.Client, stringInArray []interface{}) {
	bar := appwrite.NewBar(client)
	// Bar Service
	response, err := bar.Get("string", 123, stringInArray)
	if err != nil {
		fmt.Printf("bar.Get => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])

	response, err = bar.Post("string", 123, stringInArray)
	if err != nil {
		fmt.Printf("bar.Post => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])

	response, err = bar.Put("string", 123, stringInArray)
	if err != nil {
		fmt.Printf("bar.Put => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])

	response, err = bar.Patch("string", 123, stringInArray)
	if err != nil {
		fmt.Printf("bar.Patch => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])

	response, err = bar.Delete("string", 123, stringInArray)
	if err != nil {
		fmt.Printf("bar.Delete => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])
}

func testGeneralService(client appwrite.Client, stringInArray []interface{}) {
	general := appwrite.NewGeneral(client)
	// General Service
	response, err := general.Redirect()
	if err != nil {
		fmt.Printf("general.Redirect => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])

	testGeneralUpload(client, stringInArray)

	// Extended General Responses
	testGeneralDownload(client)

	// Large File Responses
	testLargeUpload(client, stringInArray)

	// Exception Responses
	response, err = general.Error400()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	response, err = general.Error500()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	response, err = general.Error502()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	general.Empty()

	// Final test
	response, err = general.Headers()
	if err != nil {
		fmt.Printf("general.Headers => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])
}

func testGeneralUpload(client appwrite.Client, stringInArray []interface{}) {
	general := appwrite.NewGeneral(client)
	uploadFile := path.Join("/app", "tests/resources/file.png")
	inputFile := appwrite.NewInputFile(uploadFile, "file.png")

	response, err := general.Upload("string", 123, stringInArray, inputFile)
	if err != nil {
		fmt.Printf("general.Upload => error %v", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])
}

func testGeneralDownload(client appwrite.Client) {
	general := appwrite.NewGeneral(client)
	response, err := general.Download()
	if err != nil {
		fmt.Printf("general.Download => error %v", err)
	}
	fmt.Printf("%s\n", response.Result)
}

func testLargeUpload(client appwrite.Client, stringInArray []interface{}) {
	general := appwrite.NewGeneral(client)
	uploadFile := path.Join("/app", "tests/resources/large_file.mp4")
	inputFile := appwrite.NewInputFile(uploadFile, "large_file.mp4")

	response, err := general.Upload("string", 123, stringInArray, inputFile)
	if err != nil {
		fmt.Printf("general.Upload => error %v\n", err)
	}
	fmt.Printf("%s\n", response.Result.(map[string]interface{})["result"])
}
