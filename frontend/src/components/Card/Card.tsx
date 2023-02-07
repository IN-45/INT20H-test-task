import React from 'react';
import { Link } from 'react-router-dom';

import DurationTag from '../DurationTag';

interface CardProps {
  id: string;
  image_url?: string;
  name: string;
  cooking_time_minutes?: number;
  type: 'product' | 'recipe';
}

const Card: React.FC<CardProps> = ({ id, name, image_url, cooking_time_minutes, type }) => {
  const getCustomStyles = () => {
    if (type === 'recipe') {
      return 'hover:text-orange-main';
    }
  };

  const cardContent = (
    <div className={'sm:max-w-[400px] space-y-3 relative ' + getCustomStyles()} id={id}>
      <img
        src={image_url}
        alt={name}
        title={name}
        className='xl:h-[246px] sm:h-[200px] h-[320px] w-full object-cover max-sm:brightness-50'
        onError={({ currentTarget }) => {
          currentTarget.onerror = null;
          currentTarget.src = require('../../assets/placeholder.jpeg');
        }}
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

  return (
    <div>{type === 'recipe' ? <Link to={`/recipe/${id}`}>{cardContent}</Link> : cardContent}</div>
  );
};

export default Card;
