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
        background: '#FCFCFC',
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
      spacing: {
        0.25: '0.125rem', // 2px
        1.25: '0.313rem', // 5px
        1.5: '0.375rem', // 6px
        1.75: '0.438rem', // 7px
        2.25: '0.563rem', // 9px
        2.5: '0.625rem', // 10px
        2.75: '0.688rem', // 11px,
        3.25: '0.8125rem', // 13px,
        4.5: '1.125rem', // 18px
        5: '1.25rem', // 20px
        5.5: '1.375rem', // 22px
        5.75: '1.438rem', // 23px
        6: '1.5rem', // 24px
        6.25: '1.563rem', // 25px
        7.25: '1.813rem', // 29px
        7.3: '1.875rem', // 30px
        8: '2rem', // 32px
        8.5: '2.125rem', // 34px
        8.75: '2.188rem', // 35px
        9.25: '2.313rem', // 37px
        9.5: '2.375rem', // 38px
        10: '2.5rem', // 40px
        10.5: '2.625rem', // 42px
        11.25: '2.813rem', // 45px
        11.5: '3.25rem', // 52px
        13.5: '3.375rem', // 54px
        14: '3.5rem', // 56px
        14.5: '3.625rem', // 58px
        16.5: '4.125rem', // 66px
        17.5: '4.375rem', // 70px
        18: '4.5rem', // 72px
        18.5: '4.625rem', // 74px
        22: '5.5rem', // 88px
        22.5: '5.625rem', // 90px
        26: '6.5rem', // 104px
        30: '7.5rem', // 120px
        31: '7.75rem', // 124px
        32.25: '8.063rem', // 129px
        32.5: '8.125rem', // 130px
        33.5: '8.375rem', // 134
        36.25: '9.063rem', // 145px
        37: '9.25rem', // 148px
        39.75: '9.938rem', // 159px
        40: '10rem', // 160px
        40.25: '10.063rem', // 161px
        42: '10.5rem', // 168px
        46: '11.5rem', // 184px
        55: '13.75rem', // 220px
        58: '14.5rem', // 232px
        58.75: '14.688rem', // 235px
        62.5: '15.625rem', // 250px
        68.5: '17.125rem', // 274px
        75: '18.75rem', // 300px
        78: '19.5rem', // 312px
        84: '21rem', // 336px
        85.75: '21.438rem', // 343px
        88.5: '22.125rem', // 354
        89: '22.25rem', // 356px
        113.25: '28.313rem', // 453px
        117: '29.25rem', // 468px
        120: '30rem', // 480px
        126.75: '31.688rem', // 507px
        133: '33.25rem', // 532px
        136: '34rem', // 544 px
      },
    },
  },
  plugins: [],
};
