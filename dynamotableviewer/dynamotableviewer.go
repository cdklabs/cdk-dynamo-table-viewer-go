package dynamotableviewer

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-dynamo-table-viewer.TableViewer",
		reflect.TypeOf((*TableViewer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TableViewer{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-dynamo-table-viewer.TableViewerProps",
		reflect.TypeOf((*TableViewerProps)(nil)).Elem(),
	)
}
