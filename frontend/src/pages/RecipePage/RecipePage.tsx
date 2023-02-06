import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

import AuthorTag from '../../components/AuthorTag';
import DurationTag from '../../components/DurationTag';
import ProductsList from '../../components/IngredientsList/IngredientsList';
import Instructions from '../../components/Instructions';
import TimeTag from '../../components/TimeTag';

const RecipePage: React.FC = () => {
  const { id } = useParams();
  // const [recipe, setRecipe] = useState({});

  // useEffect(() => {
  //   const apiUrl = 'http://localhost:8000/recipe/' + { id };
  //   const getRecipe = async () => {
  //     await axios.get(apiUrl).then((response) => setRecipe(response));
  //   };
  // });

  const recipe = {
    name: 'Strawberry Cream Cheesecake',
    description:
      'One thing I learned living in the Canarsie section of Brooklyn, NY was how to cook a good Italian meal. Here is a recipe I created after having this dish in a restaurant. Enjoy!',
    cooking_time_minutes: 45,
    image: require('../../assets/example.jpg'),
    author: 'Єлизавета Столярчук',
    created_at: new Date(),
  };

  const products = [
    {
      name: 'graham crackers',
      amount: '400g',
    },
    {
      name: 'unsalted butters, melted',
      amount: '150g',
    },
    {
      name: 'marshmallows',
      amount: '300g',
    },
  ];

  const instructions = [
    'To prepare crust add graham crackers to a food processor and process until you reach fine crumbs. Add melted butter and pulse 3-4 times to coat crumbs with butter.',
    'Pour mixture into a 20cm (8”) tart tin. Use the back of a spoon to firmly press the mixture out across the bottom and sides of the tart tin. Chill for 30 min.',
    'Begin by adding the marshmallows and melted butter into a microwave safe bowl. Microwave for 30 seconds and mix to combine. Set aside.',
    'Next, add the gelatine and water to a small mixing bowl and mix to combine. Microwave for 30 seconds.',
  ];

  return (
    <div>
      <h1 className='text-2xl mb-7'>{recipe.name}</h1>
      <div className='flex description mb-5'>
        <AuthorTag author={recipe.author} />
        <TimeTag time={recipe.created_at} />
        <DurationTag duration={recipe.cooking_time_minutes} />
      </div>
      <div className='mb-5 text-lg'>{recipe.description}</div>
      <img
        src={recipe.image}
        alt={recipe.name}
        className='h-[320px] md:h-[440px] lg:h-[500px] w-full object-cover'
      />
      <hr className='border-orange-main border-4 mb-9' />
      <div className='grid md:grid-cols-2 gap-5'>
        <ProductsList products={products} />
        <Instructions instructions={instructions} />
      </div>
    </div>
  );
};

export default RecipePage;
