package Controllers

import (
	"fmt"
	"net/http"
	"wan-api-spam-report/Const"
	"wan-api-spam-report/DTO"
	"wan-api-spam-report/Logic"
	"wan-api-spam-report/Utils"
	"wan-api-spam-report/ViewModels"

	"github.com/gin-gonic/gin"
)

func SaveorUpdateSpamReports(context *gin.Context) {
	var spamReport ViewModels.SpamReportViewModel
	var err error

	//Get request body
	var body DTO.SpamReportsDTO
	err = context.Bind(&body)
	//Handle when body error occurred
	if err != nil {
		spamReport.SpamReport = nil
		spamReport.Result = Const.UnSuccess
		spamReport.ErrorMessage = err.Error()
		context.JSON(http.StatusBadRequest, spamReport)
		return
	}
	fmt.Println("spamreport", body)
	//Insert new News if SpamReport ID = 0
	if body.SpamReportID == 0 {
		spamReport.SpamReport = &body
		spamReport.SpamReport, err = Logic.SpamReport_Save(spamReport.SpamReport)
		fmt.Println("saveeeeeeeee")
	} else {
		//Otherwise update News
		spamReport.IsExisted = true
		var params DTO.AddedParam = make(DTO.AddedParam)
		params["spamReportID"] = Utils.Int64ToString(body.SpamReportID)
		params["reviewID"] = SafeInt64PointerToString(body.ReviewID)
		params["commentID"] = SafeInt64PointerToString(body.CommentID)
		params["commentReplyID"] = SafeInt64PointerToString(body.CommentReplyID)
		params["name"] = body.Name
		params["movieID"] = SafeInt64PointerToString(body.MovieID)
		params["videoID"] = SafeInt64PointerToString(body.VideoID)
		params["musicID"] = SafeInt64PointerToString(body.MusicID)
		params["newsID"] = SafeInt64PointerToString(body.NewsID)
		params["reason"] = body.Reason
		params["reportUserProfileID"] = Utils.Int64ToString(body.ReportUserProfileID)
		params["photoID"] = SafeInt64PointerToString(body.PhotoID)
		params["channelProfileID"] = SafeInt64PointerToString(body.ChannelProfileID)
		params["isDisable"] = Utils.BoolToString(body.IsDisable)

		fmt.Println("update", params)

		spamReport.SpamReport, err = Logic.SpamReport_Update(params)
	}
	if err != nil {
		spamReport.SpamReport = nil
		spamReport.Result = Const.UnSuccess
		spamReport.ErrorMessage = err.Error()
		context.JSON(http.StatusBadRequest, spamReport)
		return
	}
	spamReport.SpamReportID = spamReport.SpamReport.SpamReportID
	spamReport.Result = Const.Success
	context.JSON(http.StatusOK, spamReport)
	return
}

func SafeInt64PointerToString(value *int64) string {
	if value != nil {
		return Utils.Int64ToString(*value)
	}
	return "" // hoặc giá trị mặc định khác tùy thuộc vào yêu cầu của bạn
}
