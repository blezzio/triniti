/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./assets/templates/*.gohtml"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
}

