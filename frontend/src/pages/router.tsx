import { createBrowserRouter } from 'react-router-dom';
import RootLayout from '../features/layout/RootLayout';
import Products from './Products/Products';
import SignPage from './Sign/SignPage';
import NotFound from './NotFound';
import RecipePage from './RecipePage';
import AddRecipePage from './AddRecipePage';
import Recipes from './Recipes/Recipes';

const router = createBrowserRouter([
  {
    path: '/',
    element: <RootLayout />,
    children: [
      {
        path: 'products',
        element: <Products />,
      },
      {
        path: 'recipe',
        children: [
          {
            path: 'new',
            element: <AddRecipePage />,
          },
          {
            path: ':id',
            element: <RecipePage />,
          },
        ],
      },
      {
        path: 'recipe',
        element: <Recipes />,
      },
      {
        path: 'products',
        element: <Products />,
      },
      {
        path: 'sign-in',
        element: <SignPage action={'sign-in'} />,
      },
      {
        path: 'sign-up',
        element: <SignPage action={'sign-up'} />,
      },
      {
        path: '*',
        element: <NotFound />,
      },
    ],
  },
]);

export default router;
