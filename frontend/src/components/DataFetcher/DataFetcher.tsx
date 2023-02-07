import { FC, ReactNode, useEffect, useState } from 'react';
import StateResolver from './StateResolver';
import { requestApi } from '../../config/api';

interface DataFetcherProps {
  url: string;
  children?: ReactNode | ((params: any[]) => ReactNode);
}

export enum DataState {
  loading = 'loading',
  error = 'error',
  ready = 'ready',
}

const DataFetcher: FC<DataFetcherProps> = ({ children, url }) => {
  const [data, setData] = useState([]);
  const [error, setError] = useState();
  const [state, setState] = useState<keyof typeof DataState>('loading');

  useEffect(() => {
    setState('loading');
    requestApi('GET', url)
      .then((res) => {
        setData(res.data);
        setState('ready');
      })
      .catch((err) => {
        setState('error');
        setError(err);
      });
  }, []);

  return (
    <StateResolver state={state} empty={data === null || data?.length === 0} error={error}>
      {typeof children === 'function' ? children(data) : children}
    </StateResolver>
  );
};

export default DataFetcher;