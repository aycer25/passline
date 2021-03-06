package action

import (
	"fmt"

	ucli "github.com/urfave/cli/v2"
)

const (
	// ExitOK means no error (status code 0)
	ExitOK = iota
	// ExitUnknown is used if we can't determine the exact exit cause
	ExitUnknown
	// ExitUsage is used if there was some kind of invocation error
	ExitUsage
	// Exit Duplicated is used if the item already exists
	ExitDuplicated
	// ExitAborted is used if the user willingly aborted an action
	ExitAborted
	// ExitUnsupported is used if an operation is not supported by passline
	ExitUnsupported
	// ExitAlreadyInitialized is used if someone is trying to initialize
	// an already initialized store
	ExitAlreadyInitialized
	// ExitNotInitialized is used if someone is trying to use an unitialized
	// store
	ExitNotInitialized
	// Exit Password is used if the global password is wrong
	ExitPassword
	// ExitNoName is used when no name was provided for a named entry
	ExitNoName
	// ExitNotFound is used if a requested secret is not found
	ExitNotFound
	// ExitDecrypt is used when reading/decrypting a secret failed
	ExitDecrypt
	// ExitEncrypt is used when writing/encrypting of a secret fails
	ExitEncrypt
	// ExitList is used when listing the store content fails
	ExitList
	// ExitAudit is used when audit report possible issues
	ExitAudit
	// ExitFsck is used when the integrity check fails
	ExitFsck
	// ExitConfig is used when config errors occur
	ExitConfig
	// ExitRecipients is used when a recipient operation fails
	ExitRecipients
	// ExitIO is used for misc. I/O errors
	ExitIO
)

// ExitError returns a user friendly CLI error
func ExitError(exitCode int, err error, format string, args ...interface{}) error {
	return ucli.NewExitError(fmt.Sprintf(format, args...), exitCode)
}
