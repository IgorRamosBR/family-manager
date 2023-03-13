package repositories

import (
	"context"
	"sort"
	"strconv"

	"github.com/IgorRamos/fm-transaction/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	log "github.com/sirupsen/logrus"
)

type CategoryRepository interface {
	CreateCategory(models.Category) error
	GetCategory(string, int) (models.Category, error)
	GetAllCategories() ([]models.Category, error)
	UpdateCategoryListOrder([]models.Category) error
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
		log.Errorf("Failed to marshall new category item: %s", err)
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

func (r categoryRepository) GetCategory(categoryName string, priority int) (models.Category, error) {
	categoryKey, err := attributevalue.Marshal(categoryName)
	if err != nil {
		log.Errorf("Failed to marshall Name: %s", err)
		return models.Category{}, err
	}

	priorityKey, err := attributevalue.Marshal(priority)
	if err != nil {
		log.Errorf("Failed to marshall Priority: %s", err)
		return models.Category{}, err
	}

	key := map[string]types.AttributeValue{"Name": categoryKey, "Priority": priorityKey}
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

	sort.SliceStable(categories, func(i, j int) bool {
		return categories[i].Priority < categories[j].Priority
	})

	return categories, nil
}

func (r categoryRepository) UpdateCategoryListOrder(categories []models.Category) error {
	err := r.RemoveAll()
	if err != nil {
		return err
	}

	var writeRequests []types.WriteRequest
	for _, category := range categories {
		categoryDynamo, err := attributevalue.MarshalMap(category)
		if err != nil {
			log.Errorf("Failed to marshall new category item: %s", err)
			return err
		}

		writeRequests = append(writeRequests, types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: categoryDynamo,
			},
		})
	}

	_, err = r.db.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			r.tableName: writeRequests,
		},
	})
	if err != nil {
		log.Error("Failed to update category order, error: %s", err.Error())
		return err
	}
	return nil
}

func (r categoryRepository) RemoveAll() error {
	categories, err := r.GetAllCategories()
	if err != nil {
		log.Error("Failed to get all categories, error: %s", err.Error())
		return err
	}

	var writeRequests []types.WriteRequest
	for _, category := range categories {
		writeRequests = append(writeRequests, types.WriteRequest{
			DeleteRequest: &types.DeleteRequest{
				Key: map[string]types.AttributeValue{
					"Name":     &types.AttributeValueMemberS{Value: category.Name},
					"Priority": &types.AttributeValueMemberN{Value: strconv.FormatInt(int64(category.Priority), 10)},
				},
			},
		})
	}

	_, err = r.db.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			r.tableName: writeRequests,
		},
	})
	if err != nil {
		log.Error("Failed to remove all categories, error: %s", err.Error())
		return err
	}
	return nil
}
