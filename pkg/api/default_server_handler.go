package api

import (
	"net/http"
	"update-api/pkg/models"
	"update-api/pkg/services"

	"github.com/gin-gonic/gin"
)

type DefaultServerHandler struct {
	serverService services.ServerService
}

func NewDefaultServerHandler(serverService services.ServerService) *DefaultServerHandler {
	return &DefaultServerHandler{
		serverService: serverService,
	}
}

// ShowServer godoc
// @Summary      Show server
// @Description  get server by ID
// @Tags         Server
// @Accept       json
// @Produce      json
// @Param        uuid   path      string  true  "Server UUID"
// @Success      200  	{object}  models.Response{Message=string, Data=models.Server}  "desc"
// @Failure      404  	{object}  models.Response
// @Failure      500  	{object}  models.Response
// @Router       /server/{uuid} [get]
func (inst DefaultServerHandler) Get(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, &models.Response{Message: "bad request"})
		return
	}

	server, err := inst.serverService.Get(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Response{Message: "internal error"})
		return
	}

	if server == nil {
		ctx.JSON(http.StatusNotFound, &models.Response{Message: "no entites found"})
		return
	}

	ctx.JSON(200, &models.Response{Message: "success", Data: server})
}

// ShowServers godoc
// @Summary      Show servers
// @Description  get servers
// @Tags         Server
// @Accept       json
// @Produce      json
// @Success      200  	{object}  models.Response{Message=string, Data=[]models.Server}  "desc"
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Router       /servers [get]
func (inst DefaultServerHandler) GetAll(ctx *gin.Context) {
	servers, err := inst.serverService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Response{Message: "internal error"})
		return
	}

	if len(servers) == 0 {
		ctx.JSON(http.StatusNotFound, &models.Response{Message: "no entites found"})
		return
	}

	ctx.JSON(200, &models.Response{Message: "success", Data: servers})
}

// UpdateServer godoc
// @Summary      Update server
// @Description  patch server
// @Tags         Server
// @Accept       json
// @Produce      json
// @Param data body models.Server true "Input server"
// @Success      200  	{object}  models.Response{Message=string}  "desc"
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Router       /server/{uuid} [patch]
func (inst DefaultServerHandler) Patch(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, &models.Response{Message: "bad request"})
		return
	}

	server := &models.Server{}
	if err := ctx.Bind(server); err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Response{Message: "internal error"})
		return
	}

	if err := inst.serverService.Update(server); err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Response{Message: "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{Message: "success"})
}

// DeleteServer godoc
// @Summary      Delete server
// @Description  delete server by ID
// @Tags         Server
// @Accept       json
// @Produce      json
// @Param        uuid   path      uuid  true  "Server UUID"
// @Success      200  	{object}  models.Response{Message=string}  "desc"
// @Failure      404  	{object}  models.Response
// @Failure      500  	{object}  models.Response
// @Router       /server/{uuid} [delete]
func (inst DefaultServerHandler) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, &models.Response{Message: "internal error"})
		return
	}

	if err := inst.serverService.Delete(uuid); err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Response{Message: "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{Message: "success"})
}

// CreateServer godoc
// @Summary      Create server
// @Description  post server
// @Tags         Server
// @Accept       json
// @Produce      json
// @Param data body models.Server true "Input server"
// @Success      200  	{object}  models.Response{Message=string, Data=string}  "desc"
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
// @Router       /server [post]
func (inst DefaultServerHandler) Post(ctx *gin.Context) {
	server := &models.Server{}

	if err := ctx.Bind(server); err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Response{Message: "internal error"})
		return
	}

	uuid, err := inst.serverService.Create(server)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Response{Message: "internal error"})
		return
	}

	ctx.JSON(http.StatusCreated, &models.Response{Message: "succes", Data: uuid})
}
