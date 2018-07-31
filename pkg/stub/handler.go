package stub

import (
	"context"

	"github.com/banzaicloud/logging-operator/pkg/apis/logging/v1alpha1"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
)

// NewHandler creates a new Handler struct
func NewHandler() sdk.Handler {
	return &Handler{}
}

// Handler struct
type Handler struct {
	// Fill me
}

// Handle every event set up by the watcher
func (h *Handler) Handle(ctx context.Context, event sdk.Event) (err error) {
	switch o := event.Object.(type) {
	case *v1alpha1.LoggingOperator:
		logrus.Infof("New CRD arrived %#v", o)
	}
	return
}

//
//generateFluentdConfig(*v1alpha1.LoggingOperator) {
//
//}

type loggingOperatorCRD struct {
	Input  inputFluentd
	Filter filterFluentd
}

type inputFluentd struct {
	Label string
	Value string
}

type filterFluentd struct {
	Name       string
	Format     string
	TimeFormat string
}

type outputFluentd struct {
	S3 outputS3
}

type outputS3 struct {
	Parameters Parameters
}

// Parameters struct
type Parameters struct {
	Name      string
	ValueFrom ValueFrom
	Value     string
}

// ValueFrom struct
type ValueFrom struct {
	SecretKeyRef kubernetesSecret
}

type kubernetesSecret struct {
	Name string
	Key  string
}