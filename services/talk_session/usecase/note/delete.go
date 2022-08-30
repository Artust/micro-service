package note

import (
	"avatar/services/talk_session/config"
	"avatar/services/talk_session/domain/repository"
	pb "avatar/services/talk_session/protos"
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Delete(
	ctx context.Context,
	db neo4j.Driver,
	noteRepository repository.NoteRepository,
	input *pb.DeleteNoteRequest,
) (*pb.DeleteNoteResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	var talkSessionId int64
	rowsAffectedRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		note, err := noteRepository.GetById(ctx, input.Id)
		if err != nil {
			fmt.Println("error when write transaction, error: ", err)
			return nil, err
		}
		talkSessionId = note.TalkSessionId
		rowsAffected, err := noteRepository.Delete(ctx, input.Id)
		if err != nil {
			fmt.Println("error when write transaction, error: ", err)
			return nil, err
		}
		return rowsAffected, nil
	})
	if err != nil {
		fmt.Println("error when write transaction, error: ", err)
		return nil, err
	}
	return &pb.DeleteNoteResponse{
		RowsAffected:  rowsAffectedRaw.(int64),
		TalkSessionId: talkSessionId,
	}, nil
}
