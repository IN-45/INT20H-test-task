import { FC, ReactNode } from 'react';
import { DataState } from './DataFetcher';

interface StateResolverProps {
  state: keyof typeof DataState;
  empty: boolean;
  children?: ReactNode;
  error: any;
}

const StateResolver: FC<StateResolverProps> = ({ state, children, empty, error }) => {
  if (state !== 'loading' && empty) return <h1>No data here for now...</h1>;

  switch (state) {
    case 'error':
      return (
        <div>
          <p>There seem to be an error...</p>
          <p>{error.message}</p>
        </div>
      );
    case 'loading':
      return <p>Loading...</p>;
    case 'ready':
      return <div>{children}</div>;
  }
};

export default StateResolver;