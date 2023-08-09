package applicationModel

type DefaultOperator struct {
	DataType string   `json:"data_type"`
	Operator []string `json:"operator"`
}

var GetListSkillValidOperator map[string]DefaultOperator

func InitiateDefaultOperator() {
	GetListSkillValidOperator = make(map[string]DefaultOperator)
}
