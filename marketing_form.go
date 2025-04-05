package hubspot

import (
	"fmt"
	"net/url"
)

const (
	formBasePath = "forms"
)

// Form represents a HubSpot form
type Form struct {
	ID                  string                 `json:"id,omitempty"`
	Name                string                 `json:"name"`
	CreatedAt           string                 `json:"createdAt,omitempty"`
	UpdatedAt           string                 `json:"updatedAt,omitempty"`
	Archived            bool                   `json:"archived,omitempty"`
	FieldGroups         []FormFieldGroup       `json:"fieldGroups"`
	Configuration       FormConfiguration      `json:"configuration"`
	DisplayOptions      FormDisplayOptions     `json:"displayOptions"`
	LegalConsentOptions *LegalConsentOptions   `json:"legalConsentOptions,omitempty"`
	FormType            string                 `json:"formType,omitempty"`
	Action              string                 `json:"action,omitempty"`
	Method              string                 `json:"method,omitempty"`
	RedirectURL         string                 `json:"redirectUrl,omitempty"`
	Properties          map[string]interface{} `json:"properties,omitempty"`
}

// FormFieldGroup represents a group of fields in a HubSpot form
type FormFieldGroup struct {
	GroupType    string      `json:"groupType"`
	RichTextType string      `json:"richTextType"`
	Fields       []FormField `json:"fields"`
}

// FormField represents a field in a HubSpot form
type FormField struct {
	ObjectTypeId    string                 `json:"objectTypeId"`
	Name            string                 `json:"name"`
	Label           string                 `json:"label"`
	Required        bool                   `json:"required"`
	Hidden          bool                   `json:"hidden"`
	FieldType       string                 `json:"fieldType"`
	Validation      *FormFieldValidation   `json:"validation,omitempty"`
	Placeholder     string                 `json:"placeholder,omitempty"`
	HelpText        string                 `json:"helpText,omitempty"`
	Options         []FormFieldOption      `json:"options,omitempty"`
	Properties      map[string]interface{} `json:"properties,omitempty"`
	DependentFields []string               `json:"dependentFields,omitempty"`
}

// FormFieldValidation represents validation rules for a form field
type FormFieldValidation struct {
	BlockedEmailDomains []string `json:"blockedEmailDomains,omitempty"`
	UseDefaultBlockList bool     `json:"useDefaultBlockList,omitempty"`
}

// FormFieldOption represents an option for a form field
type FormFieldOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// FormConfiguration represents the configuration of a HubSpot form
type FormConfiguration struct {
	Language                    string               `json:"language"`
	Cloneable                   bool                 `json:"cloneable"`
	PostSubmitAction            FormPostSubmitAction `json:"postSubmitAction"`
	Editable                    bool                 `json:"editable"`
	Archivable                  bool                 `json:"archivable"`
	RecaptchaEnabled            bool                 `json:"recaptchaEnabled"`
	NotifyContactOwner          bool                 `json:"notifyContactOwner"`
	NotifyRecipients            []string             `json:"notifyRecipients"`
	CreateNewContactForNewEmail bool                 `json:"createNewContactForNewEmail"`
	PrePopulateKnownValues      bool                 `json:"prePopulateKnownValues"`
	AllowLinkToResetKnownValues bool                 `json:"allowLinkToResetKnownValues"`
	LifecycleStages             []FormLifecycleStage `json:"lifecycleStages"`
}

// FormPostSubmitAction represents the action to take after form submission
type FormPostSubmitAction struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// FormLifecycleStage represents a lifecycle stage for a form
type FormLifecycleStage struct {
	ObjectTypeId string `json:"objectTypeId"`
	Value        string `json:"value"`
}

// FormDisplayOptions represents display options for a HubSpot form
type FormDisplayOptions struct {
	RenderRawHtml    bool             `json:"renderRawHtml"`
	Theme            string           `json:"theme"`
	SubmitButtonText string           `json:"submitButtonText"`
	Style            FormDisplayStyle `json:"style"`
	CSSClass         string           `json:"cssClass"`
}

// FormDisplayStyle represents the style of a HubSpot form
type FormDisplayStyle struct {
	FontFamily            string `json:"fontFamily"`
	BackgroundWidth       string `json:"backgroundWidth"`
	LabelTextColor        string `json:"labelTextColor"`
	LabelTextSize         string `json:"labelTextSize"`
	HelpTextColor         string `json:"helpTextColor"`
	HelpTextSize          string `json:"helpTextSize"`
	LegalConsentTextColor string `json:"legalConsentTextColor"`
	LegalConsentTextSize  string `json:"legalConsentTextSize"`
	SubmitColor           string `json:"submitColor"`
	SubmitAlignment       string `json:"submitAlignment"`
	SubmitFontColor       string `json:"submitFontColor"`
	SubmitSize            string `json:"submitSize"`
}

