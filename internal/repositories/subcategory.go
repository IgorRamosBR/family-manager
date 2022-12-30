package repositories

import (
	"context"

	"github.com/IgorRamos/fm-transaction/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	log "github.com/sirupsen/logrus"
)

type SubCategoryRepository interface {
	CreateSubCategory(models.SubCategory) error
	GetAllSubCategories() ([]models.SubCategory, error)
}

type subCategoryRepository struct {
	db        *dynamodb.Client
	tableName string
}

func NewSubCategoryRepository(db *dynamodb.Client, tableName string) SubCategoryRepository {
	return subCategoryRepository{
		db:        db,
		tableName: tableName,
	}
}

func (r subCategoryRepository) CreateSubCategory(subCategory models.SubCategory) error {
	subCategoryDynamo, err := attributevalue.MarshalMap(subCategory)
	if err != nil {
		log.Errorf("Failed to marshall new subcategory item: %s", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      subCategoryDynamo,
	}

	_, err = r.db.PutItem(context.TODO(), input)
	if err != nil {
		log.Errorf("Failed to put new subcategory item: %s", err)
		return err
	}

	return nil
}

func (r subCategoryRepository) GetAllSubCategories() ([]models.SubCategory, error) {
	scanInput := dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}

	scanOutput, err := r.db.Scan(context.TODO(), &scanInput)
	if err != nil {
		log.Error("Failed to get category, error: %s", err.Error())
		return []models.SubCategory{}, err
	}

	var subCategories []models.SubCategory
	err = attributevalue.UnmarshalListOfMaps(scanOutput.Items, &subCategories)
	if err != nil {
		log.Error("Failed to unmarshal subCategories, error: %s", err.Error())
		return []models.SubCategory{}, err
	}

	return subCategories, nil
}
