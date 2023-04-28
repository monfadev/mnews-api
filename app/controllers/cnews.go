package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"mnewsapi/app/models"
	"mnewsapi/app/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type newsHandler struct {
	newsService repository.Service
}

func NewNewsHandler(newsService repository.Service) *newsHandler {
	return &newsHandler{newsService}
}

func (h *newsHandler) RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Selamat datang root",
	})
}

func (h *newsHandler) GetNews(ctx *gin.Context) {
	newsHandler, err := h.newsService.FindAll()
	if err != nil { //ada error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	// mengubah struct entity menjadi response
	var newsResponse []models.NewsResponse

	for _, n := range newsHandler {
		newResponse := convertToNewsResponse(n)

		newsResponse = append(newsResponse, newResponse)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": newsResponse,
	})
}

func (h *newsHandler) GetNew(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	n, err := h.newsService.FindByID(id)
	if err != nil { //ada error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	result := convertToNewsResponse(n)

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

func (h *newsHandler) CreateNewsHandler(ctx *gin.Context) {
	var newsRequest models.NewsRequest

	err := ctx.ShouldBindJSON(&newsRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	news, err := h.newsService.Create(newsRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToNewsResponse(news),
	})
}

func (h *newsHandler) UpdateNewsHandler(ctx *gin.Context) {
	var newsUpdateRequest models.NewsUpdateRequest

	err := ctx.ShouldBindJSON(&newsUpdateRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	news, err := h.newsService.Update(id, newsUpdateRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToNewsResponse(news),
	})
}

func (h *newsHandler) DeleteNewsHandler(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	n, err := h.newsService.Delete(id)
	if err != nil { //ada error
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	result := convertToNewsResponse(n)

	ctx.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}

func convertToNewsResponse(newObject models.News) models.NewsResponse {
	return models.NewsResponse{
		ID:          newObject.ID,
		Title:       newObject.Title,
		Description: newObject.Description,
		Author:      newObject.Author,
		PhoneNumber: newObject.PhoneNumber,
	}
}
