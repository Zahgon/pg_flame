package html

import (
	"io"

	"pg_flame/pkg/plan"
)

type Flame struct {
	Name     string  `json:"name"`
	Value    float64 `json:"value"`
	Time     float64 `json:"time"`
	Detail   string  `json:"detail"`
	Color    string  `json:"color"`
	InitPlan bool    `json:"init_plan"`
	Children []Flame `json:"children"`
}

const detailSpan = "<span>%s</span>"

const colorPlan = "#00C05A"
const colorInit = "#C0C0C0"

func Generate(w io.Writer, p plan.Plan) error { _ = "STUB: not implemented"; return nil }

func buildFlame(p plan.Plan) (Flame, error) { _ = "STUB: not implemented"; return *new(Flame), nil }

func convertPlanNode(n plan.Node, color string) (Flame, error) {
	_ = "STUB: not implemented"
	return *new(Flame), nil
}

// Pass the color forward for grey InitPlan trees

// Add to the total value if the child is an InitPlan node

func name(n plan.Node) string { _ = "STUB: not implemented"; return "" }

func detail(n plan.Node) (string, error) { _ = "STUB: not implemented"; return "", nil }
