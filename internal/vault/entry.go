package vault

// Entry представляет одну запись в хранилище паролей.
type Entry struct {
	Title    string
	Username string
	Password string
}

// NewEntry создаёт новую запись с указанными данными.
func NewEntry(title, username, password string) *Entry {
	return &Entry{
		Title:    title,
		Username: username,
		Password: password,
	}
}
