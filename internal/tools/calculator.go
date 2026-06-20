package tools

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

type CalculatorTool struct{}

func NewCalculatorTool() *CalculatorTool {
	return &CalculatorTool{}
}

func (c *CalculatorTool) Name() string {
	return "calculator"
}

func (c *CalculatorTool) Execute(
	ctx context.Context,
	input string,
) (string, error) {

	input =
		strings.TrimSpace(
			strings.ReplaceAll(
				input,
				"calculate",
				"",
			),
		)

	parts := strings.Fields(input)

	if len(parts) != 3 {
		return "",
			fmt.Errorf("invalid expression")
	}

	left, err :=
		strconv.ParseFloat(
			parts[0],
			64,
		)

	if err != nil {
		return "", err
	}

	right, err :=
		strconv.ParseFloat(
			parts[2],
			64,
		)

	if err != nil {
		return "", err
	}

	var result float64

	switch parts[1] {

	case "+":
		result = left + right

	case "-":
		result = left - right

	case "*":
		result = left * right

	case "/":
		result = left / right

	default:
		return "",
			fmt.Errorf(
				"unsupported operator",
			)
	}

	return fmt.Sprintf(
		"Result = %.2f",
		result,
	), nil
}
