package models

// Code demo, xac dinh Role va Quyen
type Role interface {
	GetName() string
	GetPermissions() []string
	// Login()
	// Register()
}

type Owner struct {
}

func (o *Owner) GetName() string {
	return "Owner"
}

func (o *Owner) GetPermissions() []string {
	return []string{"read", "write", "delete"}
}

type Admin struct{}

func (a *Admin) GetName() string {
	return "Admin"
}

func (a *Admin) GetPermissions() []string {
	return []string{"read", "write"}
}

type Staff struct{}

func (u *Staff) GetName() string {
	return "User"
}

func (u *Staff) GetPermissions() []string {
	return []string{"read"}
}
