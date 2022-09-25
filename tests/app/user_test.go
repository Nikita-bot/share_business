// Code generated by mtgroup-generator.
package app

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func TestGetUser(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)

	a := New(mockRepo)

	rulesSet.EXPECT().GetUserAccessManager(gomock.Any()).Return(true)
	mockRepo.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(testUser1, nil)
	b, err := a.GetUser(profile, testUser1.ID)
	t.Nil(err)
	t.DeepEqual(testUser1, b)
}

func TestAddUser(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)

	a := New(mockRepo)

	rulesSet.EXPECT().AddUserAccessManager(gomock.Any()).Return(true)
	mockRepo.EXPECT().AddUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(testUser1, nil)
	b, err := a.AddUser(profile, testUser1)
	t.Nil(err)
	t.DeepEqual(testUser1, b)
}

func TestEditUser(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)

	a := New(mockRepo)

	rulesSet.EXPECT().EditUserAccessManager(gomock.Any()).Return(true)
	mockRepo.EXPECT().EditUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	err := a.EditUser(profile, testUser1.ID, testUser1)
	t.Nil(err)
}

func TestDeleteUser(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)

	a := New(mockRepo)

	rulesSet.EXPECT().DeleteUserAccessManager(gomock.Any()).Return(true)
	mockRepo.EXPECT().DeleteUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	err := a.DeleteUser(profile, testUser1.ID)
	t.Nil(err)
}

func TestListUser(tt *testing.T) {
	t := check.T(tt)
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepo(ctrl)

	a := New(mockRepo)

	rulesSet.EXPECT().ListUserAccessManager(gomock.Any()).Return(true)
	mockRepo.EXPECT().ListUser(gomock.Any(), gomock.Any()).Return(testUsers, []string{}, nil)
	b, _, err := a.ListUser(profile, listParams)
	t.Nil(err)
	t.DeepEqual(testUsers, b)
}
