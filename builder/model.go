package builder

type Role int

const (
	Manager Role = iota
	Member
)

type Employee struct {
	Name     string
	RoleType Role
	Salary   int
}

type Builder struct {
	e Employee
}

func NewBuilder() *Builder {
	b := new(Builder)
	//b.e = *(new(Employee))
	return b
}

func (b *Builder) Build() Employee {
	return b.e
}

func (b *Builder) name(name string) *Builder {
	b.e.Name = name
	return b
}

func (b *Builder) role(role Role) *Builder {
	b.e.RoleType = role
	return b
}

func (b *Builder) salary(salary int) *Builder {
	b.e.Salary = salary
	return b
}
