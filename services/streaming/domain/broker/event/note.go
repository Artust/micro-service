package event

type NoteEvent int64

const (
	NoteEventCreate NoteEvent = 1
	NoteEventUpdate NoteEvent = 2
	NoteEventDelete NoteEvent = 3
)
