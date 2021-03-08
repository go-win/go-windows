// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package windowsimagingcomponent implements the Windows.Win32.WindowsImagingComponent namespace.
package windowsimagingcomponent

type WICColorContextType int32

const (
	WICColorContextUninitialized  WICColorContextType = 0
	WICColorContextProfile        WICColorContextType = 1
	WICColorContextExifColorSpace WICColorContextType = 2
)

type WICBitmapCreateCacheOption int32

const (
	WICBitmapNoCache                       WICBitmapCreateCacheOption = 0
	WICBitmapCacheOnDemand                 WICBitmapCreateCacheOption = 1
	WICBitmapCacheOnLoad                   WICBitmapCreateCacheOption = 2
	WICBITMAPCREATECACHEOPTION_FORCE_DWORD WICBitmapCreateCacheOption = 2147483647
)

type WICDecodeOptions int32

const (
	WICDecodeMetadataCacheOnDemand     WICDecodeOptions = 0
	WICDecodeMetadataCacheOnLoad       WICDecodeOptions = 1
	WICMETADATACACHEOPTION_FORCE_DWORD WICDecodeOptions = 2147483647
)

type WICBitmapEncoderCacheOption int32

const (
	WICBitmapEncoderCacheInMemory           WICBitmapEncoderCacheOption = 0
	WICBitmapEncoderCacheTempFile           WICBitmapEncoderCacheOption = 1
	WICBitmapEncoderNoCache                 WICBitmapEncoderCacheOption = 2
	WICBITMAPENCODERCACHEOPTION_FORCE_DWORD WICBitmapEncoderCacheOption = 2147483647
)

type WICComponentType int32

const (
	WICDecoder                   WICComponentType = 1
	WICEncoder                   WICComponentType = 2
	WICPixelFormatConverter      WICComponentType = 4
	WICMetadataReader            WICComponentType = 8
	WICMetadataWriter            WICComponentType = 16
	WICPixelFormat               WICComponentType = 32
	WICAllComponents             WICComponentType = 63
	WICCOMPONENTTYPE_FORCE_DWORD WICComponentType = 2147483647
)

type WICComponentEnumerateOptions int32

const (
	WICComponentEnumerateDefault             WICComponentEnumerateOptions = 0
	WICComponentEnumerateRefresh             WICComponentEnumerateOptions = 1
	WICComponentEnumerateDisabled            WICComponentEnumerateOptions = -2147483648
	WICComponentEnumerateUnsigned            WICComponentEnumerateOptions = 1073741824
	WICComponentEnumerateBuiltInOnly         WICComponentEnumerateOptions = 536870912
	WICCOMPONENTENUMERATEOPTIONS_FORCE_DWORD WICComponentEnumerateOptions = 2147483647
)

type WICBitmapInterpolationMode int32

const (
	WICBitmapInterpolationModeNearestNeighbor  WICBitmapInterpolationMode = 0
	WICBitmapInterpolationModeLinear           WICBitmapInterpolationMode = 1
	WICBitmapInterpolationModeCubic            WICBitmapInterpolationMode = 2
	WICBitmapInterpolationModeFant             WICBitmapInterpolationMode = 3
	WICBitmapInterpolationModeHighQualityCubic WICBitmapInterpolationMode = 4
	WICBITMAPINTERPOLATIONMODE_FORCE_DWORD     WICBitmapInterpolationMode = 2147483647
)

type WICBitmapPaletteType int32

