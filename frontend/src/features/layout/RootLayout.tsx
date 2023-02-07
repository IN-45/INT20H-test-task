import { FC } from 'react';
import { Outlet } from 'react-router-dom';
import Navbar from '../Navbar/Navbar';
import { NavbarItemProps } from '../Navbar/NavbarItem';

const RootLayout: FC = () => {
  return (
    <div className='font-open-sans'>
      <Navbar routes={routes} />
      <div className='mx-auto pt-5 pb-12 xl:w-[1300px] lg:w-[1030px] md:w-[940px] max-md:m-5'>
        <Outlet />
      </div>
    </div>
  );
};

const routes: NavbarItemProps[] = [
  {
    to: '/products',
    text: 'Products',
  },
  {
    to: '/recipes',
    text: 'Recipes',
  },
];

export default RootLayout;
