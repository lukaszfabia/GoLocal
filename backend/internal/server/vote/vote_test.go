package vote

import (
	"backend/internal/forms"
	"backend/internal/models"
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVoteService struct {
	mock.Mock
}

func (m *MockVoteService) GetVotes(params map[string]any, limit int) ([]*models.Vote, error) {
	vote := &models.Vote{
		Text: "test vote",
	}

	voteOptions := []*models.VoteOption{}
	voteOptions = append(voteOptions, &models.VoteOption{
		Text: "test option",
		Vote: *vote,
	})

	voteOptions = append(voteOptions, &models.VoteOption{
		Text: "test option 2",
		Vote: *vote,
	})

	votes := []*models.Vote{}
	votes = append(votes, vote)

	return votes, nil
}

func (m *MockVoteService) Vote(form forms.VoteInVotingForm, user models.User) (*models.VoteAnswer, error) {
	m.Called(form)

	return &models.VoteAnswer{
		VoteOptionID: uint(form.VoteOptionID),
		VoteID:       uint(form.VoteID),
		UserID:       user.ID,
	}, nil
}

func TestVote_Success(t *testing.T) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.WriteField("voteID", "1")
	writer.WriteField("voteOptionID", "2")

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/auth/vote", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()

	resultCtx := post(w, req)

	form, ok := resultCtx.Value(_voteForm).(*forms.VoteInVotingForm)

	assert.True(t, ok, "Expected context to be returned")

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, 1, form.VoteID)
	assert.Equal(t, 2, form.VoteOptionID)

	handler := &VoteHandler{
		VoteService: new(MockVoteService),
	}

	handler.vote(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestVote_Unauthorized(t *testing.T) {
	handler := &VoteHandler{
		VoteService: new(MockVoteService),
	}

	form := &forms.VoteInVotingForm{
		VoteID:       1,
		VoteOptionID: 2,
	}

	req := httptest.NewRequest(http.MethodPost, "/api/auth/vote", nil)
	ctx := context.WithValue(req.Context(), _voteForm, form)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	handler.vote(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestVote_InvalidForm(t *testing.T) {
	mockVoteService := new(MockVoteService)
	handler := &VoteHandler{
		VoteService: mockVoteService,
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("invalidField", "invalidValue")
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/auth/vote", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	handler.vote(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockVoteService.AssertExpectations(t)
}
