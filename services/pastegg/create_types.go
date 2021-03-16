package pastegg

import "time"

/*
 * #=================#
 * # R E S P O N S E #
 * #=================#
 */

type responseCreate struct {
	Status string                `json:"status"`
	Result *responseCreateResult `json:"result"`
}

type responseCreateResult struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Visibility  string    `json:"visibility"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// optional
	Expires *time.Time            `json:"expires"`
	Files   []*responseCreateFile `json:"files"`
	// optional
	DeletionKey string `json:"deletion_key"`
}

type responseCreateFile struct {
	HighlightLanguage interface{} `json:"highlight_language"`
	ID                string      `json:"id"`
	Name              string      `json:"name"`
}

/*
 * #===============#
 * # R E Q U E S T #
 * #===============#
 */

type Create struct {
	// optional
	Name string `json:"name,omitempty"`
	// optional
	Description string `json:"description,omitempty"`
	// optional
	Visibility string `json:"visibility,omitempty"`
	// optional
	Expires *time.Time `json:"expires,omitempty"`
	// required
	Files []*CreateFile `json:"files"`
}

type CreateFile struct {
	// optional
	Name string `json:"name,omitempty"`
	// required
	Content *CreateFileContent `json:"content"`
}

type CreateFileContent struct {
	// required
	Format string `json:"format"`
	// optional
	HighlightLanguage string `json:"highlight_language,omitempty"`
	// required
	Value string `json:"value"`
}
