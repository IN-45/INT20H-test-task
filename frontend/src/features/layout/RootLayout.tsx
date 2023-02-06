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
      <div className='mx-auto mt-5 mb-12 xl:w-[1300px] lg:w-[1030px] md:w-[940px] max-md:m-5 h-screen'>
        <Outlet />
      </div>
    </div>
  );
};

export default RootLayout;
