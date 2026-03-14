package validations

import (
	"context"
	domainNewsletter "github.com/IamYGT/ygt-labs-ai-whatsapp/domains/newsletter"
	pkgError "github.com/IamYGT/ygt-labs-ai-whatsapp/pkg/error"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateUnfollowNewsletter(ctx context.Context, request domainNewsletter.UnfollowRequest) error {
	err := validation.ValidateStructWithContext(ctx, &request,
		validation.Field(&request.NewsletterID, validation.Required),
	)

	if err != nil {
		return pkgError.ValidationError(err.Error())
	}

	return nil
}
