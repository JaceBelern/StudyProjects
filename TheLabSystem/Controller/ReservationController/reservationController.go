package ReservationController

import (
	"TheLabSystem/Config/ErrorInformation"
	"TheLabSystem/Service/ReservationService"
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/GetApprovalRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/GetPersonalReservationRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/GetReservationDetailsRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/GetReservationInfoByReservationIDRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/RevertReservationRequestAndRespoonse"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/SetApprovalRequestAndResponse"
	"TheLabSystem/Types/RequestAndResponseType/Reservation/SubmitReservationRequestAndResponse"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReservationController struct {
}

func (controller ReservationController) SubmitReservation(c *gin.Context) {

	var request = &SubmitReservationRequestAndResponse.SubmitReservationRequest{}
	var response = &SubmitReservationRequestAndResponse.SubmitReservationResponse{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println(request)
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = ReservationService.ReservationService{}.SubmitReservation(cookie, request)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}

func (controller ReservationController) RevertReservation(c *gin.Context) {
	var request = &RevertReservationRequestAndRespoonse.RevertReservationRequest{}
	var response = &RevertReservationRequestAndRespoonse.RevertReservationResponse{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = ReservationService.ReservationService{}.RevertReservation(cookie, request.ReservationID)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
func (controller ReservationController) GetApproval(c *gin.Context) {
	var request = &GetApprovalRequestAndResponse.GetApprovalRequest{}
	var response = &GetApprovalRequestAndResponse.GetApprovalResponse{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data.ApprovalRes, response.Code = ReservationService.ReservationService{}.GetApproval(cookie)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
func (controller ReservationController) SetApproval(c *gin.Context) {
	var request = &SetApprovalRequestAndResponse.SetApprovalRequest{}
	var response = &SetApprovalRequestAndResponse.SetApprovalResponse{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Code = ReservationService.ReservationService{}.SetApproval(cookie, request)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
	return
}
func (controller ReservationController) GetPersonalReservations(c *gin.Context) {
	var request = &GetPersonalReservationRequestAndResponse.GetPersonalReservationRequest{}
	var response = &GetPersonalReservationRequestAndResponse.GetPersonalReservationResponse{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data.Reservations, response.Code = ReservationService.ReservationService{}.GetPersonalReservations(cookie)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}
func (controller ReservationController) GetReservationInfo(c *gin.Context) {
	var request = &GetReservationInfoByReservationIDRequestAndResponse.GetReservationInfoByReservationIDRequest{}
	var response = &GetReservationInfoByReservationIDRequestAndResponse.GetReservationInfoByReservationIDResponse{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data.Reservation, response.Data.ReservationInfos, response.Code = ReservationService.ReservationService{}.GetReservationByReservationID(cookie, request.ReservationID)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}
func (controller ReservationController) GetReservationDetailsController(c *gin.Context) {
	var request = &GetReservationDetailsRequestAndResponse.GetReservationDetailsRequest{}
	var response = &GetReservationDetailsRequestAndResponse.GetReservationDetailsResponse{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	cookie, err := c.Cookie("camp-session")
	if err != nil {
		response.Code = ErrNo.LoginRequired
		response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data.Nums, response.Code = ReservationService.ReservationService{}.GetReservationDetails(cookie, request.Day, request.DeviceType)
	response.Data.Message = ErrorInformation.GenerateErrorInformation(response.Code)
	c.JSON(http.StatusOK, response)
}
