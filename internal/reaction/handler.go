package reaction

import (
	"forum/internal/shared/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	ReactionService *ReactionService
}

func NewHandler(service *ReactionService) *Handler {
	return &Handler{ReactionService: service}
}

func (h *Handler) React(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Auth check
	userID, ok := middleware.GetUserID(r)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse form
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	reactionType := r.FormValue("type")

	// Validate reaction type
	if reactionType != "like" && reactionType != "dislike" {
		http.Error(w, "invalid reaction type", http.StatusBadRequest)
		return
	}

	var postID *int
	var commentID *int

	// Parse post_id
	if pid := r.FormValue("post_id"); pid != "" {
		if id, err := strconv.Atoi(pid); err == nil {
			postID = &id
		}
	}

	// Parse comment_id
	if cid := r.FormValue("comment_id"); cid != "" {
		if id, err := strconv.Atoi(cid); err == nil {
			commentID = &id
		}
	}

	// Must target something
	if postID == nil && commentID == nil {
		http.Error(w, "no target provided", http.StatusBadRequest)
		return
	}

	reaction := &Reaction{
		UserID:    userID,
		PostID:    postID,
		CommentID: commentID,
		Type:      reactionType,
	}

	// Business logic
	if err := h.ReactionService.React(reaction); err != nil {
		http.Error(w, "failed to react", http.StatusInternalServerError)
		return
	}

	// Redirect back
	ref := r.Referer()
	if ref == "" {
		ref = "/" // fallback
	}

	http.Redirect(w, r, ref, http.StatusSeeOther)
}