const (
	WICBitmapPaletteTypeCustom           WICBitmapPaletteType = 0
	WICBitmapPaletteTypeMedianCut        WICBitmapPaletteType = 1
	WICBitmapPaletteTypeFixedBW          WICBitmapPaletteType = 2
	WICBitmapPaletteTypeFixedHalftone8   WICBitmapPaletteType = 3
	WICBitmapPaletteTypeFixedHalftone27  WICBitmapPaletteType = 4
	WICBitmapPaletteTypeFixedHalftone64  WICBitmapPaletteType = 5
	WICBitmapPaletteTypeFixedHalftone125 WICBitmapPaletteType = 6
	WICBitmapPaletteTypeFixedHalftone216 WICBitmapPaletteType = 7
	WICBitmapPaletteTypeFixedWebPalette  WICBitmapPaletteType = 7
	WICBitmapPaletteTypeFixedHalftone252 WICBitmapPaletteType = 8
	WICBitmapPaletteTypeFixedHalftone256 WICBitmapPaletteType = 9
	WICBitmapPaletteTypeFixedGray4       WICBitmapPaletteType = 10
	WICBitmapPaletteTypeFixedGray16      WICBitmapPaletteType = 11
	WICBitmapPaletteTypeFixedGray256     WICBitmapPaletteType = 12
	WICBITMAPPALETTETYPE_FORCE_DWORD     WICBitmapPaletteType = 2147483647
)

type WICBitmapDitherType int32

const (
	WICBitmapDitherTypeNone           WICBitmapDitherType = 0
	WICBitmapDitherTypeSolid          WICBitmapDitherType = 0
	WICBitmapDitherTypeOrdered4x4     WICBitmapDitherType = 1
	WICBitmapDitherTypeOrdered8x8     WICBitmapDitherType = 2
	WICBitmapDitherTypeOrdered16x16   WICBitmapDitherType = 3
	WICBitmapDitherTypeSpiral4x4      WICBitmapDitherType = 4
	WICBitmapDitherTypeSpiral8x8      WICBitmapDitherType = 5
	WICBitmapDitherTypeDualSpiral4x4  WICBitmapDitherType = 6
	WICBitmapDitherTypeDualSpiral8x8  WICBitmapDitherType = 7
	WICBitmapDitherTypeErrorDiffusion WICBitmapDitherType = 8
	WICBITMAPDITHERTYPE_FORCE_DWORD   WICBitmapDitherType = 2147483647
)

type WICBitmapAlphaChannelOption int32

const (
	WICBitmapUseAlpha                        WICBitmapAlphaChannelOption = 0
	WICBitmapUsePremultipliedAlpha           WICBitmapAlphaChannelOption = 1
	WICBitmapIgnoreAlpha                     WICBitmapAlphaChannelOption = 2
	WICBITMAPALPHACHANNELOPTIONS_FORCE_DWORD WICBitmapAlphaChannelOption = 2147483647
)

type WICBitmapTransformOptions int32

const (
	WICBitmapTransformRotate0             WICBitmapTransformOptions = 0
	WICBitmapTransformRotate90            WICBitmapTransformOptions = 1
	WICBitmapTransformRotate180           WICBitmapTransformOptions = 2
	WICBitmapTransformRotate270           WICBitmapTransformOptions = 3
	WICBitmapTransformFlipHorizontal      WICBitmapTransformOptions = 8
	WICBitmapTransformFlipVertical        WICBitmapTransformOptions = 16
	WICBITMAPTRANSFORMOPTIONS_FORCE_DWORD WICBitmapTransformOptions = 2147483647
)

type WICBitmapLockFlags int32

const (
	WICBitmapLockRead              WICBitmapLockFlags = 1
	WICBitmapLockWrite             WICBitmapLockFlags = 2
	WICBITMAPLOCKFLAGS_FORCE_DWORD WICBitmapLockFlags = 2147483647
)

type WICBitmapDecoderCapabilities int32

const (
	WICBitmapDecoderCapabilitySameEncoder          WICBitmapDecoderCapabilities = 1
	WICBitmapDecoderCapabilityCanDecodeAllImages   WICBitmapDecoderCapabilities = 2
	WICBitmapDecoderCapabilityCanDecodeSomeImages  WICBitmapDecoderCapabilities = 4
	WICBitmapDecoderCapabilityCanEnumerateMetadata WICBitmapDecoderCapabilities = 8
	WICBitmapDecoderCapabilityCanDecodeThumbnail   WICBitmapDecoderCapabilities = 16
	WICBITMAPDECODERCAPABILITIES_FORCE_DWORD       WICBitmapDecoderCapabilities = 2147483647
)

type WICProgressOperation int32

