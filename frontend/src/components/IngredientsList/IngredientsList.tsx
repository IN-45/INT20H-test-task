import React from 'react';

import Checkbox from '../Checkbox';

interface Product {
  name: string;
  amount: string;
}

interface ProductsListProps {
  products: Product[];
}

const ProductsList: React.FC<ProductsListProps> = ({ products }) => {
  return (
    <div>
      <h3 className='text-xl mb-7'>Ingredients</h3>
      <ul className='text-lg'>
        {products.map((product, index) => {
          return (
            <li key={index} className='mb-5'>
              <Checkbox
                label={product.amount + ' ' + product.name}
                name={product.name}
                className='w-4 h-4 accent-orange-main'
              />
            </li>
          );
        })}
      </ul>
    </div>
  );
};

export default ProductsList;