// LegalConsentOptions represents the legal consent options for a form submission
type LegalConsentOptions struct {
	Type                          string                  `json:"type"`
	CommunicationConsentText      string                  `json:"communicationConsentText"`
	CommunicationsCheckboxes      []CommunicationCheckbox `json:"communicationsCheckboxes"`
	ConsentToProcessText          string                  `json:"consentToProcessText"`
	ConsentToProcessCheckboxLabel string                  `json:"consentToProcessCheckboxLabel"`
	PrivacyText                   string                  `json:"privacyText"`
	Consent                       string                  `json:"consent,omitempty"`
	ConsentToProcess              bool                    `json:"consentToProcess,omitempty"`
	Text                          string                  `json:"text,omitempty"`
	Communications                []string                `json:"communications,omitempty"`
}

// CommunicationCheckbox represents a communication consent checkbox
type CommunicationCheckbox struct {
	Required           bool   `json:"required"`
	SubscriptionTypeId int    `json:"subscriptionTypeId"`
	Label              string `json:"label"`
}

// FormResponse represents the response from the HubSpot Forms API
type FormResponse struct {
	Results []Form `json:"results"`
	Paging  struct {
		Next struct {
			After string `json:"after"`
		} `json:"next"`
	} `json:"paging"`
}

// FormSubmission represents a form submission to HubSpot
type FormSubmission struct {
	Fields              []FormSubmissionField `json:"fields"`
	Context             FormSubmissionContext `json:"context"`
	LegalConsentOptions *LegalConsentOptions  `json:"legalConsentOptions,omitempty"`
}

// FormSubmissionField represents a field in a form submission
type FormSubmissionField struct {
	ObjectTypeId string      `json:"objectTypeId"`
	Name         string      `json:"name"`
	Value        interface{} `json:"value"`
}

// FormSubmissionContext represents the context of a form submission
type FormSubmissionContext struct {
	PageURI   string `json:"pageUri"`
	PageName  string `json:"pageName"`
	IPAddress string `json:"ipAddress,omitempty"`
	UserAgent string `json:"userAgent,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Hutk      string `json:"hutk,omitempty"`
}

// FormService is an interface for the HubSpot Forms API
type FormService interface {
	Create(form *Form) (*Form, error)
	Get(formID string) (*Form, error)
	List(options *BulkRequestQueryOption) (*FormResponse, error)
	Update(formID string, form *Form) (*Form, error)
	Delete(formID string) error
	Submit(portalID string, formID string, submission *FormSubmission) error
}

// FormServiceOp handles communication with the HubSpot Forms API
type FormServiceOp struct {
	client   *Client
	formPath string
}

var _ FormService = (*FormServiceOp)(nil)

// NewFormService creates a new FormService
func NewFormService(client *Client) FormService {
	return &FormServiceOp{
		client:   client,
		formPath: fmt.Sprintf("%s/%s/%s", marketingBasePath, client.apiVersion, formBasePath),
	}
}

// Create creates a new form
func (s *FormServiceOp) Create(form *Form) (*Form, error) {
	createdForm := new(Form)
	if err := s.client.Post(s.formPath, form, createdForm); err != nil {
		return nil, err
	}
	return createdForm, nil
}

// Get retrieves a form by ID
func (s *FormServiceOp) Get(formID string) (*Form, error) {
	path := fmt.Sprintf("%s/%s", s.formPath, formID)

	form := new(Form)
	if err := s.client.Get(path, form, nil); err != nil {
		return nil, err
	}
	return form, nil
}

// List retrieves a list of forms
func (s *FormServiceOp) List(options *BulkRequestQueryOption) (*FormResponse, error) {
	response := new(FormResponse)
	if err := s.client.Get(s.formPath, response, options); err != nil {
		return nil, err
	}
	return response, nil
}

// Update updates an existing form
func (s *FormServiceOp) Update(formID string, form *Form) (*Form, error) {
	path := fmt.Sprintf("%s/%s", s.formPath, formID)
	updatedForm := new(Form)
	if err := s.client.Patch(path, form, updatedForm); err != nil {
		return nil, err
	}
	return updatedForm, nil
}

// Delete deletes a form
func (s *FormServiceOp) Delete(formID string) error {
	path := fmt.Sprintf("%s/%s", s.formPath, formID)
	return s.client.Delete(path, nil)
}

// Submit submits a form to HubSpot
func (s *FormServiceOp) Submit(portalID string, formID string, submission *FormSubmission) error {
	// HubSpot legacy-v3 and v2 forms API use a different base URL
	// Latest forms API does not support submitting forms yet
	formsClient := *s.client // Make a shallow copy of the client

	formsBaseURL, _ := url.Parse("https://api.hsforms.com")
	formsClient.baseURL = formsBaseURL

	path := fmt.Sprintf("/submissions/v3/integration/secure/submit/%s/%s", portalID, formID)

	// Use the copied client to make the request
	return formsClient.Post(path, submission, nil)
}
