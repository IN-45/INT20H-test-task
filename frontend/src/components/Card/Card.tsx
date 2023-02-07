import React from 'react';

import AuthorTag from '../AuthorTag';
import DurationTag from '../DurationTag';
import TimeTag from '../TimeTag';

interface CardProps {
  id: string;
  image_url: string;
  name: string;
  author?: string;
  cooking_time_minutes?: number;
  created_at?: Date;
}

const Card: React.FC<CardProps> = ({
  id,
  name,
  image_url,
  author,
  cooking_time_minutes,
  created_at,
}) => {
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
        {author && created_at && (
          <div className='w-full flex justify-between text-white-greyscale sm:text-black-secondary sm:opacity-50'>
            <AuthorTag author={author} />
            <TimeTag time={created_at} />
          </div>
        )}
      </div>
    </div>
  );
};

export default Card;
