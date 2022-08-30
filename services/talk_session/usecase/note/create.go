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

func Create(
	ctx context.Context,
	db neo4j.Driver,
	noteRepository repository.NoteRepository,
	input *pb.CreateNoteRequest,
) (*pb.CreateNoteResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.Note{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	createdNote, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		note, err := noteRepository.Create(ctx, &data)
		if err != nil {
			fmt.Println("error when write transaction, error: ", err)
			return nil, err
		}
		return *note, nil
	})
	if err != nil {
		fmt.Println("error when write transaction, error: ", err)
		return nil, err
	}
	note := createdNote.(entity.Note)
	var result pb.CreateNoteResponse
	err = copier.Copy(&result, note)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = note.CreatedAt.Format(time.RFC3339)
	return &result, nil
}
