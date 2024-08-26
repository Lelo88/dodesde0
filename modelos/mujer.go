package modelos

type Mujer struct {
	Hombre	// hereda toda la estructura de hombre
	EsMadre bool // le podemos añadir atributos de nuestra manera
}

func (m *Mujer) Respirar() {
	m.respirando = true
}

func (m *Mujer) Comer() {
	m.comiendo = true
}

func (m *Mujer) Pensar() {
	m.pensando = true
}

func (m *Mujer) sexo() string {
	return "mujer"
}

func (m *Mujer) Sexo() string {
	return m.sexo()
}