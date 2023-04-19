package models

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Page struct {
	Results          []Transaction `json:"results"`
	LastEvaluatedKey string        `json:"next"`
}

type NextKey struct {
	MonthYear             string
	CategoryTransactionId string
}

func DecodeLastEvaluatedKey(lastEvaluatedKey string) (map[string]types.AttributeValue, error) {
	if lastEvaluatedKey != "" {
		var nextKey NextKey
		err := decodeFromBase64(&nextKey, lastEvaluatedKey)
		if err != nil {
			return nil, err
		}

		decodedLastEvaluatedKey, err := attributevalue.MarshalMap(&nextKey)
		if err != nil {
			return nil, err
		}

		return decodedLastEvaluatedKey, nil
	}
	return nil, nil
}

func EncodeLastEvaluateKey(lastEvaluatedKey map[string]types.AttributeValue) (string, error) {
	if lastEvaluatedKey == nil {
		return "", nil
	}

	var nextKey NextKey
	err := attributevalue.UnmarshalMap(lastEvaluatedKey, &nextKey)
	if err != nil {
		return "", err
	}

	decodedLastEvaluatedKey, err := encodeToBase64(&nextKey)
	if err != nil {
		return "", err
	}

	return decodedLastEvaluatedKey, nil
}

func encodeToBase64(v interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return "", err
	}
	encoder.Close()
	return buf.String(), nil
}

func decodeFromBase64(v interface{}, enc string) error {
	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(v)
}