const (
	WICProgressOperationCopyPixels   WICProgressOperation = 1
	WICProgressOperationWritePixels  WICProgressOperation = 2
	WICProgressOperationAll          WICProgressOperation = 65535
	WICPROGRESSOPERATION_FORCE_DWORD WICProgressOperation = 2147483647
)

type WICProgressNotification int32

const (
	WICProgressNotificationBegin        WICProgressNotification = 65536
	WICProgressNotificationEnd          WICProgressNotification = 131072
	WICProgressNotificationFrequent     WICProgressNotification = 262144
	WICProgressNotificationAll          WICProgressNotification = -65536
	WICPROGRESSNOTIFICATION_FORCE_DWORD WICProgressNotification = 2147483647
)

type WICComponentSigning int32

const (
	WICComponentSigned              WICComponentSigning = 1
	WICComponentUnsigned            WICComponentSigning = 2
	WICComponentSafe                WICComponentSigning = 4
	WICComponentDisabled            WICComponentSigning = -2147483648
	WICCOMPONENTSIGNING_FORCE_DWORD WICComponentSigning = 2147483647
)

type WICGifLogicalScreenDescriptorProperties uint32

const (
	WICGifLogicalScreenSignature                        WICGifLogicalScreenDescriptorProperties = 1
	WICGifLogicalScreenDescriptorWidth                  WICGifLogicalScreenDescriptorProperties = 2
	WICGifLogicalScreenDescriptorHeight                 WICGifLogicalScreenDescriptorProperties = 3
	WICGifLogicalScreenDescriptorGlobalColorTableFlag   WICGifLogicalScreenDescriptorProperties = 4
	WICGifLogicalScreenDescriptorColorResolution        WICGifLogicalScreenDescriptorProperties = 5
	WICGifLogicalScreenDescriptorSortFlag               WICGifLogicalScreenDescriptorProperties = 6
	WICGifLogicalScreenDescriptorGlobalColorTableSize   WICGifLogicalScreenDescriptorProperties = 7
	WICGifLogicalScreenDescriptorBackgroundColorIndex   WICGifLogicalScreenDescriptorProperties = 8
	WICGifLogicalScreenDescriptorPixelAspectRatio       WICGifLogicalScreenDescriptorProperties = 9
	WICGifLogicalScreenDescriptorProperties_FORCE_DWORD WICGifLogicalScreenDescriptorProperties = 2147483647
)

type WICGifImageDescriptorProperties uint32

const (
	WICGifImageDescriptorLeft                   WICGifImageDescriptorProperties = 1
	WICGifImageDescriptorTop                    WICGifImageDescriptorProperties = 2
	WICGifImageDescriptorWidth                  WICGifImageDescriptorProperties = 3
	WICGifImageDescriptorHeight                 WICGifImageDescriptorProperties = 4
	WICGifImageDescriptorLocalColorTableFlag    WICGifImageDescriptorProperties = 5
	WICGifImageDescriptorInterlaceFlag          WICGifImageDescriptorProperties = 6
	WICGifImageDescriptorSortFlag               WICGifImageDescriptorProperties = 7
	WICGifImageDescriptorLocalColorTableSize    WICGifImageDescriptorProperties = 8
	WICGifImageDescriptorProperties_FORCE_DWORD WICGifImageDescriptorProperties = 2147483647
)

type WICGifGraphicControlExtensionProperties uint32

const (
	WICGifGraphicControlExtensionDisposal               WICGifGraphicControlExtensionProperties = 1
	WICGifGraphicControlExtensionUserInputFlag          WICGifGraphicControlExtensionProperties = 2
	WICGifGraphicControlExtensionTransparencyFlag       WICGifGraphicControlExtensionProperties = 3
	WICGifGraphicControlExtensionDelay                  WICGifGraphicControlExtensionProperties = 4
	WICGifGraphicControlExtensionTransparentColorIndex  WICGifGraphicControlExtensionProperties = 5
	WICGifGraphicControlExtensionProperties_FORCE_DWORD WICGifGraphicControlExtensionProperties = 2147483647
)

