package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func (c *Server) SaveToS3(tx neo4j.Transaction) error {
	records, err := tx.Run(`MATCH (rc:Routine) WHERE rc.name = $name AND rc.deletedAt="" RETURN count(*)`, map[string]interface{}{
		"name": "NameDefault",
	})
	if err != nil {
		log.Error("error when match routine category, error: ", err)
		return err
	}
	record, err := records.Single()
	if err != nil {
		log.Error("error when record routine category, error: ", err)
		return err
	}
	uploader := s3manager.NewUploader(c.s3Session)
	bucket := "routine"
	if record.Values[0].(int64) == 0 {
		svc := s3.New(c.s3Session)
		_, err = svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
			return err
		}
		imageData, err := os.Open(c.urlImageDefault)
		if err != nil {
			return err
		}
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(c.urlImageDefault),
			Body:   imageData,
		})
		if err != nil {
			return err
		}
		sound, err := os.Open(c.urlSound)
		if err != nil {
			return err
		}
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(c.urlSound),
			Body:   sound,
		})
		if err != nil {
			return err
		}
		animation, err := os.Open(c.urlAnimationKey)
		if err != nil {
			return err
		}
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(c.urlAnimationKey),
			Body:   animation,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
