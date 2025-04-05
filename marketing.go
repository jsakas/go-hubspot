package hubspot

import "fmt"

const (
	marketingBasePath = "marketing"
)

type Marketing struct {
	Email         MarketingEmailService
	Transactional TransactionalService
	Form          FormService
}

func newMarketing(c *Client) *Marketing {
	return &Marketing{
		Email: NewMarketingEmail(c),
		Transactional: &TransactionalServiceOp{
			client:            c,
			transactionalPath: fmt.Sprintf("%s/%s/%s", marketingBasePath, c.apiVersion, transactionalBasePath),
		},
		Form: NewFormService(c),
	}
}
