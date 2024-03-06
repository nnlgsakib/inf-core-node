package fund

import (
	"bytes"
	"fmt"

	"github.com/Infinity-Green/inf/command/helper"
	"github.com/Infinity-Green/inf/types"
)

type result struct {
	ValidatorAddr types.Address `json:"address"`
	TxHash        types.Hash    `json:"tx_hash"`
	IsMinted      bool          `json:"mint"`
}

func (r *result) GetOutput() string {
	var buffer bytes.Buffer

	vals := make([]string, 0, 3)
	vals = append(vals, fmt.Sprintf("Validator (address)|%s", r.ValidatorAddr))
	vals = append(vals, fmt.Sprintf("Transaction (hash)|%s", r.TxHash))
	vals = append(vals, fmt.Sprintf("Is minted|%v", r.IsMinted))

	buffer.WriteString("\n[ROOTCHAIN FUND]\n")
	buffer.WriteString(helper.FormatKV(vals))
	buffer.WriteString("\n")

	return buffer.String()
}
