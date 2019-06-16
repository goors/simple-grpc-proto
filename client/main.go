package main
import (
	"log"
	"service/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)
func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewDefaultAnswerServiceClient(conn);

	res, err := c.Set(context.Background(), &api.GetRequest{AdvertId: 2})
	if err != nil {
		log.Fatalf("Error when calling GetRequest: %s", err)
	}

	log.Println("Response from server: ", res.Ad)


}
