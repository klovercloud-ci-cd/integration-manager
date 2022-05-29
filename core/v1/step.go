package v1

import (
	"errors"
	"github.com/klovercloud-ci-cd/integration-manager/enums"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"reflect"
)

// Step contains pipeline step info
type Step struct {
	Name        string                       `json:"name" yaml:"name"`
	Type        enums.STEP_TYPE              `json:"type" yaml:"type"`
	Trigger     enums.TRIGGER                `json:"trigger" yaml:"trigger"`
	Params      map[enums.PARAMS]string      `json:"params" yaml:"params"`
	Next        []string                     `json:"next" yaml:"next"`
	Descriptors *[]unstructured.Unstructured `json:"descriptors" yaml:"descriptors"`
}

// StepForValidation contains pipeline step info for validation
type StepForValidation struct {
	Name    map[string]string   `json:"name" yaml:"name"`
	Type    map[string]string   `json:"type" yaml:"type"`
	Trigger map[string]string   `json:"trigger" yaml:"trigger"`
	Params  []map[string]string `json:"params" yaml:"params"`
	Next    []map[string]string `json:"next" yaml:"next"`
}

// Validate validates pipeline step
func (step Step) Validate() error {
	if step.Name == "" {
		return errors.New("step name is required")
	} else if len(step.Name) > 16 {
		return errors.New("step name length cannot be more than 16 character")
	} else {
		for i := 0; i < len(step.Name); i++ {
			if (step.Name[i] < 97 || step.Name[i] > 122) && (step.Name[i] < 48 || step.Name[i] > 57) {
				return errors.New("step name can only contain lower case characters or digits")
			}
		}
	}
	keys := reflect.ValueOf(step.Params).MapKeys()
	for i := 0; i < len(keys); i++ {
		if step.Params[enums.PARAMS(keys[i].String())] == "" {
			return errors.New("step params is missing")
		}
	}
	if step.Type == enums.BUILD || step.Type == enums.DEPLOY {
		if step.Trigger == enums.AUTO || step.Trigger == enums.MANUAL {
			return nil
		} else if step.Trigger == "" {
			return errors.New("step trigger is required")
		} else {
			return errors.New("step trigger is invalid")
		}
	} else if step.Type == "" {
		return errors.New("step type is required")
	} else {
		return errors.New("step type is invalid")
	}
}

// GetStepForValidationFromStep gets StepForValidation object from Step object
func (step Step) GetStepForValidationFromStep() StepForValidation {
	var stepForValidation StepForValidation
	stepForValidation.Name = step.GetNameWithValidation()
	stepForValidation.Type = step.GetTypeWithValidation()
	stepForValidation.Trigger = step.GetTriggerWithValidation()
	stepForValidation.Params = step.GetParamsWithValidation()
	return stepForValidation
}

func (step Step) GetNameWithValidation() map[string]string {
	nameMap := make(map[string]string)
	nameMap["name"] = "name"
	nameMap["value"] = step.Name
	nameMap["accept"] = "*"
	nameMap["validate"] = "true"
	return nameMap
}

func (step Step) GetTypeWithValidation() map[string]string {
	typeMap := make(map[string]string)
	typeMap["name"] = "type"
	typeMap["value"] = string(step.Type)
	typeMap["accept"] = string(enums.BUILD + "/" + enums.DEPLOY + "/" + enums.INTERMEDIARY + "/" + enums.JENKINS_JOB)
	if val, _ := typeMap["value"]; val == string(enums.BUILD) || val == string(enums.DEPLOY) || val == string(enums.INTERMEDIARY) || val == string(enums.JENKINS_JOB) {
		typeMap["validate"] = "true"
	} else {
		typeMap["validate"] = "false"
	}
	return typeMap
}

func (step Step) GetTriggerWithValidation() map[string]string {
	triggerMap := make(map[string]string)
	triggerMap["name"] = "trigger"
	triggerMap["value"] = string(step.Trigger)
	triggerMap["accept"] = string(enums.AUTO + "/" + enums.MANUAL)
	if val, _ := triggerMap["value"]; val == string(enums.AUTO) || val == string(enums.MANUAL) {
		triggerMap["validate"] = "true"
	} else {
		triggerMap["validate"] = "false"
	}
	return triggerMap
}

func (step Step) GetParamsWithValidation() []map[string]string {
	var paramsMap []map[string]string
	for key, val := range step.Params {
		paramMap := make(map[string]string)
		paramMap["name"] = string(key)
		paramMap["value"] = val
		if key == enums.REPOSITORY_TYPE_PARAM {
			paramMap["accept"] = "git"
		} else {
			paramMap["accept"] = "*"
		}
		if acceptValue, _ := paramMap["accept"]; acceptValue == "*" || val == acceptValue {
			paramMap["validate"] = "true"
		} else {
			paramMap["validate"] = "false"
		}
		paramsMap = append(paramsMap, paramMap)
	}
	return paramsMap
}

func (step Step) GetNextWithValidation(stepNameMap map[string]bool) []map[string]string {
	var nextMaps []map[string]string
	for _, each := range step.Next {
		nextMap := make(map[string]string)
		nextMap["name"] = "next"
		nextMap["value"] = each
		nextMap["accept"] = string(enums.BUILD + "/" + enums.DEPLOY + "/" + enums.INTERMEDIARY + "/" + enums.JENKINS_JOB)
		if _, ok := stepNameMap[each]; ok {
			nextMap["validate"] = "true"
		} else {
			nextMap["validate"] = "false"
		}
		nextMaps = append(nextMaps, nextMap)
	}
	return nextMaps
}
