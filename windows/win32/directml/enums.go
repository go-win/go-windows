// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package directml implements the Windows.Win32.DirectML namespace.
package directml

type DML_TENSOR_DATA_TYPE int32

const (
	DML_TENSOR_DATA_TYPE_UNKNOWN DML_TENSOR_DATA_TYPE = 0
	DML_TENSOR_DATA_TYPE_FLOAT32 DML_TENSOR_DATA_TYPE = 1
	DML_TENSOR_DATA_TYPE_FLOAT16 DML_TENSOR_DATA_TYPE = 2
	DML_TENSOR_DATA_TYPE_UINT32  DML_TENSOR_DATA_TYPE = 3
	DML_TENSOR_DATA_TYPE_UINT16  DML_TENSOR_DATA_TYPE = 4
	DML_TENSOR_DATA_TYPE_UINT8   DML_TENSOR_DATA_TYPE = 5
	DML_TENSOR_DATA_TYPE_INT32   DML_TENSOR_DATA_TYPE = 6
	DML_TENSOR_DATA_TYPE_INT16   DML_TENSOR_DATA_TYPE = 7
	DML_TENSOR_DATA_TYPE_INT8    DML_TENSOR_DATA_TYPE = 8
)

type DML_TENSOR_TYPE int32

const (
	DML_TENSOR_TYPE_INVALID DML_TENSOR_TYPE = 0
	DML_TENSOR_TYPE_BUFFER  DML_TENSOR_TYPE = 1
)

type DML_TENSOR_FLAGS int32

const (
	DML_TENSOR_FLAG_NONE         DML_TENSOR_FLAGS = 0
	DML_TENSOR_FLAG_OWNED_BY_DML DML_TENSOR_FLAGS = 1
)

type DML_OPERATOR_TYPE int32

