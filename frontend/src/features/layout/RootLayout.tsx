import { FC } from 'react';
import { Outlet } from 'react-router-dom';
import Navbar from '../Navbar/Navbar';
import { NavbarItemProps } from '../Navbar/NavbarItem';

const RootLayout: FC = () => {
  /* #TODO create a root layout
   * should have navigation?
   * im thinking of bar on top with all sections
   * icon ?
   * probably responsible for authorization redirection
   */


  return (
    <div className={'font-open-sans bg-[#FFC107]'}>
      <Navbar routes={routes} />
      <Outlet />
    </div>
  );
};


const routes: NavbarItemProps[] = [
  {
    to: '/',
    text: 'Home'
  },
  {
    to: '/products',
    text: 'Products'
  },
  {
    to: '/recipes',
    text: 'Recipes'
  }
]

export default RootLayout;