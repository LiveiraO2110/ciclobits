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

func BuiltApp(db *gorm.DB) (*App){
	
	estoque  		  := modules.NewEstoque(db)
	material 		  := modules.NewMaterial(db)
	pedido   		  := modules.NewPedido(db)
	produto  		  := modules.NewProduto(db)
	recurso  		  := modules.NewRecurso(db)
	ordens_producao   := modules.NewOrdensProducao(db)
	estrutura_produto := modules.NewEstruturaProduto(db)
	
	return &App{
		Estoque:  		   estoque,
		Material: 		   material,
		Pedido: 	 	   pedido,
		Produto:  		   produto,
		Recurso:  	       recurso,
		OrdensProducao:    ordens_producao,
		EstruturaProduto:  estrutura_produto,
	}
}