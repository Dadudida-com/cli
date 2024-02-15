// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package config

import (
	"sync"
)

// Ensure, that ConfigMock does implement Config.
// If this is not the case, regenerate this file with moq.
var _ Config = &ConfigMock{}

// ConfigMock is a mock implementation of Config.
//
//	func TestSomethingThatUsesConfig(t *testing.T) {
//
//		// make and configure a mocked Config
//		mockedConfig := &ConfigMock{
//			AliasesFunc: func() *AliasConfig {
//				panic("mock out the Aliases method")
//			},
//			AuthenticationFunc: func() *AuthConfig {
//				panic("mock out the Authentication method")
//			},
//			BrowserFunc: func(s string) string {
//				panic("mock out the Browser method")
//			},
//			EditorFunc: func(s string) string {
//				panic("mock out the Editor method")
//			},
//			GetOrDefaultFunc: func(s1 string, s2 string) (string, error) {
//				panic("mock out the GetOrDefault method")
//			},
//			GitProtocolFunc: func(s string) string {
//				panic("mock out the GitProtocol method")
//			},
//			HTTPUnixSocketFunc: func(s string) string {
//				panic("mock out the HTTPUnixSocket method")
//			},
//			MigrateFunc: func(migration Migration) error {
//				panic("mock out the Migrate method")
//			},
//			PagerFunc: func(s string) string {
//				panic("mock out the Pager method")
//			},
//			PromptFunc: func(s string) string {
//				panic("mock out the Prompt method")
//			},
//			SetFunc: func(s1 string, s2 string, s3 string)  {
//				panic("mock out the Set method")
//			},
//			VersionFunc: func() string {
//				panic("mock out the Version method")
//			},
//			WriteFunc: func() error {
//				panic("mock out the Write method")
//			},
//		}
//
//		// use mockedConfig in code that requires Config
//		// and then make assertions.
//
//	}
type ConfigMock struct {
	// AliasesFunc mocks the Aliases method.
	AliasesFunc func() *AliasConfig

	// AuthenticationFunc mocks the Authentication method.
	AuthenticationFunc func() *AuthConfig

	// BrowserFunc mocks the Browser method.
	BrowserFunc func(s string) string

	// EditorFunc mocks the Editor method.
	EditorFunc func(s string) string

	// GetOrDefaultFunc mocks the GetOrDefault method.
	GetOrDefaultFunc func(s1 string, s2 string) (string, error)

	// GitProtocolFunc mocks the GitProtocol method.
	GitProtocolFunc func(s string) string

	// HTTPUnixSocketFunc mocks the HTTPUnixSocket method.
	HTTPUnixSocketFunc func(s string) string

	// MigrateFunc mocks the Migrate method.
	MigrateFunc func(migration Migration) error

	// PagerFunc mocks the Pager method.
	PagerFunc func(s string) string

	// PromptFunc mocks the Prompt method.
	PromptFunc func(s string) string

	// SetFunc mocks the Set method.
	SetFunc func(s1 string, s2 string, s3 string)

	// VersionFunc mocks the Version method.
	VersionFunc func() string

	// WriteFunc mocks the Write method.
	WriteFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Aliases holds details about calls to the Aliases method.
		Aliases []struct {
		}
		// Authentication holds details about calls to the Authentication method.
		Authentication []struct {
		}
		// Browser holds details about calls to the Browser method.
		Browser []struct {
			// S is the s argument value.
			S string
		}
		// Editor holds details about calls to the Editor method.
		Editor []struct {
			// S is the s argument value.
			S string
		}
		// GetOrDefault holds details about calls to the GetOrDefault method.
		GetOrDefault []struct {
			// S1 is the s1 argument value.
			S1 string
			// S2 is the s2 argument value.
			S2 string
		}
		// GitProtocol holds details about calls to the GitProtocol method.
		GitProtocol []struct {
			// S is the s argument value.
			S string
		}
		// HTTPUnixSocket holds details about calls to the HTTPUnixSocket method.
		HTTPUnixSocket []struct {
			// S is the s argument value.
			S string
		}
		// Migrate holds details about calls to the Migrate method.
		Migrate []struct {
			// Migration is the migration argument value.
			Migration Migration
		}
		// Pager holds details about calls to the Pager method.
		Pager []struct {
			// S is the s argument value.
			S string
		}
		// Prompt holds details about calls to the Prompt method.
		Prompt []struct {
			// S is the s argument value.
			S string
		}
		// Set holds details about calls to the Set method.
		Set []struct {
			// S1 is the s1 argument value.
			S1 string
			// S2 is the s2 argument value.
			S2 string
			// S3 is the s3 argument value.
			S3 string
		}
		// Version holds details about calls to the Version method.
		Version []struct {
		}
		// Write holds details about calls to the Write method.
		Write []struct {
		}
	}
	lockAliases        sync.RWMutex
	lockAuthentication sync.RWMutex
	lockBrowser        sync.RWMutex
	lockEditor         sync.RWMutex
	lockGetOrDefault   sync.RWMutex
	lockGitProtocol    sync.RWMutex
	lockHTTPUnixSocket sync.RWMutex
	lockMigrate        sync.RWMutex
	lockPager          sync.RWMutex
	lockPrompt         sync.RWMutex
	lockSet            sync.RWMutex
	lockVersion        sync.RWMutex
	lockWrite          sync.RWMutex
}

