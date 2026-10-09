package vault

import "testing"

// TestNew проверяет создание нового пустого хранилища.
func TestNew(t *testing.T) {
	storage := New()

	if storage == nil {
		t.Fatal("New() вернула nil")
	}

	if storage.Count() != 0 {
		t.Errorf(
			"Ожидалось 0 записей, получено %d",
			storage.Count(),
		)
	}
}
