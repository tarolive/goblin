package foundation

type Spacing struct {
	XS    string
	SM    string
	MD    string
	LG    string
	XL    string
	XXL   string
	XXXL  string
	XXXXL string
}

func NewSpacing(xs string, sm string, md string, lg string, xl string, xxl string, xxxl string, xxxxl string) Spacing {

	var (
		spacing = Spacing{
			XS:    xs,
			SM:    sm,
			MD:    md,
			LG:    lg,
			XL:    xl,
			XXL:   xxl,
			XXXL:  xxxl,
			XXXXL: xxxxl,
		}
	)

	return spacing
}
