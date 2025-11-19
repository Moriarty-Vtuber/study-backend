package youtube

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/yourname/my-study-space/internal/db"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func StartMonitor(apiKey, videoID string, database *db.Client) {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Printf("Error creating YouTube service: %v", err)
		return
	}

	// Get Chat ID from Video ID
	call := service.Videos.List([]string{"liveStreamingDetails"}).Id(videoID)
	response, err := call.Do()
	if err != nil || len(response.Items) == 0 {
		log.Printf("Error finding video or no live details: %v", err)
		return
	}
	
	if response.Items[0].LiveStreamingDetails == nil {
		log.Printf("Video has no live streaming details")
		return
	}
	
	liveChatID := response.Items[0].LiveStreamingDetails.ActiveLiveChatId
    
    nextPageToken := ""

	for {
		// Poll Chat
		chatCall := service.LiveChatMessages.List(liveChatID, []string{"snippet", "authorDetails"})
        if nextPageToken != "" {
            chatCall.PageToken(nextPageToken)
        }
		chatResponse, err := chatCall.Do()
		if err != nil {
			log.Printf("Error polling chat: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

        nextPageToken = chatResponse.NextPageToken

		for _, msg := range chatResponse.Items {
			message := msg.Snippet.DisplayMessage
			userID := msg.AuthorDetails.ChannelId
			userName := msg.AuthorDetails.DisplayName
            userImage := msg.AuthorDetails.ProfileImageUrl

			// Simple Command Parsing
			if strings.Contains(message, "!in") || strings.Contains(message, "study") {
				database.EnterRoom(userID, userName, userImage)
			} else if strings.Contains(message, "!out") {
				database.ExitRoom(userID)
			}
		}

		// Wait before next poll (respect quota)
        // YouTube suggests polling interval from response, usually ~5-10s
		time.Sleep(time.Duration(chatResponse.PollingIntervalMillis) * time.Millisecond)
	}
}
