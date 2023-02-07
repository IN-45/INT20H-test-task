import { ChangeEvent, FC, useState } from 'react';
import DataFetcher from '../../components/DataFetcher/DataFetcher';
import Card from '../../components/Card';

const Products: FC = () => {
  const [category, setCategory] = useState<string>();

  const handleSelect = (evt: ChangeEvent<HTMLSelectElement>) => {
    setCategory(evt.target.value);
  };

  const getProducts = (products: any[]) => {
    let productsCards = [...products];
    if (category) {
      productsCards = products.filter((product) => product.category === category);
    }
    const Cards = productsCards.map((item) => <Card key={item.id} {...item} type={'product'} />);
    return Cards.length > 0 ? Cards : 'No products found';
  };

  return (
    <DataFetcher url={'products'}>
      {(products) => (
        <DataFetcher url={'categories'}>
          {(categories) => (
            <>
              Filter by Category:{' '}
              <select
                className={'py-2 px-4 mb-4 ml-2 bg-blue-100 rounded-full focus:outline-gray-500'}
                onChange={handleSelect}
                defaultValue={''}
              >
                <option value={''}>All Categories</option>
                {categories.map((category) => (
                  <option key={category.id}>{category.name}</option>
                ))}
              </select>
              <div className={'flex flex-col sm:flex-row gap-8'}>{getProducts(products)}</div>
            </>
          )}
        </DataFetcher>
      )}
    </DataFetcher>
  );
};

export default Products;