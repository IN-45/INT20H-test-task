import React from 'react';

const InputError: React.FC<{ children: React.ReactNode; className?: string }> = ({
  children,
  className,
}) => {
  const classList = ['text-2xs font-medium text-red-secondary'];

  if (className !== undefined) {
    classList.push(className);
  }

  return <div className={classList.join(' ')}>{children}</div>;
};

export default InputError;
