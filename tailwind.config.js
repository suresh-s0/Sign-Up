/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,jsx,css}",
  ],
  theme: {
    extend: {
      fontFamily: {
        monospace: ['"Courier New"', 'Courier', 'monospace'],  // Proper font family array
      },
    },
  },
  plugins: [],
}
