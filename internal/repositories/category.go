package repositories

import (
	"context"

	"github.com/IgorRamos/fm-transaction/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	log "github.com/sirupsen/logrus"
)

type CategoryRepository interface {
	CreateCategory(models.Category) error
	GetCategoryByName(string) (models.Category, error)
	GetAllCategories() ([]models.Category, error)
}

type categoryRepository struct {
	db        *dynamodb.Client
	tableName string
}

func NewCategoryRepository(db *dynamodb.Client, tableName string) CategoryRepository {
	return categoryRepository{
		db:        db,
		tableName: tableName,
	}
}

func (r categoryRepository) CreateCategory(category models.Category) error {
	categoryDynamo, err := attributevalue.MarshalMap(category)
	if err != nil {
		log.Errorf("Failed to marshall new transaction item: %s", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      categoryDynamo,
	}

	_, err = r.db.PutItem(context.TODO(), input)
	if err != nil {
		log.Errorf("Failed to put new category item: %s", err)
		return err
	}

	return nil
}

func (r categoryRepository) GetCategoryByName(categoryName string) (models.Category, error) {
	categoryKey, err := attributevalue.Marshal(categoryName)
	if err != nil {
		log.Errorf("Failed to marshall categoryName: %s", err)
		return models.Category{}, err
	}

	key := map[string]types.AttributeValue{"categoryName": categoryKey}
	getItemInput := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(r.tableName),
	}

	getItemOutput, err := r.db.GetItem(context.TODO(), getItemInput)
	if err != nil {
		log.Error("Failed to get category, error: %s", err.Error())
		return models.Category{}, err
	}

	var category models.Category
	err = attributevalue.UnmarshalMap(getItemOutput.Item, &category)
	if err != nil {
		log.Error("Failed to unmarshal category, error: %s", err.Error())
		return models.Category{}, err
	}

	return category, nil
}

func (r categoryRepository) GetAllCategories() ([]models.Category, error) {
	scanInput := dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}

	scanOutput, err := r.db.Scan(context.TODO(), &scanInput)
	if err != nil {
		log.Error("Failed to get all categories, error: %s", err.Error())
		return []models.Category{}, err
	}

	var categories []models.Category
	err = attributevalue.UnmarshalListOfMaps(scanOutput.Items, &categories)
	if err != nil {
		log.Error("Failed to unmarshal categories, error: %s", err.Error())
		return []models.Category{}, err
	}

	return categories, nil
}
