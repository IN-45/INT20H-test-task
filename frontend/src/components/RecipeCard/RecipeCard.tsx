import React from 'react';

import AuthorTag from '../AuthorTag';
import DurationTag from '../DurationTag';
import TimeTag from '../TimeTag';

interface RecipeCardProps {
  id: string;
  image: string;
  name: string;
  author: string;
  cooking_time_minutes: number;
  created_at: Date;
}

const RecipeCard: React.FC<RecipeCardProps> = ({
  id,
  name,
  image,
  author,
  cooking_time_minutes,
  created_at,
}) => {
  return (
    <div className='Card sm:max-w-[400px] space-y-3 hover:text-orange-main relative' id={id}>
      <img
        src={image}
        alt={name}
        title={name}
        className='xl:h-[246px] sm:h-[200px] h-[320px] w-full object-cover max-sm:brightness-50'
      />
      <div className='space-y-3 max-sm:absolute max-sm:bottom-[20px] max-sm:px-4 w-full'>
        <DurationTag duration={cooking_time_minutes} className='text-white sm:text-orange-main' />
        <h4 className='max-sm:text-white mt-0'>{name}</h4>
        <div className='w-full flex justify-between text-white sm:text-black-secondary sm:opacity-50'>
          <AuthorTag author={author} />
          <TimeTag time={created_at} />
        </div>
      </div>
    </div>
  );
};

export default RecipeCard;
