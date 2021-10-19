package v1

import (
	"errors"
	"github.com/klovercloud-ci/enums"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"reflect"
)

type Step struct {
	Name        string                       `json:"name" yaml:"name"`
	Type        enums.STEP_TYPE              `json:"type" yaml:"type"`
	Trigger     enums.TRIGGER                `json:"trigger" yaml:"trigger"`
	Params      map[enums.PARAMS]string      `json:"params" yaml:"params"`
	Next        []string                     `json:"next" yaml:"next"`
	Descriptors *[]unstructured.Unstructured `json:"descriptors" yaml:"descriptors"`
}

func (step Step) Validate() error {
	if step.Name == "" {
		return errors.New("step name is required!")
	}
	keys := reflect.ValueOf(step.Params).MapKeys()
	for i := 0; i < len(keys); i++ {
		if step.Params[enums.PARAMS(keys[i].String())] == "" {
			return errors.New("step params is missing!")
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
