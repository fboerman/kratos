package template

import (
	"context"
	"encoding/json"
	"os"
)

type (
	RecoveryInvalid struct {
		d TemplateDependencies
		m *RecoveryInvalidModel
	}
	RecoveryInvalidModel struct {
		To string
	}
)

func NewRecoveryInvalid(d TemplateDependencies, m *RecoveryInvalidModel) *RecoveryInvalid {
	return &RecoveryInvalid{d: d, m: m}
}

func (t *RecoveryInvalid) EmailRecipient() (string, error) {
	return t.m.To, nil
}

func (t *RecoveryInvalid) EmailSubject(ctx context.Context) (string, error) {
	return LoadTextTemplate(ctx, t.d, os.DirFS(t.d.CourierConfig(ctx).CourierTemplatesRoot()), "recovery/invalid/email.subject.gotmpl", "recovery/invalid/email.subject*", t.m, t.d.CourierConfig(ctx).CourierTemplatesRecoveryInvalid().Subject)
}

func (t *RecoveryInvalid) EmailBody(ctx context.Context) (string, error) {
	return LoadHTMLTemplate(ctx, t.d, os.DirFS(t.d.CourierConfig(ctx).CourierTemplatesRoot()), "recovery/invalid/email.body.gotmpl", "recovery/invalid/email.body*", t.m, t.d.CourierConfig(ctx).CourierTemplatesRecoveryInvalid().Body.HTML)
}

func (t *RecoveryInvalid) EmailBodyPlaintext(ctx context.Context) (string, error) {
	return LoadTextTemplate(ctx, t.d, os.DirFS(t.d.CourierConfig(ctx).CourierTemplatesRoot()), "recovery/invalid/email.body.plaintext.gotmpl", "recovery/invalid/email.body.plaintext*", t.m, t.d.CourierConfig(ctx).CourierTemplatesRecoveryInvalid().Body.PlainText)
}

func (t *RecoveryInvalid) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.m)
}
