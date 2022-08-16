package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	to_do "github.com/tiquiet/to-do"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item to_do.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createUsersItemQuery, listId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userId, listId int) ([]to_do.TodoItem, error) {
	var items []to_do.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti JOIN %s li ON ti.id=li.item_id 
									JOIN %s ul ON li.list_id=ul.list_id WHERE ul.list_id=$1 AND ul.user_id=$2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *TodoItemPostgres) GetItemById(userId, itemId int) (to_do.TodoItem, error) {
	var item to_do.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti JOIN %s li ON ti.id=li.item_id
								JOIN %s ul ON ul.list_id=li.list_id WHERE ul.user_id=$1 AND ti.id=$2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, userId, itemId); err != nil {
		return to_do.TodoItem{}, err
	}

	return item, nil
}

func (r *TodoItemPostgres) DeleteItem(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul WHERE ti.id=li.item_id AND li.list_id=ul.list_id AND
									ul.user_id=$1 AND ti.id=$2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	if _, err := r.db.Exec(query, userId, itemId); err != nil {
		return err
	}

	return nil
}

func (r *TodoItemPostgres) UpdateItem(userId, itemId int, input to_do.UpdateListItemInput) error {
	setValues := make([]string, 0, 3)
	args := make([]interface{}, 0, 3)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2, done=$3
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND 
									ti.id=$%d AND ul.user_id=$%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, itemId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
