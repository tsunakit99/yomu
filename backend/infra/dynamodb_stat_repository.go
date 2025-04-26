package infra

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/tsunakit99/yomu/domain/model"
	"github.com/tsunakit99/yomu/domain/repository"
)

type DynamoStatRepository struct {
	Client    *dynamodb.Client
	TableName string
}

func NewDynamoStatRepository() repository.StatRepository {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("failed to load AWS config: " + err.Error())
	}

	client := dynamodb.NewFromConfig(cfg)
	table := os.Getenv("DYNAMO_LIKE_TABLE")
	if table == "" {
		table = "article_stats"
	}

	return &DynamoStatRepository{
		Client:    client,
		TableName: table,
	}
}

func (r *DynamoStatRepository) GetStats(slug string) (*model.ArticleStat, error) {
	pvKey := map[string]types.AttributeValue{
		"slug": &types.AttributeValueMemberS{Value: slug},
		"type": &types.AttributeValueMemberS{Value: "pv"},
	}
	likeKey := map[string]types.AttributeValue{
		"slug": &types.AttributeValueMemberS{Value: slug},
		"type": &types.AttributeValueMemberS{Value: "like"},
	}

	pvItem, _ := r.Client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &r.TableName,
		Key:       pvKey,
	})

	likeItem, _ := r.Client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &r.TableName,
		Key:       likeKey,
	})

	pv := 0
	if pvAttr, ok := pvItem.Item["count"].(*types.AttributeValueMemberN); ok {
		fmt.Sscanf(pvAttr.Value, "%d", &pv)
	}

	like := 0
	if likeAttr, ok := likeItem.Item["count"].(*types.AttributeValueMemberN); ok {
		fmt.Sscanf(likeAttr.Value, "%d", &like)
	}

	return &model.ArticleStat{
		PV:   pv,
		Like: like,
	}, nil
}