const (
	DML_OPERATOR_INVALID                           DML_OPERATOR_TYPE = 0
	DML_OPERATOR_ELEMENT_WISE_IDENTITY             DML_OPERATOR_TYPE = 1
	DML_OPERATOR_ELEMENT_WISE_ABS                  DML_OPERATOR_TYPE = 2
	DML_OPERATOR_ELEMENT_WISE_ACOS                 DML_OPERATOR_TYPE = 3
	DML_OPERATOR_ELEMENT_WISE_ADD                  DML_OPERATOR_TYPE = 4
	DML_OPERATOR_ELEMENT_WISE_ASIN                 DML_OPERATOR_TYPE = 5
	DML_OPERATOR_ELEMENT_WISE_ATAN                 DML_OPERATOR_TYPE = 6
	DML_OPERATOR_ELEMENT_WISE_CEIL                 DML_OPERATOR_TYPE = 7
	DML_OPERATOR_ELEMENT_WISE_CLIP                 DML_OPERATOR_TYPE = 8
	DML_OPERATOR_ELEMENT_WISE_COS                  DML_OPERATOR_TYPE = 9
	DML_OPERATOR_ELEMENT_WISE_DIVIDE               DML_OPERATOR_TYPE = 10
	DML_OPERATOR_ELEMENT_WISE_EXP                  DML_OPERATOR_TYPE = 11
	DML_OPERATOR_ELEMENT_WISE_FLOOR                DML_OPERATOR_TYPE = 12
	DML_OPERATOR_ELEMENT_WISE_LOG                  DML_OPERATOR_TYPE = 13
	DML_OPERATOR_ELEMENT_WISE_LOGICAL_AND          DML_OPERATOR_TYPE = 14
	DML_OPERATOR_ELEMENT_WISE_LOGICAL_EQUALS       DML_OPERATOR_TYPE = 15
	DML_OPERATOR_ELEMENT_WISE_LOGICAL_GREATER_THAN DML_OPERATOR_TYPE = 16
	DML_OPERATOR_ELEMENT_WISE_LOGICAL_LESS_THAN    DML_OPERATOR_TYPE = 17
	DML_OPERATOR_ELEMENT_WISE_LOGICAL_NOT          DML_OPERATOR_TYPE = 18
	DML_OPERATOR_ELEMENT_WISE_LOGICAL_OR           DML_OPERATOR_TYPE = 19
	DML_OPERATOR_ELEMENT_WISE_LOGICAL_XOR          DML_OPERATOR_TYPE = 20
	DML_OPERATOR_ELEMENT_WISE_MAX                  DML_OPERATOR_TYPE = 21
	DML_OPERATOR_ELEMENT_WISE_MEAN                 DML_OPERATOR_TYPE = 22
	DML_OPERATOR_ELEMENT_WISE_MIN                  DML_OPERATOR_TYPE = 23
	DML_OPERATOR_ELEMENT_WISE_MULTIPLY             DML_OPERATOR_TYPE = 24
	DML_OPERATOR_ELEMENT_WISE_POW                  DML_OPERATOR_TYPE = 25
	DML_OPERATOR_ELEMENT_WISE_CONSTANT_POW         DML_OPERATOR_TYPE = 26
	DML_OPERATOR_ELEMENT_WISE_RECIP                DML_OPERATOR_TYPE = 27
	DML_OPERATOR_ELEMENT_WISE_SIN                  DML_OPERATOR_TYPE = 28
	DML_OPERATOR_ELEMENT_WISE_SQRT                 DML_OPERATOR_TYPE = 29
	DML_OPERATOR_ELEMENT_WISE_SUBTRACT             DML_OPERATOR_TYPE = 30
	DML_OPERATOR_ELEMENT_WISE_TAN                  DML_OPERATOR_TYPE = 31
	DML_OPERATOR_ELEMENT_WISE_THRESHOLD            DML_OPERATOR_TYPE = 32
	DML_OPERATOR_ELEMENT_WISE_QUANTIZE_LINEAR      DML_OPERATOR_TYPE = 33
	DML_OPERATOR_ELEMENT_WISE_DEQUANTIZE_LINEAR    DML_OPERATOR_TYPE = 34
	DML_OPERATOR_ACTIVATION_ELU                    DML_OPERATOR_TYPE = 35
	DML_OPERATOR_ACTIVATION_HARDMAX                DML_OPERATOR_TYPE = 36
	DML_OPERATOR_ACTIVATION_HARD_SIGMOID           DML_OPERATOR_TYPE = 37
	DML_OPERATOR_ACTIVATION_IDENTITY               DML_OPERATOR_TYPE = 38
	DML_OPERATOR_ACTIVATION_LEAKY_RELU             DML_OPERATOR_TYPE = 39
	DML_OPERATOR_ACTIVATION_LINEAR                 DML_OPERATOR_TYPE = 40
	DML_OPERATOR_ACTIVATION_LOG_SOFTMAX            DML_OPERATOR_TYPE = 41
	DML_OPERATOR_ACTIVATION_PARAMETERIZED_RELU     DML_OPERATOR_TYPE = 42
	DML_OPERATOR_ACTIVATION_PARAMETRIC_SOFTPLUS    DML_OPERATOR_TYPE = 43
	DML_OPERATOR_ACTIVATION_RELU                   DML_OPERATOR_TYPE = 44
	DML_OPERATOR_ACTIVATION_SCALED_ELU             DML_OPERATOR_TYPE = 45
	DML_OPERATOR_ACTIVATION_SCALED_TANH            DML_OPERATOR_TYPE = 46
	DML_OPERATOR_ACTIVATION_SIGMOID                DML_OPERATOR_TYPE = 47
	DML_OPERATOR_ACTIVATION_SOFTMAX                DML_OPERATOR_TYPE = 48
	DML_OPERATOR_ACTIVATION_SOFTPLUS               DML_OPERATOR_TYPE = 49
	DML_OPERATOR_ACTIVATION_SOFTSIGN               DML_OPERATOR_TYPE = 50
	DML_OPERATOR_ACTIVATION_TANH                   DML_OPERATOR_TYPE = 51
	DML_OPERATOR_ACTIVATION_THRESHOLDED_RELU       DML_OPERATOR_TYPE = 52
	DML_OPERATOR_CONVOLUTION                       DML_OPERATOR_TYPE = 53
	DML_OPERATOR_GEMM                              DML_OPERATOR_TYPE = 54
	DML_OPERATOR_REDUCE                            DML_OPERATOR_TYPE = 55
	DML_OPERATOR_AVERAGE_POOLING                   DML_OPERATOR_TYPE = 56
	DML_OPERATOR_LP_POOLING                        DML_OPERATOR_TYPE = 57
	DML_OPERATOR_MAX_POOLING                       DML_OPERATOR_TYPE = 58
	DML_OPERATOR_ROI_POOLING                       DML_OPERATOR_TYPE = 59
	DML_OPERATOR_SLICE                             DML_OPERATOR_TYPE = 60
	DML_OPERATOR_CAST                              DML_OPERATOR_TYPE = 61
	DML_OPERATOR_SPLIT                             DML_OPERATOR_TYPE = 62
	DML_OPERATOR_JOIN                              DML_OPERATOR_TYPE = 63
	DML_OPERATOR_PADDING                           DML_OPERATOR_TYPE = 64
	DML_OPERATOR_VALUE_SCALE_2D                    DML_OPERATOR_TYPE = 65
	DML_OPERATOR_UPSAMPLE_2D                       DML_OPERATOR_TYPE = 66
	DML_OPERATOR_GATHER                            DML_OPERATOR_TYPE = 67
	DML_OPERATOR_SPACE_TO_DEPTH                    DML_OPERATOR_TYPE = 68
	DML_OPERATOR_DEPTH_TO_SPACE                    DML_OPERATOR_TYPE = 69
	DML_OPERATOR_TILE                              DML_OPERATOR_TYPE = 70
	DML_OPERATOR_TOP_K                             DML_OPERATOR_TYPE = 71
	DML_OPERATOR_BATCH_NORMALIZATION               DML_OPERATOR_TYPE = 72
	DML_OPERATOR_MEAN_VARIANCE_NORMALIZATION       DML_OPERATOR_TYPE = 73
	DML_OPERATOR_LOCAL_RESPONSE_NORMALIZATION      DML_OPERATOR_TYPE = 74
	DML_OPERATOR_LP_NORMALIZATION                  DML_OPERATOR_TYPE = 75
	DML_OPERATOR_RNN                               DML_OPERATOR_TYPE = 76
	DML_OPERATOR_LSTM                              DML_OPERATOR_TYPE = 77
	DML_OPERATOR_GRU                               DML_OPERATOR_TYPE = 78
	DML_OPERATOR_ELEMENT_WISE_SIGN                 DML_OPERATOR_TYPE = 79
	DML_OPERATOR_ELEMENT_WISE_IS_NAN               DML_OPERATOR_TYPE = 80
	DML_OPERATOR_ELEMENT_WISE_ERF                  DML_OPERATOR_TYPE = 81
	DML_OPERATOR_ELEMENT_WISE_SINH                 DML_OPERATOR_TYPE = 82
	DML_OPERATOR_ELEMENT_WISE_COSH                 DML_OPERATOR_TYPE = 83
	DML_OPERATOR_ELEMENT_WISE_TANH                 DML_OPERATOR_TYPE = 84
	DML_OPERATOR_ELEMENT_WISE_ASINH                DML_OPERATOR_TYPE = 85
	DML_OPERATOR_ELEMENT_WISE_ACOSH                DML_OPERATOR_TYPE = 86
	DML_OPERATOR_ELEMENT_WISE_ATANH                DML_OPERATOR_TYPE = 87
	DML_OPERATOR_ELEMENT_WISE_IF                   DML_OPERATOR_TYPE = 88
	DML_OPERATOR_ELEMENT_WISE_ADD1                 DML_OPERATOR_TYPE = 89
	DML_OPERATOR_ACTIVATION_SHRINK                 DML_OPERATOR_TYPE = 90
	DML_OPERATOR_MAX_POOLING1                      DML_OPERATOR_TYPE = 91
	DML_OPERATOR_MAX_UNPOOLING                     DML_OPERATOR_TYPE = 92
	DML_OPERATOR_DIAGONAL_MATRIX                   DML_OPERATOR_TYPE = 93
	DML_OPERATOR_SCATTER                           DML_OPERATOR_TYPE = 94
	DML_OPERATOR_ONE_HOT                           DML_OPERATOR_TYPE = 95
	DML_OPERATOR_RESAMPLE                          DML_OPERATOR_TYPE = 96
)