type WICGifApplicationExtensionProperties uint32

const (
	WICGifApplicationExtensionApplication            WICGifApplicationExtensionProperties = 1
	WICGifApplicationExtensionData                   WICGifApplicationExtensionProperties = 2
	WICGifApplicationExtensionProperties_FORCE_DWORD WICGifApplicationExtensionProperties = 2147483647
)

type WICGifCommentExtensionProperties uint32

const (
	WICGifCommentExtensionText                   WICGifCommentExtensionProperties = 1
	WICGifCommentExtensionProperties_FORCE_DWORD WICGifCommentExtensionProperties = 2147483647
)

type WICJpegCommentProperties uint32

const (
	WICJpegCommentText                   WICJpegCommentProperties = 1
	WICJpegCommentProperties_FORCE_DWORD WICJpegCommentProperties = 2147483647
)

type WICJpegLuminanceProperties uint32

const (
	WICJpegLuminanceTable                  WICJpegLuminanceProperties = 1
	WICJpegLuminanceProperties_FORCE_DWORD WICJpegLuminanceProperties = 2147483647
)

type WICJpegChrominanceProperties uint32

const (
	WICJpegChrominanceTable                  WICJpegChrominanceProperties = 1
	WICJpegChrominanceProperties_FORCE_DWORD WICJpegChrominanceProperties = 2147483647
)

type WIC8BIMIptcProperties uint32

const (
	WIC8BIMIptcPString                WIC8BIMIptcProperties = 0
	WIC8BIMIptcEmbeddedIPTC           WIC8BIMIptcProperties = 1
	WIC8BIMIptcProperties_FORCE_DWORD WIC8BIMIptcProperties = 2147483647
)

type WIC8BIMResolutionInfoProperties uint32

const (
	WIC8BIMResolutionInfoPString                WIC8BIMResolutionInfoProperties = 1
	WIC8BIMResolutionInfoHResolution            WIC8BIMResolutionInfoProperties = 2
	WIC8BIMResolutionInfoHResolutionUnit        WIC8BIMResolutionInfoProperties = 3
	WIC8BIMResolutionInfoWidthUnit              WIC8BIMResolutionInfoProperties = 4
	WIC8BIMResolutionInfoVResolution            WIC8BIMResolutionInfoProperties = 5
	WIC8BIMResolutionInfoVResolutionUnit        WIC8BIMResolutionInfoProperties = 6
	WIC8BIMResolutionInfoHeightUnit             WIC8BIMResolutionInfoProperties = 7
	WIC8BIMResolutionInfoProperties_FORCE_DWORD WIC8BIMResolutionInfoProperties = 2147483647
)

type WIC8BIMIptcDigestProperties uint32

const (
	WIC8BIMIptcDigestPString                WIC8BIMIptcDigestProperties = 1
	WIC8BIMIptcDigestIptcDigest             WIC8BIMIptcDigestProperties = 2
	WIC8BIMIptcDigestProperties_FORCE_DWORD WIC8BIMIptcDigestProperties = 2147483647
)

type WICPngGamaProperties uint32

const (
	WICPngGamaGamma                  WICPngGamaProperties = 1
	WICPngGamaProperties_FORCE_DWORD WICPngGamaProperties = 2147483647
)

type WICPngBkgdProperties uint32

const (
	WICPngBkgdBackgroundColor        WICPngBkgdProperties = 1
	WICPngBkgdProperties_FORCE_DWORD WICPngBkgdProperties = 2147483647
)

type WICPngItxtProperties uint32

const (
	WICPngItxtKeyword                WICPngItxtProperties = 1
	WICPngItxtCompressionFlag        WICPngItxtProperties = 2
	WICPngItxtLanguageTag            WICPngItxtProperties = 3
	WICPngItxtTranslatedKeyword      WICPngItxtProperties = 4
	WICPngItxtText                   WICPngItxtProperties = 5
	WICPngItxtProperties_FORCE_DWORD WICPngItxtProperties = 2147483647
)

type WICPngChrmProperties uint32

