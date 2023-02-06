import React from 'react';

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
    <div className={'TimeTag ' + className}>
      <p className='text-xs '>
        {day}.{month}.{year}
      </p>
    </div>
  );
};

export default TimeTag;