type DML_REDUCE_FUNCTION int32

const (
	DML_REDUCE_FUNCTION_ARGMAX      DML_REDUCE_FUNCTION = 0
	DML_REDUCE_FUNCTION_ARGMIN      DML_REDUCE_FUNCTION = 1
	DML_REDUCE_FUNCTION_AVERAGE     DML_REDUCE_FUNCTION = 2
	DML_REDUCE_FUNCTION_L1          DML_REDUCE_FUNCTION = 3
	DML_REDUCE_FUNCTION_L2          DML_REDUCE_FUNCTION = 4
	DML_REDUCE_FUNCTION_LOG_SUM     DML_REDUCE_FUNCTION = 5
	DML_REDUCE_FUNCTION_LOG_SUM_EXP DML_REDUCE_FUNCTION = 6
	DML_REDUCE_FUNCTION_MAX         DML_REDUCE_FUNCTION = 7
	DML_REDUCE_FUNCTION_MIN         DML_REDUCE_FUNCTION = 8
	DML_REDUCE_FUNCTION_MULTIPLY    DML_REDUCE_FUNCTION = 9
	DML_REDUCE_FUNCTION_SUM         DML_REDUCE_FUNCTION = 10
	DML_REDUCE_FUNCTION_SUM_SQUARE  DML_REDUCE_FUNCTION = 11
)

