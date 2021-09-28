package v1

import (
	"github.com/klovercloud-ci/enums"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
	return nil
}
