/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,jsx,ts,tsx}", "node_modules/flowbite-react/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        'inputColor': 'rgba(43, 47, 54, 0.8)',
      },
    },
    fontFamily:{
      main:["Poppins","sans-serif"],
    },
  },
   plugins: [
        require('flowbite/plugin')
    ],
};
