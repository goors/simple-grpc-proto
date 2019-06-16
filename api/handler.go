package api
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"golang.org/x/net/context"
	"strconv"
)
// Server represents the gRPC server
type Server struct {
}

func (s *Server) Set(ctx context.Context, c *GetRequest) (*Response, error) {

	adId := strconv.FormatInt(c.GetAdvertId(), 16)

	var (
		title string
		description string
		first_name string
		last_name string
		email string
		user_title string
	)
	dbConf := GetConfig();

	db, err := sql.Open("mysql", dbConf.Username + ":" + dbConf.Password + "@tcp(" + dbConf.Host + ":3306)/" + dbConf.Database)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select a.title as title, a.description as description, u.first_name as first_name, u.last_name as last_name, u.email as email, u.title as user_title from ad a left join user u on a.u = u.id where a.id = ?", adId)
	if err != nil {
		log.Fatal(err)
	}

	answer := DefaultAnswer{}
	answer.Message = "We got your ad."
	ad := Ad{}

	for rows.Next() {
		err := rows.Scan(&title, &description, &first_name, &last_name, &email, &user_title)
		if err != nil {
			log.Fatal(err)
		}
		ad.Description = description
		ad.Title = title
		ad.User = &User{};
		ad.User.Title = user_title
		ad.User.Email = email
		ad.User.FirstName = first_name
		ad.User.LastName = last_name
	}

	defer db.Close()
	return &Response{Answer: &answer, Ad: &ad}, nil

}
