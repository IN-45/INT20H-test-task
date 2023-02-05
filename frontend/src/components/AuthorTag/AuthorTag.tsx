import React from 'react';
import { ReactComponent as ProfileIcon } from '../../icons/profile.svg';

interface AuthorTagProps {
  author: string;
  className?: string;
}

const AuthorTag: React.FC<AuthorTagProps> = ({ author, className }) => {
  return (
    <div className={'AuthorTag flex ' + className}>
      <ProfileIcon />
      <p className='text-sm font-medium pl-1.5'>{author}</p>
    </div>
  );
};

export default AuthorTag;
