package test

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
	"testing"
	"tradealgo/pkg/maxdiff"
	"tradealgo/service"
)

func Test_Service(b *testing.T) {
	var buffer bytes.Buffer

	const Max = 1000000
	const BaseTick = 100000
	const BasePrice = 100
	for index := 0; index < Max; index++ {
		buffer.WriteString(fmt.Sprintf("%d,%d", BasePrice+index, BaseTick+index))
		buffer.WriteString(fmt.Sprintln(""))
	}

	reader := csv.NewReader(strings.NewReader(buffer.String()))

	service.MaxProfit(reader, maxdiff.MaxDiffCompute)

}