const (
	WICPngChrmWhitePointX            WICPngChrmProperties = 1
	WICPngChrmWhitePointY            WICPngChrmProperties = 2
	WICPngChrmRedX                   WICPngChrmProperties = 3
	WICPngChrmRedY                   WICPngChrmProperties = 4
	WICPngChrmGreenX                 WICPngChrmProperties = 5
	WICPngChrmGreenY                 WICPngChrmProperties = 6
	WICPngChrmBlueX                  WICPngChrmProperties = 7
	WICPngChrmBlueY                  WICPngChrmProperties = 8
	WICPngChrmProperties_FORCE_DWORD WICPngChrmProperties = 2147483647
)

type WICPngHistProperties uint32

const (
	WICPngHistFrequencies            WICPngHistProperties = 1
	WICPngHistProperties_FORCE_DWORD WICPngHistProperties = 2147483647
)

type WICPngIccpProperties uint32

const (
	WICPngIccpProfileName            WICPngIccpProperties = 1
	WICPngIccpProfileData            WICPngIccpProperties = 2
	WICPngIccpProperties_FORCE_DWORD WICPngIccpProperties = 2147483647
)

type WICPngSrgbProperties uint32

const (
	WICPngSrgbRenderingIntent        WICPngSrgbProperties = 1
	WICPngSrgbProperties_FORCE_DWORD WICPngSrgbProperties = 2147483647
)

type WICPngTimeProperties uint32

const (
	WICPngTimeYear                   WICPngTimeProperties = 1
	WICPngTimeMonth                  WICPngTimeProperties = 2
	WICPngTimeDay                    WICPngTimeProperties = 3
	WICPngTimeHour                   WICPngTimeProperties = 4
	WICPngTimeMinute                 WICPngTimeProperties = 5
	WICPngTimeSecond                 WICPngTimeProperties = 6
	WICPngTimeProperties_FORCE_DWORD WICPngTimeProperties = 2147483647
)

type WICHeifProperties uint32

const (
	WICHeifOrientation            WICHeifProperties = 1
	WICHeifProperties_FORCE_DWORD WICHeifProperties = 2147483647
)

type WICHeifHdrProperties uint32

const (
	WICHeifHdrMaximumLuminanceLevel                 WICHeifHdrProperties = 1
	WICHeifHdrMaximumFrameAverageLuminanceLevel     WICHeifHdrProperties = 2
	WICHeifHdrMinimumMasteringDisplayLuminanceLevel WICHeifHdrProperties = 3
	WICHeifHdrMaximumMasteringDisplayLuminanceLevel WICHeifHdrProperties = 4
	WICHeifHdrCustomVideoPrimaries                  WICHeifHdrProperties = 5
	WICHeifHdrProperties_FORCE_DWORD                WICHeifHdrProperties = 2147483647
)

type WICWebpAnimProperties uint32

const (
	WICWebpAnimLoopCount              WICWebpAnimProperties = 1
	WICWebpAnimProperties_FORCE_DWORD WICWebpAnimProperties = 2147483647
)

type WICWebpAnmfProperties uint32

const (
	WICWebpAnmfFrameDuration          WICWebpAnmfProperties = 1
	WICWebpAnmfProperties_FORCE_DWORD WICWebpAnmfProperties = 2147483647
)

type WICSectionAccessLevel uint32

const (
	WICSectionAccessLevelRead         WICSectionAccessLevel = 1
	WICSectionAccessLevelReadWrite    WICSectionAccessLevel = 3
	WICSectionAccessLevel_FORCE_DWORD WICSectionAccessLevel = 2147483647
)

type WICPixelFormatNumericRepresentation uint32

const (
	WICPixelFormatNumericRepresentationUnspecified     WICPixelFormatNumericRepresentation = 0
	WICPixelFormatNumericRepresentationIndexed         WICPixelFormatNumericRepresentation = 1
	WICPixelFormatNumericRepresentationUnsignedInteger WICPixelFormatNumericRepresentation = 2
	WICPixelFormatNumericRepresentationSignedInteger   WICPixelFormatNumericRepresentation = 3
	WICPixelFormatNumericRepresentationFixed           WICPixelFormatNumericRepresentation = 4
	WICPixelFormatNumericRepresentationFloat           WICPixelFormatNumericRepresentation = 5
	WICPixelFormatNumericRepresentation_FORCE_DWORD    WICPixelFormatNumericRepresentation = 2147483647
)

