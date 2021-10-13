package main
 
import (
	"os/exec"
  "github.com/aws/aws-lambda-go/lambda"
)

type Output struct {
  Payload string `json:"payload"`
}

type Input struct {
	Cmd string
	Nonce string
}
 
func HandleLambdaEvent(input Input) (Output, error) {
	cmd := exec.Command("bash", "-c", input.Cmd)
	cmdOut, cmdErr := cmd.Output()
	if cmdErr != nil {
		panic(cmdErr)
	}
  return Output{string(cmdOut)}, nil
}
 
func main() {
  lambda.Start(HandleLambdaEvent)
}
