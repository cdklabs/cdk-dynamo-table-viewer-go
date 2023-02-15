// An AWS CDK construct which exposes an endpoint with the contents of a DynamoDB table
package dynamotableviewer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-dynamo-table-viewer-go/dynamotableviewer/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-dynamo-table-viewer-go/dynamotableviewer/internal"
)

// Installs an endpoint in your stack that allows users to view the contents of a DynamoDB table through their browser.
type TableViewer interface {
	constructs.Construct
	Endpoint() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for TableViewer
type jsiiProxy_TableViewer struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_TableViewer) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TableViewer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewTableViewer(parent constructs.Construct, id *string, props *TableViewerProps) TableViewer {
	_init_.Initialize()

	if err := validateNewTableViewerParameters(parent, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TableViewer{}

	_jsii_.Create(
		"cdk-dynamo-table-viewer.TableViewer",
		[]interface{}{parent, id, props},
		&j,
	)

	return &j
}

func NewTableViewer_Override(t TableViewer, parent constructs.Construct, id *string, props *TableViewerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-dynamo-table-viewer.TableViewer",
		[]interface{}{parent, id, props},
		t,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func TableViewer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTableViewer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-dynamo-table-viewer.TableViewer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableViewer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

