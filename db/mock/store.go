// Code generated by MockGen. DO NOT EDIT.
// Source: db/tutorial/store.go
//
// Generated by this command:
//
//	mockgen -source db/tutorial/store.go -destination db/mock/store.go
//
// Package mock_tutorial is a generated GoMock package.
package mock_tutorial

import (
	context "context"
	tutorial "go-practice/db/tutorial"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(ctx context.Context, arg tutorial.CreateAccountParams) (tutorial.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, arg)
	ret0, _ := ret[0].(tutorial.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), ctx, arg)
}

// CreateEntry mocks base method.
func (m *MockStore) CreateEntry(ctx context.Context, arg tutorial.CreateEntryParams) (tutorial.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", ctx, arg)
	ret0, _ := ret[0].(tutorial.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockStoreMockRecorder) CreateEntry(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockStore)(nil).CreateEntry), ctx, arg)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(ctx context.Context, arg tutorial.CreateSessionParams) (tutorial.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, arg)
	ret0, _ := ret[0].(tutorial.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), ctx, arg)
}

// CreateTransfer mocks base method.
func (m *MockStore) CreateTransfer(ctx context.Context, arg tutorial.CreateTransferParams) (tutorial.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", ctx, arg)
	ret0, _ := ret[0].(tutorial.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockStoreMockRecorder) CreateTransfer(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockStore)(nil).CreateTransfer), ctx, arg)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(ctx context.Context, arg tutorial.CreateUserParams) (tutorial.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(tutorial.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), ctx, arg)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), ctx, id)
}

// DeleteEntry mocks base method.
func (m *MockStore) DeleteEntry(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntry", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntry indicates an expected call of DeleteEntry.
func (mr *MockStoreMockRecorder) DeleteEntry(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntry", reflect.TypeOf((*MockStore)(nil).DeleteEntry), ctx, id)
}

// DeleteTransfer mocks base method.
func (m *MockStore) DeleteTransfer(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransfer", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransfer indicates an expected call of DeleteTransfer.
func (mr *MockStoreMockRecorder) DeleteTransfer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransfer", reflect.TypeOf((*MockStore)(nil).DeleteTransfer), ctx, id)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(ctx context.Context, arg tutorial.ListAccountsParams) ([]tutorial.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", ctx, arg)
	ret0, _ := ret[0].([]tutorial.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), ctx, arg)
}

// ListEntries mocks base method.
func (m *MockStore) ListEntries(ctx context.Context) ([]tutorial.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntries", ctx)
	ret0, _ := ret[0].([]tutorial.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntries indicates an expected call of ListEntries.
func (mr *MockStoreMockRecorder) ListEntries(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntries", reflect.TypeOf((*MockStore)(nil).ListEntries), ctx)
}

// ListTransfers mocks base method.
func (m *MockStore) ListTransfers(ctx context.Context) ([]tutorial.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfers", ctx)
	ret0, _ := ret[0].([]tutorial.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfers indicates an expected call of ListTransfers.
func (mr *MockStoreMockRecorder) ListTransfers(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfers", reflect.TypeOf((*MockStore)(nil).ListTransfers), ctx)
}

// SelectAccount mocks base method.
func (m *MockStore) SelectAccount(ctx context.Context, id int64) (tutorial.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAccount", ctx, id)
	ret0, _ := ret[0].(tutorial.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAccount indicates an expected call of SelectAccount.
func (mr *MockStoreMockRecorder) SelectAccount(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAccount", reflect.TypeOf((*MockStore)(nil).SelectAccount), ctx, id)
}

// SelectAccountForUpdate mocks base method.
func (m *MockStore) SelectAccountForUpdate(ctx context.Context, id int64) (tutorial.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAccountForUpdate", ctx, id)
	ret0, _ := ret[0].(tutorial.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAccountForUpdate indicates an expected call of SelectAccountForUpdate.
func (mr *MockStoreMockRecorder) SelectAccountForUpdate(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAccountForUpdate", reflect.TypeOf((*MockStore)(nil).SelectAccountForUpdate), ctx, id)
}

// SelectEntry mocks base method.
func (m *MockStore) SelectEntry(ctx context.Context, id int64) (tutorial.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectEntry", ctx, id)
	ret0, _ := ret[0].(tutorial.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectEntry indicates an expected call of SelectEntry.
func (mr *MockStoreMockRecorder) SelectEntry(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectEntry", reflect.TypeOf((*MockStore)(nil).SelectEntry), ctx, id)
}

// SelectSession mocks base method.
func (m *MockStore) SelectSession(ctx context.Context, username string) (tutorial.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectSession", ctx, username)
	ret0, _ := ret[0].(tutorial.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectSession indicates an expected call of SelectSession.
func (mr *MockStoreMockRecorder) SelectSession(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectSession", reflect.TypeOf((*MockStore)(nil).SelectSession), ctx, username)
}

// SelectTransfer mocks base method.
func (m *MockStore) SelectTransfer(ctx context.Context, id int64) (tutorial.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectTransfer", ctx, id)
	ret0, _ := ret[0].(tutorial.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectTransfer indicates an expected call of SelectTransfer.
func (mr *MockStoreMockRecorder) SelectTransfer(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectTransfer", reflect.TypeOf((*MockStore)(nil).SelectTransfer), ctx, id)
}

// SelectUser mocks base method.
func (m *MockStore) SelectUser(ctx context.Context, username string) (tutorial.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUser", ctx, username)
	ret0, _ := ret[0].(tutorial.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUser indicates an expected call of SelectUser.
func (mr *MockStoreMockRecorder) SelectUser(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUser", reflect.TypeOf((*MockStore)(nil).SelectUser), ctx, username)
}

// TransferTx mocks base method.
func (m *MockStore) TransferTx(ctx context.Context, arg tutorial.TransferTxParams) (tutorial.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", ctx, arg)
	ret0, _ := ret[0].(tutorial.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockStoreMockRecorder) TransferTx(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockStore)(nil).TransferTx), ctx, arg)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(ctx context.Context, arg tutorial.UpdateAccountParams) (tutorial.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", ctx, arg)
	ret0, _ := ret[0].(tutorial.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), ctx, arg)
}
