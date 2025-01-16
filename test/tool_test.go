package test

import (
	"fmt"
	"hookfunc/pkg/helper/uuid"
	"testing"
)

func TestTool(t *testing.T) {
	fmt.Println(uuid.GenerateSMSCode())
}