type DML_MATRIX_TRANSFORM int32

const (
	DML_MATRIX_TRANSFORM_NONE      DML_MATRIX_TRANSFORM = 0
	DML_MATRIX_TRANSFORM_TRANSPOSE DML_MATRIX_TRANSFORM = 1
)

type DML_CONVOLUTION_MODE int32

const (
	DML_CONVOLUTION_MODE_CONVOLUTION       DML_CONVOLUTION_MODE = 0
	DML_CONVOLUTION_MODE_CROSS_CORRELATION DML_CONVOLUTION_MODE = 1
)

type DML_CONVOLUTION_DIRECTION int32

const (
	DML_CONVOLUTION_DIRECTION_FORWARD  DML_CONVOLUTION_DIRECTION = 0
	DML_CONVOLUTION_DIRECTION_BACKWARD DML_CONVOLUTION_DIRECTION = 1
)

type DML_PADDING_MODE int32

const (
	DML_PADDING_MODE_CONSTANT   DML_PADDING_MODE = 0
	DML_PADDING_MODE_EDGE       DML_PADDING_MODE = 1
	DML_PADDING_MODE_REFLECTION DML_PADDING_MODE = 2
)

type DML_INTERPOLATION_MODE int32

const (
	DML_INTERPOLATION_MODE_NEAREST_NEIGHBOR DML_INTERPOLATION_MODE = 0
	DML_INTERPOLATION_MODE_LINEAR           DML_INTERPOLATION_MODE = 1
)

type DML_RECURRENT_NETWORK_DIRECTION int32

const (
	DML_RECURRENT_NETWORK_DIRECTION_FORWARD       DML_RECURRENT_NETWORK_DIRECTION = 0
	DML_RECURRENT_NETWORK_DIRECTION_BACKWARD      DML_RECURRENT_NETWORK_DIRECTION = 1
	DML_RECURRENT_NETWORK_DIRECTION_BIDIRECTIONAL DML_RECURRENT_NETWORK_DIRECTION = 2
)

type DML_FEATURE_LEVEL int32

const (
	DML_FEATURE_LEVEL_1_0 DML_FEATURE_LEVEL = 4096
	DML_FEATURE_LEVEL_2_0 DML_FEATURE_LEVEL = 8192
)

type DML_FEATURE int32

const (
	DML_FEATURE_TENSOR_DATA_TYPE_SUPPORT DML_FEATURE = 0
	DML_FEATURE_FEATURE_LEVELS           DML_FEATURE = 1
)

type DML_EXECUTION_FLAGS int32

const (
	DML_EXECUTION_FLAG_NONE                             DML_EXECUTION_FLAGS = 0
	DML_EXECUTION_FLAG_ALLOW_HALF_PRECISION_COMPUTATION DML_EXECUTION_FLAGS = 1
	DML_EXECUTION_FLAG_DISABLE_META_COMMANDS            DML_EXECUTION_FLAGS = 2
	DML_EXECUTION_FLAG_DESCRIPTORS_VOLATILE             DML_EXECUTION_FLAGS = 4
)

type DML_CREATE_DEVICE_FLAGS int32

const (
	DML_CREATE_DEVICE_FLAG_NONE  DML_CREATE_DEVICE_FLAGS = 0
	DML_CREATE_DEVICE_FLAG_DEBUG DML_CREATE_DEVICE_FLAGS = 1
)

type DML_BINDING_TYPE int32

const (
	DML_BINDING_TYPE_NONE         DML_BINDING_TYPE = 0
	DML_BINDING_TYPE_BUFFER       DML_BINDING_TYPE = 1
	DML_BINDING_TYPE_BUFFER_ARRAY DML_BINDING_TYPE = 2
)
