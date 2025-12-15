package service

import (
	"errors"

	"example.com/notes-api/internal/core"
	"example.com/notes-api/internal/repo"
)

type NoteService struct {
	repo *repo.NoteRepoMem
}

func NewNoteService(repo *repo.NoteRepoMem) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) CreateNote(title, content string) (*core.Note, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if len(title) > 100 {
		return nil, errors.New("title too long")
	}
	if len(content) > 5000 {
		return nil, errors.New("content too long")
	}

	note := core.Note{
		Title:   title,
		Content: content,
	}

	id, err := s.repo.Create(note)
	if err != nil {
		return nil, err
	}

	return s.repo.GetByID(id)
}

func (s *NoteService) GetAllNotes() ([]core.Note, error) {
	return s.repo.GetAll()
}

func (s *NoteService) GetNoteByID(id int64) (*core.Note, error) {
	return s.repo.GetByID(id)
}

func (s *NoteService) UpdateNote(id int64, title, content string) (*core.Note, error) {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err 
	}

	if title != "" && len(title) > 100 {
		return nil, errors.New("title too long")
	}
	if content != "" && len(content) > 5000 {
		return nil, errors.New("content too long")
	}

	updates := core.Note{}
	if title != "" {
		updates.Title = title
	}
	if content != "" {
		updates.Content = content
	}

	err = s.repo.Update(id, updates)
	if err != nil {
		return nil, err
	}

	return s.repo.GetByID(id)
}

func (s *NoteService) DeleteNote(id int64) error {
	return s.repo.Delete(id)
}