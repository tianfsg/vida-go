// vida-go/static/lang/lang.go

package lang

type Language struct {
	Title                    string             `json:"title"`
	Description              string             `json:"description"`
	Keywords                 string             `json:"keywords"`
	SiteName                 string             `json:"siteName"`
	Header                   string             `json:"header"`
	Name                     string             `json:"name"`
	Role                     string             `json:"role"`
	AboutDescription         string             `json:"aboutDescription"`
	PreviewCV                string             `json:"previewCV"`
	TalkButton               string             `json:"talkButton"`
	SiteKey                  string             `json:"sitekey"`
	Lang                     string             `json:"lang"`
	Flag                     string             `json:"flag"`
	MenuHome                 string             `json:"menuHome"`
	MenuArt                  string             `json:"menuArt"`
	MenuAbout                string             `json:"menuAbout"`
	MenuSkills               string             `json:"menuSkills"`
	MenuPortfolio            string             `json:"menuPortfolio"`
	MenuContact              string             `json:"menuContact"`
	JourneyTitle             string             `json:"journeyTitle"`
	EducationTitle           string             `json:"educationTitle"`
	EducationItemTitle       string             `json:"educationItemTitle"`
	EducationItemDesc        string             `json:"educationItemDescription"`
	ExperienceTitle          string             `json:"experienceTitle"`
	ExperienceItem1Title     string             `json:"experienceItem1Title"`
	ExperienceItem1Desc      string             `json:"experienceItem1Description"`
	ExperienceItem2Title     string             `json:"experienceItem2Title"`
	ExperienceItem2Desc      string             `json:"experienceItem2Description"`
	ExperienceItem3Title     string             `json:"experienceItem3Title"`
	ExperienceItem3Desc      string             `json:"experienceItem3Description"`
	SkillsTitle              string             `json:"skillsTitle"`
	TechnicalSkillsTitle     string             `json:"technicalSkillsTitle"`
	ComplementarySkillsTitle string             `json:"complementarySkillsTitle"`
	TechnicalSkill1          string             `json:"technicalSkill1"`
	TechnicalSkill2          string             `json:"technicalSkill2"`
	TechnicalSkill3          string             `json:"technicalSkill3"`
	TechnicalSkill4          string             `json:"technicalSkill4"`
	ComplementarySkill1      string             `json:"complementarySkill1"`
	ComplementarySkill2      string             `json:"complementarySkill2"`
	ComplementarySkill3      string             `json:"complementarySkill3"`
	ComplementarySkill4      string             `json:"complementarySkill4"`
	ContactTitle             string             `json:"contactTitle"`
	ContactSubtitle          string             `json:"contactSubtitle"`
	EmailLabel               string             `json:"emailLabel"`
	PhoneLabel               string             `json:"phoneLabel"`
	FooterPrivacy            string             `json:"footerPrivacy"`
	FooterCookies            string             `json:"footerCookies"`
	FooterTerms              string             `json:"footerTerms"`
	FooterLegal              string             `json:"footerLegal"`
	FooterSignature          string             `json:"footerSignature"`
	CookiesPolicy            CookiesPolicy      `json:"cookies_policy"`
	LegalNotice              LegalNotice        `json:"legal_notice"`
	TermsAndConditions       TermsAndConditions `json:"terms_and_conditions"`
	PrivacyPolicy            PrivacyPolicy      `json:"privacy_policy"`
}

type CookiesPolicy struct {
	Title                  string `json:"title"`
	WhatAreCookies         string `json:"what_are_cookies"`
	WhatAreCookiesDetails  string `json:"what_are_cookies_details"`
	HowWeUseCookies        string `json:"how_we_use_cookies"`
	HowWeUseCookiesDetails string `json:"how_we_use_cookies_details"`
	TypesOfCookies         string `json:"types_of_cookies"`
	TypesOfCookiesDetails  string `json:"types_of_cookies_details"`
	YourChoices            string `json:"your_choices"`
	YourChoicesDetails     string `json:"your_choices_details"`
	ContactUs              string `json:"contact_us"`
	ContactUsDetails       string `json:"contact_us_details"`
}

type LegalNotice struct {
	Title               string `json:"title"`
	Content             string `json:"content"`
	WebsiteOwner        string `json:"website_owner"`
	WebsiteOwnerDetails string `json:"website_owner_details"`
	Disclaimer          string `json:"disclaimer"`
	DisclaimerDetails   string `json:"disclaimer_details"`
	GoverningLaw        string `json:"governing_law"`
	GoverningLawDetails string `json:"governing_law_details"`
}

type TermsAndConditions struct {
	Title                        string `json:"title"`
	Content                      string `json:"content"`
	UseOfSite                    string `json:"use_of_site"`
	UseOfSiteDetails             string `json:"use_of_site_details"`
	IntellectualProperty         string `json:"intellectual_property"`
	IntellectualPropertyDetails  string `json:"intellectual_property_details"`
	LimitationOfLiability        string `json:"limitation_of_liability"`
	LimitationOfLiabilityDetails string `json:"limitation_of_liability_details"`
	ChangesToTerms               string `json:"changes_to_terms"`
	ChangesToTermsDetails        string `json:"changes_to_terms_details"`
}

type PrivacyPolicy struct {
	Title                       string `json:"title"`
	Content                     string `json:"content"`
	InformationWeCollect        string `json:"information_we_collect"`
	InformationWeCollectDetails string `json:"information_we_collect_details"`
	HowWeUseInformation         string `json:"how_we_use_information"`
	HowWeUseInformationDetails  string `json:"how_we_use_information_details"`
	Cookies                     string `json:"cookies"`
	CookiesDetails              string `json:"cookies_details"`
	YourRights                  string `json:"your_rights"`
	YourRightsDetails           string `json:"your_rights_details"`
}

var supportedLanguages = map[string]func(*Language) Language{
	"EN": (*Language).returnEnglish,
	"ES": (*Language).returnSpanish,
	"DE": (*Language).returnGerman,
}

func GetLanguage(langCode string) Language {
	if getLangFunc, exists := supportedLanguages[langCode]; exists {
		var l Language
		return getLangFunc(&l)
	}
	// Default to English if language not found
	var l Language
	return l.returnEnglish()
}
