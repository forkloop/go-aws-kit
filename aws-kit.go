package main

import (
    "fmt"
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)

func retrieveQueueUrl(queueName string) (string) {
    svc := sqs.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

    res, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
        QueueName: aws.String(queueName),
    })

    if err != nil {
        panic(err)
    }

    return *res.QueueUrl
}

func main() {
    queueName := os.Args[1]

    svc := sqs.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

    queueUrl := retrieveQueueUrl(queueName)

    res, err := svc.GetQueueAttributes(&sqs.GetQueueAttributesInput{
        QueueUrl: aws.String(queueUrl),
        AttributeNames: []*string{
            //aws.String("ApproximateNumberOfMessages"),
            aws.String("All"),
        },
    })


    if err == nil {
        fmt.Println(res)
    } else {
        fmt.Println(err.Error())
    }
}

