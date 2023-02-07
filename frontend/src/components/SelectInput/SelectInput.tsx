import React, { useState, useRef, useEffect, useCallback, useMemo } from 'react';
import { FieldProps } from 'formik';

import { ReactComponent as ChevronLeftIcon } from '../../icons/chevron.left.svg';
import TextInput from '../TextInput';
import InputError from '../InputError';

export interface OptionConfig {
  label: string;
  value?: string;
}
export interface SelectInputProps {
  options: OptionConfig[];
  value?: string;
  onChange?: (value: string) => void;
  placeholder?: string;
  name?: string;
  className?: string;
  id?: string;
  label?: string;
  isUnstyled?: boolean;
  onFocus?: React.FocusEventHandler;
  onBlur?: React.FocusEventHandler;
  isHovered?: boolean;
  error?: string;
  isMiddle?: boolean;
}

export function generateOptionsFromValues(o: Record<string, unknown>): OptionConfig[] {
  return Object.values(o).map((value) => ({
    label: String(value),
  }));
}

const SelectInput: React.FC<SelectInputProps & Partial<FieldProps>> = ({
  options,
  value,
  onChange,
  placeholder,
  name,
  className,
  id,
  label,
  isUnstyled,
  field,
  form,
  onFocus,
  onBlur,
  isHovered = false,
  error,
  isMiddle = false,
}) => {
  const [visibleText, setVisibleText] = useState('');
  const [searchTerm, setSearchTerm] = useState('');
  const [isMenuOpened, setIsMenuOpened] = useState(false);

  const rootRef = useRef<HTMLDivElement>(null);
  const visibleInputRef = useRef<HTMLInputElement>(null);

  const meta = field?.name !== undefined ? form?.getFieldMeta(field.name) : undefined;

  const anyValue = value || String(meta?.value);
  const anyError = error || (meta?.touched ? meta?.error : undefined);

  const setValue = (value: string) => {
    onChange?.(value);
    if (field?.name !== undefined) {
      form?.setFieldValue(field.name, value);
    }
  };

  const definedValueOptions = useMemo(
    () =>
      options.map((option) => ({
        label: option.label,
        value: option.value || option.label,
      })),
    [options],
  );

  const selectedOption = useMemo(
    () => definedValueOptions.find((option) => option.value === anyValue),
    [definedValueOptions, anyValue],
  );

  const makeValueVisible = useCallback(() => {
    if (selectedOption !== undefined) {
      setVisibleText(selectedOption.label);
    }
  }, [setVisibleText, selectedOption]);

  const onClickAbroad = useCallback(
    (e: MouseEvent) => {
      if (!rootRef.current?.contains(e.target as any)) {
        makeValueVisible();
        setIsMenuOpened(false);
      }
    },
    [makeValueVisible, setIsMenuOpened],
  );

  useEffect(() => {
    document.body.addEventListener('click', onClickAbroad);
    return () => document.body.removeEventListener('click', onClickAbroad);
  }, [onClickAbroad]);

  useEffect(() => {
    makeValueVisible();
  }, [makeValueVisible]);

  const getVisibleOptions = () =>
    definedValueOptions.filter(
      (option) =>
        searchTerm === '' || option.label.toLowerCase().includes(searchTerm.toLowerCase()),
    );

  const chevronClassList = ['transition-transform'];

  if (isMenuOpened) {
    chevronClassList.push('rotate-180');
  }

  const rootClassList = ['relative font-medium text-sm'];

  if (className !== undefined) {
    rootClassList.push(className);
  }

  return (
    <div>
      <div className={rootClassList.join(' ')} id={id} ref={rootRef}>
        <TextInput
          type='text'
          placeholder={placeholder}
          onChange={(e) => {
            setVisibleText(e.target.value);
            setSearchTerm(e.target.value);
          }}
          value={visibleText}
          onFocus={(event) => {
            setIsMenuOpened(true);
            onFocus?.(event);
          }}
          onBlur={onBlur}
          label={label}
          isUnstyled={isUnstyled}
          isHovered={isHovered}
          isMiddle={isMiddle}
          ref={visibleInputRef}
        >
          <button
            type='button'
            onClick={() => {
              if (isMenuOpened) {
                visibleInputRef.current?.blur();
                setIsMenuOpened(false);
                makeValueVisible();
              } else {
                visibleInputRef.current?.focus();
              }
            }}
          >
            <ChevronLeftIcon className={chevronClassList.join(' ')} />
          </button>
        </TextInput>
        {isMenuOpened && getVisibleOptions().length > 0 && (
          <div
            className={[
              'absolute top-full mt-1 z-10 max-h-96 w-full overflow-y-auto rounded-md',
              'bg-background shadow-main-box-shadow p-1.5',
            ].join(' ')}
          >
            {getVisibleOptions().map((option) => {
              const buttonClassList = [
                'block w-full text-left rounded-lg py-3 px-1.5 hover:bg-chosen',
              ];

              if (option.value === anyValue) {
                buttonClassList.push('bg-chosen');
              }

              return (
                <button
                  type='button'
                  className={buttonClassList.join(' ')}
                  key={option.value}
                  onClick={() => {
                    setValue(option.value);
                    setIsMenuOpened(false);
                  }}
                >
                  {option.label}
                </button>
              );
            })}
          </div>
        )}
      </div>
      {anyError !== undefined && <InputError className='mt-0.25'>{anyError}</InputError>}
      {value !== undefined ? (
        <input hidden readOnly name={name} value={value} />
      ) : (
        <input hidden {...field} />
      )}
    </div>
  );
};

export default SelectInput;
