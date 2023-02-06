import { createBrowserRouter } from 'react-router-dom';
import RootLayout from '../features/layout/RootLayout';

import NotFound from './NotFound';
import RecipePage from './RecipePage';

/* #TODO implement all the routes here
 * descriptions of paths as well as suggestions should go here
 * (dishes/recipes)/:id for dishes/recipes(choose one of the names)
 * /products to view entire list of available products or just the products that user have(choose one or implement both?)
 * /(dishes/recipes)/available for viewing what we can make with available products
 * OPTIONAL BELOW ***
 * /(dishes/recipes)/add for adding new dishes/recipes
 * /products/add to add new products
 * ?? sort by difficulty
 * ?? look for closest shops
 *  */

const router = createBrowserRouter([
  {
    path: '/',
    element: <RootLayout />,
    children: [
      {
        path: 'recipe/:id',
        element: <RecipePage />,
      },
      {
        path: '*',
        element: <NotFound />,
      },
    ],
  },
]);

export default router;
