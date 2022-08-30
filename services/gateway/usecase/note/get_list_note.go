package note

import (
	pb "avatar/services/gateway/protos/talk_session"
	"context"
)

type GetListNoteInput struct {
	TalkSessionId int64 `form:"talkSessionId"`
	Page          int64 `form:"page"`
	PerPage       int64 `form:"perPage"`
}

type GetListNoteOutput struct {
	Notes []*CreateNoteOutput `json:"results"`
}

func GetListNote(input *GetListNoteInput, talkSessionClient pb.TalkSessionClient) (*GetListNoteOutput, error) {
	data := &pb.GetListNoteRequest{
		TalkSessionId: int64(input.TalkSessionId),
		Page:          input.Page,
		PerPage:       input.PerPage,
	}
	response, err := talkSessionClient.GetListNote(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &GetListNoteOutput{}
	output.Notes = make([]*CreateNoteOutput, 0)
	for _, v := range response.Notes {
		note := &CreateNoteOutput{
			Id:            v.Id,
			TalkSessionId: v.TalkSessionId,
			Content:       v.Content,
			IsGuest:       &v.IsGuest,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		}
		output.Notes = append(output.Notes, note)
		note = nil
	}
	return output, nil
}
