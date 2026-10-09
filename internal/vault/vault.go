package vault

// Vault представляет хранилище записей.
// На текущем этапе структура содержит только количество записей.
// Механизмы шифрования и сохранения будут добавлены отдельно.
type Vault struct {
	count int
}

// New создаёт новое пустое хранилище.
func New() *Vault {
	return &Vault{}
}

// Count возвращает количество записей в хранилище.
func (v *Vault) Count() int {
	return v.count
}
