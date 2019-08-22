package charts

const (
	POINT   = "point"
	NEAREST = "nearest"
	INDEX   = "index"
	DATASET = "dataset"
	X       = "x"
	Y       = "y"
)

var MODE = []string{POINT, NEAREST, INDEX, DATASET, X, Y}

const (
	TOP    = "top"
	LEFT   = "left"
	BOTTOM = "bottom"
	RIGHT  = "right"
)

var POSITION = []string{TOP, LEFT, BOTTOM, RIGHT}

const (
	NEUE      = "Helvetica Neue"
	HELVETICA = "Helvetica"
	ARIAL     = "Arial"
	SANS      = "sans-serif"
)

var FontFamily = []string{NEUE, HELVETICA, ARIAL, SANS}

const (
	FILLORIGIN = "origin"
	FILLFALSE  = "false"
	FILLSTART  = "start"
	FILLEND    = "end"
)

var Fill = []string{FILLORIGIN, FILLFALSE, FILLSTART, FILLEND}
