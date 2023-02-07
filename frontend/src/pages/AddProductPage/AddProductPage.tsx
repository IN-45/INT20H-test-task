import React from 'react';

import { Field, Form, Formik } from 'formik';
import TextInput from '../../components/TextInput';
import Button from '../../components/Button';

const AddRecipePage: React.FC = () => {
  const initialValues = {
    name: '',
  };

  const handleSubmitForm = (values: any) => {
    console.log(values);
  };

  return (
    <div>
      <h1 className='text-xl my-7 text-center'>ADD NEW PRODUCT</h1>
      <div className='mt-6 max-w-[770px] mx-auto'>
        <Formik initialValues={initialValues} onSubmit={handleSubmitForm} validateOnMount>
          {({ isSubmitting, isValid }) => (
            <Form>
              <div className='space-y-6 '>
                <Field
                  id='name'
                  component={TextInput}
                  name='name'
                  placeholder={'e.g. Hot Pot'}
                  label={'Recipe title'}
                />
              </div>
              <Button
                type='submit'
                disabled={isSubmitting || !isValid}
                label='Save'
                className='px-7 py-1 mt-7'
              />
            </Form>
          )}
        </Formik>
      </div>
    </div>
  );
};

export default AddRecipePage;
