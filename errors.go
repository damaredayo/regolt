package regolt

var (
	ErrLabelMe               = "LabelMe"               // uncategorised error
	ErrInternalError         = "InternalError"         // the server ran into an issue
	ErrInvalidSession        = "InvalidSession"        // authentication details are incorrect
	ErrOnboardingNotFinished = "OnboardingNotFinished" // user has not chosen a username
	ErrAlreadyAuthenticated  = "AlreadyAuthenticated"  // this connection is already authenticated
)
