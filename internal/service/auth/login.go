package auth

import (
	"context"
	"fmt"
	// Import del pacchetto interno del progetto
)

// domain contains the dependencies required by the authentication service.
func (d *domain) Login(ctx context.Context, email, password string) (string, error) {
	// Check if the user exists
	user, err := d.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	// Verify the password
	if !d.passwordHasher.Verify(password, user.Password) {
		return "", fmt.Errorf("invalid credentials")
	}

	// Generate a token
	token, err := d.tokenGenerator.Generate(ctx, user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
