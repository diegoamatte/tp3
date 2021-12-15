package biblioteca

type Cola interface{
	EstaVacia()(bool)
	Encolar(valor interface{});
	VerPrimero()interface{};
	Desencolar()interface{}
}