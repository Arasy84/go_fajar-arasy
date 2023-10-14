package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	openai "github.com/sashabaranov/go-openai"
)

type LaptopHandler struct {
	LaptopUsecase LaptopUsecase 
}

type LaptopResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type LaptopUsecase interface {
	RecommendLaptop(userInput, brand, ram, cpu, screenSize, openAIKey string) (string, error)
}

type laptopUsecase struct{}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	laptopUsecase := NewLaptopUsecase()
	laptopHandler := &LaptopHandler{
		LaptopUsecase: laptopUsecase,
	}

	e.POST("/laptop-recomendation", laptopHandler.RecommendLaptop)

	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	e.Start(port)
}

func NewLaptopUsecase() LaptopUsecase {
	return &laptopUsecase{}
}

func (uc *laptopUsecase) RecommendLaptop(userInput, brand, ram, cpu, screenSize, openAIKey string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Halo, perkenalkan saya sistem untuk rekomendasi laptop",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		},
	}

	resp, err := uc.getCompletionFromMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := resp.Choices[0].Message.Content
	return answer, nil
}

func (uc *laptopUsecase) getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}

func (h *LaptopHandler) RecommendLaptop(c echo.Context) error {
	var requestData map[string]interface{}
	err := c.Bind(&requestData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON format")
	}

	budget, okBudget := requestData["budget"].(float64)
	purpose, okPurpose := requestData["purpose"].(string)
	brand, okBrand := requestData["brand"].(string)
	ram, okRAM := requestData["ram"].(string)
	cpu, okCPU := requestData["cpu"].(string)
	screenSize, okScreenSize := requestData["screen_size"].(string)

	if !okBudget || !okPurpose || !okBrand || !okRAM || !okCPU || !okScreenSize {
		return c.JSON(http.StatusBadRequest, "Invalid request format")
	}

	userInput := fmt.Sprintf("Rekomendasi Laptop merk %s untuk %s dengan spesifikasi %s ram %s cpu %s ukuran layar dan dengan budget sebesar Rp %.0f.", brand, purpose, ram, cpu, screenSize, budget)

	answer, err := h.LaptopUsecase.RecommendLaptop(userInput, brand, ram, cpu, screenSize, os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error dalam membuat rekomendasi")
	}

	responseData := LaptopResponse{
		Status: "success",
		Data:   answer,
	}

	return c.JSON(http.StatusOK, responseData)
}
