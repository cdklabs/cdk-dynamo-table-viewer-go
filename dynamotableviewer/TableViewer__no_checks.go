//go:build no_runtime_type_checking

package dynamotableviewer

// Building without runtime type checking enabled, so all the below just return nil

func validateTableViewer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTableViewerParameters(parent constructs.Construct, id *string, props *TableViewerProps) error {
	return nil
}

