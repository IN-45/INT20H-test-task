/** @type {import('tailwindcss').Config} */
// eslint-disable-next-line no-undef
module.exports = {
  content: ['./src/**/*.{js,jsx,ts,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        'open-sans': ['Open Sans', 'sans-serif'],
      },
      boxShadow: {
        'main-box-shadow': '0px 4px 16px rgba(29, 101, 137, 0.15)',
      },
      backgroundColor: {
        'navbar-link': {
          active: '#008080',
        },
        navbar: '#001f3f',
      },
      colors: {
        transparent: '#00000000',
        black: {
          main: '#2D2727',
          secondary: '#393333',
        },
        grey: {
          greyscale: '#909590',
          'greyscale-2': '#D4D5D4',
        },
        orange: {
          main: '#FF845E',
        },
        red: {
          secondary: '#D34E24',
        },
        white: {
          greyscale: '#FFFFFF',
          disabled: '#F2F2F2',
        },
        chosen: '#F2F5F7',
      },
      fontSize: {
        '2xs': ['13px', '20px'],
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
