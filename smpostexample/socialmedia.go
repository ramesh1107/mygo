package smpostexample

import (
	"time"
)
{
	"time"
}
type MoodState int

const{
	MoodStateNuetral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopefull
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateCloudNine
	
}
type AuditableContent struct {
	TimeCreated time.Time
	TimeModified time.Time
	CreatedBy string
	ModifiedBy string
}

type Post struct {
	TimeCreated time.Time
	TimeModified time.Time
	Caption string
	MessageBody string
	URL string
	ImageUrI string
	ThumbnailURI string
	Keywords	[]string
	Likers	[]string
	AuthorMood	MoodState
}

var Moods map[string] MoodState

func init()  {
	Moods= map[string]MoodState{"nuetral":MoodStateNuetral,"happy":MoodStateHappy,"sad":MoodStateSad,
		"Angry":MoodStateAngry,"Hopeful":MoodStateHopefull,"Thrilled":MoodStateHopefull,
		"Bored":MoodStateBored,"shy":MoodStateShy,"comical":MoodStateComical,
		"cloudnine":MoodStateCloudNine}
	
}

func NewPost(username string, mood MoodState,caption string, messageBody string,url string, ImageUrI string, ThumbnailURI string,Keywords[]string)* Post{
	auditableContent:=AuditableContent{CreatedBy:username,TimeCreated:time.Now(),}
	return &Post{Caption: caption, MessageBody: messageBody, URL: url,ImageUrI:ImageUrI,ThumbnailURI:ThumbnailURI,AuthorMood:mood,Keywords:Keywords,AuditableContent:auditableContent
	}

	
}