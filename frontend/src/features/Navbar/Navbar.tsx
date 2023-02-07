import { FC, useState } from 'react';
import NavbarItem, { NavbarItemProps } from './NavbarItem';
import { NavLink, useNavigate } from 'react-router-dom';
import { useCookies } from 'react-cookie';

interface NavbarProps {
  routes: NavbarItemProps[];
}

const Navbar: FC<NavbarProps> = ({ routes }) => {
  const [menuIsShown, setMenuIsShown] = useState(true);
  const [cookies, , removeCookies] = useCookies(['token', 'user_id']);
  const navigate = useNavigate();
  const isLogged = cookies.token !== undefined;

  const loginAction = {
    login: () => {
      navigate('/sign-in');
    },
    logout: () => {
      removeCookies('token');
      removeCookies('user_id');
      navigate('/');
    },
  };

  const tweakMenu = () => {
    setMenuIsShown((prevState) => !prevState);
  };

  return (
    <nav className='bg-navbar flex items-center justify-between flex-wrap p-6'>
      <NavLink to={'/'}>
        <div className='flex items-center flex-shrink-0 text-white mr-6'>
          <span className='font-semibold text-xl tracking-tight'>🍕 Food App</span>
        </div>
      </NavLink>
      <div className='block lg:hidden' onClick={tweakMenu}>
        <button className='flex items-center px-3 py-2 border rounded text-white border-white hover:text-orange-main hover:border-orange-main'>
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
      <div
        className={`${
          menuIsShown ? 'visible' : 'hidden'
        } w-full block flex-grow lg:visible lg:flex lg:items-center lg:w-auto`}
      >
        <div className='text-lg lg:flex-grow'>
          {routes.map((route, index) => (
            <NavbarItem key={index} text={route.text} to={route.to} />
          ))}
        </div>
        <div className={'ml-2 mt-4 lg:mt-0'}>
          <button
            onClick={() => {
              loginAction[isLogged ? 'logout' : 'login']();
            }}
            className={
              'bg-blue-500 hover:bg-blue-600 active:bg-blue-700 text-white font-bold py-2 px-4 rounded select-none'
            }
          >
            {isLogged ? 'Logout' : 'Login'}
          </button>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
