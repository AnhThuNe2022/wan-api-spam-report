package DTO

import "time"

type SpamReportsDTO struct {
	SpamReportID            int64     `json:"spamReportID" gorm:"column:SpamReportID"`
	ReviewID                *int64    `json:"reviewID" gorm:"column:ReviewID"`
	CommentID               *int64    `json:"commentID" gorm:"column:CommentID"`
	CommentReplyID          *int64    `json:"commentReplyID" gorm:"column:CommentReplyID"`
	Name                    string    `json:"name" gorm:"column:Name"`
	Active                  bool      `json:"active" gorm:"column:Active"`
	CreatedBy               string    `json:"createdBy" gorm:"column:CreatedBy"`
	CreatedDate             time.Time `json:"createdDate" gorm:"column:CreatedDate"`
	CreatedDateS            string    `json:"createdDateS" gorm:"column:CreatedDateS"`
	ModifiedBy              string    `json:"modifiedBy" gorm:"column:ModifiedBy"`
	ModifiedDate            time.Time `json:"modifiedDate" gorm:"column:ModifiedDate"`
	ModifiedDateS           string    `json:"modifiedDateS" gorm:"column:ModifiedDateS"`
	MovieID                 *int64    `json:"movieID" gorm:"column:MovieID"`
	VideoID                 *int64    `json:"videoID" gorm:"column:VideoID"`
	MusicID                 *int64    `json:"musicID" gorm:"column:MusicID"`
	NewsID                  *int64    `json:"newsID" gorm:"column:NewsID"`
	Reason                  string    `json:"reason" gorm:"column:Reason"`
	ReportUserProfileID     int64     `json:"reportUserProfileID" gorm:"column:ReportUserProfileID"`
	PhotoID                 *int64    `json:"photoID" gorm:"column:PhotoID"`
	ChannelProfileID        *int64    `json:"channelProfileID" gorm:"column:ChannelProfileID"`
	IsDisable               bool      `json:"isDisable" gorm:"column:IsDisable"`
	EditMode                string    `json:"editMode" gorm:"column:EditMode"`
	Guid                    string    `json:"guid" gorm:"column:Guid"`
	ReviewName              string    `json:"reviewName" gorm:"column:ReviewName"`
	CommentName             string    `json:"commentName" gorm:"column:CommentName"`
	Content                 string    `json:"content" gorm:"column:Content"`
	CommentReplyName        string    `json:"commentReplyName" gorm:"column:CommentReplyName"`
	UserName                string    `json:"userName" gorm:"column:UserName"`
	ReportedUserProfileName string    `json:"reportedUserProfileName" gorm:"column:ReportedUserProfileName"`
	ReportedUserImageName   string    `json:"reportedUserImageName" gorm:"column:ReportedUserImageName"`
	PostedDate              time.Time `json:"postedDate" gorm:"column:PostedDate"`
	UserImage               string    `json:"userImage" gorm:"column:UserImage"`
	VideoName               string    `json:"videoName" gorm:"column:VideoName"`
	MovieName               string    `json:"movieName" gorm:"column:MovieName"`
	NewsName                string    `json:"newsName" gorm:"column:NewsName"`
	MusicName               string    `json:"musicName" gorm:"column:MusicName"`
	ChannelName             string    `json:"channelName" gorm:"column:ChannelName"`
	ReviewlName             string    `json:"reviewlName" gorm:"column:ReviewlName"`
	PhotoName               string    `json:"photoName" gorm:"column:PhotoName"`
	ReportUserProfileName   string    `json:"reportUserProfileName" gorm:"column:ReportUserProfileName"`
	ReportUserProfileImage  string    `json:"reportUserProfileImage" gorm:"column:ReportUserProfileImage"`
	ReportDate              time.Time `json:"reportDate" gorm:"column:ReportDate"`
	InfoMovie               string    `json:"infoMovie" gorm:"column:InfoMovie"`
	InfoVideo               string    `json:"infoVideo" gorm:"column:InfoVideo"`
	InfoMusic               string    `json:"infoMusic" gorm:"column:InfoMusic"`
	InfoNews                string    `json:"infoNews" gorm:"column:InfoNews"`
	InfoPhoto               string    `json:"infoPhoto" gorm:"column:InfoPhoto"`
	ContentCommentReply     string    `json:"contentCommentReply" gorm:"column:ContentCommentReply"`
	ContentType             string    `json:"contentType" gorm:"column:ContentType"`
	ContentUrl              string    `json:"contentUrl" gorm:"column:ContentUrl"`
	ContentID               string    `json:"contentID" gorm:"column:ContentID"`
}
