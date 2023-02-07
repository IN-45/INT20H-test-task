import { ChangeEvent, FC, useState } from 'react';
import Card from '../../components/Card';
import DataFetcher from '../../components/DataFetcher/DataFetcher';
import products from '../Products/Products';

const Recipes: FC = () => {
  const [isFiltered, setIsFiltered] = useState(false);

  const handleCheck = () => {
    setIsFiltered((prevState) => !prevState);
  };

  const getAvailability = (products: any[]) => {
    return products.some((product: { [x: string]: number }) => product['missed_amount'] > 0);
  };

  return (
    <>
      Filter by availability:
      <input type={'checkbox'} onChange={handleCheck} className={'mb-4 ml-1'} />
      <DataFetcher url={isFiltered ? 'filter-recipes' : 'recipe'}>
        {(recipes) => (
          <div className={'flex flex-col sm:flex-row gap-8'}>
            {recipes.map((recipe) => (
              <Card
                key={recipe.id}
                {...recipe}
                type={'recipe'}
                unavailable={getAvailability(recipe.products)}
              />
            ))}
          </div>
        )}
      </DataFetcher>
    </>
  );
};

export default Recipes;