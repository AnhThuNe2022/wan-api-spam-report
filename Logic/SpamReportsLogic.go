package Logic

import (
	"fmt"
	"log"
	"time"
	"wan-api-spam-report/DTO"
	"wan-api-spam-report/Initializers"
	"wan-api-spam-report/Models"
	"wan-api-spam-report/Utils"
)

func SpamReport_Save(item *DTO.SpamReportsDTO) (*DTO.SpamReportsDTO, error) {
	var spamReportItem = &Models.SpamReports{}
	var err error

	//Create new data to prepare to save
	ConvertSpamReportDTOToData(item, spamReportItem)
	spamReportItem.CreatedDate = time.Now()
	spamReportItem.CreatedBy = item.CreatedBy
	spamReportItem.Active = true

	//Save to DB

	//Handle when save data error occurred
	if err = Initializers.DB.Save(&spamReportItem).Error; err != nil {
		return item, err
	}

	//Return data when data saved successfully
	item.SpamReportID = spamReportItem.SpamReportID
	item.CreatedDate = spamReportItem.CreatedDate
	return item, err
}
func SpamReport_Update(params DTO.AddedParam) (*DTO.SpamReportsDTO, error) {
	var spamReportItem = &Models.SpamReports{}
	var item = &DTO.SpamReportsDTO{}
	var err error

	// Verify spamReportID ID
	var id int64
	id = Utils.StringToInt64(params["spamReportID"])
	if err = Initializers.DB.First(&spamReportItem, id).Error; err != nil {
		return item, err
	}
	fmt.Println("FK_SpamReports_ChannelProfiles_ChannelProfileIaaaaD", spamReportItem)
	delete(params, "spamReportID")
	//delete(params, "channelConfigID")

	//Update fields
	for key, value := range params {
		switch key {
		case "reviewID":
			if value != "" {
				var reviewID int64 = Utils.StringToInt64(value)
				spamReportItem.ReviewID = &reviewID
			}
			break
		case "commentID":
			if value != "" {
				var commentID int64 = Utils.StringToInt64(value)
				spamReportItem.CommentID = &commentID

			}
			break

		case "commentReplyID":
			if value != "" {
				var commentReplyID int64 = Utils.StringToInt64(value)
				spamReportItem.CommentReplyID = &commentReplyID
			}
			break

		case "name":
			if value != "" {
				var name string = value
				spamReportItem.Name = name
			}
			break

		case "movieID":
			if value != "" {
				var movieID int64 = Utils.StringToInt64(value)
				spamReportItem.MovieID = &movieID
			}
			break

		case "videoID":
			if value != "" {
				var videoID int64 = Utils.StringToInt64(value)
				spamReportItem.VideoID = &videoID
				fmt.Println("FK_ssrofileIaamovieIDaaD", videoID)

			}
			break

		case "musicID":
			if value != "" {
				var musicID int64 = Utils.StringToInt64(value)
				spamReportItem.MusicID = &musicID
			}
			break

		case "newsID":
			if value != "" {
				var newsID int64 = Utils.StringToInt64(value)
				spamReportItem.NewsID = &newsID
			}
			break

		case "reason":
			if value != "" {
				var reason string = value
				spamReportItem.Reason = reason
			}
			break

		case "reportUserProfileID":
			if value != "" {
				var reportUserProfileID int64 = Utils.StringToInt64(value)
				spamReportItem.ReportUserProfileID = reportUserProfileID
			}
			break

		case "photoID":
			if value != "" {
				var photoID int64 = Utils.StringToInt64(value)
				spamReportItem.PhotoID = &photoID
			}
			break

		// case "channelProfileID":
		// 	if value != "0" {
		// 		var channelProfileID int64 = Utils.StringToInt64(value)
		// 		spamReportItem.ChannelProfileID = channelProfileID
		// 		fmt.Println("FK_SpamReports_ChannelProfiles_ChannelProfileID", value)
		// 	}
		// 	fmt.Println("FK_SpamReports_ChannelProfiles_ChannelProfileIaaaaD", value)
		// 	break

		case "isDisable":
			if value != "" {
				var isDisable bool = Utils.StringToBool(value)
				spamReportItem.IsDisable = isDisable
			}
			break

		default:
			log.Println("The field key doesn't exist. Re-enter the key.")
		}
	}
	spamReportItem.ModifiedDate = time.Now()

	//Save to database
	if err = Initializers.DB.Save(&spamReportItem).Error; err != nil {
		return item, err
	}
	item = ConvertToSpamReportDTO(spamReportItem)
	return item, err

}

func ConvertSpamReportDTOToData(item *DTO.SpamReportsDTO, data *Models.SpamReports) {
	data.SpamReportID = item.SpamReportID
	data.ReviewID = item.ReviewID
	data.CommentID = item.CommentID
	data.CommentReplyID = item.CommentReplyID
	data.MovieID = item.MovieID
	data.VideoID = item.VideoID
	data.MusicID = item.MusicID
	data.NewsID = item.NewsID
	data.PhotoID = item.PhotoID
	// data.PostID = item.PostID
	data.ChannelProfileID = item.ChannelProfileID
	data.Name = item.Name
	data.Reason = item.Reason
	data.ReportUserProfileID = item.ReportUserProfileID
	data.IsDisable = item.IsDisable
	// data.IsSensitive = item.IsSensitive
	data.Active = item.Active
	data.CreatedBy = item.CreatedBy
	data.CreatedDate = item.CreatedDate
	data.ModifiedBy = item.ModifiedBy
	data.ModifiedDate = item.ModifiedDate
}

func ConvertToSpamReportDTO(item *Models.SpamReports) *DTO.SpamReportsDTO {
	var dto DTO.SpamReportsDTO
	dto.SpamReportID = item.SpamReportID
	dto.ReviewID = item.ReviewID
	dto.CommentID = item.CommentID
	dto.CommentReplyID = item.CommentReplyID
	dto.MovieID = item.MovieID
	dto.VideoID = item.VideoID
	dto.MusicID = item.MusicID
	dto.NewsID = item.NewsID
	dto.PhotoID = item.PhotoID

	// dto.PostID = item.PostID
	dto.ChannelProfileID = item.ChannelProfileID
	dto.Name = item.Name
	dto.Reason = item.Reason
	dto.ReportUserProfileID = item.ReportUserProfileID
	dto.IsDisable = item.IsDisable
	// dto.IsSensitive = item.IsSensitive
	dto.Active = item.Active
	dto.CreatedBy = item.CreatedBy
	dto.CreatedDate = item.CreatedDate
	dto.ModifiedBy = item.ModifiedBy
	dto.ModifiedDate = item.ModifiedDate

	return &dto
}
