package cautils

import (
	"fmt"
	"strings"
)

func GetControlLink(controlID string) string {
	// For CIS Controls, cis-1.1.3 will be transformed to cis-1-1-3 in documentation link.
	docLinkID := strings.ReplaceAll(controlID, ".", "-")
	return fmt.Sprintf("https://github.com/5GSEC/5G-Blueprint-Controls/blob/main/Controls/%s.md", strings.ToLower(docLinkID))
}
