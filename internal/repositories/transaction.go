package repositories

import (
	"context"

	"github.com/IgorRamos/fm-create-transaction/internal/models"
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
	Offset           string
	Limit            int32
	lastEvaluatedKey map[string]types.AttributeValue
}
type paginateID struct {
	ID string `dynamodbav:"id" json:"id"`
}

func NewTransactionRepository(db *dynamodb.Client, tableName string) TransactionRepository {
	return transactionRepository{
		db:        db,
		tableName: tableName,
	}
}

func (r transactionRepository) CreateTransaction(transaction models.Transaction) error {
	transaction.TransactionId = uuid.New().String()
	transaction.CategorySubcategoryId = transaction.Category + transaction.Subcatetegory + transaction.TransactionId
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
	if params.Offset != "" {
		lastKey := paginateID{ID: params.Offset}
		key, err := attributevalue.MarshalMap(lastKey)
		if err != nil {
			log.Error("Failed to marshal last evaluated key, error: %s", err.Error())
			return models.Page{}, err
		}
		params.lastEvaluatedKey = key
	}

	if params.Limit == 0 {
		params.Limit = DEFAULT_LIMIT
	}

	scanInput := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		Limit:            aws.Int32(params.Limit),
		FilterExpression: aws.String("MonthYear = :period"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":period": &types.AttributeValueMemberS{Value: params.Period},
		},
		ExclusiveStartKey: params.lastEvaluatedKey,
	}
	scanOutput, err := r.db.Scan(context.TODO(), scanInput)
	if err != nil {
		log.Error("Failed to scan transactions, error: %s", err.Error())
		return models.Page{}, err
	}

	transactions := make([]models.Transaction, scanOutput.Count)
	err = attributevalue.UnmarshalListOfMaps(scanOutput.Items, &transactions)
	if err != nil {
		log.Error("Failed to unmarshal transaction results, error: %s", err.Error())
		return models.Page{}, err
	}

	nextKey := paginateID{}
	err = attributevalue.UnmarshalMap(scanOutput.LastEvaluatedKey, &nextKey)
	if err != nil {
		log.Error("Failed to unmarshal next key, error: %s", err.Error())
		return models.Page{}, err
	}

	page = models.Page{
		Results: transactions,
		Next:    nextKey.ID,
	}
	return page, nil
}
