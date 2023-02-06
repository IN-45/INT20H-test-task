import React from 'react';

interface InstructionsProps {
  instructions: string[];
}

const Instructions: React.FC<InstructionsProps> = ({ instructions }) => {
  return (
    <div>
      <h3 className='text-xl mb-5'>Instructions</h3>
      <ol className='text-lg'>
        {instructions.map((instruction, index) => {
          return (
            <li key={index} className='flex mb-7'>
              <div className='mr-3 py-1 px-2 font-bold bg-orange-main rounded-full h-[28px] text-md text-white'>
                {index + 1}
              </div>
              <div>{instruction}</div>
            </li>
          );
        })}
      </ol>
    </div>
  );
};

export default Instructions;
