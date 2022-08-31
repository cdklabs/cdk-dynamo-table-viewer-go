//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// An AWS CDK construct which exposes an endpoint with the contents of a DynamoDB table
package dynamotableviewer

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateTableViewer_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewTableViewerParameters(parent constructs.Construct, id *string, props *TableViewerProps) error {
	if parent == nil {
		return fmt.Errorf("parameter parent is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

