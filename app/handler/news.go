package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"mnewsapi/app/news"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type newsHandler struct {
	newsService news.Service
}

func NewNewsHandler(newsService news.Service) *newsHandler {
	return &newsHandler{newsService}
}

func (h *newsHandler) RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":       "Bekasi Memiliki Internet Tercepat Se-Indonesia",
		"description": "Laborum ipsum est officia anim reprehenderit in ad occaecat excepteur.",
	})
}

func (h *newsHandler) HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "Hellow",
	})
}

func (h *newsHandler) NewsHandler(ctx *gin.Context) {
	title := ctx.Param("title")
	desc := ctx.Param("desc")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"desc":  desc,
	})

}

func (h *newsHandler) QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	desc := ctx.Query("desc")

	ctx.JSON(http.StatusOK, gin.H{
		"title": title,
		"desc":  desc,
	})
}

func (h *newsHandler) PostNewsHandler(ctx *gin.Context) {
	var newsRequest news.NewsRequest

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
		"data": news,
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
	var newsResponse []news.NewsResponse

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
	var newsRequest news.NewsRequest

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
	var newsUpdateRequest news.NewsUpdateRequest

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

func convertToNewsResponse(newObject news.News) news.NewsResponse {
	return news.NewsResponse{
		ID:          newObject.ID,
		Title:       newObject.Title,
		Description: newObject.Description,
		Author:      newObject.Author,
		PhoneNumber: newObject.PhoneNumber,
	}
}