type WICPlanarOptions int32

const (
	WICPlanarOptionsDefault             WICPlanarOptions = 0
	WICPlanarOptionsPreserveSubsampling WICPlanarOptions = 1
	WICPLANAROPTIONS_FORCE_DWORD        WICPlanarOptions = 2147483647
)

type WICJpegIndexingOptions uint32

const (
	WICJpegIndexingOptionsGenerateOnDemand WICJpegIndexingOptions = 0
	WICJpegIndexingOptionsGenerateOnLoad   WICJpegIndexingOptions = 1
	WICJpegIndexingOptions_FORCE_DWORD     WICJpegIndexingOptions = 2147483647
)

type WICJpegTransferMatrix uint32

const (
	WICJpegTransferMatrixIdentity     WICJpegTransferMatrix = 0
	WICJpegTransferMatrixBT601        WICJpegTransferMatrix = 1
	WICJpegTransferMatrix_FORCE_DWORD WICJpegTransferMatrix = 2147483647
)

type WICJpegScanType uint32

const (
	WICJpegScanTypeInterleaved      WICJpegScanType = 0
	WICJpegScanTypePlanarComponents WICJpegScanType = 1
	WICJpegScanTypeProgressive      WICJpegScanType = 2
	WICJpegScanType_FORCE_DWORD     WICJpegScanType = 2147483647
)

type WICTiffCompressionOption int32

const (
	WICTiffCompressionDontCare           WICTiffCompressionOption = 0
	WICTiffCompressionNone               WICTiffCompressionOption = 1
	WICTiffCompressionCCITT3             WICTiffCompressionOption = 2
	WICTiffCompressionCCITT4             WICTiffCompressionOption = 3
	WICTiffCompressionLZW                WICTiffCompressionOption = 4
	WICTiffCompressionRLE                WICTiffCompressionOption = 5
	WICTiffCompressionZIP                WICTiffCompressionOption = 6
	WICTiffCompressionLZWHDifferencing   WICTiffCompressionOption = 7
	WICTIFFCOMPRESSIONOPTION_FORCE_DWORD WICTiffCompressionOption = 2147483647
)

type WICJpegYCrCbSubsamplingOption int32

const (
	WICJpegYCrCbSubsamplingDefault      WICJpegYCrCbSubsamplingOption = 0
	WICJpegYCrCbSubsampling420          WICJpegYCrCbSubsamplingOption = 1
	WICJpegYCrCbSubsampling422          WICJpegYCrCbSubsamplingOption = 2
	WICJpegYCrCbSubsampling444          WICJpegYCrCbSubsamplingOption = 3
	WICJpegYCrCbSubsampling440          WICJpegYCrCbSubsamplingOption = 4
	WICJPEGYCRCBSUBSAMPLING_FORCE_DWORD WICJpegYCrCbSubsamplingOption = 2147483647
)

type WICPngFilterOption int32

const (
	WICPngFilterUnspecified        WICPngFilterOption = 0
	WICPngFilterNone               WICPngFilterOption = 1
	WICPngFilterSub                WICPngFilterOption = 2
	WICPngFilterUp                 WICPngFilterOption = 3
	WICPngFilterAverage            WICPngFilterOption = 4
	WICPngFilterPaeth              WICPngFilterOption = 5
	WICPngFilterAdaptive           WICPngFilterOption = 6
	WICPNGFILTEROPTION_FORCE_DWORD WICPngFilterOption = 2147483647
)

type WICNamedWhitePoint int32

