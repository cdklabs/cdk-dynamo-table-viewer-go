package dynamotableviewer

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
)

// Experimental.
type TableViewerProps struct {
	// The DynamoDB table to view.
	//
	// Note that all contents of this table will be
	// visible to the public.
	// Experimental.
	Table awsdynamodb.ITable `field:"required" json:"table" yaml:"table"`
	// The endpoint type of the [LambdaRestApi](https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_aws-apigateway.LambdaRestApi.html) that will be created.
	// Default: - EDGE.
	//
	// Experimental.
	EndpointType awsapigateway.EndpointType `field:"optional" json:"endpointType" yaml:"endpointType"`
	// Name of the column to sort by, prefix with "-" for descending order.
	// Default: - No sort.
	//
	// Experimental.
	SortBy *string `field:"optional" json:"sortBy" yaml:"sortBy"`
	// The web page title.
	// Default: - No title.
	//
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