// Aliases calls AliasesFunc.
func (mock *ConfigMock) Aliases() *AliasConfig {
	if mock.AliasesFunc == nil {
		panic("ConfigMock.AliasesFunc: method is nil but Config.Aliases was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAliases.Lock()
	mock.calls.Aliases = append(mock.calls.Aliases, callInfo)
	mock.lockAliases.Unlock()
	return mock.AliasesFunc()
}

// AliasesCalls gets all the calls that were made to Aliases.
// Check the length with:
//
//	len(mockedConfig.AliasesCalls())
func (mock *ConfigMock) AliasesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAliases.RLock()
	calls = mock.calls.Aliases
	mock.lockAliases.RUnlock()
	return calls
}

// Authentication calls AuthenticationFunc.
func (mock *ConfigMock) Authentication() *AuthConfig {
	if mock.AuthenticationFunc == nil {
		panic("ConfigMock.AuthenticationFunc: method is nil but Config.Authentication was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAuthentication.Lock()
	mock.calls.Authentication = append(mock.calls.Authentication, callInfo)
	mock.lockAuthentication.Unlock()
	return mock.AuthenticationFunc()
}

// AuthenticationCalls gets all the calls that were made to Authentication.
// Check the length with:
//
//	len(mockedConfig.AuthenticationCalls())
func (mock *ConfigMock) AuthenticationCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAuthentication.RLock()
	calls = mock.calls.Authentication
	mock.lockAuthentication.RUnlock()
	return calls
}

// Browser calls BrowserFunc.
func (mock *ConfigMock) Browser(s string) string {
	if mock.BrowserFunc == nil {
		panic("ConfigMock.BrowserFunc: method is nil but Config.Browser was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockBrowser.Lock()
	mock.calls.Browser = append(mock.calls.Browser, callInfo)
	mock.lockBrowser.Unlock()
	return mock.BrowserFunc(s)
}

// BrowserCalls gets all the calls that were made to Browser.
// Check the length with:
//
//	len(mockedConfig.BrowserCalls())
func (mock *ConfigMock) BrowserCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockBrowser.RLock()
	calls = mock.calls.Browser
	mock.lockBrowser.RUnlock()
	return calls
}

// Editor calls EditorFunc.
func (mock *ConfigMock) Editor(s string) string {
	if mock.EditorFunc == nil {
		panic("ConfigMock.EditorFunc: method is nil but Config.Editor was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockEditor.Lock()
	mock.calls.Editor = append(mock.calls.Editor, callInfo)
	mock.lockEditor.Unlock()
	return mock.EditorFunc(s)
}

// EditorCalls gets all the calls that were made to Editor.
// Check the length with:
//
//	len(mockedConfig.EditorCalls())
func (mock *ConfigMock) EditorCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockEditor.RLock()
	calls = mock.calls.Editor
	mock.lockEditor.RUnlock()
	return calls
}

// GetOrDefault calls GetOrDefaultFunc.
func (mock *ConfigMock) GetOrDefault(s1 string, s2 string) (string, error) {
	if mock.GetOrDefaultFunc == nil {
		panic("ConfigMock.GetOrDefaultFunc: method is nil but Config.GetOrDefault was just called")
	}
	callInfo := struct {
		S1 string
		S2 string
	}{
		S1: s1,
		S2: s2,
	}
	mock.lockGetOrDefault.Lock()
	mock.calls.GetOrDefault = append(mock.calls.GetOrDefault, callInfo)
	mock.lockGetOrDefault.Unlock()
	return mock.GetOrDefaultFunc(s1, s2)
}

// GetOrDefaultCalls gets all the calls that were made to GetOrDefault.
// Check the length with:
//
//	len(mockedConfig.GetOrDefaultCalls())
func (mock *ConfigMock) GetOrDefaultCalls() []struct {
	S1 string
	S2 string
} {
	var calls []struct {
		S1 string
		S2 string
	}
	mock.lockGetOrDefault.RLock()
	calls = mock.calls.GetOrDefault
	mock.lockGetOrDefault.RUnlock()
	return calls
}

// GitProtocol calls GitProtocolFunc.
func (mock *ConfigMock) GitProtocol(s string) string {
	if mock.GitProtocolFunc == nil {
		panic("ConfigMock.GitProtocolFunc: method is nil but Config.GitProtocol was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockGitProtocol.Lock()
	mock.calls.GitProtocol = append(mock.calls.GitProtocol, callInfo)
	mock.lockGitProtocol.Unlock()
	return mock.GitProtocolFunc(s)
}

// GitProtocolCalls gets all the calls that were made to GitProtocol.
// Check the length with:
//
//	len(mockedConfig.GitProtocolCalls())
func (mock *ConfigMock) GitProtocolCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockGitProtocol.RLock()
	calls = mock.calls.GitProtocol
	mock.lockGitProtocol.RUnlock()
	return calls
}

// HTTPUnixSocket calls HTTPUnixSocketFunc.
func (mock *ConfigMock) HTTPUnixSocket(s string) string {
	if mock.HTTPUnixSocketFunc == nil {
		panic("ConfigMock.HTTPUnixSocketFunc: method is nil but Config.HTTPUnixSocket was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockHTTPUnixSocket.Lock()
	mock.calls.HTTPUnixSocket = append(mock.calls.HTTPUnixSocket, callInfo)
	mock.lockHTTPUnixSocket.Unlock()
	return mock.HTTPUnixSocketFunc(s)
}

// HTTPUnixSocketCalls gets all the calls that were made to HTTPUnixSocket.
// Check the length with:
//
//	len(mockedConfig.HTTPUnixSocketCalls())
func (mock *ConfigMock) HTTPUnixSocketCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockHTTPUnixSocket.RLock()
	calls = mock.calls.HTTPUnixSocket
	mock.lockHTTPUnixSocket.RUnlock()
	return calls
}

// Migrate calls MigrateFunc.
func (mock *ConfigMock) Migrate(migration Migration) error {
	if mock.MigrateFunc == nil {
		panic("ConfigMock.MigrateFunc: method is nil but Config.Migrate was just called")
	}
	callInfo := struct {
		Migration Migration
	}{
		Migration: migration,
	}
	mock.lockMigrate.Lock()
	mock.calls.Migrate = append(mock.calls.Migrate, callInfo)
	mock.lockMigrate.Unlock()
	return mock.MigrateFunc(migration)
}

// MigrateCalls gets all the calls that were made to Migrate.
// Check the length with:
//
//	len(mockedConfig.MigrateCalls())
func (mock *ConfigMock) MigrateCalls() []struct {
	Migration Migration
} {
	var calls []struct {
		Migration Migration
	}
	mock.lockMigrate.RLock()
	calls = mock.calls.Migrate
	mock.lockMigrate.RUnlock()
	return calls
}

// Pager calls PagerFunc.
func (mock *ConfigMock) Pager(s string) string {
	if mock.PagerFunc == nil {
		panic("ConfigMock.PagerFunc: method is nil but Config.Pager was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockPager.Lock()
	mock.calls.Pager = append(mock.calls.Pager, callInfo)
	mock.lockPager.Unlock()
	return mock.PagerFunc(s)
}

// PagerCalls gets all the calls that were made to Pager.
// Check the length with:
//
//	len(mockedConfig.PagerCalls())
func (mock *ConfigMock) PagerCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockPager.RLock()
	calls = mock.calls.Pager
	mock.lockPager.RUnlock()
	return calls
}

// Prompt calls PromptFunc.
func (mock *ConfigMock) Prompt(s string) string {
	if mock.PromptFunc == nil {
		panic("ConfigMock.PromptFunc: method is nil but Config.Prompt was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockPrompt.Lock()
	mock.calls.Prompt = append(mock.calls.Prompt, callInfo)
	mock.lockPrompt.Unlock()
	return mock.PromptFunc(s)
}

// PromptCalls gets all the calls that were made to Prompt.
// Check the length with:
//
//	len(mockedConfig.PromptCalls())
func (mock *ConfigMock) PromptCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockPrompt.RLock()
	calls = mock.calls.Prompt
	mock.lockPrompt.RUnlock()
	return calls
}

// Set calls SetFunc.
func (mock *ConfigMock) Set(s1 string, s2 string, s3 string) {
	if mock.SetFunc == nil {
		panic("ConfigMock.SetFunc: method is nil but Config.Set was just called")
	}
	callInfo := struct {
		S1 string
		S2 string
		S3 string
	}{
		S1: s1,
		S2: s2,
		S3: s3,
	}
	mock.lockSet.Lock()
	mock.calls.Set = append(mock.calls.Set, callInfo)
	mock.lockSet.Unlock()
	mock.SetFunc(s1, s2, s3)
}

// SetCalls gets all the calls that were made to Set.
// Check the length with:
//
//	len(mockedConfig.SetCalls())
func (mock *ConfigMock) SetCalls() []struct {
	S1 string
	S2 string
	S3 string
} {
	var calls []struct {
		S1 string
		S2 string
		S3 string
	}
	mock.lockSet.RLock()
	calls = mock.calls.Set
	mock.lockSet.RUnlock()
	return calls
}

// Version calls VersionFunc.
func (mock *ConfigMock) Version() string {
	if mock.VersionFunc == nil {
		panic("ConfigMock.VersionFunc: method is nil but Config.Version was just called")
	}
	callInfo := struct {
	}{}
	mock.lockVersion.Lock()
	mock.calls.Version = append(mock.calls.Version, callInfo)
	mock.lockVersion.Unlock()
	return mock.VersionFunc()
}

// VersionCalls gets all the calls that were made to Version.
// Check the length with:
//
//	len(mockedConfig.VersionCalls())
func (mock *ConfigMock) VersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockVersion.RLock()
	calls = mock.calls.Version
	mock.lockVersion.RUnlock()
	return calls
}

// Write calls WriteFunc.
func (mock *ConfigMock) Write() error {
	if mock.WriteFunc == nil {
		panic("ConfigMock.WriteFunc: method is nil but Config.Write was just called")
	}
	callInfo := struct {
	}{}
	mock.lockWrite.Lock()
	mock.calls.Write = append(mock.calls.Write, callInfo)
	mock.lockWrite.Unlock()
	return mock.WriteFunc()
}

// WriteCalls gets all the calls that were made to Write.
// Check the length with:
//
//	len(mockedConfig.WriteCalls())
func (mock *ConfigMock) WriteCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockWrite.RLock()
	calls = mock.calls.Write
	mock.lockWrite.RUnlock()
	return calls
}
