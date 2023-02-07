import React, { FC, FormEvent, useState } from 'react';
import { requestApi } from '../../config/api';
import { Link, useNavigate } from 'react-router-dom';
import TextInput from '../../components/TextInput';
import { ErrorMessage } from 'formik';

interface SignPageProps {
  action: string;
}

const SignPage: FC<SignPageProps> = ({ action }) => {
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string>();

  const handleSubmit = (evt: FormEvent<HTMLFormElement>, type: string) => {
    setError(undefined);
    evt.preventDefault();
    setIsLoading(true);
    if (type === 'sign-up') {
      requestApi('POST', 'sign-up', { data: evt.target })
        .then(() => handleSubmit(evt, 'sign-in'))
        .catch((err) => setError(err.response.data));
    } else {
      requestApi('POST', 'sign-in', { data: evt.target })
        .then(() => {
          navigate('/products');
        })
        .catch((err) => setError(err.response.data));
    }
    setIsLoading(false);
  };

  return (
    <form
      className={'w-1/2 mx-auto'}
      onSubmit={(evt) => {
        handleSubmit(evt, action);
      }}
    >
      <TextInput
        label={'Email'}
        name={'email'}
        id={'email'}
        type={'email'}
        className={'mb-4'}
        required
      />
      <TextInput
        label={'Password'}
        name={'password'}
        id={'password'}
        type={'password'}
        className={'mb-4'}
        required
      />
      {!!error && <p className={'text-red-500'}>{error}</p>}
      <div className={'text-center'}>
        <button
          className={
            'bg-blue-500 hover:bg-blue-600 active:bg-blue-700 text-white font-bold py-2 px-4 rounded select-none'
          }
          disabled={isLoading}
        >
          {action === 'sign-in' ? 'Login' : 'Sign Up'}
        </button>
      </div>
      {action === 'sign-in' && (
        <p className={'text-gray-500 mt-4'}>
          Don&apos;t have an account ?{' '}
          <Link className={'text-blue-700 hover:text-blue-800 underline'} to={'/sign-up'}>
            Create new account here
          </Link>
        </p>
      )}
    </form>
  );
};

export default SignPage;