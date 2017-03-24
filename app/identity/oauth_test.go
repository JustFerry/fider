package identity_test

import (
	"testing"

	"github.com/WeCanHearYou/wechy/app/identity"
	. "github.com/onsi/gomega"
)

func TestGetAuthURL_Facebook(t *testing.T) {
	RegisterTestingT(t)

	oauth := &identity.HTTPOAuthService{}
	authURL := oauth.GetAuthURL(identity.OAuthFacebookProvider, "")

	Expect(authURL).To(Equal("https://www.facebook.com/dialog/oauth?client_id=&redirect_uri=http%3A%2F%2Flogin.test.canhearyou.com%3A3000%2Foauth%2Ffacebook%2Fcallback&response_type=code&scope=public_profile+email&state="))
}

func TestGetAuthURL_Google(t *testing.T) {
	RegisterTestingT(t)

	oauth := &identity.HTTPOAuthService{}
	authURL := oauth.GetAuthURL(identity.OAuthGoogleProvider, "")

	Expect(authURL).To(Equal("https://accounts.google.com/o/oauth2/auth?client_id=&redirect_uri=http%3A%2F%2Flogin.test.canhearyou.com%3A3000%2Foauth%2Fgoogle%2Fcallback&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email&state="))
}
