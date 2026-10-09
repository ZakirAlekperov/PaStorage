package vault

import "testing"

// TestNewEntry проверяет создание записи с заданными данными.
func TestNewEntry(t *testing.T) {
	title := "GitHub"
	username := "user@example.com"
	password := "test-password"

	entry := NewEntry(title, username, password)

	if entry.Title != title {
		t.Errorf(
			"Название: ожидалось %q, получено %q",
			title,
			entry.Title,
		)
	}

	if entry.Username != username {
		t.Errorf(
			"Пользователь: ожидалось %q, получено %q",
			username,
			entry.Username,
		)
	}

	if entry.Password != password {
		t.Errorf(
			"Пароль не соответствует переданному значению",
		)
	}
}
