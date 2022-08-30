package note

import (
	"avatar/services/talk_session/config"
	"avatar/services/talk_session/domain/entity"
	"avatar/services/talk_session/domain/repository"
	pb "avatar/services/talk_session/protos"
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetList(
	ctx context.Context,
	db neo4j.Driver,
	noteRepository repository.NoteRepository,
	input *pb.GetListNoteRequest,
) (*pb.GetListNoteResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListNoteOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	notesRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		notes, err := noteRepository.GetList(ctx, data)
		if err != nil {
			fmt.Println("error when write transaction, error: ", err)
			return nil, err
		}
		return notes, nil
	})
	if err != nil {
		fmt.Println("error when write transaction, error: ", err)
		return nil, err
	}
	notes := notesRaw.([]*entity.Note)
	var results pb.GetListNoteResponse
	results.Notes = make([]*pb.CreateNoteResponse, 0)
	for _, note := range notes {
		var response pb.CreateNoteResponse
		err = copier.Copy(&response, note)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = note.CreatedAt.Format(time.RFC3339)
		if note.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = note.UpdatedAt.Format(time.RFC3339)
		}
		results.Notes = append(results.Notes, &response)
	}
	return &results, nil
}
