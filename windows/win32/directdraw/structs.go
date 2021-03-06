// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package directdraw implements the Windows.Win32.DirectDraw namespace.
package directdraw

type DDFXROP struct {
}

type DDARGB struct {
	BLUE  int
	GREEN int
	RED   int
	ALPHA int
}

type DDRGBA struct {
	RED   int
	GREEN int
	BLUE  int
	ALPHA int
}

type DDBLTFX struct {
	DWSIZE                   int
	DWDDFX                   int
	DWROP                    int
	DWDDROP                  int
	DWROTATIONANGLE          int
	DWZBUFFEROPCODE          int
	DWZBUFFERLOW             int
	DWZBUFFERHIGH            int
	DWZBUFFERBASEDEST        int
	DWZDESTCONSTBITDEPTH     int
	Anonymous1               int
	DWZSRCCONSTBITDEPTH      int
	Anonymous2               int
	DWALPHAEDGEBLENDBITDEPTH int
	DWALPHAEDGEBLEND         int
	DWRESERVED               int
	DWALPHADESTCONSTBITDEPTH int
	Anonymous3               int
	DWALPHASRCCONSTBITDEPTH  int
	Anonymous4               int
	Anonymous5               int
	DDCKDESTCOLORKEY         int
	DDCKSRCCOLORKEY          int
}

type DDSCAPS struct {
	DWCAPS int
}

type DDOSCAPS struct {
	DWCAPS int
}

type DDSCAPSEX struct {
	DWCAPS2   int
	DWCAPS3   int
	Anonymous int
}

type DDSCAPS2 struct {
	DWCAPS    int
	DWCAPS2   int
	DWCAPS3   int
	Anonymous int
}

type DDCAPS_DX1 struct {
	DWSIZE                         int
	DWCAPS                         int
	DWCAPS2                        int
	DWCKEYCAPS                     int
	DWFXCAPS                       int
	DWFXALPHACAPS                  int
	DWPALCAPS                      int
	DWSVCAPS                       int
	DWALPHABLTCONSTBITDEPTHS       int
	DWALPHABLTPIXELBITDEPTHS       int
	DWALPHABLTSURFACEBITDEPTHS     int
	DWALPHAOVERLAYCONSTBITDEPTHS   int
	DWALPHAOVERLAYPIXELBITDEPTHS   int
	DWALPHAOVERLAYSURFACEBITDEPTHS int
	DWZBUFFERBITDEPTHS             int
	DWVIDMEMTOTAL                  int
	DWVIDMEMFREE                   int
	DWMAXVISIBLEOVERLAYS           int
	DWCURRVISIBLEOVERLAYS          int
	DWNUMFOURCCCODES               int
	DWALIGNBOUNDARYSRC             int
	DWALIGNSIZESRC                 int
	DWALIGNBOUNDARYDEST            int
	DWALIGNSIZEDEST                int
	DWALIGNSTRIDEALIGN             int
	DWROPS                         int
	DDSCAPS                        int
	DWMINOVERLAYSTRETCH            int
	DWMAXOVERLAYSTRETCH            int
	DWMINLIVEVIDEOSTRETCH          int
	DWMAXLIVEVIDEOSTRETCH          int
	DWMINHWCODECSTRETCH            int
	DWMAXHWCODECSTRETCH            int
	DWRESERVED1                    int
	DWRESERVED2                    int
	DWRESERVED3                    int
}

