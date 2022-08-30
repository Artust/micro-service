package routine_center

import (
	"avatar/services/gateway/config"
	upload "avatar/services/gateway/infra/upload/respository"
	"avatar/services/gateway/pkg/util"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
	"mime/multipart"

	"golang.org/x/sync/errgroup"
)

type CreateRoutineCenterInput struct {
	Name              string                `form:"name" binding:"required"`
	Detail            string                `form:"detail" binding:"required"`
	AnimationFile     *multipart.FileHeader `form:"animationFile" binding:"required"`
	ImageFile         *multipart.FileHeader `form:"imageFile"`
	SoundFile         *multipart.FileHeader `form:"soundFile" binding:"required"`
	StartDate         string                `form:"startDate" binding:"required"`
	EndDate           string                `form:"endDate" binding:"required"`
	CategoryId        int64                 `form:"categoryId" binding:"required"`
	Gender            int64                 `form:"gender" binding:"min=0,max=1"`
}
type CreateRoutineCenterOutput struct {
	Id                int64  `json:"id"`
	Name              string `json:"name"`
	Detail            string `json:"detail"`
	AnimationFile     string `json:"animationFile"`
	ImageFile         string `json:"imageFile"`
	SoundFile         string `json:"soundFile"`
	StartDate         string `json:"startDate"`
	EndDate           string `json:"endDate"`
	CategoryId        int64  `json:"categoryId"`
	Gender            int64  `json:"gender"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}

func CreateRoutineCenter(input *CreateRoutineCenterInput,
	centerClient pb.CenterClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) (*CreateRoutineCenterOutput, error) {
	var err error
	ctx := context.Background()
	eg, _ := errgroup.WithContext(ctx)
	animationFileName, _ := util.GenerateFileName(input.AnimationFile.Filename, input.Name)
	input.AnimationFile.Filename = animationFileName
	eg.Go(func() error {
		animationFileName, err = upload.UploadToS3(input.AnimationFile, config.RoutineBucketName)
		if err != nil {
			return err
		}
		return nil
	})
	soundFileName, _ := util.GenerateFileName(input.SoundFile.Filename, input.Name)
	input.SoundFile.Filename = soundFileName
	eg.Go(func() error {
		soundFileName, err = upload.UploadToS3(input.SoundFile, config.RoutineBucketName)
		if err != nil {
			return err
		}
		return nil
	})
	var imageFileName string
	if input.ImageFile != nil {
		imageFileName, _ = util.GenerateFileName(input.ImageFile.Filename, input.Name)
		input.ImageFile.Filename = imageFileName
		eg.Go(func() error {
			imageFileName, err = upload.UploadToS3(input.ImageFile, config.RoutineBucketName)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if eg.Wait() != nil {
		return nil, eg.Wait()
	}
	if imageFileName == "" {
		imageFileName = fmt.Sprintf("/%s/%s", config.RoutineBucketName, "defaultRoutineImage.png")
	}
	createRoutineCenterInput := &pb.CreateRoutineRequest{
		Name:              input.Name,
		Detail:            input.Detail,
		AnimationFile:     animationFileName,
		ImageFile:         imageFileName,
		SoundFile:         soundFileName,
		StartDate:         input.StartDate,
		EndDate:           input.EndDate,
		CategoryId:        input.CategoryId,
		Gender:            input.Gender,
	}
	response, err := centerClient.CreateRoutine(ctx, createRoutineCenterInput)
	if err != nil {
		return nil, err
	}
	createRoutineCenterOutput := &CreateRoutineCenterOutput{
		Id:                response.Id,
		Name:              response.Name,
		Detail:            response.Detail,
		AnimationFile:     fmt.Sprintf("%s%s", cfg.S3Uri, response.AnimationFile),
		ImageFile:         fmt.Sprintf("%s%s", cfg.S3Uri, response.ImageFile),
		SoundFile:         fmt.Sprintf("%s%s", cfg.S3Uri, response.SoundFile),
		StartDate:         response.StartDate,
		EndDate:           response.EndDate,
		Gender:            response.Gender,
		CategoryId:        response.CategoryId,
		CreatedAt:         response.CreatedAt,
		UpdatedAt:         response.UpdatedAt,
	}

	return createRoutineCenterOutput, nil
}
