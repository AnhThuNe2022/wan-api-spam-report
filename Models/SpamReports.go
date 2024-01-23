package Models

import (
	"time"
	"wan-api-spam-report/Const"
)

type SpamReports struct {
	SpamReportID        int64     `json:"spamReportID" gorm:"primaryKey; column:SpamReportID" `
	ReviewID            *int64    `json:"reviewID" gorm:"default:null;column:ReviewID"`
	CommentID           *int64    `json:"commentID" gorm:"default:null;column:CommentID"`
	CommentReplyID      *int64    `json:"commentReplyID" gorm:"default:null;column:CommentReplyID"`
	MovieID             *int64    `json:"movieID" gorm:"default:null;column:MovieID"`
	VideoID             *int64    `json:"videoID" gorm:"default:null;column:VideoID"`
	MusicID             *int64    `json:"musicID" gorm:"default:null;column:MusicID"`
	NewsID              *int64    `json:"newsID" gorm:"default:null; column:NewsID"`
	PhotoID             *int64    `json:"photoID" gorm:"default:null;column:PhotoID"`
	PostID              *int64    `json:"postID" gorm:"default:null;column:PostID"`
	ChannelProfileID    *int64    `json:"channelProfileID" gorm:"default:null;column:ChannelProfileID"`
	Name                string    `json:"name" gorm:"column:Name"`
	Reason              string    `json:"reason{" gorm:"column:Reason"`
	ReportUserProfileID int64     `json:"reportUserProfileID" gorm:"column:ReportUserProfileID"`
	IsDisable           bool      `json:"isDisable" gorm:"column:IsDisable"`
	IsSensitive         bool      `json:"isSensitive" gorm:"column:IsSensitive"`
	Active              bool      `json:"active" gorm:"column:Active"`
	CreatedBy           string    `json:"createdBy" gorm:"column:CreatedBy"`
	CreatedDate         time.Time `json:"createdDate" gorm:"column:CreatedDate"`
	ModifiedBy          string    `json:"modifiedBy" gorm:"column:ModifiedBy"`
	ModifiedDate        time.Time `json:"modifiedDate" gorm:"column:ModifiedDate"`
	// Comment             virtual   `json:"comment" gorm:"column:Comment"`
	// CommentReply        virtual   `json:"commentReply" gorm:"column:CommentReply"`
	// Review              virtual   `json:"review" gorm:"column:Review"`
	// Movie               virtual   `json:"movie" gorm:"column:Movie"`
	// Video               virtual   `json:"video" gorm:"column:Video"`
	// Music               virtual   `json:"music" gorm:"column:Music"`
	// News                virtual   `json:"news" gorm:"column:News"`
	// Photo               virtual   `json:"photo" gorm:"column:Photo"`
	// Post                virtual   `json:"post" gorm:"column:Post"`
	// ChannelProfile      virtual   `json:"channelProfile" gorm:"column:ChannelProfile"`
	// UserProfile         virtual   `json:"userProfile" gorm:"column:UserProfile"`
}

func (SpamReports) TableName() string {
	return Const.TABLE_SpamReports
}
