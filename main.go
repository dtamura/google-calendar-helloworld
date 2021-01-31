package main

import (
	"context"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func main() {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Panic(err)
	}

	svc, err := calendar.New(client)
	if err != nil {
		log.Panic(err)
	}
	t := time.Now().Format(time.RFC3339)
	events, err := svc.Events.List(os.Getenv("GOOGLE_CALENDAR_ID")).ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	log.Println("Upcoming events:")
	if len(events.Items) == 0 {
		log.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			log.Printf("%v (%v)\n", item.Summary, date)
		}
	}
}
