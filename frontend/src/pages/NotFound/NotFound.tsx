import React, { useEffect } from 'react';

const NotFound = () => {
  useEffect(() => {
    document.title = 'Сторінку не знайдено';
  }, []);

  return (
    <div className='text-center absolute top-1/2 right-1/2 -translate-y-1/2 translate-x-1/2'>
      <h1 className='text-2xl text-purple-main'>Помилка 404</h1>
      <h2 className='text-xl pt-3'>Схоже, що адреса введена хибно або сторінка була видалена</h2>
    </div>
  );
};

export default NotFound;
