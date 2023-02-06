import { FC, useState } from 'react';
import NavbarItem, { NavbarItemProps } from './NavbarItem';
import { NavLink } from 'react-router-dom';

interface NavbarProps {
  routes: NavbarItemProps[];
}

const Navbar: FC<NavbarProps> = ({ routes }) => {
  const [menuIsShown, setMenuIsShown] = useState(true);

  const tweakMenu = () => {
    setMenuIsShown((prevState) => !prevState);
  };

  return (
    <nav className='flex items-center justify-between flex-wrap bg-navbar p-6'>
      <div className='flex items-center flex-shrink-0 text-white mr-6'>
        <span className='font-semibold text-xl tracking-tight'>🍕 Food App</span>
      </div>
      <div className='block lg:hidden' onClick={tweakMenu}>
        <button className='flex items-center px-3 py-2 border rounded text-teal-200 border-teal-400 hover:text-white hover:border-white'>
          <svg
            className='fill-current h-3 w-3'
            viewBox='0 0 20 20'
            xmlns='http://www.w3.org/2000/svg'
          >
            <title>Menu</title>
            <path d='M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z' />
          </svg>
        </button>
      </div>
      {menuIsShown && (
        <div className='w-full block flex-grow lg:flex lg:items-center lg:w-auto'>
          <div className='text-lg lg:flex-grow'>
            {routes.map((route, index) => (
              <NavbarItem key={index} text={route.text} to={route.to} />
            ))}
          </div>
          <div className={'ml-2 mt-4'}>
            <NavLink
              to={'/login'}
              className={({ isActive }) =>
                'bg-blue-500 hover:bg-blue-600 active:bg-blue-700 text-white font-bold py-2 px-4 rounded select-none' +
                (isActive ? ' bg-blue-800' : '')
              }
            >
              Login
            </NavLink>
          </div>
        </div>
      )}
    </nav>
  );
};

export default Navbar;