type DDCAPS_DX3 struct {
	DWSIZE                         int
	DWCAPS                         int
	DWCAPS2                        int
	DWCKEYCAPS                     int
	DWFXCAPS                       int
	DWFXALPHACAPS                  int
	DWPALCAPS                      int
	DWSVCAPS                       int
	DWALPHABLTCONSTBITDEPTHS       int
	DWALPHABLTPIXELBITDEPTHS       int
	DWALPHABLTSURFACEBITDEPTHS     int
	DWALPHAOVERLAYCONSTBITDEPTHS   int
	DWALPHAOVERLAYPIXELBITDEPTHS   int
	DWALPHAOVERLAYSURFACEBITDEPTHS int
	DWZBUFFERBITDEPTHS             int
	DWVIDMEMTOTAL                  int
	DWVIDMEMFREE                   int
	DWMAXVISIBLEOVERLAYS           int
	DWCURRVISIBLEOVERLAYS          int
	DWNUMFOURCCCODES               int
	DWALIGNBOUNDARYSRC             int
	DWALIGNSIZESRC                 int
	DWALIGNBOUNDARYDEST            int
	DWALIGNSIZEDEST                int
	DWALIGNSTRIDEALIGN             int
	DWROPS                         int
	DDSCAPS                        int
	DWMINOVERLAYSTRETCH            int
	DWMAXOVERLAYSTRETCH            int
	DWMINLIVEVIDEOSTRETCH          int
	DWMAXLIVEVIDEOSTRETCH          int
	DWMINHWCODECSTRETCH            int
	DWMAXHWCODECSTRETCH            int
	DWRESERVED1                    int
	DWRESERVED2                    int
	DWRESERVED3                    int
	DWSVBCAPS                      int
	DWSVBCKEYCAPS                  int
	DWSVBFXCAPS                    int
	DWSVBROPS                      int
	DWVSBCAPS                      int
	DWVSBCKEYCAPS                  int
	DWVSBFXCAPS                    int
	DWVSBROPS                      int
	DWSSBCAPS                      int
	DWSSBCKEYCAPS                  int
	DWSSBFXCAPS                    int
	DWSSBROPS                      int
	DWRESERVED4                    int
	DWRESERVED5                    int
	DWRESERVED6                    int
}

type DDCAPS_DX5 struct {
	DWSIZE                         int
	DWCAPS                         int
	DWCAPS2                        int
	DWCKEYCAPS                     int
	DWFXCAPS                       int
	DWFXALPHACAPS                  int
	DWPALCAPS                      int
	DWSVCAPS                       int
	DWALPHABLTCONSTBITDEPTHS       int
	DWALPHABLTPIXELBITDEPTHS       int
	DWALPHABLTSURFACEBITDEPTHS     int
	DWALPHAOVERLAYCONSTBITDEPTHS   int
	DWALPHAOVERLAYPIXELBITDEPTHS   int
	DWALPHAOVERLAYSURFACEBITDEPTHS int
	DWZBUFFERBITDEPTHS             int
	DWVIDMEMTOTAL                  int
	DWVIDMEMFREE                   int
	DWMAXVISIBLEOVERLAYS           int
	DWCURRVISIBLEOVERLAYS          int
	DWNUMFOURCCCODES               int
	DWALIGNBOUNDARYSRC             int
	DWALIGNSIZESRC                 int
	DWALIGNBOUNDARYDEST            int
	DWALIGNSIZEDEST                int
	DWALIGNSTRIDEALIGN             int
	DWROPS                         int
	DDSCAPS                        int
	DWMINOVERLAYSTRETCH            int
	DWMAXOVERLAYSTRETCH            int
	DWMINLIVEVIDEOSTRETCH          int
	DWMAXLIVEVIDEOSTRETCH          int
	DWMINHWCODECSTRETCH            int
	DWMAXHWCODECSTRETCH            int
	DWRESERVED1                    int
	DWRESERVED2                    int
	DWRESERVED3                    int
	DWSVBCAPS                      int
	DWSVBCKEYCAPS                  int
	DWSVBFXCAPS                    int
	DWSVBROPS                      int
	DWVSBCAPS                      int
	DWVSBCKEYCAPS                  int
	DWVSBFXCAPS                    int
	DWVSBROPS                      int
	DWSSBCAPS                      int
	DWSSBCKEYCAPS                  int
	DWSSBFXCAPS                    int
	DWSSBROPS                      int
	DWMAXVIDEOPORTS                int
	DWCURRVIDEOPORTS               int
	DWSVBCAPS2                     int
	DWNLVBCAPS                     int
	DWNLVBCAPS2                    int
	DWNLVBCKEYCAPS                 int
	DWNLVBFXCAPS                   int
	DWNLVBROPS                     int
}

