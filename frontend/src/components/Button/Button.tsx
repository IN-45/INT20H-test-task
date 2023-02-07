import React from 'react';

interface ButtonProps {
  type: 'button' | 'submit' | 'reset';
  label: string;
  disabled?: boolean;
  className?: string;
  onClick?: () => void;
}

const Button: React.FC<ButtonProps> = ({ type, label, disabled = false, className, onClick }) => {
  return (
    <button
      type={type}
      disabled={disabled}
      onClick={onClick}
      className={
        'bg-orange-main font-bold text-white-greyscale rounded-lg border-2 border-orange-main hover:bg-transparent hover:text-orange-main transition-colors disabled:bg-orange-disabled disabled:border-orange-disabled disabled:text-white-greyscale ' +
        className
      }
    >
      {label}
    </button>
  );
};

export default Button;
