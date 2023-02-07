import { FC } from 'react';
import DataFetcher from '../../components/DataFetcher/DataFetcher';
import Card from '../../components/Card';

const Products: FC = () => {
  return (
    <DataFetcher url={'products'}>
      {(data) => data.map((item) => <Card key={item.id} {...item} />)}
    </DataFetcher>
  );
};

export default Products;