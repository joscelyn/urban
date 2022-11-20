package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joscelyn/urban"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var listFlag bool
var whoisFlag string

var client = func() urban.UrbanClient {
	conn, err := grpc.Dial(
		"0.0.0.0:8000",
		grpc.WithInsecure(),
		grpc.WithTimeout(1*time.Second),
	)

	if err != nil {
		log.Fatal(err)
	}

	return urban.NewUrbanClient(conn)
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "client that access to the hacker news proxy server",
	Run: func(cmd *cobra.Command, args []string) {
		if !listFlag && whoisFlag == "" {
			log.Fatal("please use either --list flag or --whois flag")
		}

		if listFlag {
			ctx := context.Background()
			res, err := client().GetTopStories(ctx, &emptypb.Empty{})
			if err != nil {
				log.Print(err)
				os.Exit(-1)
			}

			for _, item := range res.Posts {
				fmt.Print("– ", item.Title, "\n   ", item.Url, "\n\n")
			}
		} else {
			ctx := context.Background()
			res, err := client().Whois(ctx, &urban.WhoisRequest{Username: whoisFlag})
			if err != nil {
				log.Print(err)
				os.Exit(-1)
			}

			fmt.Print("User: ", res.User, "\n", "Karma: ", res.Karma, "\n", "About: ", res.About, "\n", "Joined: ", res.Joined.AsTime().Format("2006-01-02"), "\n")
		}
	},
}

func init() {
	RootCmd.AddCommand(clientCmd)
	clientCmd.Flags().BoolVarP(&listFlag, "list", "l", false, "list of the latest news")
	clientCmd.Flags().StringVarP(&whoisFlag, "whois", "w", "", "get informations of an user")
}
