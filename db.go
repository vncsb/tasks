package tasksdb

import (
	"encoding/binary"
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
)

type Task struct {
	ID          int
	Description string
	Done        bool
}

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func AddTask(description string) error {
	db, err := setupDB()
	if err != nil {
		return errors.New("Could not open database file")
	}

	task := Task{
		Description: description,
		Done:        false,
	}

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		id64, _ := b.NextSequence()
		id := int(id64)

		buf, err := json.Marshal(task)
		if err != nil {
			return errors.New("Could not encode task")
		}

		return b.Put(itob(id), buf)
	})
}

func DoTask(id int) error {
	db, err := setupDB()
	if err != nil {
		return errors.New("Could not open database file")
	}

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		v := b.Get(itob(id))

		if v == nil {
			return errors.New("Task not found")
		}

		var task Task
		err := json.Unmarshal(v, task)
		if err != nil {
			return errors.New("Could not decode task")
		}

		task.Done = true

		v, err = json.Marshal(task)
		if err != nil {
			return errors.New("Could not encode task")
		}

		return b.Put(itob(task.ID), v)
	})
}

func ListTasks() ([]Task, error) {
	db, err := setupDB()
	if err != nil {
		return nil, errors.New("Could not open database file")
	}

	tasks := make([]Task, 0)

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		return b.ForEach(func(k, v []byte) error {
			var task Task
			err := json.Unmarshal(v, task)
			if err != nil {
				return errors.New("Could note decode task")
			}

			if !task.Done {
				tasks = append(tasks, task)
			}

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
