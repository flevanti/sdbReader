package main

import (
	"github.com/aws/aws-sdk-go/service/simpledb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	totalRecordsRead := 0
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION"))},
	)

	sdbObj := simpledb.New(sess)
	selectI := simpledb.SelectInput{
		ConsistentRead:   nil,
		NextToken:        nil,
		SelectExpression: aws.String("select itemName() from dev_entries limit 2500"),
	}

	for {
		selectO, err := sdbObj.Select(&selectI)

		if err != nil {
			fmt.Println(err)
			break
		} else {
			//fmt.Println(selectO)
			totalRecordsRead += len(selectO.Items)
			fmt.Println("records extracted: ", len(selectO.Items), " - Total records extracted: ", totalRecordsRead)
			if selectO.NextToken != nil {
				selectI.SetNextToken(*selectO.NextToken)
			} else {
				break
			}
		}
	} //end infinite loop

}
