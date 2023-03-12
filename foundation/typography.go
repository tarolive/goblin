package foundation

type HeadlineTypography struct {
	XS  string
	SM  string
	MD  string
	LG  string
	XL  string
	XXL string
}

func NewHeadlineTypography(xs string, sm string, md string, lg string, xl string, xxl string) HeadlineTypography {

	var (
		headlineTypography = HeadlineTypography{
			XS:  xs,
			SM:  sm,
			MD:  md,
			LG:  lg,
			XL:  xl,
			XXL: xxl,
		}
	)

	return headlineTypography
}
