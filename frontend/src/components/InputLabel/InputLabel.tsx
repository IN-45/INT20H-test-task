import React from 'react';

const InputLabel: React.FC<{ children: React.ReactNode; className?: string }> = ({
  children,
  className,
}) => {
  const classList = ['form-label text-2xs font-medium text-grey-greyscale'];

  if (className !== undefined) {
    classList.push(className);
  }

  return <div className={classList.join(' ')}>{children}</div>;
};

export default InputLabel;
