package http

import (
	"encoding/json"
	"evo_fintech/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type DataService interface {
	Upload(data *multipart.File) error
	Download(filter *entity.Filter) ([]*entity.Data, error)
}

type DataController struct {
	service DataService
}

func newDataController(service DataService) *DataController {
	return &DataController{service: service}
}

// @Tags        upload
// @Description upload csv file, parsing it and saving the parsing results to the database
// @Param       file formData file true "csv file"
// @Success     202
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /api/upload [post]
func (d *DataController) upload(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	data, err := file.Open()
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "can't open the file")
		return
	}
	err = d.service.Upload(&data)
	if err != nil {
		if strings.Contains(err.Error(), "parse error") {
			errorResponse(ctx, http.StatusBadRequest, "invalid file")
			return
		}
		if strings.Contains(err.Error(), "duplicate") {
			errorResponse(ctx, http.StatusBadRequest, "duplicate primary key")
			return
		}
		if strings.Contains(err.Error(), "No connection") {
			errorResponse(ctx, http.StatusInternalServerError, "database problems")
			return
		}
		errorResponse(ctx, http.StatusBadRequest, "incorrect csv")
		return
	}
	ctx.Status(http.StatusAccepted)
}

// @Tags        download
// @Description Download in json or csv format with filters
// @Param       format path string true "download format: json or csv"
// @Param       transaction_id query int false "transaction id: n or 1, 2, 3, ..., n"
// @Param       terminal_id query string false "terminal id"
// @Param       status query string false "status: accepted or declined"
// @Param       payment_type query string false "payment type: cash or card"
// @Param       date_post query string false "date post: from yyyy-mm-dd, to yyyy-mm-dd"
// @Param       payment_narrative query string false "payment narrative"
// @Success     200
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /api/download/{format} [get]
func (d *DataController) download(ctx *gin.Context) {
	format := ctx.Param("format")
	transactionId, err := strconv.Atoi(ctx.Query("transaction_id"))
	if err != nil && ctx.Query("transaction_id") != "" {
		errorResponse(ctx, http.StatusBadRequest, "transaction_id is incorrect")
		return
	}
	terminalId := ctx.Query("terminal_id")
	if terminalId != "" {
		if _, err := strconv.Atoi(terminalId); err != nil {
			if ok, _ := regexp.MatchString(`^\d+(?:[ \t]*,[ \t]*\d+)+$`, terminalId); !ok {
				errorResponse(ctx, http.StatusBadRequest, "terminal_id is incorrect")
				return
			}
		}
	}
	status := ctx.Query("status")
	if status != "accepted" && status != "declined" && status != "" {
		errorResponse(ctx, http.StatusBadRequest, "status is incorrect")
		return
	}
	paymentType := ctx.Query("payment_type")
	if paymentType != "cash" && paymentType != "card" && paymentType != "" {
		errorResponse(ctx, http.StatusBadRequest, "payment_type is incorrect")
		return
	}
	datePost := ctx.Query("date_post")
	var datePostSlice []string
	if datePost != "" {
		if ok, _ := regexp.MatchString(`^(from)\s+(19|20)\d\d[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01]), (to)\s+(19|20)\d\d[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])$`, datePost); !ok {
			errorResponse(ctx, http.StatusBadRequest, "date_post is incorrect")
			return
		} else {
			datePost = strings.ReplaceAll(datePost, "from", "")
			datePost = strings.ReplaceAll(datePost, "to", "")
			datePost = strings.ReplaceAll(datePost, " ", "")
			datePostSlice = strings.Split(datePost, ",")
		}
	}
	paymentNarrative := ctx.Query("payment_narrative")

	data, err := d.service.Download(&entity.Filter{
		TransactionId:    transactionId,
		TerminalId:       terminalId,
		Status:           status,
		PaymentType:      paymentType,
		DatePost:         datePostSlice,
		PaymentNarrative: paymentNarrative,
	})
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, "database problems")
		return
	}
	var out []byte
	switch format {
	case "json":
		out, err = json.Marshal(data)
		if err != nil {
			errorResponse(ctx, http.StatusInternalServerError, "json encoding error")
			return
		}
		ctx.Data(http.StatusOK, "application/json", out)
	case "csv":
		out, err = gocsv.MarshalBytes(data)
		if err != nil {
			errorResponse(ctx, http.StatusInternalServerError, "csv encoding error")
			return
		}
		ctx.Data(http.StatusOK, "text/csv", out)
	default:
		errorResponse(ctx, http.StatusBadRequest, "invalid download format")
	}
}
