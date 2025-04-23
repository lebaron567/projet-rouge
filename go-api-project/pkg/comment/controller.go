package comment

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/config"
	"api/database/dbmodel"
	model "api/pkg/models"

	"github.com/go-chi/chi/v5"
)

// AddComment godoc
// @Summary Add a comment to a post
// @Description Add a comment to a post
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 201
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to add comment"
// @Router /posts/{id}/comments [post]
func AddComment(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		var comment dbmodel.Comment
		if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		comment.IDPost = uint(postID)
		if _, err := cfg.CommentRepository.Create(&comment); err != nil {
			http.Error(w, "Failed to add comment", http.StatusInternalServerError)
			return
		}

		response := model.Comment{
			ID:      comment.ID,
			IDUser:  comment.IDUser,
			IDPost:  comment.IDPost,
			Content: comment.Content,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

// GetComments godoc
// @Summary Get comments for a post
// @Description Get comments for a post
// @Tags comments
// @Produce json
// @Param id path int true "Post ID"
// @Success 200
// @Failure 400 {string} string "Invalid post ID"
// @Failure 500 {string} string "Failed to get comments"
// @Router /posts/{id}/comments [get]
func GetComments(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		comments, err := cfg.CommentRepository.FindByPostID(postID)
		if err != nil {
			http.Error(w, "Failed to get comments", http.StatusInternalServerError)
			return
		}

		var response []model.Comment
		for _, comment := range comments {
			response = append(response, model.Comment{
				ID:      comment.ID,
				IDUser:  comment.IDUser,
				IDPost:  comment.IDPost,
				Content: comment.Content,
			})
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

// DeleteComment godoc
// @Summary Delete a comment by ID
// @Description Delete a comment by ID
// @Tags comments
// @Param id path int true "Comment ID"
// @Success 204
// @Failure 400 {string} string "Invalid comment ID"
// @Failure 500 {string} string "Failed to delete comment"
// @Router /comments/{id} [delete]
func DeleteComment(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		commentID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid comment ID", http.StatusBadRequest)
			return
		}

		if err := cfg.CommentRepository.Delete(commentID); err != nil {
			http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
