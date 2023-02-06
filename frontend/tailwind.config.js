/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{js,jsx,ts,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        'open-sans': '"Open Sans", sans-serif',
      },
      colors: {
        transparent: '#00000000',
        black: {
          main: '#2D2727',
          secondary: '#393333',
        },
        orange: {
          main: '#FF845E',
        },
      },
      fontSize: {
        sm: ['12px', '15px'],
        md: ['16px', '20px'],
        lg: ['20px', '30px'],
        xl: ['32px', '36px'],
        '2xl': ['60px', '70px'],
      },
      screens: {
        xs: '660px',
        sm: '770px',
        md: '980px',
        lg: '1100px',
        xl: '1380px',
      },
    },
  },
  plugins: [],
};
