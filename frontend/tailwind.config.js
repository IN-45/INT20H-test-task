/** @type {import('tailwindcss').Config} */
// eslint-disable-next-line no-undef
module.exports = {
  content: ['./src/**/*.{js,jsx,ts,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        'open-sans': ['Open Sans', 'sans-serif'],
      },
      backgroundColor: {
        'navbar-link':
          {
            'active': '#008080',
          },
        navbar: '#001f3f',
      },
    },
  },
  plugins: [],
};