type DDCAPS_DX6 struct {
	DWSIZE                         int
	DWCAPS                         int
	DWCAPS2                        int
	DWCKEYCAPS                     int
	DWFXCAPS                       int
	DWFXALPHACAPS                  int
	DWPALCAPS                      int
	DWSVCAPS                       int
	DWALPHABLTCONSTBITDEPTHS       int
	DWALPHABLTPIXELBITDEPTHS       int
	DWALPHABLTSURFACEBITDEPTHS     int
	DWALPHAOVERLAYCONSTBITDEPTHS   int
	DWALPHAOVERLAYPIXELBITDEPTHS   int
	DWALPHAOVERLAYSURFACEBITDEPTHS int
	DWZBUFFERBITDEPTHS             int
	DWVIDMEMTOTAL                  int
	DWVIDMEMFREE                   int
	DWMAXVISIBLEOVERLAYS           int
	DWCURRVISIBLEOVERLAYS          int
	DWNUMFOURCCCODES               int
	DWALIGNBOUNDARYSRC             int
	DWALIGNSIZESRC                 int
	DWALIGNBOUNDARYDEST            int
	DWALIGNSIZEDEST                int
	DWALIGNSTRIDEALIGN             int
	DWROPS                         int
	DDSOLDCAPS                     int
	DWMINOVERLAYSTRETCH            int
	DWMAXOVERLAYSTRETCH            int
	DWMINLIVEVIDEOSTRETCH          int
	DWMAXLIVEVIDEOSTRETCH          int
	DWMINHWCODECSTRETCH            int
	DWMAXHWCODECSTRETCH            int
	DWRESERVED1                    int
	DWRESERVED2                    int
	DWRESERVED3                    int
	DWSVBCAPS                      int
	DWSVBCKEYCAPS                  int
	DWSVBFXCAPS                    int
	DWSVBROPS                      int
	DWVSBCAPS                      int
	DWVSBCKEYCAPS                  int
	DWVSBFXCAPS                    int
	DWVSBROPS                      int
	DWSSBCAPS                      int
	DWSSBCKEYCAPS                  int
	DWSSBFXCAPS                    int
	DWSSBROPS                      int
	DWMAXVIDEOPORTS                int
	DWCURRVIDEOPORTS               int
	DWSVBCAPS2                     int
	DWNLVBCAPS                     int
	DWNLVBCAPS2                    int
	DWNLVBCKEYCAPS                 int
	DWNLVBFXCAPS                   int
	DWNLVBROPS                     int
	DDSCAPS                        int
}

type DDCAPS_DX7 struct {
	DWSIZE                         int
	DWCAPS                         int
	DWCAPS2                        int
	DWCKEYCAPS                     int
	DWFXCAPS                       int
	DWFXALPHACAPS                  int
	DWPALCAPS                      int
	DWSVCAPS                       int
	DWALPHABLTCONSTBITDEPTHS       int
	DWALPHABLTPIXELBITDEPTHS       int
	DWALPHABLTSURFACEBITDEPTHS     int
	DWALPHAOVERLAYCONSTBITDEPTHS   int
	DWALPHAOVERLAYPIXELBITDEPTHS   int
	DWALPHAOVERLAYSURFACEBITDEPTHS int
	DWZBUFFERBITDEPTHS             int
	DWVIDMEMTOTAL                  int
	DWVIDMEMFREE                   int
	DWMAXVISIBLEOVERLAYS           int
	DWCURRVISIBLEOVERLAYS          int
	DWNUMFOURCCCODES               int
	DWALIGNBOUNDARYSRC             int
	DWALIGNSIZESRC                 int
	DWALIGNBOUNDARYDEST            int
	DWALIGNSIZEDEST                int
	DWALIGNSTRIDEALIGN             int
	DWROPS                         int
	DDSOLDCAPS                     int
	DWMINOVERLAYSTRETCH            int
	DWMAXOVERLAYSTRETCH            int
	DWMINLIVEVIDEOSTRETCH          int
	DWMAXLIVEVIDEOSTRETCH          int
	DWMINHWCODECSTRETCH            int
	DWMAXHWCODECSTRETCH            int
	DWRESERVED1                    int
	DWRESERVED2                    int
	DWRESERVED3                    int
	DWSVBCAPS                      int
	DWSVBCKEYCAPS                  int
	DWSVBFXCAPS                    int
	DWSVBROPS                      int
	DWVSBCAPS                      int
	DWVSBCKEYCAPS                  int
	DWVSBFXCAPS                    int
	DWVSBROPS                      int
	DWSSBCAPS                      int
	DWSSBCKEYCAPS                  int
	DWSSBFXCAPS                    int
	DWSSBROPS                      int
	DWMAXVIDEOPORTS                int
	DWCURRVIDEOPORTS               int
	DWSVBCAPS2                     int
	DWNLVBCAPS                     int
	DWNLVBCAPS2                    int
	DWNLVBCKEYCAPS                 int
	DWNLVBFXCAPS                   int
	DWNLVBROPS                     int
	DDSCAPS                        int
}

