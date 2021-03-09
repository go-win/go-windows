// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package machinelearning implements the Windows.Win32.MachineLearning namespace.
package machinelearning

type WINML_TENSOR_BINDING_DESC struct {
	DataType      int
	NumDimensions int
	PSHAPE        int
	DataSize      int
	PDATA         int
}

type WINML_SEQUENCE_BINDING_DESC struct {
	ElementCount int
	ElementType  int
	Anonymous    int
}

type WINML_MAP_BINDING_DESC struct {
	ElementCount int
	KeyType      int
	Anonymous1   int
	Fields       int
	Anonymous2   int
}

type WINML_IMAGE_BINDING_DESC struct {
	ElementType   int
	NumDimensions int
	PSHAPE        int
	DataSize      int
	PDATA         int
}

type WINML_RESOURCE_BINDING_DESC struct {
	ElementType   int
	NumDimensions int
	PSHAPE        int
	PRESOURCE     int
}

type WINML_BINDING_DESC struct {
	Name      int
	BindType  int
	Anonymous int
}

type WINML_TENSOR_VARIABLE_DESC struct {
	ElementType   int
	NumDimensions int
	PSHAPE        int
}

type WINML_SEQUENCE_VARIABLE_DESC struct {
	ElementType int
}

type WINML_MAP_VARIABLE_DESC struct {
	KeyType int
	Fields  int
}

type WINML_IMAGE_VARIABLE_DESC struct {
	ElementType   int
	NumDimensions int
	PSHAPE        int
}

type WINML_VARIABLE_DESC struct {
	Name        int
	Description int
	FeatureType int
	Required    int
	Anonymous   int
}

type WINML_MODEL_DESC struct {
	Author      int
	Name        int
	Domain      int
	Description int
	Version     int
}

type MLOperatorEdgeDescription struct {
	EDGETYPE  int
	Anonymous int
}

type MLOperatorSchemaEdgeDescription struct {
	OPTIONS    int
	TYPEFORMAT int
	Anonymous  int
}

type MLOperatorEdgeTypeConstraint struct {
	TYPELABEL        int
	ALLOWEDTYPES     int
	ALLOWEDTYPECOUNT int
}

type MLOperatorAttribute struct {
	NAME     int
	TYPE     int
	REQUIRED int
}

type MLOperatorAttributeNameValue struct {
	NAME       int
	TYPE       int
	VALUECOUNT int
	Anonymous  int
}

type MLOperatorSchemaDescription struct {
	NAME                           int
	OPERATORSETVERSIONATLASTCHANGE int
	INPUTS                         int
	INPUTCOUNT                     int
	OUTPUTS                        int
	OUTPUTCOUNT                    int
	TYPECONSTRAINTS                int
	TYPECONSTRAINTCOUNT            int
	ATTRIBUTES                     int
	ATTRIBUTECOUNT                 int
	DEFAULTATTRIBUTES              int
	DEFAULTATTRIBUTECOUNT          int
}

type MLOperatorSetId struct {
	DOMAIN  int
	VERSION int
}

type MLOperatorKernelDescription struct {
	DOMAIN                    int
	NAME                      int
	MINIMUMOPERATORSETVERSION int
	EXECUTIONTYPE             int
	TYPECONSTRAINTS           int
	TYPECONSTRAINTCOUNT       int
	DEFAULTATTRIBUTES         int
	DEFAULTATTRIBUTECOUNT     int
	OPTIONS                   int
	EXECUTIONOPTIONS          int
}
