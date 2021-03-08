// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package js implements the Windows.Win32.Js namespace.
package js

type JsRuntimeVersion int32

const (
	JsRuntimeVersion10 = 0
	JsRuntimeVersion11 = 1
	JsRuntimeVersionEdge = -1
)

type JsErrorCode uint32

const (
	JsNoError = 0
	JsErrorCategoryUsage = 65536
	JsErrorInvalidArgument = 65537
	JsErrorNullArgument = 65538
	JsErrorNoCurrentContext = 65539
	JsErrorInExceptionState = 65540
	JsErrorNotImplemented = 65541
	JsErrorWrongThread = 65542
	JsErrorRuntimeInUse = 65543
	JsErrorBadSerializedScript = 65544
	JsErrorInDisabledState = 65545
	JsErrorCannotDisableExecution = 65546
	JsErrorHeapEnumInProgress = 65547
	JsErrorArgumentNotObject = 65548
	JsErrorInProfileCallback = 65549
	JsErrorInThreadServiceCallback = 65550
	JsErrorCannotSerializeDebugScript = 65551
	JsErrorAlreadyDebuggingContext = 65552
	JsErrorAlreadyProfilingContext = 65553
	JsErrorIdleNotEnabled = 65554
	JsErrorCategoryEngine = 131072
	JsErrorOutOfMemory = 131073
	JsErrorCategoryScript = 196608
	JsErrorScriptException = 196609
	JsErrorScriptCompile = 196610
	JsErrorScriptTerminated = 196611
	JsErrorScriptEvalDisabled = 196612
	JsErrorCategoryFatal = 262144
	JsErrorFatal = 262145
)

type JsRuntimeAttributes int32

const (
	JsRuntimeAttributeNone = 0
	JsRuntimeAttributeDisableBackgroundWork = 1
	JsRuntimeAttributeAllowScriptInterrupt = 2
	JsRuntimeAttributeEnableIdleProcessing = 4
	JsRuntimeAttributeDisableNativeCodeGeneration = 8
	JsRuntimeAttributeDisableEval = 16
)

type JsMemoryEventType int32

const (
	JsMemoryAllocate = 0
	JsMemoryFree = 1
	JsMemoryFailure = 2
)

type JsValueType int32

const (
	JsUndefined = 0
	JsNull = 1
	JsNumber = 2
	JsString = 3
	JsBoolean = 4
	JsObject = 5
	JsFunction = 6
	JsError = 7
	JsArray = 8
)
