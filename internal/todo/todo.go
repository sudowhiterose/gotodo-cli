package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type Store struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

func defaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".gotodo.json"), nil
}

func Load() (*Store, error) {
	path, err := defaultPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Store{NextID: 1}, nil
		}
		return nil, err
	}

	var s Store
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл задач: %w", err)
	}
	return &s, nil
}

func (s *Store) Save() error {
	path, err := defaultPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

func (s *Store) Add(text string) Task {
	task := Task{
		ID:        s.NextID,
		Text:      text,
		Done:      false,
		CreatedAt: time.Now(),
	}
	s.Tasks = append(s.Tasks, task)
	s.NextID++
	return task
}

func (s *Store) Find(id int) (*Task, error) {
	for i := range s.Tasks {
		if s.Tasks[i].ID == id {
			return &s.Tasks[i], nil
		}
	}
	return nil, errors.New("задача не найдена")
}

func (s *Store) Delete(id int) error {
	for i, t := range s.Tasks {
		if t.ID == id {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("задача не найдена")
}

func (s *Store) ClearDone() int {
	var remaining []Task
	count := 0
	for _, t := range s.Tasks {
		if t.Done {
			count++
			continue
		}
		remaining = append(remaining, t)
	}
	s.Tasks = remaining
	return count
}
