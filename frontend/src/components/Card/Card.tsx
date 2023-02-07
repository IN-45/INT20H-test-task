import React from 'react';

import DurationTag from '../DurationTag';

interface CardProps {
  id: string;
  image_url: string;
  name: string;
  cooking_time_minutes?: number;
}

const Card: React.FC<CardProps> = ({ id, name, image_url, cooking_time_minutes }) => {
  return (
    <div className='sm:max-w-[400px] space-y-3 hover:text-orange-main relative' id={id}>
      <img
        src={image_url}
        alt={name}
        title={name}
        className='xl:h-[246px] sm:h-[200px] h-[320px] w-full object-cover max-sm:brightness-50'
      />
      <div className='space-y-3 max-sm:absolute max-sm:bottom-[20px] max-sm:px-4 w-full'>
        {cooking_time_minutes && (
          <DurationTag
            duration={cooking_time_minutes}
            className='text-white-greyscale sm:text-orange-main'
          />
        )}
        <h4 className='max-sm:text-white-greyscale mt-0'>{name}</h4>
      </div>
    </div>
  );
};

export default Card;
