package ilbm

type Masking uint8

const (
	MaskNone                = Masking(iota) // No mask
	MaskHasMask                             // Has mask
	MaskHasTransparentColor                 // Has transparent color
	MaskHasLasso                            // Has lasso
)

type Compression uint8

const (
	CompressionNone     = Compression(iota) // No compression
	CompressionByteRun1                     // Run Length encoding
)

type Header struct {
	Width, Height         uint16      // Raster width & height in pixels
	X, Y                  int16       // Pixel position within image
	Planes                uint8       // Number of bit planes
	Masking               Masking     // Masking value
	Compression           Compression // Compression value
	padding1              uint8       // padding.
	TransparentColor      uint16      // Index of transparent color (sort of)
	XAspect, YAspect      uint8       // Pixel aspect ratio
	PageWidth, PageHeight int16       // source page size in pixels
}
