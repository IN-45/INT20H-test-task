import { FC } from 'react';
import { NavLink } from 'react-router-dom';

export interface NavbarItemProps {
  text: string;
  to: string;
}

const NavbarItem: FC<NavbarItemProps> = ({ to, text }) => {
  return (
    <NavLink
      end
      to={to}
      className={({ isActive }) =>
        'block w-fit mt-4 lg:inline-block lg:mt-0 text-orange-main hover:bg-orange-main hover:text-white active:bg-orange-main active:text-white active:outline-none py-2 px-4 rounded mx-2 select-none' +
        (isActive ? ' py-2 px-4 bg-navbar-link-active' : '')
      }
    >
      {text}
    </NavLink>
  );
};

export default NavbarItem;
