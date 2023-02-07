import React, { ChangeEvent, useEffect, useRef, useState } from 'react';
import { ReactComponent as EditIcon } from '../../icons/edit.svg';
import { ReactComponent as DocumentIcon } from '../../icons/document.svg';
import { ReactComponent as LoadingIcon } from '../../icons/loading-file.svg';
import {
  checkFile,
  getButtonStyles,
  getImgStyles,
  getStylesByError,
  getStylesByState,
  getStylesByType,
  getTextStyles,
} from './helpers';

interface UploadFileProps {
  initialFile?: string | File;
  onChange?: (image: File) => void;
  type?: 'avatar' | 'document';
  text?: string;
}

const UploadFile = ({ initialFile, onChange, type = 'document', text = '' }: UploadFileProps) => {
  const inputRef = useRef<HTMLInputElement>(null);
  const [pdfIsUploaded, setPdfIsUploaded] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(false);
  const [chosenFile, setChosenFile] = useState('');
  const fileIsChosen = chosenFile !== '';

  const getFileInput = () => {
    !isLoading && !error && inputRef?.current?.click();
  };

  useEffect(() => {
    if (typeof initialFile === 'object') {
      if (initialFile.type.includes('pdf')) {
        setPdfIsUploaded(true);
        setChosenFile(initialFile.name);
      } else setChosenFile(URL.createObjectURL(initialFile));
    } else setChosenFile(initialFile ?? '');
  }, [initialFile]);

  useEffect(() => {
    const timeout = setTimeout(() => {
      setError(false);
    }, 500);
    return () => {
      clearTimeout(timeout);
    };
  }, [error]);

  const setImg = (evt: ChangeEvent<HTMLInputElement>) => {
    if (evt.target.files && evt.target.files.length === 1) {
      const file = evt.target.files[0];
      setIsLoading(true);
      if (checkFile(file, type)) {
        setPdfIsUploaded(false);
        onChange?.(file);
        if (file.type === 'application/pdf') {
          setPdfIsUploaded(true);
          setChosenFile(file.name);
        } else setChosenFile(URL.createObjectURL(file));
      } else setError(true);
      setIsLoading(false);
    }
  };

  return (
    <>
      <div
        className={`select-none relative overflow-hidden flex flex-col items-center justify-center ${getStylesByError(
          error,
        )} ${getStylesByType(type)} ${getStylesByState(fileIsChosen, type)}`}
        onClick={getFileInput}
      >
        {!isLoading && (
          <>
            {type === 'document' && !fileIsChosen && (
              <DocumentIcon className={'fill-orange-main stroke-transparent'} />
            )}
            {!fileIsChosen && <p className={`text-center ${getTextStyles(type)}`}>{text}</p>}
            {pdfIsUploaded && fileIsChosen && (
              <p className={'absolute bottom-10 font-semibold text-md'}>{chosenFile}</p>
            )}
            {!pdfIsUploaded && fileIsChosen && (
              <img className={`${getImgStyles(type)} h-full w-full`} src={chosenFile} alt={''} />
            )}
            {fileIsChosen && (
              <button
                className={`${getButtonStyles(type)} z-10 rounded-full absolute`}
                type='button'
              >
                <EditIcon className={'stroke-transparent'} />
              </button>
            )}
          </>
        )}
        {isLoading && (
          <>
            <LoadingIcon className={'animate-spin'} />
            <p className={'absolute bottom-10 font-semibold text-md animate-pulse'}>Loading file</p>
          </>
        )}
      </div>
      <input type='file' hidden={true} ref={inputRef} accept='image/*' onChange={setImg} />
    </>
  );
};

export default UploadFile;
