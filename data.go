package tripismapiutilities

import "net/url"

// CleanWebsiteQueryString removes unused / tracking query string parameters from website URLs
//	for example, it deletes all "utm_source", "utm_campaign" and "utm_medium" Google Analytcs parameters
func CleanWebsiteQueryString(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return uri
	}

	q := u.Query()
	q.Del("scid")
	q.Del("cid")
	q.Del("y_source")
	q.Del("utm_source")
	q.Del("utm_campaign")
	q.Del("utm_medium")
	q.Del("utm_term")
	q.Del("WT.mc_id")
	q.Del("SEO_id")
	q.Del("s_cid")
	q.Del("cm_mmc")

	u.RawQuery = q.Encode()
	return u.String()
}
