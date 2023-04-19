package repositories

import (
	"context"
	"fmt"

	"github.com/IgorRamos/fm-transaction/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const DEFAULT_LIMIT = 10

type TransactionRepository interface {
	CreateTransaction(models.Transaction) error
	GetTransactions(params QueryParameters) (page models.Page, err error)
}

type transactionRepository struct {
	db        *dynamodb.Client
	tableName string
}

type QueryParameters struct {
	Period           string
	LastEvaluatedKey string
	Limit            int32
}

func NewTransactionRepository(db *dynamodb.Client, tableName string) TransactionRepository {
	return transactionRepository{
		db:        db,
		tableName: tableName,
	}
}

func (r transactionRepository) CreateTransaction(transaction models.Transaction) error {
	transaction.TransactionId = uuid.New().String()
	transaction.CategoryTransactionId = fmt.Sprintf("%s#%s", transaction.Category, transaction.TransactionId)
	transactionDynamo, err := attributevalue.MarshalMap(transaction)
	if err != nil {
		log.Errorf("Failed to marshall new transaction item: %s", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      transactionDynamo,
	}

	_, err = r.db.PutItem(context.TODO(), input)
	if err != nil {
		log.Errorf("Failed to put new transaction item: %s", err)
		return err
	}

	return nil
}

func (r transactionRepository) GetTransactions(params QueryParameters) (page models.Page, err error) {
	var lastEvaluatedKey map[string]types.AttributeValue

	if params.LastEvaluatedKey != "" {
		lastEvaluatedKey, err = models.DecodeLastEvaluatedKey(params.LastEvaluatedKey)
		if err != nil {
			log.Error("Failed to decode last evaluated key, error: %s", err.Error())
			return models.Page{}, err
		}
	}

	if params.Limit == 0 {
		params.Limit = DEFAULT_LIMIT
	}

	keyConditionExpression := "MonthYear = :period"
	expressionAttributeValues := map[string]types.AttributeValue{
		":period": &types.AttributeValueMemberS{Value: params.Period},
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(r.tableName),
		Limit:                     aws.Int32(params.Limit),
		KeyConditionExpression:    &keyConditionExpression,
		ExpressionAttributeValues: expressionAttributeValues,
		ExclusiveStartKey:         lastEvaluatedKey,
	}
	queryOutput, err := r.db.Query(context.TODO(), queryInput)
	if err != nil {
		log.Error("Failed to scan transactions, error: %s", err.Error())
		return models.Page{}, err
	}

	transactions := make([]models.Transaction, queryOutput.Count)
	err = attributevalue.UnmarshalListOfMaps(queryOutput.Items, &transactions)
	if err != nil {
		log.Error("Failed to unmarshal transaction results, error: %s", err.Error())
		return models.Page{}, err
	}

	encodedLastEvaluateKey, err := models.EncodeLastEvaluateKey(queryOutput.LastEvaluatedKey)
	if err != nil {
		log.Error("Failed to encoded lastEvaluateKey, error: %s", err.Error())
		return models.Page{}, err
	}

	if encodedLastEvaluateKey == params.LastEvaluatedKey {
		encodedLastEvaluateKey = ""
	}

	page = models.Page{
		Results:          transactions,
		LastEvaluatedKey: encodedLastEvaluateKey,
	}
	return page, nil
}
