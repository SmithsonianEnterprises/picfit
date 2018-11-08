package engine

type Operation string

func (o Operation) String() string {
	return string(o)
}

const (
	Resize              = Operation("resize")
	Thumbnail           = Operation("thumbnail")
	ThumbnailWithAnchor = Operation("focusedthumbnail")
	Rotate              = Operation("rotate")
	Flip                = Operation("flip")
	Fit                 = Operation("fit")
	Noop                = Operation("noop")
)

var Operations = map[string]Operation{
	Resize.String():              Resize,
	Thumbnail.String():           Thumbnail,
	ThumbnailWithAnchor.String(): ThumbnailWithAnchor,
	Flip.String():                Flip,
	Rotate.String():              Rotate,
	Fit.String():                 Fit,
	Noop.String():                Noop,
}
