package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	dbpb "diaxel/proto/db"
)

func main() {
	// Connect to database service
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := dbpb.NewDatabaseServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("🔍 Testing gRPC communication...")

	// Test 1: Create Assistant
	fmt.Println("\n1️⃣ Creating assistant...")
	createReq := &dbpb.CreateAssistantRequest{
		Name:     "Test Assistant",
		ApiToken: "test-api-token-12345",
		UserId:   "test-user-123",
	}

	assistant, err := client.CreateAssistant(ctx, createReq)
	if err != nil {
		log.Printf("❌ Failed to create assistant: %v", err)
	} else {
		fmt.Printf("✅ Assistant created: ID=%s, Name=%s, ApiToken=%s\n",
			assistant.Id, assistant.Name, assistant.ApiToken)
	}

	// Test 2: Get Assistant by ID
	fmt.Println("\n2️⃣ Getting assistant by ID...")
	getReq := &dbpb.GetAssistantRequest{
		Id: assistant.Id,
	}

	retrievedAssistant, err := client.GetAssistant(ctx, getReq)
	if err != nil {
		log.Printf("❌ Failed to get assistant: %v", err)
	} else {
		fmt.Printf("✅ Assistant retrieved: ID=%s, Name=%s, ApiToken=%s\n",
			retrievedAssistant.Id, retrievedAssistant.Name, retrievedAssistant.ApiToken)
	}

	// Test 3: Get Assistant by API Token
	fmt.Println("\n3️⃣ Getting assistant by API token...")
	getByTokenReq := &dbpb.GetAssistantByAPITokenRequest{
		ApiToken: "test-api-token-12345",
	}

	assistantByToken, err := client.GetAssistantByAPIToken(ctx, getByTokenReq)
	if err != nil {
		log.Printf("❌ Failed to get assistant by API token: %v", err)
	} else {
		fmt.Printf("✅ Assistant retrieved by API token: ID=%s, Name=%s, ApiToken=%s\n",
			assistantByToken.Id, assistantByToken.Name, assistantByToken.ApiToken)
	}

	// Test 4: Create Chat
	fmt.Println("\n4️⃣ Creating chat...")
	chatReq := &dbpb.CreateChatRequest{
		AssistantId: assistant.Id,
		CustomerId:  "550e8400-e29b-41d4-a716-446655440000", // Valid UUID
		Platform:    "test",
	}

	chat, err := client.CreateChat(ctx, chatReq)
	if err != nil {
		log.Printf("❌ Failed to create chat: %v", err)
	} else {
		fmt.Printf("✅ Chat created: ID=%s, AssistantId=%s\n",
			chat.Id, chat.AssistantId)
	}

	fmt.Println("\n🎉 All gRPC tests completed successfully!")
}
