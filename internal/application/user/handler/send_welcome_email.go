package handler

import (
	"context"
	"errors"
	"funding/internal/application/event"
	"funding/internal/domain/user"
	"log"
)

type Mailer interface {
	Send(ctx context.Context, msg WelcomeEmail) error
}

type WelcomeEmail struct {
	Email string
	Name  string
}

type SendWelcomeEmailHandler struct {
	mailer Mailer
}

var _ event.Handler = (*SendWelcomeEmailHandler)(nil)

func NewSendWelcomeEmailHandler(mailer Mailer) *SendWelcomeEmailHandler {
	return &SendWelcomeEmailHandler{mailer: mailer}
}

func (h *SendWelcomeEmailHandler) Handle(ctx context.Context, event any) error {
	e, ok := event.(user.UserRegistered)
	if !ok {
		return errors.New("invalid event type")
	}

	msg := WelcomeEmail{
		Email: e.Email,
		Name:  e.Name,
	}
	return h.mailer.Send(ctx, msg)
}

func (h *SendWelcomeEmailHandler) Send(ctx context.Context, msg WelcomeEmail) error {
	log.Printf("[MockMailer] Sending welcome email to %s (%s)\n", msg.Name, msg.Email)
	return nil
}
