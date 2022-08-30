package handler

import (
	pb "avatar/services/talk_session/protos"
	"avatar/services/talk_session/usecase/note"
	"context"
)

func (s *Server) CreateNote(ctx context.Context, input *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	return note.Create(ctx, s.neo4jDriver, s.noteRepository, input)
}

func (s *Server) UpdateNote(ctx context.Context, input *pb.UpdateNoteRequest) (*pb.CreateNoteResponse, error) {
	return note.Update(ctx, s.neo4jDriver, s.noteRepository, input)
}

func (s *Server) DeleteNote(ctx context.Context, input *pb.DeleteNoteRequest) (*pb.DeleteNoteResponse, error) {
	return note.Delete(ctx, s.neo4jDriver, s.noteRepository, input)
}

func (s *Server) GetListNote(ctx context.Context, input *pb.GetListNoteRequest) (*pb.GetListNoteResponse, error) {
	return note.GetList(ctx, s.neo4jDriver, s.noteRepository, input)
}