type DDPIXELFORMAT struct {
	DWSIZE     int
	DWFLAGS    int
	DWFOURCC   int
	Anonymous1 int
	Anonymous2 int
	Anonymous3 int
	Anonymous4 int
	Anonymous5 int
}

type DDOVERLAYFX struct {
	DWSIZE                   int
	DWALPHAEDGEBLENDBITDEPTH int
	DWALPHAEDGEBLEND         int
	DWRESERVED               int
	DWALPHADESTCONSTBITDEPTH int
	Anonymous1               int
	DWALPHASRCCONSTBITDEPTH  int
	Anonymous2               int
	DCKDESTCOLORKEY          int
	DCKSRCCOLORKEY           int
	DWDDFX                   int
	DWFLAGS                  int
}

type DDBLTBATCH struct {
	LPRDEST   int
	LPDDSSRC  int
	LPRSRC    int
	DWFLAGS   int
	LPDDBLTFX int
}

type DDGAMMARAMP struct {
	RED   int
	GREEN int
	BLUE  int
}

type DDDEVICEIDENTIFIER struct {
	SZDRIVER             int
	SZDESCRIPTION        int
	LIDRIVERVERSION      int
	DWVENDORID           int
	DWDEVICEID           int
	DWSUBSYSID           int
	DWREVISION           int
	GUIDDEVICEIDENTIFIER int
}

type DDDEVICEIDENTIFIER2 struct {
	SZDRIVER             int
	SZDESCRIPTION        int
	LIDRIVERVERSION      int
	DWVENDORID           int
	DWDEVICEID           int
	DWSUBSYSID           int
	DWREVISION           int
	GUIDDEVICEIDENTIFIER int
	DWWHQLLEVEL          int
}

type DDSURFACEDESC struct {
	DWSIZE            int
	DWFLAGS           int
	DWHEIGHT          int
	DWWIDTH           int
	Anonymous1        int
	DWBACKBUFFERCOUNT int
	Anonymous2        int
	DWALPHABITDEPTH   int
	DWRESERVED        int
	LPSURFACE         int
	DDCKCKDESTOVERLAY int
	DDCKCKDESTBLT     int
	DDCKCKSRCOVERLAY  int
	DDCKCKSRCBLT      int
	DDPFPIXELFORMAT   int
	DDSCAPS           int
}

type DDSURFACEDESC2 struct {
	DWSIZE           int
	DWFLAGS          int
	DWHEIGHT         int
	DWWIDTH          int
	Anonymous1       int
	Anonymous2       int
	Anonymous3       int
	DWALPHABITDEPTH  int
	DWRESERVED       int
	LPSURFACE        int
	Anonymous4       int
	DDCKCKDESTBLT    int
	DDCKCKSRCOVERLAY int
	DDCKCKSRCBLT     int
	Anonymous5       int
	DDSCAPS          int
	DWTEXTURESTAGE   int
}

type DDOPTSURFACEDESC struct {
	DWSIZE             int
	DWFLAGS            int
	DDSCAPS            int
	DDOSCAPS           int
	GUID               int
	DWCOMPRESSIONRATIO int
}

type DDCOLORCONTROL struct {
	DWSIZE       int
	DWFLAGS      int
	LBRIGHTNESS  int
	LCONTRAST    int
	LHUE         int
	LSATURATION  int
	LSHARPNESS   int
	LGAMMA       int
	LCOLORENABLE int
	DWRESERVED1  int
}
