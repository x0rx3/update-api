package api

import (
	"net/http"
	"update-api/pkg/models"
	"update-api/pkg/services"

	"github.com/gin-gonic/gin"
)

type DefaultResultHandler struct {
	resultService services.ResultService
}

func NewDefaultResultHandler(resultService services.ResultService) *DefaultResultHandler {
	return &DefaultResultHandler{
		resultService: resultService,
	}
}

// ShowResult godoc
// @Summary      Show Result
// @Description  get result by ID
// @Tags         Result
// @Accept       json
// @Produce      json
// @Param        uuid   path      string  true  "Result UUID"
// @Success      200  	{object}  models.Response{Message=string, Data=models.Result}  "desc"
// @Failure      404  	{object}  models.Response
// @Failure      500  	{object}  models.Response
// @Router       /result/{uuid} [get]
func (inst DefaultResultHandler) Get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, &models.Response{Message: "bad request"})
		return
	}

	result, err := inst.resultService.Get(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Response{Message: "internal error"})
		return
	}

	if result == nil {
		ctx.JSON(http.StatusNotFound, &models.Response{Message: "no entites found"})
		return
	}

	ctx.JSON(200, &models.Response{Message: "success", Data: result})
}

// ShowResults godoc
// @Summary      Show results
// @Description  get results
// @Tags         Result
// @Accept       json
// @Produce      json
// @Success      200  	{object}  models.Response{Message=string, Data=[]models.Result}  "desc"
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Router       /results [get]
func (inst DefaultResultHandler) GetAll(ctx *gin.Context) {
	results, err := inst.resultService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Response{Message: "internal error"})
		return
	}

	if len(results) == 0 {
		ctx.JSON(http.StatusNotFound, &models.Response{Message: "no entites found"})
		return
	}

	ctx.JSON(200, &models.Response{Message: "success", Data: results})
}
