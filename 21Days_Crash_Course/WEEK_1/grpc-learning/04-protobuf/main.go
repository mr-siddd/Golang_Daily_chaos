package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"

	positionpb "github.com/mr-siddd/grpc-learning/04-protobuf/proto"
)

func main() {
	// ---- the SAME data, built two ways ----

	// 1. As a protobuf message (using the generated struct)
	pbPosition := &positionpb.Position{
		Symbol:    "TCS",
		RiskScore: 80.5,
		Timestamp: "1719800000",
	}

	// 2. As a plain Go struct for JSON (same data)
	type JSONPosition struct {
		Symbol    string  `json:"symbol"`
		RiskScore float64 `json:"risk_score"`
		Timestamp int64   `json:"timestamp"`
	}
	jsonPosition := JSONPosition{
		Symbol:    "TCS",
		RiskScore: 80.5,
		Timestamp: 1719800000,
	}

	// ---- serialize both ----

	// protobuf → binary bytes
	pbBytes, err := proto.Marshal(pbPosition)
	if err != nil {
		panic(err)
	}

	// json → text bytes
	jsonBytes, err := json.Marshal(jsonPosition)
	if err != nil {
		panic(err)
	}

	// ---- compare ----
	fmt.Println("=== SAME DATA, TWO FORMATS ===")
	fmt.Println()

	fmt.Println("--- JSON ---")
	fmt.Printf("readable: %s\n", string(jsonBytes))
	fmt.Printf("size:     %d bytes\n", len(jsonBytes))
	fmt.Println()

	fmt.Println("--- Protobuf ---")
	fmt.Printf("raw bytes: %v\n", pbBytes)
	fmt.Printf("hex:       % x\n", pbBytes)
	fmt.Printf("size:      %d bytes\n", len(pbBytes))
	fmt.Println()

	// the win
	saved := len(jsonBytes) - len(pbBytes)
	pct := float64(saved) / float64(len(jsonBytes)) * 100
	fmt.Printf("protobuf saved %d bytes (%.0f%% smaller)\n", saved, pct)
}
