package handler

import (
	"fmt"
	"net/http"

	"github.com/abraaoduarte/goopportunities/schemas"
	"github.com/gin-gonic/gin"
)


func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening  := schemas.Opening{}
	// Find
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	// Delete
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)


}