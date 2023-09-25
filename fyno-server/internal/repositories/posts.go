package repositories

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"fyno/server/internal/models"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type PostImages models.PostImages

func (pi *PostImages) Scan(src interface{}) error {
	fmt.Println("Scanning")
	if src == nil {
		return errors.New("PostImages Scan: null value")
	}

	// The src value returned by pq is of type []byte, so we need to convert it to a string
	srcStr, ok := src.([]byte)
	if !ok {
		return errors.New("PostImages Scan: invalid type")
	}
	fmt.Println("srcStr", srcStr)
	srcStr = bytes.Trim(srcStr, "{}")
	fmt.Println("srcStr", srcStr)
	// Split the string into URL and rank using the separator ","
	parts := strings.Split(string(srcStr), ",")
	if len(parts) != 2 {
		return errors.New("PostImages Scan: invalid format")
	}

	// Parse the URL and rank values from the string parts
	url := strings.Trim(parts[0], "\"()")
	fmt.Println("url", url)

	rank, err := strconv.Atoi(strings.Trim(parts[1], "\"()"))
	fmt.Println("rank", rank)
	if err != nil {
		fmt.Println("error", err)
		return errors.New("PostImages Scan: invalid rank value")
	}

	// Set the values of the PostImages struct
	pi.Url = url
	pi.Rank = rank

	return nil
}

type postRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) models.PostRepository {
	return &postRepository{
		DB: db,
	}
}

func (pr *postRepository) GetAll(userID uuid.UUID) ([]*models.Post, error) {
	query := `SELECT p.id, p.user_id, p.kind, p.name, p.age, p.gender, p.content, 
					l.id AS location_id, l.name AS location_name, 
					c.id AS category_id, c.name AS category_name, 
					p.created_at,
					ARRAY_AGG((pi.url, pi.rank)) AS post_images
				FROM posts AS p
				JOIN locations AS l ON p.location_id = l.id 
				JOIN categories As c ON p.category_id = c.id
				LEFT JOIN post_images AS pi ON p.id = pi.post_id`

	var args []interface{}

	// if userID is not nil, add a WHERE clause to the query
	if userID != uuid.Nil {
		query += " WHERE p.user_id = $1"
		args = append(args, userID.String())
	}

	query += " GROUP BY p.id, l.id, c.id"
	query += " ORDER BY p.created_at DESC"

	rows, err := pr.DB.Query(query, args...)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var p models.Post
		var postImages []PostImages
		err := rows.Scan(&p.ID, &p.UserID, &p.Kind, &p.Name, &p.Age, &p.Gender, &p.Content, &p.Location.ID, &p.Location.Name, &p.Category.ID, &p.Category.Name, &p.CreatedAt, pq.Array(&postImages))
		if err != nil {
			fmt.Println("error", err)
			return nil, err
		}
		p.PostImages = make([]models.PostImages, len(postImages))
		for i, pi := range postImages {
			p.PostImages[i] = models.PostImages(pi)
		}
		posts = append(posts, &p)
	}
	fmt.Println("posts", posts)
	return posts, nil
}

func (pr *postRepository) Get(id uuid.UUID) (*models.Post, error) {
	fmt.Println("id", id)
	query := `SELECT u.name, u.avatar_url, 
					p.id, p.user_id, p.kind, p.name, p.age, p.gender, p.content, p.created_at, 
					l.id AS location_id, l.name AS location_name, 
					c.id AS category_id, c.name AS category_name, 
					ARRAY_AGG((pi.url, pi.rank)) AS post_images
				FROM posts AS p
				JOIN users AS u ON p.user_id = u.id
				JOIN locations AS l ON p.location_id = l.id 
				JOIN categories As c ON p.category_id = c.id
				LEFT JOIN post_images AS pi ON p.id = pi.post_id
				WHERE p.id = $1
				GROUP BY p.id, l.id, c.id, u.name, u.avatar_url`

	row := pr.DB.QueryRow(query, id)

	var p models.Post
	var postImages []PostImages
	err := row.Scan(&p.Username, &p.AvatarURL, &p.ID, &p.UserID, &p.Kind, &p.Name, &p.Age, &p.Gender, &p.Content, &p.CreatedAt, &p.Location.ID, &p.Location.Name, &p.Category.ID, &p.Category.Name, pq.Array(&postImages))
	if err != nil {
		fmt.Println("error", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println("postImages", postImages)
	fmt.Println("p", p)
	p.PostImages = make([]models.PostImages, len(postImages))
	for i, pi := range postImages {
		p.PostImages[i] = models.PostImages(pi)
	}
	return &p, nil
}

func (pr *postRepository) Create(p *models.Post) (uuid.UUID, error) {
	query := `INSERT INTO posts (id, user_id, kind, name, age, gender, content, location_id, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := pr.DB.Exec(query, p.ID, p.UserID, p.Kind, p.Name, p.Age, p.Gender, p.Content, p.Location.ID, p.Category.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return p.ID, nil
}

func (pr *postRepository) CreatePostImage(id uuid.UUID, p models.PostImages, postID uuid.UUID) error {
	fmt.Println("postID", postID)
	query := `INSERT INTO post_images (id, url, rank, post_id) VALUES ($1, $2, $3, $4)`
	_, err := pr.DB.Exec(query, id, p.Url, p.Rank, postID)
	if err != nil {
		return err
	}

	return nil
}

func (pr *postRepository) DeleteAll() error {
	query := `DELETE FROM posts`
	_, err := pr.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