const (
	WICWhitePointDefault           WICNamedWhitePoint = 1
	WICWhitePointDaylight          WICNamedWhitePoint = 2
	WICWhitePointCloudy            WICNamedWhitePoint = 4
	WICWhitePointShade             WICNamedWhitePoint = 8
	WICWhitePointTungsten          WICNamedWhitePoint = 16
	WICWhitePointFluorescent       WICNamedWhitePoint = 32
	WICWhitePointFlash             WICNamedWhitePoint = 64
	WICWhitePointUnderwater        WICNamedWhitePoint = 128
	WICWhitePointCustom            WICNamedWhitePoint = 256
	WICWhitePointAutoWhiteBalance  WICNamedWhitePoint = 512
	WICWhitePointAsShot            WICNamedWhitePoint = 1
	WICNAMEDWHITEPOINT_FORCE_DWORD WICNamedWhitePoint = 2147483647
)

type WICRawCapabilities int32

const (
	WICRawCapabilityNotSupported   WICRawCapabilities = 0
	WICRawCapabilityGetSupported   WICRawCapabilities = 1
	WICRawCapabilityFullySupported WICRawCapabilities = 2
	WICRAWCAPABILITIES_FORCE_DWORD WICRawCapabilities = 2147483647
)

type WICRawRotationCapabilities int32

const (
	WICRawRotationCapabilityNotSupported           WICRawRotationCapabilities = 0
	WICRawRotationCapabilityGetSupported           WICRawRotationCapabilities = 1
	WICRawRotationCapabilityNinetyDegreesSupported WICRawRotationCapabilities = 2
	WICRawRotationCapabilityFullySupported         WICRawRotationCapabilities = 3
	WICRAWROTATIONCAPABILITIES_FORCE_DWORD         WICRawRotationCapabilities = 2147483647
)

type WICRawParameterSet int32

const (
	WICAsShotParameterSet          WICRawParameterSet = 1
	WICUserAdjustedParameterSet    WICRawParameterSet = 2
	WICAutoAdjustedParameterSet    WICRawParameterSet = 3
	WICRAWPARAMETERSET_FORCE_DWORD WICRawParameterSet = 2147483647
)

type WICRawRenderMode int32

const (
	WICRawRenderModeDraft        WICRawRenderMode = 1
	WICRawRenderModeNormal       WICRawRenderMode = 2
	WICRawRenderModeBestQuality  WICRawRenderMode = 3
	WICRAWRENDERMODE_FORCE_DWORD WICRawRenderMode = 2147483647
)

type WICDdsDimension int32

const (
	WICDdsTexture1D           WICDdsDimension = 0
	WICDdsTexture2D           WICDdsDimension = 1
	WICDdsTexture3D           WICDdsDimension = 2
	WICDdsTextureCube         WICDdsDimension = 3
	WICDDSTEXTURE_FORCE_DWORD WICDdsDimension = 2147483647
)

type WICDdsAlphaMode int32

const (
	WICDdsAlphaModeUnknown       WICDdsAlphaMode = 0
	WICDdsAlphaModeStraight      WICDdsAlphaMode = 1
	WICDdsAlphaModePremultiplied WICDdsAlphaMode = 2
	WICDdsAlphaModeOpaque        WICDdsAlphaMode = 3
	WICDdsAlphaModeCustom        WICDdsAlphaMode = 4
	WICDDSALPHAMODE_FORCE_DWORD  WICDdsAlphaMode = 2147483647
)

type WICMetadataCreationOptions int32

const (
	WICMetadataCreationDefault      WICMetadataCreationOptions = 0
	WICMetadataCreationAllowUnknown WICMetadataCreationOptions = 0
	WICMetadataCreationFailUnknown  WICMetadataCreationOptions = 65536
	WICMetadataCreationMask         WICMetadataCreationOptions = -65536
)

type WICPersistOptions int32

const (
	WICPersistOptionDefault       WICPersistOptions = 0
	WICPersistOptionLittleEndian  WICPersistOptions = 0
	WICPersistOptionBigEndian     WICPersistOptions = 1
	WICPersistOptionStrictFormat  WICPersistOptions = 2
	WICPersistOptionNoCacheStream WICPersistOptions = 4
	WICPersistOptionPreferUTF8    WICPersistOptions = 8
	WICPersistOptionMask          WICPersistOptions = 65535
)
