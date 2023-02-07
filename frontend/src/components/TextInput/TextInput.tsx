import React from 'react';
import { FieldProps } from 'formik';

import InputLabel from '../InputLabel';
import InputError from '../InputError';

export type TextInputProps = React.PropsWithChildren<
  React.InputHTMLAttributes<HTMLInputElement> &
    Partial<FieldProps> & {
      label?: string;
      error?: string;
      isUnstyled?: boolean;
      isHovered?: boolean;
      isMiddle?: boolean;
    }
>;

const TextInput = React.forwardRef<HTMLInputElement, TextInputProps>(
  (
    {
      field,
      form,
      label,
      error,
      children,
      className,
      isUnstyled = false,
      isHovered = false,
      isMiddle = false,
      ...props
    },
    ref,
  ) => {
    const anyError =
      (field !== undefined &&
        form !== undefined &&
        form.touched[field.name] &&
        form.errors[field.name]) ||
      error;

    const inputClassList = [
      'focus:outline-none focus:ring-0 block w-full text-sm',
      'placeholder:text-sm placeholder:text-placeholder:text-grey-greyscale placeholder:opacity-70',
    ];

    const rootClassList = [];

    if (!isUnstyled) {
      inputClassList.push(
        'form-control text-black-main text-sm p-3',
        'bg-white-input border border-solid rounded-lg',
        '[&:not(:focus)]:hover:placeholder:text-orange-main [&:not(:focus)]:hover:text-orange-main',
        'focus:border-orange-main',
        'caret-orange-main transition',
        anyError !== undefined ? 'border-red-secondary' : 'border-grey-greyscale',
      );
      rootClassList.push('text-grey-greyscale font-medium');
    } else {
      inputClassList.push('border-0 bg-transparent p-0');
    }

    if (isHovered) {
      inputClassList.push('placeholder:text-orange-main');
    }

    if (isMiddle) {
      inputClassList.push('shadow-border');
    }

    if (className !== undefined) {
      rootClassList.push(className);
    }

    return (
      <div className={rootClassList.join(' ')}>
        {label !== undefined && <InputLabel className='mb-0.25'>{label}</InputLabel>}
        <div className='relative'>
          <input className={inputClassList.join(' ')} ref={ref} {...props} {...field} />
          <div className='absolute top-0 bottom-0 right-3.25 flex items-center'>{children}</div>
        </div>
        {anyError !== undefined && typeof anyError === 'string' && (
          <InputError className='mt-0.25'>{anyError}</InputError>
        )}
      </div>
    );
  },
);

TextInput.displayName = 'TextInput';

export default TextInput;
