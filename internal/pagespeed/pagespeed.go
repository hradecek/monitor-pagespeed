package main

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/pagespeedonline/v5"
	"log"
)

func main() {
	service, err := pagespeedonline.NewService(context.Background(), option.WithAPIKey(API_KEY))
	if err != nil {
		log.Fatal(err)
	}

	pagespeedService := pagespeedonline.NewPagespeedapiService(service)
	pagespeedCall := pagespeedService.Runpagespeed().Url(URL)

	pagespeedResponse, err := pagespeedCall.Do()
	if err != nil {
		log.Fatal(err)
	}
	print(pagespeedResponse)
}

