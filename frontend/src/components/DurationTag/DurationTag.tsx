import React from 'react';
import { ReactComponent as ClockIcon } from '../../icons/clock.svg';

interface CategoryTagProps {
  duration: number;
  className?: string;
}

const CategoryTag: React.FC<CategoryTagProps> = ({ duration, className }) => {
  return (
    <div className={'CategoryTag flex ' + className}>
      <ClockIcon className='w-[15px] h-[15px]' />
      <p className='text-sm font-medium ml-2'>{duration} хв</p>
    </div>
  );
};

export default CategoryTag;
