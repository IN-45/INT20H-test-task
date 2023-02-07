import * as yup from 'yup';

export const RecipeSchema = yup.object().shape({
  name: yup
    .string()
    .required('Please provide a recipe name')
    .max(100, 'Title should not exceed 100 characters')
    .matches(/^((?![<>]).)*$/, "Recipe title can't contain special characters < >"),
  description: yup
    .string()
    .required('Please provide a short description')
    .max(200, 'Description should not exceed 200 characters')
    .matches(/^((?![<>]).)*$/, "Recipe description address can't contain special characters < >"),
  cooking_time_minutes: yup
    .number()
    .typeError('Cooking time must be a number')
    .required('Please provide a cooking time duration')
    .min(0, 'Cooking time should be more than or equal to 0'),
  instructions: yup.array().of(yup.string().required('Please provide a step description')),
  products: yup.array().of(
    yup.object({
      product_id: yup.string().required('Please choose a product'),
      amount: yup
        .number()
        .typeError('Amount must be a number')
        .required('Please provide amount of a product')
        .min(0),
    }),
  ),
});
