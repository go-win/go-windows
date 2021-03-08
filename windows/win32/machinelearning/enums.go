// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package machinelearning implements the Windows.Win32.MachineLearning namespace.
package machinelearning

type WINML_TENSOR_DATA_TYPE int32

const (
	WINML_TENSOR_UNDEFINED = 0
	WINML_TENSOR_FLOAT = 1
	WINML_TENSOR_UINT8 = 2
	WINML_TENSOR_INT8 = 3
	WINML_TENSOR_UINT16 = 4
	WINML_TENSOR_INT16 = 5
	WINML_TENSOR_INT32 = 6
	WINML_TENSOR_INT64 = 7
	WINML_TENSOR_STRING = 8
	WINML_TENSOR_BOOLEAN = 9
	WINML_TENSOR_FLOAT16 = 10
	WINML_TENSOR_DOUBLE = 11
	WINML_TENSOR_UINT32 = 12
	WINML_TENSOR_UINT64 = 13
	WINML_TENSOR_COMPLEX64 = 14
	WINML_TENSOR_COMPLEX128 = 15
)

type WINML_FEATURE_TYPE int32

const (
	WINML_FEATURE_UNDEFINED = 0
	WINML_FEATURE_TENSOR = 1
	WINML_FEATURE_SEQUENCE = 2
	WINML_FEATURE_MAP = 3
	WINML_FEATURE_IMAGE = 4
)

type WINML_BINDING_TYPE int32

const (
	WINML_BINDING_UNDEFINED = 0
	WINML_BINDING_TENSOR = 1
	WINML_BINDING_SEQUENCE = 2
	WINML_BINDING_MAP = 3
	WINML_BINDING_IMAGE = 4
	WINML_BINDING_RESOURCE = 5
)

type WINML_RUNTIME_TYPE int32

const (
	WINML_RUNTIME_CNTK = 0
)

type MLOperatorAttributeType uint32

const (
	MLOperatorAttributeType_Undefined = 0
	MLOperatorAttributeType_Float = 2
	MLOperatorAttributeType_Int = 3
	MLOperatorAttributeType_String = 4
	MLOperatorAttributeType_FloatArray = 7
	MLOperatorAttributeType_IntArray = 8
	MLOperatorAttributeType_StringArray = 9
)

type MLOperatorTensorDataType uint32

const (
	MLOperatorTensorDataType_Undefined = 0
	MLOperatorTensorDataType_Float = 1
	MLOperatorTensorDataType_UInt8 = 2
	MLOperatorTensorDataType_Int8 = 3
	MLOperatorTensorDataType_UInt16 = 4
	MLOperatorTensorDataType_Int16 = 5
	MLOperatorTensorDataType_Int32 = 6
	MLOperatorTensorDataType_Int64 = 7
	MLOperatorTensorDataType_String = 8
	MLOperatorTensorDataType_Bool = 9
	MLOperatorTensorDataType_Float16 = 10
	MLOperatorTensorDataType_Double = 11
	MLOperatorTensorDataType_UInt32 = 12
	MLOperatorTensorDataType_UInt64 = 13
	MLOperatorTensorDataType_Complex64 = 14
	MLOperatorTensorDataType_Complex128 = 15
)

type MLOperatorEdgeType uint32

const (
	MLOperatorEdgeType_Undefined = 0
	MLOperatorEdgeType_Tensor = 1
)

type MLOperatorParameterOptions uint32

const (
	Single = 0
	Optional = 1
	Variadic = 2
)

type MLOperatorSchemaEdgeTypeFormat int32

const (
	EdgeDescription = 0
	Label = 1
)

type MLOperatorKernelOptions uint32

const (
	None = 0
	AllowDynamicInputShapes = 1
)

type MLOperatorExecutionType uint32

const (
	MLOperatorExecutionType_Undefined = 0
	MLOperatorExecutionType_Cpu = 1
	MLOperatorExecutionType_D3D12 = 2
)

