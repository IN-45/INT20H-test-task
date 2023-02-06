import React from 'react';

interface CheckboxProps {
  label: string;
  name: string;
  className?: string;
}

const Checkbox: React.FC<CheckboxProps> = ({ label, name, className }) => {
  return (
    <div>
      <input
        type='checkbox'
        name={name.split(' ').join('-')}
        id={name.split(' ').join('-')}
        className={className}
      />
      <label htmlFor={name.split(' ').join('-')} className='pl-2 strikethrough'>
        {label}
      </label>
    </div>
  );
};

export default Checkbox;
