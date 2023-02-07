import React, { useEffect, useState } from 'react';

import { Field, FieldArray, Form, Formik } from 'formik';
import TextInput from '../../components/TextInput';
import { RecipeSchema } from './validationSchema';
import Button from '../../components/Button';
import DataFetcher from '../../components/DataFetcher/DataFetcher';
import SelectInput from '../../components/SelectInput';

interface RecipeForm {
  name: string;
  description: string;
  cooking_time_minutes: string;
  instructions: string[];
  products: Product[];
}

interface Product {
  product_id: string;
  amount: string;
  amount_type: string;
}

interface ProductRaw {
  amount_type: string;
  category: string;
  id: string;
  image_url: string;
  name: string;
}

const AddRecipePage: React.FC = () => {
  const [initialValues, setInitialValues] = useState({
    name: '',
    description: '',
    cooking_time_minutes: '',
    instructions: [''],
    products: [{ product_id: '', amount: '', amount_type: '' }],
  });

  const handleSubmitForm = (values: RecipeForm) => {
    const recipeData = {
      name: values.name,
      description: values.description,
      cooking_time_minutes: parseInt(values.cooking_time_minutes),
      instructions: values.instructions.map((instruction, index) => {
        return {
          priority: index + 1,
          description: instruction,
        };
      }),
      products: values.products.map((product) => {
        return {
          product_id: product.product_id,
          amount: parseFloat(product.amount),
          amount_type: product.amount_type,
        };
      }),
    };

    console.log(recipeData);
  };

  const addInstruction = (values: RecipeForm) => {
    const current = values;
    current.instructions.push('');
    setInitialValues({ ...current });
  };

  const addProduct = (values: RecipeForm) => {
    const current = values;
    current.products.push({ product_id: '', amount: '', amount_type: '' });
    setInitialValues({ ...current });
  };

  const deleteInstruction = (values: RecipeForm) => {
    const current = values;
    current.instructions.pop();
    setInitialValues({ ...current });
  };

  const deleteProduct = (values: RecipeForm) => {
    const current = values;
    current.products.pop();
    setInitialValues({ ...current });
  };

  const convertProducts = (products: ProductRaw[]) => {
    return products.map((product) => {
      return {
        value: product.id,
        label: product.name,
      };
    });
  };

  const findAmountType = (products: ProductRaw[], id: string) => {
    console.log(products.filter((product) => product.id === id)[0]?.amount_type);
    return products.filter((product) => product.id === id)[0]?.amount_type || '';
  };

  return (
    <div>
      <h1 className='text-xl my-7 text-center'>ADD NEW RECIPE</h1>
      <div className='mt-6 max-w-[770px] mx-auto'>
        <Formik
          initialValues={initialValues}
          onSubmit={handleSubmitForm}
          validationSchema={RecipeSchema}
          enableReinitialize={true}
          validateOnMount
        >
          {({ isSubmitting, isValid, values }) => (
            <Form>
              <div className='space-y-6'>
                <Field
                  id='name'
                  component={TextInput}
                  name='name'
                  placeholder={'e.g. Hot Pot'}
                  label={'Recipe title'}
                />
                <Field
                  id='description'
                  component={TextInput}
                  name='description'
                  placeholder={'e.g. Best Chinese food dish!'}
                  label={'Recipe description'}
                />
                <Field
                  id='cooking_time_minutes'
                  component={TextInput}
                  name='cooking_time_minutes'
                  placeholder={'e.g. 15'}
                  label={'Cooking time (in minutes)'}
                />
                <div className='font-bold text-lg'>Recipe Instructions</div>
                <FieldArray name='instructions'>
                  {() =>
                    values.instructions.map((instruction, index) => {
                      return (
                        <div key={index}>
                          <Field
                            id={`instructions.${index}`}
                            component={TextInput}
                            name={`instructions.${index}`}
                            placeholder={'e.g. Boil for 15 mins'}
                            label={`Step №${index + 1} description`}
                          />
                        </div>
                      );
                    })
                  }
                </FieldArray>
                <Button
                  type='button'
                  label='+ add step'
                  className='px-2'
                  onClick={() => addInstruction(values)}
                />
                {initialValues.instructions.length > 1 && (
                  <Button
                    type='button'
                    label='х delete last step'
                    className='px-2 ml-3 bg-red-secondary border-red-secondary hover:text-red-secondary'
                    onClick={() => deleteInstruction(values)}
                  />
                )}
                <div className='font-bold text-lg'>Products List</div>
                <DataFetcher url={'products'}>
                  {(products) => (
                    <FieldArray name='products'>
                      {() =>
                        values.products.map((product, index) => {
                          return (
                            <div key={index} className='space-y-2'>
                              <div className='font-bold text-md mb-2'>Product {index + 1}</div>
                              <Field
                                id={`products.${index}.product_id`}
                                component={SelectInput}
                                options={convertProducts(products)}
                                name={`products.${index}.product_id`}
                                placeholder={'e.g. Sugar'}
                                label='Product'
                              />
                              <Field
                                id={`products.${index}.amount`}
                                component={TextInput}
                                name={`products.${index}.amount`}
                                placeholder={'e.g. 2'}
                                label={'Amount of product (in ' + findAmountType(products) + ')'}
                              />
                              <Button
                                type='button'
                                label='hhohohooh'
                                onClick={() => console.log(products)}
                              />
                            </div>
                          );
                        })
                      }
                    </FieldArray>
                  )}
                </DataFetcher>
                <Button
                  type='button'
                  label='+ add product'
                  className='px-2'
                  onClick={() => addProduct(values)}
                />
                {initialValues.products.length > 1 && (
                  <Button
                    type='button'
                    label='х delete last product'
                    className='px-2 ml-3 bg-red-secondary border-red-secondary hover:text-red-secondary'
                    onClick={() => deleteProduct(values)}
                  />
                )}
              </div>
              <Button
                type='submit'
                disabled={isSubmitting || !isValid}
                label='Save'
                className='w-full mt-16 py-2'
              />
            </Form>
          )}
        </Formik>
      </div>
    </div>
  );
};

export default AddRecipePage;
