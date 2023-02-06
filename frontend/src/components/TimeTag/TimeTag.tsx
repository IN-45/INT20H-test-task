import React from 'react';
import { ReactComponent as CalendarIcon } from '../../icons/calendar.svg';

interface TimeTagProps {
  time: Date;
  className?: string;
}

const TimeTag: React.FC<TimeTagProps> = ({ time, className }) => {
  let day = time.getDate().toString();
  if (parseInt(day) < 10) {
    day = `0${day}`;
  }

  let month = (time.getMonth() + 1).toString();
  if (parseInt(month) < 10) {
    month = `0${month}`;
  }

  const year = time.getFullYear();

  return (
    <div className={'flex ' + className}>
      <CalendarIcon className='w-[15px] h-[15px] mr-1' />
      <p className='text-xs '>
        {day}.{month}.{year}
      </p>
    </div>
  );
};

export default TimeTag;
