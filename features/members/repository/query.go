package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	members "eigen-backend-test-case/features/members"

	"github.com/jackc/pgx/v5/pgxpool"
)

type membersRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewMembersRepository(db *pgxpool.Pool, ctx context.Context) members.RepositoryInterface {
	return &membersRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *membersRepository) ListMembersWithBorrowedAmount() (res []members.ListOfMembers, err error) {
	query := `
	SELECT 
		m.*,
		COUNT(bb.*) AS borrowed_amount
	FROM members m
	LEFT JOIN borrowed_books bb ON m.id = bb.member_id AND bb.is_returned = FALSE
	GROUP BY m.id
	ORDER BY id ASC`

	rows, err := r.db.Query(r.ctx, query)
	if err != nil {
		errMsg := fmt.Errorf("error list members with borrowed amount")
		log.Printf("%v, err: %v\n", errMsg, err)
		return nil, errMsg
	}
	defer rows.Close()

	for rows.Next() {
		var temp members.ListOfMembers
		err := rows.Scan(&temp.Member.ID, &temp.Member.Code, &temp.Member.Name, &temp.BorrowedAmount)
		if err != nil {
			errMsg := fmt.Errorf("error scanning list members with borrowed amount")
			log.Printf("%v, err: %v\n", errMsg, err)
			return nil, errMsg
		}

		res = append(res, temp)
	}

	return
}

func (r *membersRepository) InsertListOfMembers(input []members.Member) (err error) {
	query := `INSERT INTO members(code, name) VALUES`

	var placeholder []string
	var arguments []interface{}

	for i, v := range input {
		params := fmt.Sprintf("($%d, $%d)", (i*2)+1, (i*2)+2)
		placeholder = append(placeholder, params)
		arguments = append(arguments, v.Code, v.Name)
	}

	queryPlaceholder := strings.Join(placeholder, ",")
	query += queryPlaceholder

	res, err := r.db.Exec(r.ctx, query, arguments...)
	if err != nil {
		errMsg := fmt.Errorf("error insert list of members")
		log.Printf("%v, err:%v\n", errMsg, err)
		return errMsg
	}

	affectedRows := res.RowsAffected()
	if affectedRows <= 0 {
		errMsg := fmt.Errorf("no members added to database for insert list of members")
		log.Println("no rows are affected for insert list of members")
		return errMsg
	}

	return err
}
