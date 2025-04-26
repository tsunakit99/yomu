package infra

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/tsunakit99/yomu/domain/repository"
)

type DynamoLikeRepository struct {
	Client    *dynamodb.Client
	TableName string
}

func NewDynamoLikeRepository() repository.LikeRepository {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("failed to load AWS config: " + err.Error())
	}

	client := dynamodb.NewFromConfig(cfg)
	table := os.Getenv("DYNAMO_LIKE_TABLE")
	if table == "" {
		table = "article_stats" // fallback
	}

	return &DynamoLikeRepository{
		Client:    client,
		TableName: table,
	}
}

func (r *DynamoLikeRepository) IncrementLike(slug string) error {
	_, err := r.Client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: &r.TableName,
		Key: map[string]types.AttributeValue{
			"slug": &types.AttributeValueMemberS{Value: slug},
			"type": &types.AttributeValueMemberS{Value: "like"},
		},
		UpdateExpression: aws.String("ADD #count :inc"),
		ExpressionAttributeNames: map[string]string{
			"#count": "count",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":inc": &types.AttributeValueMemberN{Value: "1"},
		},
		ReturnValues: types.ReturnValueUpdatedNew,
	})
	return err
}
