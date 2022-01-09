package speedtrackertypes

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetDatabaseClient() (dbClient *dynamodb.Client) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "eu-west-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}

func IsExistingUser(dbClient *dynamodb.Client, newUserEmailAddress string) (isExistingUser bool) {
	out, err := dbClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String("sst-user-data-4aace0e"),
		KeyConditionExpression: aws.String("#DDB_PK = :pkey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pkey": &types.AttributeValueMemberS{Value: fmt.Sprintf("USER#%s", newUserEmailAddress)},
		},
		ExpressionAttributeNames: map[string]string{
			"#DDB_PK": "PK",
		},
		ScanIndexForward: aws.Bool(true),
	})

	if out.Count != 0 {
		return true
	}

	if err != nil {
		panic(err)
	}

	return false
}

func CreateAttributeValueSliceFromPersonalBestSlice(inputMap []PersonalBest) (attributeValueSlice []types.AttributeValue) {
	var returnVal []types.AttributeValue

	for i := 0; i < len(inputMap); i++ {
		newObj, _ := attributevalue.Marshal(inputMap[i])
		returnVal = append(returnVal, newObj)
	}

	return returnVal
}

func RecordPBInHistory(dbClient *dynamodb.Client, userEmailAddress string, pb PersonalBest) {
	pbAttributeValue, marshallingErr := attributevalue.MarshalMap(pb)

	if marshallingErr != nil {
		log.Fatalf("Marshalling Err in RecordReplacedPB: %s", marshallingErr.Error())
	}

	_, err := dbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("sst-user-data-4aace0e"),
		Item: map[string]types.AttributeValue{
			"PK":           &types.AttributeValueMemberS{Value: fmt.Sprintf("USER#%s", userEmailAddress)},
			"SK":           &types.AttributeValueMemberS{Value: fmt.Sprintf("PB#%s#%s#%s#%s", strings.ToUpper(pb.Swing.Colour), strings.ToUpper(pb.Swing.Position), strings.ToUpper(pb.Swing.Side), pb.Date.Format(time.RFC3339Nano))},
			"speed":        &types.AttributeValueMemberN{Value: strconv.Itoa(pb.Swing.Speed)},
			"personalBest": &types.AttributeValueMemberM{Value: pbAttributeValue},
		},
	})

	if err != nil {
		log.Panic(err)
		panic(err)
	}
}

func UpdateUserCurrentPBs(dbClient *dynamodb.Client, userEmailAddress string, pbs []PersonalBest) {
	pbSlice := CreateAttributeValueSliceFromPersonalBestSlice(pbs)

	_, err := dbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("sst-user-data-4aace0e"),
		Item: map[string]types.AttributeValue{
			"PK":            &types.AttributeValueMemberS{Value: fmt.Sprintf("USER#%s", userEmailAddress)},
			"SK":            &types.AttributeValueMemberS{Value: "CURRENTPBS"},
			"personalBests": &types.AttributeValueMemberL{Value: pbSlice},
		},
	})

	if err != nil {
		log.Panic(err)
		panic(err)
	}
}

func GetUserCurrentPBs(dbClient *dynamodb.Client, userEmailAddress string) (userSessions []PersonalBest) {
	out, err := dbClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String("sst-user-data-4aace0e"),
		KeyConditionExpression: aws.String("#DDB_PK = :pkey and #DDB_SK = :skey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pkey": &types.AttributeValueMemberS{Value: fmt.Sprintf("USER#%s", userEmailAddress)},
			":skey": &types.AttributeValueMemberS{Value: "CURRENTPBS"},
		},
		ExpressionAttributeNames: map[string]string{
			"#DDB_PK": "PK",
			"#DDB_SK": "SK",
		},
		ScanIndexForward: aws.Bool(true),
	})

	if err != nil {
		log.Fatalf("User current pbs database query error: %s", err.Error())
		panic(err)
	}

	if out.Count == 0 {
		return nil
	}

	queryResult := out.Items[0]
	existingPersonalBests := queryResult["personalBests"]

	retrievedPBs := []PersonalBest{}
	unmarshalError := attributevalue.Unmarshal(existingPersonalBests, &retrievedPBs)

	if unmarshalError != nil {
		log.Fatalf("User current pbs unmarshalling error: %s", unmarshalError.Error())
	}

	return retrievedPBs
}
