// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/hootsuite/atlantis/server/events (interfaces: ProjectConfigReader)

package mocks

import (
	"reflect"

	events "github.com/hootsuite/atlantis/server/events"
	pegomock "github.com/petergtz/pegomock"
)

type MockProjectConfigReader struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectConfigReader() *MockProjectConfigReader {
	return &MockProjectConfigReader{fail: pegomock.GlobalFailHandler}
}

func (mock *MockProjectConfigReader) Exists(projectPath string) bool {
	params := []pegomock.Param{projectPath}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Exists", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockProjectConfigReader) Read(projectPath string) (events.ProjectConfig, error) {
	params := []pegomock.Param{projectPath}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Read", params, []reflect.Type{reflect.TypeOf((*events.ProjectConfig)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 events.ProjectConfig
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(events.ProjectConfig)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockProjectConfigReader) VerifyWasCalledOnce() *VerifierProjectConfigReader {
	return &VerifierProjectConfigReader{mock, pegomock.Times(1), nil}
}

func (mock *MockProjectConfigReader) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierProjectConfigReader {
	return &VerifierProjectConfigReader{mock, invocationCountMatcher, nil}
}

func (mock *MockProjectConfigReader) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierProjectConfigReader {
	return &VerifierProjectConfigReader{mock, invocationCountMatcher, inOrderContext}
}

type VerifierProjectConfigReader struct {
	mock                   *MockProjectConfigReader
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierProjectConfigReader) Exists(projectPath string) *ProjectConfigReader_Exists_OngoingVerification {
	params := []pegomock.Param{projectPath}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Exists", params)
	return &ProjectConfigReader_Exists_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ProjectConfigReader_Exists_OngoingVerification struct {
	mock              *MockProjectConfigReader
	methodInvocations []pegomock.MethodInvocation
}

func (c *ProjectConfigReader_Exists_OngoingVerification) GetCapturedArguments() string {
	projectPath := c.GetAllCapturedArguments()
	return projectPath[len(projectPath)-1]
}

func (c *ProjectConfigReader_Exists_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierProjectConfigReader) Read(projectPath string) *ProjectConfigReader_Read_OngoingVerification {
	params := []pegomock.Param{projectPath}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Read", params)
	return &ProjectConfigReader_Read_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ProjectConfigReader_Read_OngoingVerification struct {
	mock              *MockProjectConfigReader
	methodInvocations []pegomock.MethodInvocation
}

func (c *ProjectConfigReader_Read_OngoingVerification) GetCapturedArguments() string {
	projectPath := c.GetAllCapturedArguments()
	return projectPath[len(projectPath)-1]
}

func (c *ProjectConfigReader_Read_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}
