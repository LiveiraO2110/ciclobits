package bootstrap

import (
	"ciclobits/internal/bootstrap/modules"

	"gorm.io/gorm"
)

type App struct{
	Estoque  		 *modules.Estoque
	Material 		 *modules.Material
	Pedido 	 		 *modules.Pedido
	Produto  		 *modules.Produto
	Recurso  	     *modules.Recurso
	OrdensProducao   *modules.OrdensProducao
	EstruturaProduto *modules.EstruturaProduto
}

func BuiltApp(db *gorm.DB) (App){
	
}