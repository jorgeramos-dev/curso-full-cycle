import ProductAdmFacade from "../facade/product-adm.facade";
import ProductRepository from "../repository/product.repository";
import AddProductUseCase from "../usecase/add-product/add-product.usecase";
import CheckStockUseCase from "../usecase/check-stock/check-stock.usecase";

export default class ProductAdmFacadeFactory {
    static create(): ProductAdmFacade {
        //Esse cara fica extretamente acoplado com o Sequelize
        const productRepository = new ProductRepository();
        const addProductUseCase = new AddProductUseCase(productRepository);
        const checkStockUseCase = new CheckStockUseCase(productRepository);
        const productFacade = new ProductAdmFacade({
            addProductUseCase: addProductUseCase,
            checkStockUseCase: checkStockUseCase
        });

        return productFacade;
    }
}