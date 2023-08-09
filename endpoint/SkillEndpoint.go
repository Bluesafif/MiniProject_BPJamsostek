package endpoint

import (
	"MiniProjectBPJamsostek/service/Skill"
	"net/http"
)

type skillEndpoint struct {
	FileName string
}

var SkillEndpoint = skillEndpoint{}.New()

func (input skillEndpoint) New() (output skillEndpoint) {
	output.FileName = "SkillEndpoint.go"
	return
}

func (input skillEndpoint) SkillWithParam(response http.ResponseWriter, request *http.Request) {
	var err error
	if request.Method == "POST" {
		err = Skill.SkillService.InsertSkill(request, response)
	} else if request.Method == "GET" {
		err = Skill.SkillService.GetListSkill(request, response)
	} else if request.Method == "DELETE" {
		err = Skill.SkillService.DeleteSkill(request, response)
	}

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
