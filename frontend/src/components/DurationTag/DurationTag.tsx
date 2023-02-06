import React from 'react';
import { ReactComponent as ClockIcon } from '../../icons/clock.svg';

interface DurationTagProps {
  duration: number;
  className?: string;
}

const DurationTag: React.FC<DurationTagProps> = ({ duration, className }) => {
  return (
    <div className={'flex ' + className}>
      <ClockIcon className='w-[15px] h-[15px]' />
      <p className='text-sm font-medium ml-1'>{duration} хв</p>
    </div>
  );
};

export default DurationTag;
