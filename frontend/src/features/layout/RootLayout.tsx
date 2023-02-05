import { FC } from 'react';
import { Outlet } from 'react-router-dom';

const RootLayout: FC = () => {
  /* #TODO create a root layout
   * should have navigation?
   * im thinking of bar on top with all sections
   * icon ?
   * probably responsible for authorization redirection
   */

  return (
    <div>
      <h1>RootLayout</h1>
      <Outlet />
    </div>
  );
};

export default RootLayout;