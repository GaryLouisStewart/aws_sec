package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func createIamUsers() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.

	// create a new session in eu-west-2 london region, change this to the region that matches your aws credenetials file ~/.aws/credentials or ~/.aws/profile file(s)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2")},
	)

	// Create a IAM service client.
	svc := iam.New(sess)

	// take the filename as an argument that is passed when you run the script using os.args
	file := os.Args[1]

	// read the file passed into the function used the following: https://golangdocs.com/golang-read-file-line-by-line to figure this out.
	readFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	// create an array to append each line from the file into and use later on to iterate over.
	var linesInFile []string

	for scanner.Scan() {
		linesInFile = append(linesInFile, scanner.Text())
	}

	for _, line := range linesInFile {

		_, err = svc.GetUser(&iam.GetUserInput{
			UserName: &line,
		})

		if awserr, ok := err.(awserr.Error); ok && awserr.Code() == iam.ErrCodeNoSuchEntityException {
			result, err := svc.CreateUser(&iam.CreateUserInput{
				UserName: &line,
			})

			if err != nil {
				fmt.Println("CreateUser Error", err)
				return
			}
			fmt.Println("Success", result)
		} else {
			fmt.Println("GetUser Error", err)
		}

	}

}

func main() {
	createIamUsers()
}
