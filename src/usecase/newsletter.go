package usecase

import (
	"context"

	domainNewsletter "github.com/IamYGT/ygtwa/domains/newsletter"
	"github.com/IamYGT/ygtwa/infrastructure/whatsapp"
	pkgError "github.com/IamYGT/ygtwa/pkg/error"
	"github.com/IamYGT/ygtwa/pkg/utils"
	"github.com/IamYGT/ygtwa/validations"
)

type serviceNewsletter struct{}

func NewNewsletterService() domainNewsletter.INewsletterUsecase {
	return &serviceNewsletter{}
}

func (service serviceNewsletter) Unfollow(ctx context.Context, request domainNewsletter.UnfollowRequest) (err error) {
	if err = validations.ValidateUnfollowNewsletter(ctx, request); err != nil {
		return err
	}

	client := whatsapp.ClientFromContext(ctx)
	if client == nil {
		return pkgError.ErrWaCLI
	}

	JID, err := utils.ValidateJidWithLogin(client, request.NewsletterID)
	if err != nil {
		return err
	}

	return client.UnfollowNewsletter(ctx, JID)
}
